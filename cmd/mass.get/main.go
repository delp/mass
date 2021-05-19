package main

import (
	"fmt"
	"os"
	"strconv"

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
	port, err := strconv.Atoi(address)
	check(err)
	port16 := uint16(port)

	listener, err := noise.NewNode(noise.WithNodeBindPort(port16))
	check(err)

	defer listener.Close()

	listener.Handle(func(ctx noise.HandlerContext) error {
		if !ctx.IsRequest() {
			return nil
		}

		fmt.Printf("Got a message: '%s'\n", string(ctx.Data()))

		return ctx.Send([]byte("Hiiiii"))
	})

	check(listener.Listen())

	fmt.Println("listening on", listener.Addr())

}
