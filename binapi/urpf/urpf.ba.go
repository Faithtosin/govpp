// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              22.02-rc0~165-g4b8b64f97
// source: /usr/share/vpp/api/plugins/urpf.api.json

// Package urpf contains generated bindings for API file urpf.api.
//
// Contents:
//   1 enum
//   2 messages
//
package urpf

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	_ "github.com/edwarnicke/govpp/binapi/fib_types"
	interface_types "github.com/edwarnicke/govpp/binapi/interface_types"
	ip_types "github.com/edwarnicke/govpp/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "urpf"
	APIVersion = "1.0.0"
	VersionCrc = 0xb2bfd2c4
)

// UrpfMode defines enum 'urpf_mode'.
type UrpfMode uint8

const (
	URPF_API_MODE_OFF    UrpfMode = 1
	URPF_API_MODE_LOOSE  UrpfMode = 2
	URPF_API_MODE_STRICT UrpfMode = 3
)

var (
	UrpfMode_name = map[uint8]string{
		1: "URPF_API_MODE_OFF",
		2: "URPF_API_MODE_LOOSE",
		3: "URPF_API_MODE_STRICT",
	}
	UrpfMode_value = map[string]uint8{
		"URPF_API_MODE_OFF":    1,
		"URPF_API_MODE_LOOSE":  2,
		"URPF_API_MODE_STRICT": 3,
	}
)

func (x UrpfMode) String() string {
	s, ok := UrpfMode_name[uint8(x)]
	if ok {
		return s
	}
	return "UrpfMode(" + strconv.Itoa(int(x)) + ")"
}

// UrpfUpdate defines message 'urpf_update'.
type UrpfUpdate struct {
	IsInput   bool                           `binapi:"bool,name=is_input,default=true" json:"is_input,omitempty"`
	Mode      UrpfMode                       `binapi:"urpf_mode,name=mode" json:"mode,omitempty"`
	Af        ip_types.AddressFamily         `binapi:"address_family,name=af" json:"af,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *UrpfUpdate) Reset()               { *m = UrpfUpdate{} }
func (*UrpfUpdate) GetMessageName() string { return "urpf_update" }
func (*UrpfUpdate) GetCrcString() string   { return "2bf8a77c" }
func (*UrpfUpdate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *UrpfUpdate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsInput
	size += 1 // m.Mode
	size += 1 // m.Af
	size += 4 // m.SwIfIndex
	return size
}
func (m *UrpfUpdate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsInput)
	buf.EncodeUint8(uint8(m.Mode))
	buf.EncodeUint8(uint8(m.Af))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *UrpfUpdate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsInput = buf.DecodeBool()
	m.Mode = UrpfMode(buf.DecodeUint8())
	m.Af = ip_types.AddressFamily(buf.DecodeUint8())
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// UrpfUpdateReply defines message 'urpf_update_reply'.
type UrpfUpdateReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *UrpfUpdateReply) Reset()               { *m = UrpfUpdateReply{} }
func (*UrpfUpdateReply) GetMessageName() string { return "urpf_update_reply" }
func (*UrpfUpdateReply) GetCrcString() string   { return "e8d4e804" }
func (*UrpfUpdateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *UrpfUpdateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *UrpfUpdateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *UrpfUpdateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_urpf_binapi_init() }
func file_urpf_binapi_init() {
	api.RegisterMessage((*UrpfUpdate)(nil), "urpf_update_2bf8a77c")
	api.RegisterMessage((*UrpfUpdateReply)(nil), "urpf_update_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*UrpfUpdate)(nil),
		(*UrpfUpdateReply)(nil),
	}
}
