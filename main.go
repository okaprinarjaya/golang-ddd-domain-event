package main

import (
	"fmt"

	"github.com/asaskevich/EventBus"
)

func main() {
	bus := EventBus.New()
	bus.Subscribe("main.calculator", calculator)
	bus.Subscribe("main.calculator", calculator1)
	bus.Subscribe("main.calculator", calculator2)

	bus.Publish("main.calculator", 20, 49)

	bus.Unsubscribe("main.calculator", calculator)

	fmt.Println("Hello Oprex Event Bus")
}

func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func calculator1(a int, b int) {
	fmt.Printf("%d\n", a*b)
}

func calculator2(a int, b int) {
	fmt.Printf("%d\n", a-b)
}
