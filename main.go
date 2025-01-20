package main

import (
	"fmt"
	"sluzov/greeting" // Импортируйте ваш пакет по имени из go.mod
)

func main() {
	fmt.Println("This is main!")
	greeting.Hello()
	greeting.Hi()
}
