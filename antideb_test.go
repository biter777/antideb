package antideb

import (
	"fmt"
	"testing"
)

// TestDetect - detect a debugger / anti-debugger
func TestDetect(t *testing.T) {
	fmt.Printf("DetectVDSO: %v\n", DetectVDSO())
	fmt.Printf("DetectNoASLR: %v\n", DetectNoASLR())
	fmt.Printf("DetectDebugEnv: %v\n", DetectDebugEnv())
	fmt.Printf("DetectLdHook: %v\n", DetectLdHook())
	fmt.Printf("DetectNearHeap: %v\n", DetectNearHeap())
	fmt.Printf("DetectParent: %v\n", DetectParent())
	fmt.Printf("DetectParent2: %v\n", DetectParent2())
	fmt.Printf("DetectPtrace: %v\n", DetectPtrace())
	fmt.Printf("DetectBreakpoints: %v\n", DetectBreakpoints())
	fmt.Printf("Detect: %v\n", Detect(false))
}
