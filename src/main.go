package main

import (
	"fmt"

	// hook "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("--- Initiating ---")

	registerHandles()
}

func registerHandles() {
	fmt.Println("--- Registering Handles ---")

	hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
		fmt.Println("[KeyDown] Keychar: ", e.Keychar)
	})

	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		fmt.Println("[MouseDown] Keychar: ", e.Keychar)
	})

	hook.Register(hook.MouseMove, []string{}, func(e hook.Event) {
		fmt.Println("[MouseMove] x, y: ", e.X, e.Y)
	})

	hook.Register(hook.MouseDrag, []string{}, func(e hook.Event) {
		fmt.Println("[MouseDrag] x, y: ", e.X, e.Y)
	})

	hook.Register(hook.MouseWheel, []string{}, func(e hook.Event) {
		fmt.Println("[MouseWheel] Amount: ", e.Amount)
		fmt.Println("[MouseWheel] Clicks: ", e.Clicks)
		fmt.Println("[MouseWheel] Direction: ", e.Direction)
		fmt.Println("[MouseWheel] Rotation: ", e.Rotation)
	})

	s := hook.Start()
	<-hook.Process(s)
}
