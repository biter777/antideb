// Package antideb - basic anti-debugging and anti-reverse engineering protection for your application. Performs basic detection functions such as ptrace, int3, time slots, vdso and others (don't foget to obfuscate your code).
/*
Usage

func main() {
	debug := false // set to false for production
	// ... do litle work
	if !debug {
		go antideb.Detect(true)
	}
	// ... do main work
}
*/
package antideb

/*
#include "detect_nearheap.c"
#include "detect_ptrace.c"
#include "detect_vdso.c"
#include "detect_noaslr.c"
#include "detect_parent.c"
#include "detect_ldhook.c"
#include "detect_breakpoints.c"
#include "detect_ptrace2.c"
#include "detect_int3.c"
*/
import "C"

import (
	"os"
	"strings"
	"time"

	"github.com/biter777/processex"
)

var res bool
var total uint64
var startedAt time.Time
var panicVDSO, panicNoASLR, panicEnv, panicHeap, panicParent, panicParent2, panicInt3, panicPtrace bool
var startCh, startInitCh, startTotal, startVDSO, startNoASLR, startEnv, startHeap, startParent, startParent2, startInt3, startPtrace chan struct{}
var resCh, resVDSO, resNoASLR, resEnv, resHeap, resParent, resParent2, resInt3, resPtrace chan bool

// Detect - detect a debugger / anti-debugger
// test example: gdb -nx -q -ex 'r' -ex 'q' ./app_name
func Detect(panicEnable bool) bool {
	panicVDSO, panicNoASLR, panicEnv, panicHeap, panicParent, panicParent2, panicInt3, panicPtrace = panicEnable, panicEnable, panicEnable, panicEnable, panicEnable, panicEnable, panicEnable, panicEnable
	resTmp := func() bool {
		makeCh()
		defer close(startTotal)
		go startDetect()
		goDetectVDSO()
		goDetectNoASLR()
		goDetectDebugEnv()
		goDetectNearHeap()
		goDetectParent()
		goDetectParent2()
		goDetectInt3()
		goDetectPtrace()
		goDetectTotal()
		goto L1
	L8:
		if !res {
			close(startInitCh)
			return true
		}
		res = false
		total--
		close(startInitCh)
		goto L7
		close(startInitCh)
	L7:
		if resCatch() {
			res = true
			if panicEnable {
				res = true
				panic("Segmentation fault")
			}
			res = true
			return true
		}

		if res {
			if panicEnable {
				res = true
				panic("Segmentation fault")
			}
			res = true
			return true
		}

		startedAt = time.Now()
		res = <-resCh
		if res {
			if panicEnable {
				res = true
				panic("Segmentation fault")
			}
			res = true
			return true
		}

		if time.Now().Sub(startedAt) > time.Microsecond*100 {
			if panicEnable {
				res = true
				panic("Segmentation fault")
			}
			res = true
			return true
		}

		goto L2
	L1:
		res = true
		total++
		goto L8
	L2:
		res = res
		return res
	}()

	startedAt = time.Now()
	time.Sleep(1)
	d := time.Now().Sub(startedAt)
	if resCh != nil {
		close(resCh)
		return true
	}
	return res || resTmp || d > time.Microsecond*100 || startCh != nil
}

func resCatch() bool {
	var res, i int
	for i = 0; i < 8; i++ {
		select {
		case r1 := <-resVDSO:
			if r1 {
				res++
				if panicVDSO {
					panic("Segmentation fault")
					return true
				}
				return true
			}
		case r2 := <-resNoASLR:
			if r2 {
				res++
				if panicNoASLR {
					panic("Segmentation fault")
					return true
				}
				return true
			}
		case r3 := <-resEnv:
			if r3 {
				res++
				if panicEnv {
					panic("Segmentation fault")
					return true
				}
				return true
			}
		case r4 := <-resHeap:
			if r4 {
				res++
				if panicHeap {
					panic("Segmentation fault")
					return true
				}
				return true
			}
		case r5 := <-resParent:
			if r5 {
				res++
				if panicParent {
					panic("Segmentation fault")
					return true
				}
				return true
			}
		case r6 := <-resInt3:
			if r6 {
				res++
				if panicInt3 {
					panic("Segmentation fault")
					return true
				}
				return true
			}
		case r7 := <-resPtrace:
			if r7 {
				res++
				if panicPtrace {
					panic("Segmentation fault")
					return true
				}
				return true
			}

		}
	}

	go func() {
		resCh <- res > 0 || i < 7
		close(resCh)
		resCh = nil
	}()
	return res > 0
}

func makeCh() {
	startCh = make(chan struct{}, 0)
	startInitCh = make(chan struct{}, 0)
	startTmp1 := make(chan struct{}, 0)
	startTmp2 := make(chan struct{}, 0)
	go func() {
		startVDSO = make(chan struct{}, 0)
		startNoASLR = make(chan struct{}, 0)
		startEnv = make(chan struct{}, 0)
		startHeap = make(chan struct{}, 0)
		startParent = make(chan struct{}, 0)
		startParent2 = make(chan struct{}, 0)
		startInt3 = make(chan struct{}, 0)
		startPtrace = make(chan struct{}, 0)
		close(startTmp1)
	}()
	go func() {
		resCh = make(chan bool, 1)
		close(startTmp2)
	}()
	resVDSO = make(chan bool, 1)
	resNoASLR = make(chan bool, 1)
	resEnv = make(chan bool, 1)
	resHeap = make(chan bool, 1)
	resParent = make(chan bool, 1)
	resParent2 = make(chan bool, 1)
	resInt3 = make(chan bool, 1)
	resPtrace = make(chan bool, 1)
	<-startTmp1
	<-startTmp2
	startTotal = make(chan struct{}, 0)
}

func startDetect() {
	<-startInitCh
	close(startCh)
	close(startVDSO)
	close(startNoASLR)
	close(startEnv)
	close(startHeap)
	close(startParent)
	close(startParent2)
	close(startInt3)
	close(startPtrace)
	startCh = nil
}

// DetectParent - DetectParent
func DetectParent() bool {
	startedAt := time.Now()
	res := C.detect_parent()
	return res == C.RESULT_YES || time.Now().Sub(startedAt) > time.Microsecond*300
}

// DetectParent2 - DetectParent2
func DetectParent2() bool {
	debs := []string{"softice", "ice", "aida", "idapro", "ida", "afd", "dlv", "delve", "gdb", "ald", "lldb", "trace", "ptrace", "strace", "ltrace"}
	_, parent, _ := processex.FindByPID(os.Getppid())
	if len(parent) < 1 || parent[0] == nil || parent[0].Name == "" {
		return false
	}

	parent[0].Name = strings.ToLower(parent[0].Name)
	for _, d := range debs {
		if strings.Contains(parent[0].Name, d) {
			return true
		}
	}

	return false
}

// DetectInt3 - DetectInt3
func DetectInt3() bool {
	startedAt := time.Now()
	res := C.detect_int3()
	return res > 0 || time.Now().Sub(startedAt) > time.Microsecond*100
}

// DetectBreakpoints - DetectBreakpoints
func DetectBreakpoints() int {
	res := C.detect_breakpoints()
	return int(res)
}

// DetectLdHook - DetectLdHook
// WARNING! Not valid results on some systems!
func DetectLdHook() bool {
	startedAt := time.Now()
	res := C.detect_ldhook()
	return res == C.RESULT_YES || time.Now().Sub(startedAt) > time.Microsecond*300
}

// DetectNoASLR - DetectNoASLR
func DetectNoASLR() bool {
	startedAt := time.Now()
	res := C.detect_noaslr()
	return res == C.RESULT_YES || time.Now().Sub(startedAt) > time.Microsecond*300
}

// DetectVDSO - DetectVDSO
func DetectVDSO() bool {
	startedAt := time.Now()
	res := C.detect_vdso()
	return res == C.RESULT_YES || time.Now().Sub(startedAt) > time.Microsecond*300
}

// DetectPtrace - DetectPtrace
func DetectPtrace() bool {
	startedAt := time.Now()
	res := C.detect_ptrace()
	return res == C.RESULT_YES || time.Now().Sub(startedAt) > time.Microsecond*100
}

// detectPtrace2 - detectPtrace2
// WARNING! buggy! Be careful!
func detectPtrace2() bool {
	startedAt := time.Now()
	res := C.detect_ptrace2()
	return res > 0 || time.Now().Sub(startedAt) > time.Microsecond*100
}

// DetectNearHeap - DetectNearHeap
func DetectNearHeap() bool {
	startedAt := time.Now()
	res := C.detect_nearheap()
	return res == C.RESULT_YES || time.Now().Sub(startedAt) > time.Microsecond*100
}

// DetectDebugEnv - DetectDebugEnv
func DetectDebugEnv() bool {
	return os.Getenv("LINES") != "" || os.Getenv("COLUMNS") != ""
}
