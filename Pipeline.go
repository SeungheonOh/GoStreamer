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

type GstPipeline struct {
	GstBin
}

func (p *C.GstPipeline) impl() *GstPipeline {
	return (*GstPipeline)(unsafe.Pointer(p))
}

func (p *GstPipeline) native() *C.GstPipeline {
	return (*C.GstPipeline)(unsafe.Pointer(p))
}

func (e *GstElement) ToPipeline() *GstPipeline {
	return (*GstPipeline)(unsafe.Pointer(e))
}

func PipelineNew(name string) *GstPipeline {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	p := C.gst_pipeline_new((*C.gchar)(n))
	return (*C.GstPipeline)(unsafe.Pointer(p)).impl()
}

func (p *GstPipeline) AutoClock() {
	C.gst_pipeline_auto_clock(p.native())
}

func (p *GstPipeline) GetBus() *GstBus {
	return C.gst_pipeline_get_bus(p.native()).impl()
}

func (p *GstPipeline) GetClock() *GstClock {
	return C.gst_pipeline_get_clock(p.native()).impl()
}

func (p *GstPipeline) GetDelay() GstClockTime {
	return GstClockTime(C.gst_pipeline_get_delay(p.native()))
}

func (p *GstPipeline) GetLatency() GstClockTime {
	return GstClockTime(C.gst_pipeline_get_latency(p.native()))
}

func (p *GstPipeline) SetAutoFlushBus(auto_flush bool) {
	a := 0
	if auto_flush {
		a = 1
	}
	C.gst_pipeline_set_auto_flush_bus(p.native(), C.gboolean(a))
}

func (p *GstPipeline) SetClock(clock *GstClock) bool {
	return C.gst_pipeline_set_clock(p.native(), clock.native()) != 0
}

func (p *GstPipeline) SetDelay(delay GstClockTime) {
	C.gst_pipeline_set_delay(p.native(), delay.native())
}

func (p *GstPipeline) SetLatency(latency GstClockTime) {
	C.gst_pipeline_set_latency(p.native(), latency.native())
}

func (p *GstPipeline) UseClock(clock *GstClock) {
	C.gst_pipeline_use_clock(p.native(), clock.native())
}
