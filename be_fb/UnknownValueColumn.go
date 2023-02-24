// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UnknownValueColumn struct {
	_tab flatbuffers.Table
}

func GetRootAsUnknownValueColumn(buf []byte, offset flatbuffers.UOffsetT) *UnknownValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UnknownValueColumn{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUnknownValueColumn(buf []byte, offset flatbuffers.UOffsetT) *UnknownValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UnknownValueColumn{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *UnknownValueColumn) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UnknownValueColumn) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *UnknownValueColumn) Value() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *UnknownValueColumn) MutateValue(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func UnknownValueColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func UnknownValueColumnAddValue(builder *flatbuffers.Builder, value int8) {
	builder.PrependInt8Slot(0, value, 0)
}
func UnknownValueColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
