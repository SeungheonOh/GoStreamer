package gostreamer

import (
	"unsafe"
)

/*
#cgo pkg-config: gstreamer-1.0
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstBus struct {
	GstObject
}

func (b *C.GstBus) impl() *GstBus {
	return (*GstBus)(unsafe.Pointer(b))
}

func (b *GstBus) native() *C.GstBus {
	return (*C.GstBus)(unsafe.Pointer(b))
}

func BusNew() *GstBus {
	return C.gst_bus_new().impl()
}

func (b *GstBus) AddSignalWatch() {
	C.gst_bus_add_signal_watch(b.native())
}

func (b *GstBus) AddSignalWatchFull(priority int) {
	C.gst_bus_add_signal_watch_full(b.native(), C.int(priority))
}

// Few missing methods

func (b *GstBus) Peek() *GstMessage {
	return C.gst_bus_peek(b.native()).impl()
}

func (b *GstBus) Poll(events GstMessageType, timeout GstClockTime) *GstMessage {
	return C.gst_bus_poll(b.native(), C.GstMessageType(events), C.GstClockTime(timeout)).impl()
}

func (b *GstBus) Pop() *GstMessage {
	return C.gst_bus_pop(b.native()).impl()
}

func (b *GstBus) PopFiltered(types GstMessageType) *GstMessage {
	return C.gst_bus_pop_filtered(b.native(), C.GstMessageType(types)).impl()
}

func (b *GstBus) Post(message *GstMessage) bool {
	return C.gst_bus_post(b.native(), message.native()) != 0
}

func (b *GstBus) RemoveSignalWatch() {
	C.gst_bus_remove_signal_watch(b.native())
}

func (b *GstBus) RemoveWatch() {
	C.gst_bus_remove_watch(b.native())
}

func (b *GstBus) SetFlushing(flushing bool) {
	f := 0
	if flushing {
		f = 1
	}
	C.gst_bus_set_flushing(b.native(), C.gboolean(f))
}

// Some more missing

func (b *GstBus) TimedPop(timeout GstClockTime) *GstMessage {
	return C.gst_bus_timed_pop(b.native(), C.GstClockTime(timeout)).impl()
}

func (b *GstBus) TimedPopFilter(timeout GstClockTime, types GstMessageType) *GstMessage {
	return C.gst_bus_timed_pop_filtered(b.native(), C.GstClockTime(timeout), C.GstMessageType(types)).impl()
}
