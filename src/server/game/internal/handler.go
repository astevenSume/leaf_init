package internal

import (
	"fmt"
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

func init() {
	// handler(&msg.Work{}, handleWork)
	// handler(&msg.HelloNo{}, handleHello)

	handler(&msg.Hello{}, handleProtoHello)
	handler(&msg.Test{}, HandleTest)

}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
func HandleTest(args []interface{}) {
	m := args[0].(*msg.Test)
	a := args[1].(gate.Agent)

	fmt.Println("server  protobuf test ============= > ", *m.Id)

	a.WriteMsg(&msg.Hello{
		Name: "received hello ok ..... success ",
	})
}

func handleProtoHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	fmt.Println("server  protobuf hello ============= > ", m.Name)

	a.WriteMsg(&msg.Hello{
		Name: "received hello ok ..... success ",
	})
}

// func handleWork(args []interface{}) {
// 	m := args[0].(*msg.Work)
// 	a := args[1].(gate.Agent)

// 	fmt.Println("handle work....", m.Number)

// 	a.WriteMsg(&msg.Work{
// 		Number: "received work ok ..... success ",
// 	})
// }

// func handleHello(args []interface{}) {
// 	m := args[0].(*msg.HelloNo)
// 	a := args[1].(gate.Agent)

// 	fmt.Println("handle hello....", m.Name)

// 	a.WriteMsg(&msg.HelloNo{
// 		Name: "received hello ok ..... success ",
// 	})
// }

// func handleHello(args []interface{}) {
// 	m := args[0].(*msg.Hello)
// 	a := args[1].(gate.Agent)

// 	fmt.Println("handle Hello....")

// 	llog.Debug(" -- : %v", m.Name)

// 	a.WriteMsg(&msg.Hello{
// 		Name: proto.String("received ok ..... success "),
// 	})
// }
