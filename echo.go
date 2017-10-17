package main

import (
  "github.com/senior-buddy/buddy"
  "fmt"
)

/*
  Requests are in the form of

  {
      "endpoint": "test",
      "params": {
          "test1": "result1",
          "test2": "result2"
      }
  }

 */


// Controller implementation
type EchoController struct{}

func (c *EchoController) Echo(req *buddy.Request) {
  fmt.Println("I'm in the Echo method!")

  // respond here once the responder is implemented!
}

func (c *EchoController) PrintParams(req *buddy.Request) {
  fmt.Println(req.Params)
}

func main() {

  // Register endpoints here in the form of endpoint, controller, method
  buddy.Add("echo", EchoController{}, "Echo")
  buddy.Add("print_params", EchoController{}, "PrintParams")

  // Run the server and serve traffic
  buddy.Run()
}
