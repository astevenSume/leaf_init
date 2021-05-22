package msg

import (
	"fmt"

	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

// var Processor2 = json.NewProcessor()

func init() {
	// idJson := Processor2.Register(&Work{})
	// fmt.Println("res id json : ", idJson)
	// Processor2.Register(&HelloNo{})

	id2 := Processor.Register(&Test{})
	fmt.Println("res id2 : ", id2)

	id := Processor.Register(&Hello{})
	fmt.Println("res id : ", id)
}

// type HelloNo struct {
// 	Name string
// }
