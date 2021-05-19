package main

import (
	"context"
	"fmt"
	"os"

	"github.com/perlin-network/noise"
)

type Request struct {
	Name string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	if len(os.Args) != 2 {
		panic("ahhh")
	}

	address := os.Args[1]
	sender, err := noise.NewNode()
	check(err)

	defer sender.Close()

	check(sender.Listen())

	//check(sender.Send(context.TODO(), address, []byte("Hi Bob!")))
	foo, err := sender.Request(context.TODO(), address, []byte("hi"))
	fmt.Println(foo)
	check(err)
}
