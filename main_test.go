package gostreamer

import (
	"fmt"
)

func TestA() {
	fmt.Println(GstElementStateGetName(GST_STATE_NULL))
	fmt.Println(GST_STATE_READY.GetName())
}
