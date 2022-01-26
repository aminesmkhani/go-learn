package hello

import (
	"fmt"
	"sectron.ir/greetings"
)

func main() {
	message := greetings.Hello("Amin")
	fmt.Println(message)
}
