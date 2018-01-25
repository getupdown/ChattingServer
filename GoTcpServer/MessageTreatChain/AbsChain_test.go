package MessageTreatChain

import (
	"testing"
	"unsafe"
	"sync"
	"fmt"
)

func TestSize(t *testing.T) {
	x := unsafe.Sizeof(sync.RWMutex{})
	fmt.Println(x)
	var y int32
	x = unsafe.Sizeof(y)
	fmt.Println(x)
}

