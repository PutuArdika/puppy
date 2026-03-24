package puppy

import "github.com/GoesToEleven/dog"

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof Woof"
}

func BigBarks() string {
	return "And when grown up it will " + dog.WhenGrownUp(Barks())
}