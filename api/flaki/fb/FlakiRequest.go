// automatically generated by the FlatBuffers compiler, do not modify

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FlakiRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsFlakiRequest(buf []byte, offset flatbuffers.UOffsetT) *FlakiRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FlakiRequest{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FlakiRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FlakiRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func FlakiRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func FlakiRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
