package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("18.1 - Interface de tipo generico")
	generica("oie")
	generica(true)
	generica(1.2)
	generica(1)
}
