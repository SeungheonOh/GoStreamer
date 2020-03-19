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

type GstTask struct {
	GstObject
}
type GstTaskState C.GstTaskState

const (
	GST_TASK_STARTED GstTaskState = 0
	GST_TASK_STOPPED GstTaskState = 1
	GST_TASK_PAUSED  GstTaskState = 2
)

func (t *C.GstTask) impl() *GstTask {
	return (*GstTask)(unsafe.Pointer(t))
}

func (t *GstTask) native() *C.GstTask {
	return (*C.GstTask)(unsafe.Pointer(t))
}

// Constructors

//TaskNew: Callback

// Methods

func (t *GstTask) GetPool() *GstTaskPool {
	return C.gst_task_get_pool(t.native()).impl()
}

func (t *GstTask) GetState() GstTaskState {
	return GstTaskState(C.gst_task_get_state(t.native()))
}

func (t *GstTask) Join() bool {
	return C.gst_task_join(t.native()) != 0
}

func (t *GstTask) Pause() bool {
	return C.gst_task_pause(t.native()) != 0
}

// SetEnterCallback: callback

// SetLeaveCallback: callback

// SetLock: GRecMutex

func (t *GstTask) SetPool(pool *GstTaskPool) {
	C.gst_task_set_pool(t.native(), pool.native())
}

func (t *GstTask) SetState(state GstTaskState) bool {
	return C.gst_task_set_state(t.native(), C.GstTaskState(state)) != 0
}

func (t *GstTask) Start() bool {
	return C.gst_task_start(t.native()) != 0
}

func (t *GstTask) Stop() bool {
	return C.gst_task_stop(t.native()) != 0
}

// Functions

func TaskCleanupAll() {
	C.gst_task_cleanup_all()
}
