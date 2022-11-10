package main

import (
	"asjade"
	"go-jiajia/server/rest"
)

// Hello .
type Hello struct{}

// Hello .
func (h *Hello) Hello(c *rest.Context) {}

// Routes .
func (h *Hello) Routes() []*rest.Route {
	return []*rest.Route{}
}

func main() {
	asjade.RegisteSchema("rest", &Hello{})
	asjade.Start()
}
