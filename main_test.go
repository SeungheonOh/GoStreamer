package gostreamer

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println(GstElementStateGetName(GST_STATE_NULL))
	fmt.Println(GST_STATE_READY.GetName())

	test := CapsNewAny()
	fmt.Println(test.IsFixed())
	fmt.Println(test.GetSize())
}
