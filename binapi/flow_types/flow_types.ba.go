// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.09-release
// source: /usr/share/vpp/api/core/flow_types.api.json

// Package flow_types contains generated bindings for API file flow_types.api.
//
// Contents:
//   2 enums
//  17 structs
//   1 union
//
package flow_types

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	ethernet_types "github.com/edwarnicke/govpp/binapi/ethernet_types"
	ip_types "github.com/edwarnicke/govpp/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

// FlowAction defines enum 'flow_action'.
type FlowAction uint32

const (
	FLOW_ACTION_COUNT             FlowAction = 1
	FLOW_ACTION_MARK              FlowAction = 2
	FLOW_ACTION_BUFFER_ADVANCE    FlowAction = 4
	FLOW_ACTION_REDIRECT_TO_NODE  FlowAction = 8
	FLOW_ACTION_REDIRECT_TO_QUEUE FlowAction = 16
	FLOW_ACTION_DROP              FlowAction = 64
)

var (
	FlowAction_name = map[uint32]string{
		1:  "FLOW_ACTION_COUNT",
		2:  "FLOW_ACTION_MARK",
		4:  "FLOW_ACTION_BUFFER_ADVANCE",
		8:  "FLOW_ACTION_REDIRECT_TO_NODE",
		16: "FLOW_ACTION_REDIRECT_TO_QUEUE",
		64: "FLOW_ACTION_DROP",
	}
	FlowAction_value = map[string]uint32{
		"FLOW_ACTION_COUNT":             1,
		"FLOW_ACTION_MARK":              2,
		"FLOW_ACTION_BUFFER_ADVANCE":    4,
		"FLOW_ACTION_REDIRECT_TO_NODE":  8,
		"FLOW_ACTION_REDIRECT_TO_QUEUE": 16,
		"FLOW_ACTION_DROP":              64,
	}
)

func (x FlowAction) String() string {
	s, ok := FlowAction_name[uint32(x)]
	if ok {
		return s
	}
	return "FlowAction(" + strconv.Itoa(int(x)) + ")"
}

// FlowType defines enum 'flow_type'.
type FlowType uint32

const (
	FLOW_TYPE_ETHERNET           FlowType = 1
	FLOW_TYPE_IP4                FlowType = 2
	FLOW_TYPE_IP6                FlowType = 3
	FLOW_TYPE_IP4_L2TPV3OIP      FlowType = 4
	FLOW_TYPE_IP4_IPSEC_ESP      FlowType = 5
	FLOW_TYPE_IP4_IPSEC_AH       FlowType = 6
	FLOW_TYPE_IP4_N_TUPLE        FlowType = 7
	FLOW_TYPE_IP6_N_TUPLE        FlowType = 8
	FLOW_TYPE_IP4_N_TUPLE_TAGGED FlowType = 9
	FLOW_TYPE_IP6_N_TUPLE_TAGGED FlowType = 10
	FLOW_TYPE_IP4_VXLAN          FlowType = 11
	FLOW_TYPE_IP6_VXLAN          FlowType = 12
	FLOW_TYPE_IP4_GTPC           FlowType = 13
	FLOW_TYPE_IP4_GTPU           FlowType = 14
)

var (
	FlowType_name = map[uint32]string{
		1:  "FLOW_TYPE_ETHERNET",
		2:  "FLOW_TYPE_IP4",
		3:  "FLOW_TYPE_IP6",
		4:  "FLOW_TYPE_IP4_L2TPV3OIP",
		5:  "FLOW_TYPE_IP4_IPSEC_ESP",
		6:  "FLOW_TYPE_IP4_IPSEC_AH",
		7:  "FLOW_TYPE_IP4_N_TUPLE",
		8:  "FLOW_TYPE_IP6_N_TUPLE",
		9:  "FLOW_TYPE_IP4_N_TUPLE_TAGGED",
		10: "FLOW_TYPE_IP6_N_TUPLE_TAGGED",
		11: "FLOW_TYPE_IP4_VXLAN",
		12: "FLOW_TYPE_IP6_VXLAN",
		13: "FLOW_TYPE_IP4_GTPC",
		14: "FLOW_TYPE_IP4_GTPU",
	}
	FlowType_value = map[string]uint32{
		"FLOW_TYPE_ETHERNET":           1,
		"FLOW_TYPE_IP4":                2,
		"FLOW_TYPE_IP6":                3,
		"FLOW_TYPE_IP4_L2TPV3OIP":      4,
		"FLOW_TYPE_IP4_IPSEC_ESP":      5,
		"FLOW_TYPE_IP4_IPSEC_AH":       6,
		"FLOW_TYPE_IP4_N_TUPLE":        7,
		"FLOW_TYPE_IP6_N_TUPLE":        8,
		"FLOW_TYPE_IP4_N_TUPLE_TAGGED": 9,
		"FLOW_TYPE_IP6_N_TUPLE_TAGGED": 10,
		"FLOW_TYPE_IP4_VXLAN":          11,
		"FLOW_TYPE_IP6_VXLAN":          12,
		"FLOW_TYPE_IP4_GTPC":           13,
		"FLOW_TYPE_IP4_GTPU":           14,
	}
)

func (x FlowType) String() string {
	s, ok := FlowType_name[uint32(x)]
	if ok {
		return s
	}
	return "FlowType(" + strconv.Itoa(int(x)) + ")"
}

// FlowEthernet defines type 'flow_ethernet'.
type FlowEthernet struct {
	Foo     int32                     `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr ethernet_types.MacAddress `binapi:"mac_address,name=src_addr" json:"src_addr,omitempty"`
	DstAddr ethernet_types.MacAddress `binapi:"mac_address,name=dst_addr" json:"dst_addr,omitempty"`
	Type    uint16                    `binapi:"u16,name=type" json:"type,omitempty"`
}

// FlowIP4 defines type 'flow_ip4'.
type FlowIP4 struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
}

// FlowIP4Gtpc defines type 'flow_ip4_gtpc'.
type FlowIP4Gtpc struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
	SrcPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=src_port" json:"src_port,omitempty"`
	DstPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=dst_port" json:"dst_port,omitempty"`
	Teid     uint32                     `binapi:"u32,name=teid" json:"teid,omitempty"`
}

// FlowIP4Gtpu defines type 'flow_ip4_gtpu'.
type FlowIP4Gtpu struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
	SrcPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=src_port" json:"src_port,omitempty"`
	DstPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=dst_port" json:"dst_port,omitempty"`
	Teid     uint32                     `binapi:"u32,name=teid" json:"teid,omitempty"`
}

// FlowIP4IpsecAh defines type 'flow_ip4_ipsec_ah'.
type FlowIP4IpsecAh struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
	Spi      uint32                     `binapi:"u32,name=spi" json:"spi,omitempty"`
}

// FlowIP4IpsecEsp defines type 'flow_ip4_ipsec_esp'.
type FlowIP4IpsecEsp struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
	Spi      uint32                     `binapi:"u32,name=spi" json:"spi,omitempty"`
}

// FlowIP4L2tpv3oip defines type 'flow_ip4_l2tpv3oip'.
type FlowIP4L2tpv3oip struct {
	Foo       int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr   ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr   ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol  IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
	SessionID uint32                     `binapi:"u32,name=session_id" json:"session_id,omitempty"`
}

// FlowIP4NTuple defines type 'flow_ip4_n_tuple'.
type FlowIP4NTuple struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
	SrcPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=src_port" json:"src_port,omitempty"`
	DstPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=dst_port" json:"dst_port,omitempty"`
}

// FlowIP4NTupleTagged defines type 'flow_ip4_n_tuple_tagged'.
type FlowIP4NTupleTagged struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
	SrcPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=src_port" json:"src_port,omitempty"`
	DstPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=dst_port" json:"dst_port,omitempty"`
}

// FlowIP4Vxlan defines type 'flow_ip4_vxlan'.
type FlowIP4Vxlan struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP4AddressAndMask `binapi:"ip4_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
	SrcPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=src_port" json:"src_port,omitempty"`
	DstPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=dst_port" json:"dst_port,omitempty"`
	Vni      uint16                     `binapi:"u16,name=vni" json:"vni,omitempty"`
}

// FlowIP6 defines type 'flow_ip6'.
type FlowIP6 struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP6AddressAndMask `binapi:"ip6_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP6AddressAndMask `binapi:"ip6_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
}

// FlowIP6NTuple defines type 'flow_ip6_n_tuple'.
type FlowIP6NTuple struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP6AddressAndMask `binapi:"ip6_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP6AddressAndMask `binapi:"ip6_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
	SrcPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=src_port" json:"src_port,omitempty"`
	DstPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=dst_port" json:"dst_port,omitempty"`
}

// FlowIP6NTupleTagged defines type 'flow_ip6_n_tuple_tagged'.
type FlowIP6NTupleTagged struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP6AddressAndMask `binapi:"ip6_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP6AddressAndMask `binapi:"ip6_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
	SrcPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=src_port" json:"src_port,omitempty"`
	DstPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=dst_port" json:"dst_port,omitempty"`
}

// FlowIP6Vxlan defines type 'flow_ip6_vxlan'.
type FlowIP6Vxlan struct {
	Foo      int32                      `binapi:"i32,name=foo" json:"foo,omitempty"`
	SrcAddr  ip_types.IP6AddressAndMask `binapi:"ip6_address_and_mask,name=src_addr" json:"src_addr,omitempty"`
	DstAddr  ip_types.IP6AddressAndMask `binapi:"ip6_address_and_mask,name=dst_addr" json:"dst_addr,omitempty"`
	Protocol IPProtAndMask              `binapi:"ip_prot_and_mask,name=protocol" json:"protocol,omitempty"`
	SrcPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=src_port" json:"src_port,omitempty"`
	DstPort  IPPortAndMask              `binapi:"ip_port_and_mask,name=dst_port" json:"dst_port,omitempty"`
	Vni      uint16                     `binapi:"u16,name=vni" json:"vni,omitempty"`
}

// FlowRule defines type 'flow_rule'.
type FlowRule struct {
	Type                         FlowType   `binapi:"flow_type,name=type" json:"type,omitempty"`
	Index                        uint32     `binapi:"u32,name=index" json:"index,omitempty"`
	Actions                      FlowAction `binapi:"flow_action,name=actions" json:"actions,omitempty"`
	MarkFlowID                   uint32     `binapi:"u32,name=mark_flow_id" json:"mark_flow_id,omitempty"`
	RedirectNodeIndex            uint32     `binapi:"u32,name=redirect_node_index" json:"redirect_node_index,omitempty"`
	RedirectDeviceInputNextIndex uint32     `binapi:"u32,name=redirect_device_input_next_index" json:"redirect_device_input_next_index,omitempty"`
	RedirectQueue                uint32     `binapi:"u32,name=redirect_queue" json:"redirect_queue,omitempty"`
	BufferAdvance                int32      `binapi:"i32,name=buffer_advance" json:"buffer_advance,omitempty"`
	Flow                         FlowUnion  `binapi:"flow,name=flow" json:"flow,omitempty"`
}

// IPPortAndMask defines type 'ip_port_and_mask'.
type IPPortAndMask struct {
	Port uint16 `binapi:"u16,name=port" json:"port,omitempty"`
	Mask uint16 `binapi:"u16,name=mask" json:"mask,omitempty"`
}

// IPProtAndMask defines type 'ip_prot_and_mask'.
type IPProtAndMask struct {
	Prot ip_types.IPProto `binapi:"ip_proto,name=prot" json:"prot,omitempty"`
	Mask uint8            `binapi:"u8,name=mask" json:"mask,omitempty"`
}

// FlowUnion defines union 'flow'.
type FlowUnion struct {
	// FlowUnion can be one of:
	// - Ethernet *FlowEthernet
	// - IP4 *FlowIP4
	// - IP6 *FlowIP6
	// - IP4L2tpv3oip *FlowIP4L2tpv3oip
	// - IP4IpsecEsp *FlowIP4IpsecEsp
	// - IP4IpsecAh *FlowIP4IpsecAh
	// - IP4NTuple *FlowIP4NTuple
	// - IP6NTuple *FlowIP6NTuple
	// - IP4NTupleTagged *FlowIP4NTupleTagged
	// - IP6NTupleTagged *FlowIP6NTupleTagged
	// - IP4Vxlan *FlowIP4Vxlan
	// - IP6Vxlan *FlowIP6Vxlan
	// - IP4Gtpc *FlowIP4Gtpc
	// - IP4Gtpu *FlowIP4Gtpu
	XXX_UnionData [80]byte
}

func FlowUnionEthernet(a FlowEthernet) (u FlowUnion) {
	u.SetEthernet(a)
	return
}
func (u *FlowUnion) SetEthernet(a FlowEthernet) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr[:], 6)
	buf.EncodeBytes(a.DstAddr[:], 6)
	buf.EncodeUint16(a.Type)
}
func (u *FlowUnion) GetEthernet() (a FlowEthernet) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr[:], buf.DecodeBytes(6))
	copy(a.DstAddr[:], buf.DecodeBytes(6))
	a.Type = buf.DecodeUint16()
	return
}

func FlowUnionIP4(a FlowIP4) (u FlowUnion) {
	u.SetIP4(a)
	return
}
func (u *FlowUnion) SetIP4(a FlowIP4) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 4)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 4)
	buf.EncodeBytes(a.DstAddr.Addr[:], 4)
	buf.EncodeBytes(a.DstAddr.Mask[:], 4)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
}
func (u *FlowUnion) GetIP4() (a FlowIP4) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(4))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	return
}

func FlowUnionIP6(a FlowIP6) (u FlowUnion) {
	u.SetIP6(a)
	return
}
func (u *FlowUnion) SetIP6(a FlowIP6) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 16)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 16)
	buf.EncodeBytes(a.DstAddr.Addr[:], 16)
	buf.EncodeBytes(a.DstAddr.Mask[:], 16)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
}
func (u *FlowUnion) GetIP6() (a FlowIP6) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(16))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(16))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(16))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(16))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	return
}

func FlowUnionIP4L2tpv3oip(a FlowIP4L2tpv3oip) (u FlowUnion) {
	u.SetIP4L2tpv3oip(a)
	return
}
func (u *FlowUnion) SetIP4L2tpv3oip(a FlowIP4L2tpv3oip) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 4)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 4)
	buf.EncodeBytes(a.DstAddr.Addr[:], 4)
	buf.EncodeBytes(a.DstAddr.Mask[:], 4)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
	buf.EncodeUint32(a.SessionID)
}
func (u *FlowUnion) GetIP4L2tpv3oip() (a FlowIP4L2tpv3oip) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(4))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	a.SessionID = buf.DecodeUint32()
	return
}

func FlowUnionIP4IpsecEsp(a FlowIP4IpsecEsp) (u FlowUnion) {
	u.SetIP4IpsecEsp(a)
	return
}
func (u *FlowUnion) SetIP4IpsecEsp(a FlowIP4IpsecEsp) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 4)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 4)
	buf.EncodeBytes(a.DstAddr.Addr[:], 4)
	buf.EncodeBytes(a.DstAddr.Mask[:], 4)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
	buf.EncodeUint32(a.Spi)
}
func (u *FlowUnion) GetIP4IpsecEsp() (a FlowIP4IpsecEsp) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(4))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	a.Spi = buf.DecodeUint32()
	return
}

func FlowUnionIP4IpsecAh(a FlowIP4IpsecAh) (u FlowUnion) {
	u.SetIP4IpsecAh(a)
	return
}
func (u *FlowUnion) SetIP4IpsecAh(a FlowIP4IpsecAh) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 4)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 4)
	buf.EncodeBytes(a.DstAddr.Addr[:], 4)
	buf.EncodeBytes(a.DstAddr.Mask[:], 4)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
	buf.EncodeUint32(a.Spi)
}
func (u *FlowUnion) GetIP4IpsecAh() (a FlowIP4IpsecAh) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(4))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	a.Spi = buf.DecodeUint32()
	return
}

func FlowUnionIP4NTuple(a FlowIP4NTuple) (u FlowUnion) {
	u.SetIP4NTuple(a)
	return
}
func (u *FlowUnion) SetIP4NTuple(a FlowIP4NTuple) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 4)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 4)
	buf.EncodeBytes(a.DstAddr.Addr[:], 4)
	buf.EncodeBytes(a.DstAddr.Mask[:], 4)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
	buf.EncodeUint16(a.SrcPort.Port)
	buf.EncodeUint16(a.SrcPort.Mask)
	buf.EncodeUint16(a.DstPort.Port)
	buf.EncodeUint16(a.DstPort.Mask)
}
func (u *FlowUnion) GetIP4NTuple() (a FlowIP4NTuple) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(4))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	a.SrcPort.Port = buf.DecodeUint16()
	a.SrcPort.Mask = buf.DecodeUint16()
	a.DstPort.Port = buf.DecodeUint16()
	a.DstPort.Mask = buf.DecodeUint16()
	return
}

func FlowUnionIP6NTuple(a FlowIP6NTuple) (u FlowUnion) {
	u.SetIP6NTuple(a)
	return
}
func (u *FlowUnion) SetIP6NTuple(a FlowIP6NTuple) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 16)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 16)
	buf.EncodeBytes(a.DstAddr.Addr[:], 16)
	buf.EncodeBytes(a.DstAddr.Mask[:], 16)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
	buf.EncodeUint16(a.SrcPort.Port)
	buf.EncodeUint16(a.SrcPort.Mask)
	buf.EncodeUint16(a.DstPort.Port)
	buf.EncodeUint16(a.DstPort.Mask)
}
func (u *FlowUnion) GetIP6NTuple() (a FlowIP6NTuple) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(16))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(16))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(16))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(16))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	a.SrcPort.Port = buf.DecodeUint16()
	a.SrcPort.Mask = buf.DecodeUint16()
	a.DstPort.Port = buf.DecodeUint16()
	a.DstPort.Mask = buf.DecodeUint16()
	return
}

func FlowUnionIP4NTupleTagged(a FlowIP4NTupleTagged) (u FlowUnion) {
	u.SetIP4NTupleTagged(a)
	return
}
func (u *FlowUnion) SetIP4NTupleTagged(a FlowIP4NTupleTagged) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 4)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 4)
	buf.EncodeBytes(a.DstAddr.Addr[:], 4)
	buf.EncodeBytes(a.DstAddr.Mask[:], 4)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
	buf.EncodeUint16(a.SrcPort.Port)
	buf.EncodeUint16(a.SrcPort.Mask)
	buf.EncodeUint16(a.DstPort.Port)
	buf.EncodeUint16(a.DstPort.Mask)
}
func (u *FlowUnion) GetIP4NTupleTagged() (a FlowIP4NTupleTagged) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(4))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	a.SrcPort.Port = buf.DecodeUint16()
	a.SrcPort.Mask = buf.DecodeUint16()
	a.DstPort.Port = buf.DecodeUint16()
	a.DstPort.Mask = buf.DecodeUint16()
	return
}

func FlowUnionIP6NTupleTagged(a FlowIP6NTupleTagged) (u FlowUnion) {
	u.SetIP6NTupleTagged(a)
	return
}
func (u *FlowUnion) SetIP6NTupleTagged(a FlowIP6NTupleTagged) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 16)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 16)
	buf.EncodeBytes(a.DstAddr.Addr[:], 16)
	buf.EncodeBytes(a.DstAddr.Mask[:], 16)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
	buf.EncodeUint16(a.SrcPort.Port)
	buf.EncodeUint16(a.SrcPort.Mask)
	buf.EncodeUint16(a.DstPort.Port)
	buf.EncodeUint16(a.DstPort.Mask)
}
func (u *FlowUnion) GetIP6NTupleTagged() (a FlowIP6NTupleTagged) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(16))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(16))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(16))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(16))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	a.SrcPort.Port = buf.DecodeUint16()
	a.SrcPort.Mask = buf.DecodeUint16()
	a.DstPort.Port = buf.DecodeUint16()
	a.DstPort.Mask = buf.DecodeUint16()
	return
}

func FlowUnionIP4Vxlan(a FlowIP4Vxlan) (u FlowUnion) {
	u.SetIP4Vxlan(a)
	return
}
func (u *FlowUnion) SetIP4Vxlan(a FlowIP4Vxlan) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 4)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 4)
	buf.EncodeBytes(a.DstAddr.Addr[:], 4)
	buf.EncodeBytes(a.DstAddr.Mask[:], 4)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
	buf.EncodeUint16(a.SrcPort.Port)
	buf.EncodeUint16(a.SrcPort.Mask)
	buf.EncodeUint16(a.DstPort.Port)
	buf.EncodeUint16(a.DstPort.Mask)
	buf.EncodeUint16(a.Vni)
}
func (u *FlowUnion) GetIP4Vxlan() (a FlowIP4Vxlan) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(4))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	a.SrcPort.Port = buf.DecodeUint16()
	a.SrcPort.Mask = buf.DecodeUint16()
	a.DstPort.Port = buf.DecodeUint16()
	a.DstPort.Mask = buf.DecodeUint16()
	a.Vni = buf.DecodeUint16()
	return
}

func FlowUnionIP6Vxlan(a FlowIP6Vxlan) (u FlowUnion) {
	u.SetIP6Vxlan(a)
	return
}
func (u *FlowUnion) SetIP6Vxlan(a FlowIP6Vxlan) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 16)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 16)
	buf.EncodeBytes(a.DstAddr.Addr[:], 16)
	buf.EncodeBytes(a.DstAddr.Mask[:], 16)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
	buf.EncodeUint16(a.SrcPort.Port)
	buf.EncodeUint16(a.SrcPort.Mask)
	buf.EncodeUint16(a.DstPort.Port)
	buf.EncodeUint16(a.DstPort.Mask)
	buf.EncodeUint16(a.Vni)
}
func (u *FlowUnion) GetIP6Vxlan() (a FlowIP6Vxlan) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(16))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(16))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(16))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(16))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	a.SrcPort.Port = buf.DecodeUint16()
	a.SrcPort.Mask = buf.DecodeUint16()
	a.DstPort.Port = buf.DecodeUint16()
	a.DstPort.Mask = buf.DecodeUint16()
	a.Vni = buf.DecodeUint16()
	return
}

func FlowUnionIP4Gtpc(a FlowIP4Gtpc) (u FlowUnion) {
	u.SetIP4Gtpc(a)
	return
}
func (u *FlowUnion) SetIP4Gtpc(a FlowIP4Gtpc) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 4)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 4)
	buf.EncodeBytes(a.DstAddr.Addr[:], 4)
	buf.EncodeBytes(a.DstAddr.Mask[:], 4)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
	buf.EncodeUint16(a.SrcPort.Port)
	buf.EncodeUint16(a.SrcPort.Mask)
	buf.EncodeUint16(a.DstPort.Port)
	buf.EncodeUint16(a.DstPort.Mask)
	buf.EncodeUint32(a.Teid)
}
func (u *FlowUnion) GetIP4Gtpc() (a FlowIP4Gtpc) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(4))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	a.SrcPort.Port = buf.DecodeUint16()
	a.SrcPort.Mask = buf.DecodeUint16()
	a.DstPort.Port = buf.DecodeUint16()
	a.DstPort.Mask = buf.DecodeUint16()
	a.Teid = buf.DecodeUint32()
	return
}

func FlowUnionIP4Gtpu(a FlowIP4Gtpu) (u FlowUnion) {
	u.SetIP4Gtpu(a)
	return
}
func (u *FlowUnion) SetIP4Gtpu(a FlowIP4Gtpu) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeInt32(a.Foo)
	buf.EncodeBytes(a.SrcAddr.Addr[:], 4)
	buf.EncodeBytes(a.SrcAddr.Mask[:], 4)
	buf.EncodeBytes(a.DstAddr.Addr[:], 4)
	buf.EncodeBytes(a.DstAddr.Mask[:], 4)
	buf.EncodeUint8(uint8(a.Protocol.Prot))
	buf.EncodeUint8(a.Protocol.Mask)
	buf.EncodeUint16(a.SrcPort.Port)
	buf.EncodeUint16(a.SrcPort.Mask)
	buf.EncodeUint16(a.DstPort.Port)
	buf.EncodeUint16(a.DstPort.Mask)
	buf.EncodeUint32(a.Teid)
}
func (u *FlowUnion) GetIP4Gtpu() (a FlowIP4Gtpu) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Foo = buf.DecodeInt32()
	copy(a.SrcAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.SrcAddr.Mask[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Addr[:], buf.DecodeBytes(4))
	copy(a.DstAddr.Mask[:], buf.DecodeBytes(4))
	a.Protocol.Prot = ip_types.IPProto(buf.DecodeUint8())
	a.Protocol.Mask = buf.DecodeUint8()
	a.SrcPort.Port = buf.DecodeUint16()
	a.SrcPort.Mask = buf.DecodeUint16()
	a.DstPort.Port = buf.DecodeUint16()
	a.DstPort.Mask = buf.DecodeUint16()
	a.Teid = buf.DecodeUint32()
	return
}
