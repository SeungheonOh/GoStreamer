package gostreamer

import (
	"fmt"
	"os"
	"testing"
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
	return
	fmt.Println("=====Simple=====")

	src := ElementFactoryMake("videotestsrc", "src")
	sink := ElementFactoryMake("xvimagesink", "sink")

	pipe := PipelineNew("pipeline")
	pipe.Add(src, sink)
	src.Link(sink)

	pipe.SetState(GST_STATE_PLAYING)

	bus := pipe.GetBus()
	msg := bus.TimedPopFiltered(GST_CLOCK_TIME_NONE, GST_MESSAGE_ERROR|GST_MESSAGE_EOS)
	fmt.Println("Bus message", msg.Type().GetName())
	msg.Unref()
}

func TestNotSimple(t *testing.T) {
	fmt.Println("=====NotSimple=====")
	source := ElementFactoryMake("uridecodebin", "source")
	convert := ElementFactoryMake("audioconvert", "convert")
	resample := ElementFactoryMake("audioresample", "resmaple")
	sink := ElementFactoryMake("autoaudiosink", "sink")

	pipeline := PipelineNew("test-pipe")

	if source == nil || convert == nil || resample == nil || sink == nil || pipeline == nil {
		panic("Failed to initialize elements")
	}

	if !pipeline.Add(source, convert, resample, sink) {
		panic("Failed to add")
	}
	if !source.Link(convert, resample, sink) {
		panic("Failed to link")
	}

	change_state := pipeline.SetState(GST_STATE_PLAYING)
	if change_state == GST_STATE_CHANGE_FAILURE {
		panic("failed to play")
	}

	bus := pipeline.GetBus()
	msg := bus.TimedPopFiltered(GST_CLOCK_TIME_NONE, GST_MESSAGE_ERROR|GST_MESSAGE_EOS)
	fmt.Println("Bus message", msg.Type().GetName())
	msg.Unref()
}
