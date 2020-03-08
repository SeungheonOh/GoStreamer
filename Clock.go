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

//type GstClock C.GstClock
type GstClock struct {
	GstObject
}
type GstClockID C.GstClockID
type GstClockTime C.GstClockTime
type GstClockTimeDiff C.GstClockTimeDiff
type GstClockEntryType C.GstClockEntryType
type GstClockFlags C.GstClockFlags
type GstClockReturn C.GstClockReturn

const (
	GST_CLOCK_ENTRY_SINGLE               GstClockEntryType = 0
	GST_CLOCK_ENTRY_PERIODIC             GstClockEntryType = 1
	GST_CLOCK_FLAG_CAN_DO_SINGLE_SYNC    GstClockFlags     = 16
	GST_CLOCK_FLAG_CAN_DO_SINGLE_ASYNC   GstClockFlags     = 32
	GST_CLOCK_FLAG_CAN_DO_PERIODIC_SYNC  GstClockFlags     = 64
	GST_CLOCK_FLAG_CAN_DO_PERIODIC_ASYNC GstClockFlags     = 128
	GST_CLOCK_FLAG_CAN_SET_RESOLUTION    GstClockFlags     = 256
	GST_CLOCK_FLAG_CAN_SET_MASTER        GstClockFlags     = 512
	GST_CLOCK_FLAG_NEEDS_STARTUP_SYNC    GstClockFlags     = 1024
	GST_CLOCK_FLAG_LAST                  GstClockFlags     = 4096
	GST_CLOCK_OK                         GstClockReturn    = 0
	GST_CLOCK_EARLY                      GstClockReturn    = 1
	GST_CLOCK_UNSCHEDULED                GstClockReturn    = 2
	GST_CLOCK_BUSY                       GstClockReturn    = 3
	GST_CLOCK_BADTIME                    GstClockReturn    = 4
	GST_CLOCK_ERROR                      GstClockReturn    = 5
	GST_CLOCK_UNSUPPORTED                GstClockReturn    = 6
	GST_CLOCK_DONE                       GstClockReturn    = 7
)

func (c *C.GstClock) impl() *GstClock {
	return (*GstClock)(unsafe.Pointer(c))
}

func (c *GstClock) native() *C.GstClock {
	return (*C.GstClock)(unsafe.Pointer(c))
}

func (c *GstClock) AddObservation(slave, master GstClockTime, r_squared *float64) bool {
	return C.gst_clock_add_observation(c.native(), slave.native(), master.native(), (*C.double)(r_squared)) != 0
}

func (c *GstClock) AdjustUnlocked(internal GstClockTime) GstClockTime {
	return C.gst_clock_adjust_unlocked(c.native(), internal.native()).impl()
}

func (c *GstClock) GetMaster() *GstClock {
	return C.gst_clock_get_master(c.native()).impl()
}

func (c *GstClock) GetResolution() GstClockTime {
	return C.gst_clock_get_resolution(c.native()).impl()
}

func (c *GstClock) GetTime() GstClockTime {
	return C.gst_clock_get_time(c.native()).impl()
}

func (c *GstClock) GetTimeout() GstClockTime {
	return C.gst_clock_get_timeout(c.native()).impl()
}

func (c *GstClock) IsSynced() bool {
	return C.gst_clock_is_synced(c.native()) != 0
}

func (c *GstClock) NewPeriodicId(start_time, interval GstClockTime) GstClockID {
	return GstClockID(C.gst_clock_new_periodic_id(c.native(), start_time.native(), interval.native()))
}

func (c *GstClock) NewSingleShotId(start_time GstClockTime) GstClockID {
	return GstClockID(C.gst_clock_new_single_shot_id(c.native(), start_time.native()))
}

func (c *GstClock) SetCalibration(internal, external, rate_num, rate_denom GstClockTime) {
	C.gst_clock_set_calibration(c.native(), internal.native(), external.native(), rate_num.native(), rate_denom.native())
}

func (c *GstClock) SetMaster(master *GstClock) bool {
	return C.gst_clock_set_master(c.native(), master.native()) != 0
}

func (c *GstClock) SetResolution(resolution GstClockTime) GstClockTime {
	return C.gst_clock_set_resolution(c.native(), resolution.native()).impl()
}

func (c *GstClock) SetSynced(synced bool) {
	b := 0
	if synced {
		b = 1
	}
	C.gst_clock_set_synced(c.native(), C.gboolean(b))
}

func (c *GstClock) SetTimeout(timeout GstClockTime) {
	C.gst_clock_set_timeout(c.native(), timeout.native())
}

func (c *GstClock) SingleShotIdReinit(id GstClockID, time GstClockTime) bool {
	return C.gst_clock_single_shot_id_reinit(c.native(), C.GstClockID(id), time.native()) != 0
}

func (c *GstClock) UnadjustUnlocked(external GstClockTime) GstClockTime {
	return GstClockTime(C.gst_clock_unadjust_unlocked(c.native(), external.native()))
}

func (c *GstClock) UnadjustWithCalibration(external_target, cinternal, cexternal, cnum, cdenom GstClockTime) GstClockTime {
	return C.gst_clock_unadjust_with_calibration(c.native(), external_target.native(), cinternal.native(), cexternal.native(),
		cnum.native(), cdenom.native()).impl()
}

func (c *GstClock) WaitForSync(timeout GstClockTime) bool {
	return C.gst_clock_wait_for_sync(c.native(), timeout.native()) != 0
}

// Functions
func ClockIdCompareFunc(id1, id2 GstClockID) int {
	return int(C.gst_clock_id_compare_func(C.gconstpointer(id1), C.gconstpointer(id2)))
}

func ClockIdGetClock(id GstClockID) *GstClock {
	return C.gst_clock_id_get_clock(C.GstClockID(id)).impl()
}

func ClockIdGetTime(id GstClockID) GstClockTime {
	return C.gst_clock_id_get_time(C.GstClockID(id)).impl()
}

func ClockIdRef(id GstClockID) GstClockID {
	return GstClockID(C.gst_clock_id_ref(C.GstClockID(id)))
}

func ClockIdUnref(id GstClockID) {
	C.gst_clock_id_unref(C.GstClockID(id))
}

func ClockIdUnschedule(id GstClockID) {
	C.gst_clock_id_unschedule(C.GstClockID(id))
}

func ClockIdUsesClock(id GstClockID, clock *GstClock) bool {
	return C.gst_clock_id_uses_clock(C.GstClockID(id), clock.native()) != 0
}

func ClockIdWait(id GstClockID, jitter *GstClockTimeDiff) GstClockReturn {
	return GstClockReturn(C.gst_clock_id_wait(C.GstClockID(id), (*C.GstClockTimeDiff)(jitter)))
}

func (c C.GstClockTime) impl() GstClockTime {
	return (GstClockTime)(c)
}

func (c GstClockTime) native() C.GstClockTime {
	return (C.GstClockTime)(c)
}
