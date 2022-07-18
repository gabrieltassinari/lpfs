package lpfs

import (
	"fmt"
	"testing"
)

//	TestLoadAverage tests all functions that get data from /proc/loadavg.
func TestLoadAverage(t *testing.T) {
	l1, err := GetLoadAverage1()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("LoadAverage1(): %v, err: %v\n", l1, err)

	l5, err := GetLoadAverage5()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("LoadAverage5(): %v, err: %v\n", l5, err)

	l15, err := GetLoadAverage15()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("LoadAverage15(): %v, err: %v\n", l15, err)

	tskq, err := GetTaskQueueSize()
	if err != nil {
		t.Errorf("%v", err)
	}

	fmt.Printf("GetTaskQueueSize(): %v, err: %v\n", tskq, err)
	rszq, err := GetRunnableQueueSize()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetRunnableQueueSize(): %v, err: %v\n", rszq, err)

	pid, err := GetMostRecentPid()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetMostRecentPid(): %v, err: %v\n", pid, err)
}

//	TestSwaps tests all functions that get data from /proc/swaps.
func TestSwaps(t *testing.T) {
	sf, err := GetSwapFilename()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetSwapFilename(): %v, err: %v\n", sf, err)

	st, err := GetSwapType()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetSwapType(): %v, err: %v\n", st, err)

	ss, err := GetSwapSize()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetSwapSize(): %v, err: %v\n", ss, err)

	su, err := GetSwapUsed()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetSwapUsed(): %v, err: %v\n", su, err)

	sp, err := GetSwapPriority()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("GetSwapPriority(): %v, err: %v\n", sp, err)
}
