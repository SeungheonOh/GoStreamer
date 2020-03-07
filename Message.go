package main

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstMessageType C.GstMessageType
type GstProgressType C.GstProgressType
type GstStreamStatusType C.GstStreamStatusType
type GstStructureChangeType C.GstStructureChangeType

type GstMessage C.GstMessage

const (
	GST_MESSAGE_UNKNOWN          GstMessageType = 0
	GST_MESSAGE_EOS              GstMessageType = 1
	GST_MESSAGE_ERROR            GstMessageType = 2
	GST_MESSAGE_WARNING          GstMessageType = 4
	GST_MESSAGE_INFO             GstMessageType = 8
	GST_MESSAGE_TAG              GstMessageType = 16
	GST_MESSAGE_BUFFERING        GstMessageType = 32
	GST_MESSAGE_STATE_CHANGED    GstMessageType = 64
	GST_MESSAGE_STATE_DIRTY      GstMessageType = 128
	GST_MESSAGE_STEP_DONE        GstMessageType = 256
	GST_MESSAGE_CLOCK_PROVIDE    GstMessageType = 512
	GST_MESSAGE_CLOCK_LOST       GstMessageType = 1024
	GST_MESSAGE_NEW_CLOCK        GstMessageType = 2048
	GST_MESSAGE_STRUCTURE_CHANGE GstMessageType = 4096
	GST_MESSAGE_STREAM_STATUS    GstMessageType = 8192
	GST_MESSAGE_APPLICATION      GstMessageType = 16384
	GST_MESSAGE_ELEMENT          GstMessageType = 32768
	GST_MESSAGE_SEGMENT_START    GstMessageType = 65536
	GST_MESSAGE_SEGMENT_DONE     GstMessageType = 131072
	GST_MESSAGE_DURATION_CHANGED GstMessageType = 262144
	GST_MESSAGE_LATENCY          GstMessageType = 524288
	GST_MESSAGE_ASYNC_START      GstMessageType = 1048576
	GST_MESSAGE_ASYNC_DONE       GstMessageType = 2097152
	GST_MESSAGE_REQUEST_STATE    GstMessageType = 4194304
	GST_MESSAGE_STEP_START       GstMessageType = 8388608
	GST_MESSAGE_QOS              GstMessageType = 16777216
	GST_MESSAGE_ANY              GstMessageType = 429496729
)

func GstMessageNewApplication(src *GstObject, structure *GstStructure) *GstMessage {
	return (*GstMessage)(C.gst_message_new_application(src.native(), structure.native()))
}

func (m *GstMessage) native() *C.GstMessage {
	return (*C.GstMessage)(m)
}

func (m *GstMessage) GstMessageUnref() {
	C.gst_message_unref(m.native())
}
