package main

import (
	"fmt"
	"os"

	"github.com/Shopify/go-lua"
)

func main() {
	l := lua.NewStateEx()

	lua.OpenLibraries(l)

	libOpen := func(state *lua.State) int {
		lua.NewLibrary(state, lib)
		return 1
	}
	lua.Require(l, "glua", libOpen, false)
	l.Pop(1)

	if err := lua.DoFile(l, "init.lua"); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

var lib = []lua.RegistryFunction{
	{"hello", hello},
}

func hello(l *lua.State) int {
	name := lua.CheckString(l, 1)
	fmt.Println("hello", name)
	return 1
}
