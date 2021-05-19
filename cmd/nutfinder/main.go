package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/perlin-network/noise"
)

type Offer struct {
	Desc    string
	Contact string
}

type Request struct {
	Search string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checkArgs() {
	if len(os.Args) != 2 {
		panic("ahhh")
	}
}

func getPortArg() uint16 {
	address := os.Args[1]
	port, err := strconv.Atoi(address)
	check(err)
	return uint16(port)
}

func main() {

	checkArgs()
	port := getPortArg()

	listener, err := noise.NewNode(noise.WithNodeBindPort(port))
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
