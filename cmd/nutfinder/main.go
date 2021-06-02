package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/delp/mass"
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

func genericTest() {

	alice, err := noise.NewNode()
	check(err)

	bob, err := noise.NewNode()
	check(err)

	// Gracefully release resources for Alice and Bob at the end of the example.

	defer alice.Close()
	defer bob.Close()

	// When Bob gets a message from Alice, print it out and respond to Alice with 'Hi Alice!'

	bob.Handle(func(ctx noise.HandlerContext) error {
		if !ctx.IsRequest() {
			return nil
		}

		fmt.Printf("Got a message from Alice: '%s'\n", string(ctx.Data()))

		return ctx.Send([]byte("Hi Alice!"))
	})

	// Have Alice and Bob start listening for new peers.

	check(alice.Listen())
	check(bob.Listen())

	// Have Alice send Bob a request with the message 'Hi Bob!'

	res, err := alice.Request(context.TODO(), bob.Addr(), []byte("Hi Bob!"))
	check(err)

	// Print out the response Bob got from Alice.

	fmt.Printf("Got a message from Bob: '%s'\n", string(res))

	// Output:
	// Got a message from Alice: 'Hi Bob!'
	// Got a message from Bob: 'Hi Alice!'
}

func main() {
	// Let there be nodes Alice and Bob.

	//genericTest()

	seeker, err := noise.NewNode()
	check(err)

	holder, err := noise.NewNode()
	check(err)

	// Gracefully release resources for Alice and Bob at the end of the example.

	defer seeker.Close()
	defer holder.Close()

	ask := mass.Ask{
		Abstract: "nuts",
	}

	payload, err := json.Marshal(ask)

	holder.Handle(func(ctx noise.HandlerContext) error {
		if !ctx.IsRequest() {
			return nil
		}

		fmt.Printf("Got a message from Alice: '%s'\n", string(ctx.Data()))

		return ctx.Send([]byte("Hi Alice!"))
	})

	check(seeker.Listen())
	check(holder.Listen())

	res, err := seeker.Request(context.TODO(), holder.Addr(), payload)
	check(err)

	// Print out the response Bob got from Alice.

	fmt.Printf("Got a message from Bob: '%s'\n", string(res))

}
