package internal

import (
	"fmt"
	"server/base"

	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	fmt.Print("module init..............")
}

func (m *Module) OnDestroy() {
	fmt.Print("module destory..............")
}
