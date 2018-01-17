package MessageTreatChain

import (
	"testing"
	"GoTcpServer/Info"
)

func TestAbsChain(t *testing.T) {
	abs := &AbsChain{}
	abs.Treat(&Info.MessageInfo{})
}

