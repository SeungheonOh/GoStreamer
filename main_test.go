package gostreamer

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	argv := os.Args[2:]
	GstInit(&argv)
}

func TestVersion(t *testing.T) {
	fmt.Println(GstGetVersion())
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
