// +build windows

package conntrackwrapper

type conntrack struct {
}

func New() ConntrackWrapper {
	return &conntrack{}
}

func (c *conntrack) ConntrackTableUpdateMark(ipSrc, ipDst string, protonum uint8, srcport, dstport uint16, newmark uint32) error {
	return nil
}
