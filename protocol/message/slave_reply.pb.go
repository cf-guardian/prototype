// Code generated by protoc-gen-go.
// source: slave_reply.proto
// DO NOT EDIT!

package message

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SlaveReply struct {
	Reply            *string `protobuf:"bytes,1,req" json:"Reply,omitempty"`
	Error            *string `protobuf:"bytes,2,opt" json:"Error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SlaveReply) Reset()         { *m = SlaveReply{} }
func (m *SlaveReply) String() string { return proto.CompactTextString(m) }
func (*SlaveReply) ProtoMessage()    {}

func (m *SlaveReply) GetReply() string {
	if m != nil && m.Reply != nil {
		return *m.Reply
	}
	return ""
}

func (m *SlaveReply) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func init() {
}
