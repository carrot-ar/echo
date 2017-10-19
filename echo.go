package main

import (
	"fmt"
	"github.com/senior-buddy/buddy"
)

/*
  Requests are in the form of

  {
			"session_token: "de0u21ejdd12d",
			"endpoint": "test",
      "params": {
          "test1": "result1",
          "test2": "result2"
      }
  }

*/

// Controller implementation
type EchoController struct{}

func (c *EchoController) Echo(req *buddy.Request, res *buddy.Responder) {
	fmt.Println("I'm in the Echo method!")

	res.Broadcast <- []byte([]byte(fmt.Sprintf("%v",req.Origin)))
}

func (c *EchoController) PrintParams(req *buddy.Request, res *buddy.Responder) {
	fmt.Println(req.Params)
}

type EchoStreamController struct {
	count int
}

func (c *EchoStreamController) EchoStreamer(req *buddy.Request, res *buddy.Responder) {
	fmt.Println(c.count)
	c.count += 1
}

func main() {

	// Register endpoints here in the form of endpoint, controller, method
	buddy.Add("echo", EchoController{}, "Echo", false)
	buddy.Add("print_params", EchoController{}, "PrintParams", false)
	buddy.Add("echo_streamer", EchoController{}, "EchoStreamer", true)

	// Run the server and serve traffic
	buddy.Run()
}
