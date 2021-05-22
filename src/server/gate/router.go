package gate

import (
	"server/game"
	"server/msg"
)

func init() {
	// msg.Processor2.SetRouter(&msg.Work{}, game.ChanRPC)
	// msg.Processor2.SetRouter(&msg.HelloNo{}, game.ChanRPC)

	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Test{}, game.ChanRPC)
}
