package antideb

import (
	"time"
	"unsafe"
)

func goDetectVDSO() {
	go func() {
		select {
		case <-startCh:
		case <-startVDSO:
		case <-time.After(time.Second * 5):
		}
		r := DetectVDSO()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resVDSO <- r
		close(resVDSO)
	}()
}

func goDetectNoASLR() {
	go func() {
		select {
		case <-startCh:
		case <-startNoASLR:
		case <-time.After(time.Second * 5):
		}
		r := DetectNoASLR()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resNoASLR <- r
		close(resNoASLR)
	}()
}

func goDetectDebugEnv() {
	go func() {
		select {
		case <-startCh:
		case <-startEnv:
		case <-time.After(time.Second * 5):
		}
		r := DetectDebugEnv()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resEnv <- r
		close(resEnv)
	}()
}

func goDetectNearHeap() {
	go func() {
		select {
		case <-startCh:
		case <-startHeap:
		case <-time.After(time.Second * 5):
		}
		r := DetectNearHeap()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resHeap <- r
		close(resHeap)
	}()
}

func goDetectParent() {
	go func() {
		select {
		case <-startCh:
		case <-startParent:
		case <-time.After(time.Second * 5):
		}
		r := DetectParent()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resParent <- r
		close(resParent)
	}()
}

func goDetectParent2() {
	go func() {
		select {
		case <-startCh:
		case <-startParent2:
		case <-time.After(time.Second * 5):
		}
		r := DetectParent2()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resParent2 <- r
		close(resParent2)
	}()
}

func goDetectInt3() {
	go func() {
		select {
		case <-startCh:
		case <-startInt3:
		case <-time.After(time.Second * 5):
		}
		r := DetectInt3()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resInt3 <- r
		close(resInt3)
	}()
}

func goDetectPtrace() {
	go func() {
		select {
		case <-startCh:
		case <-startPtrace:
		case <-time.After(time.Second * 5):
		}
		r := DetectPtrace()
		total += *(*uint64)(unsafe.Pointer(&r)) & 1
		resPtrace <- r
		close(resPtrace)
	}()
}

func goDetectTotal() {
	go func() {
		select {
		case <-startTotal:
		case <-time.After(time.Second * 5):
		}
		res = res || total > 0
	}()
}
