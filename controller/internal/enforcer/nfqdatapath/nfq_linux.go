// +build linux

package nfqdatapath

// Go libraries
import (
	"context"
	"fmt"
	"strconv"
	"time"

	nfqueue "go.aporeto.io/netlink-go/nfqueue"
	"go.aporeto.io/trireme-lib/controller/pkg/claimsheader"
	"go.aporeto.io/trireme-lib/controller/pkg/connection"
	"go.aporeto.io/trireme-lib/controller/pkg/packet"
	markconstants "go.aporeto.io/trireme-lib/utils/constants"
	"go.uber.org/zap"
)

func errorCallback(err error, _ interface{}) {
	zap.L().Error("Error while processing packets on queue", zap.Error(err))
}
func networkCallback(packet *nfqueue.NFPacket, d interface{}) {
	d.(*Datapath).processNetworkPacketsFromNFQ(packet)
}

func appCallBack(packet *nfqueue.NFPacket, d interface{}) {
	d.(*Datapath).processApplicationPacketsFromNFQ(packet)
}

// startNetworkInterceptor will the process that processes  packets from the network
// Still has one more copy than needed. Can be improved.
func (d *Datapath) startNetworkInterceptor(ctx context.Context) {
	var err error

	nfq := make([]nfqueue.Verdict, d.filterQueue.GetNumNetworkQueues())

	for i := uint16(0); i < d.filterQueue.GetNumNetworkQueues(); i++ {
		// Initialize all the queues
		nfq[i], err = nfqueue.CreateAndStartNfQueue(ctx, d.filterQueue.GetNetworkQueueStart()+i, d.filterQueue.GetNetworkQueueSize(), nfqueue.NfDefaultPacketSize, networkCallback, errorCallback, d)
		if err != nil {
			for retry := 0; retry < 5 && err != nil; retry++ {
				nfq[i], err = nfqueue.CreateAndStartNfQueue(ctx, d.filterQueue.GetNetworkQueueStart()+i, d.filterQueue.GetNetworkQueueSize(), nfqueue.NfDefaultPacketSize, networkCallback, errorCallback, d)
				<-time.After(3 * time.Second)
			}
			if err != nil {
				zap.L().Fatal("Unable to initialize netfilter queue", zap.Error(err))
			}
		}
	}
}

// startApplicationInterceptor will create a interceptor that processes
// packets originated from a local application
func (d *Datapath) startApplicationInterceptor(ctx context.Context) {
	var err error

	nfq := make([]nfqueue.Verdict, d.filterQueue.GetNumApplicationQueues())

	for i := uint16(0); i < d.filterQueue.GetNumApplicationQueues(); i++ {
		nfq[i], err = nfqueue.CreateAndStartNfQueue(ctx, d.filterQueue.GetApplicationQueueStart()+i, d.filterQueue.GetApplicationQueueSize(), nfqueue.NfDefaultPacketSize, appCallBack, errorCallback, d)

		if err != nil {
			for retry := 0; retry < 5 && err != nil; retry++ {
				nfq[i], err = nfqueue.CreateAndStartNfQueue(ctx, d.filterQueue.GetApplicationQueueStart()+i, d.filterQueue.GetApplicationQueueSize(), nfqueue.NfDefaultPacketSize, appCallBack, errorCallback, d)
				<-time.After(3 * time.Second)
			}
			if err != nil {
				zap.L().Fatal("Unable to initialize netfilter queue", zap.Int("QueueNum", int(d.filterQueue.GetNetworkQueueStart()+i)), zap.Error(err))
			}

		}
	}
}

// processNetworkPacketsFromNFQ processes packets arriving from the network in an NF queue
func (d *Datapath) processNetworkPacketsFromNFQ(p *nfqueue.NFPacket) {
	// Parse the packet - drop if parsing fails
	netPacket, err := packet.New(packet.PacketTypeNetwork, p.Buffer, strconv.Itoa(p.Mark), true)
	var processError error
	var tcpConn *connection.TCPConnection
	var udpConn *connection.UDPConnection
	if err != nil {
		netPacket.Print(packet.PacketFailureCreate, d.packetLogs)
	} else if netPacket.IPProto() == packet.IPProtocolTCP {
		tcpConn, processError = d.processNetworkTCPPackets(netPacket)
	} else if netPacket.IPProto() == packet.IPProtocolUDP {
		udpConn, processError = d.ProcessNetworkUDPPacket(netPacket)
	} else {
		processError = fmt.Errorf("invalid ip protocol: %d", netPacket.IPProto())

	}

	// TODO: Use error types and handle it in switch case here
	if processError != nil {
		zap.L().Debug("Dropping packet on network path",
			zap.Error(processError),
			zap.String("SourceIP", netPacket.SourceAddress().String()),
			zap.String("DestiatnionIP", netPacket.DestinationAddress().String()),
			zap.Int("SourcePort", int(netPacket.SourcePort())),
			zap.Int("DestinationPort", int(netPacket.DestPort())),
			zap.Int("Protocol", int(netPacket.IPProto())),
			zap.String("Flags", packet.TCPFlagsToStr(netPacket.GetTCPFlags())),
		)
		length := uint32(len(p.Buffer))
		buffer := p.Buffer
		p.QueueHandle.SetVerdict2(uint32(p.QueueHandle.QueueNum), 0, uint32(p.Mark)&^(markconstants.NFQueueMask|markconstants.NFSetMarkMask), length, uint32(p.ID), buffer)
		if netPacket.IPProto() == packet.IPProtocolTCP {
			d.collectTCPPacket(&debugpacketmessage{
				Mark:    p.Mark,
				p:       netPacket,
				tcpConn: tcpConn,
				udpConn: nil,
				err:     processError,
				network: true,
			})
		} else if netPacket.IPProto() == packet.IPProtocolUDP {
			d.collectUDPPacket(&debugpacketmessage{
				Mark:    p.Mark,
				p:       netPacket,
				tcpConn: nil,
				udpConn: udpConn,
				err:     processError,
				network: true,
			})
		}

		return
	}

	v := uint32(1)
	if tcpConn != nil {
		if !tcpConn.PingConfig.Passthrough && tcpConn.PingConfig.Type != claimsheader.PingTypeNone {
			v = uint32(0)
		}
	}

	if netPacket.IPProto() == packet.IPProtocolTCP {
		// // Accept the packet
		buffer := make([]byte, netPacket.IPTotalLen())
		copyIndex := copy(buffer, netPacket.GetBuffer(0))
		copyIndex += copy(buffer[copyIndex:], netPacket.GetTCPOptions())
		copyIndex += copy(buffer[copyIndex:], netPacket.GetTCPData())

		p.QueueHandle.SetVerdict2(uint32(p.QueueHandle.QueueNum), v, uint32(p.Mark)&^(markconstants.NFQueueMask|markconstants.NFSetMarkMask), uint32(copyIndex), uint32(p.ID), buffer)
	} else {
		p.QueueHandle.SetVerdict2(uint32(p.QueueHandle.QueueNum), v, uint32(p.Mark)&^(markconstants.NFQueueMask|markconstants.NFSetMarkMask), uint32(len(netPacket.GetBuffer(0))), uint32(p.ID), netPacket.GetBuffer(0))
	}

	if netPacket.IPProto() == packet.IPProtocolTCP {
		d.collectTCPPacket(&debugpacketmessage{
			Mark:    p.Mark,
			p:       netPacket,
			tcpConn: tcpConn,
			udpConn: nil,
			err:     nil,
			network: true,
		})
	} else if netPacket.IPProto() == packet.IPProtocolUDP {
		d.collectUDPPacket(&debugpacketmessage{
			Mark:    p.Mark,
			p:       netPacket,
			tcpConn: nil,
			udpConn: udpConn,
			err:     nil,
			network: true,
		})
	}

}

// processApplicationPackets processes packets arriving from an application and are destined to the network
func (d *Datapath) processApplicationPacketsFromNFQ(p *nfqueue.NFPacket) {

	// Being liberal on what we transmit - malformed TCP packets are let go
	// We are strict on what we accept on the other side, but we don't block
	// lots of things at the ingress to the network
	appPacket, err := packet.New(packet.PacketTypeApplication, p.Buffer, strconv.Itoa(p.Mark), true)

	var processError error
	var tcpConn *connection.TCPConnection
	var udpConn *connection.UDPConnection
	if err != nil {
		appPacket.Print(packet.PacketFailureCreate, d.packetLogs)
	} else if appPacket.IPProto() == packet.IPProtocolTCP {
		tcpConn, processError = d.processApplicationTCPPackets(appPacket)
	} else if appPacket.IPProto() == packet.IPProtocolUDP {
		udpConn, processError = d.ProcessApplicationUDPPacket(appPacket)
	} else {
		processError = fmt.Errorf("invalid ip protocol: %d", appPacket.IPProto())
	}
	if processError != nil {
		zap.L().Debug("Dropping packet on app path",
			zap.Error(processError),
			zap.String("SourceIP", appPacket.SourceAddress().String()),
			zap.String("DestiatnionIP", appPacket.DestinationAddress().String()),
			zap.Int("SourcePort", int(appPacket.SourcePort())),
			zap.Int("DestinationPort", int(appPacket.DestPort())),
			zap.Int("Protocol", int(appPacket.IPProto())),
			zap.String("Flags", packet.TCPFlagsToStr(appPacket.GetTCPFlags())),
		)

		length := uint32(len(p.Buffer))
		buffer := p.Buffer
		p.QueueHandle.SetVerdict2(uint32(p.QueueHandle.QueueNum), 0, uint32(p.Mark)&^(markconstants.NFQueueMask|markconstants.NFSetMarkMask), length, uint32(p.ID), buffer)
		if appPacket.IPProto() == packet.IPProtocolTCP {

			d.collectTCPPacket(&debugpacketmessage{
				Mark:    p.Mark,
				p:       appPacket,
				tcpConn: tcpConn,
				udpConn: nil,
				err:     processError,
				network: false,
			})
		} else if appPacket.IPProto() == packet.IPProtocolUDP {
			d.collectUDPPacket(&debugpacketmessage{
				Mark:    p.Mark,
				p:       appPacket,
				tcpConn: nil,
				udpConn: udpConn,
				err:     processError,
				network: false,
			})
		}
		return
	}

	if appPacket.IPProto() == packet.IPProtocolTCP {
		// Accept the packet
		buffer := make([]byte, appPacket.IPTotalLen())
		copyIndex := copy(buffer, appPacket.GetBuffer(0))
		copyIndex += copy(buffer[copyIndex:], appPacket.GetTCPOptions())
		copyIndex += copy(buffer[copyIndex:], appPacket.GetTCPData())

		p.QueueHandle.SetVerdict2(uint32(p.QueueHandle.QueueNum), 1, uint32(p.Mark)&^(markconstants.NFQueueMask|markconstants.NFSetMarkMask), uint32(copyIndex), uint32(p.ID), buffer)

	} else {
		p.QueueHandle.SetVerdict2(uint32(p.QueueHandle.QueueNum), 1, uint32(p.Mark)&^(markconstants.NFQueueMask|markconstants.NFSetMarkMask), uint32(len(appPacket.GetBuffer(0))), uint32(p.ID), appPacket.GetBuffer(0))
	}

	if appPacket.IPProto() == packet.IPProtocolTCP {
		d.collectTCPPacket(&debugpacketmessage{
			Mark:    p.Mark,
			p:       appPacket,
			tcpConn: tcpConn,
			udpConn: nil,
			err:     nil,
			network: false,
		})
	} else if appPacket.IPProto() == packet.IPProtocolUDP {
		d.collectUDPPacket(&debugpacketmessage{
			Mark:    p.Mark,
			p:       appPacket,
			tcpConn: nil,
			udpConn: udpConn,
			err:     nil,
			network: false,
		})
	}

}

func (d *Datapath) cleanupPlatform() {}
