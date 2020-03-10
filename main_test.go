package gostreamer

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func init() {
	argv := os.Args
	GstInit(&argv)
}

func TestVersion(t *testing.T) {
	fmt.Println("=====Version=====")
	fmt.Println(GstVersion())
	fmt.Println(GstVersionString())
}

func TestElement(t *testing.T) {
	fmt.Println("=====Element=====")
	fmt.Println(GstElementStateGetName(GST_STATE_NULL))
	fmt.Println(GST_STATE_READY.GetName())
}

func TestCaps(t *testing.T) {
	fmt.Println("=====Caps=====")
	test := CapsFromString("video/x-raw")
	fmt.Println("Caps is fixed: ", test.IsFixed())
	fmt.Println("Caps size: ", test.GetSize())

	capsfeature := CapsFeaturesNewAny()
	fmt.Println("Caps Features Any to string: ", capsfeature.ToString())
}

func TestElementFactory(t *testing.T) {
	fmt.Println("=====ElementFactory=====")
	e := ElementFactoryMake("videotestsrc", "MyVideoSrc")
	state, _, _ := e.GetState(1000)
	fmt.Println(state.GetName())
}

func TestSimple(t *testing.T) {
	fmt.Println("=====Simple=====")

	src := ElementFactoryMake("videotestsrc", "src")
	sink := ElementFactoryMake("xvimagesink", "sink")

	src.Set("uri", "https://www.freedesktop.org/software/gstreamer-sdk/data/media/sintel_trailer-480p.webm")

	pipe := PipelineNew("pipeline")
	pipe.Add(src, sink)
	src.Link(sink)

	bus := pipe.GetBus()
	for {
		pipe.SetState(GST_STATE_PAUSED)
		msg := bus.TimedPopFiltered(GST_CLOCK_TIME_NONE, GST_MESSAGE_ELEMENT)
		fmt.Println("Bus message", msg.Type().GetName())
		msg.Unref()
		pipe.SetState(GST_STATE_PLAYING)
		time.Sleep(time.Millisecond * 100)
	}

}
