package main

import "fmt"

// Register hooks
var Hooks = map[string]func(){
	"OnStart":     HelloWorld,
	"PageInstall": Message,
}

func HelloWorld() {
	fmt.Println("Hello World")
}

func Message() {
	fmt.Println("We're now in install page!")
}
