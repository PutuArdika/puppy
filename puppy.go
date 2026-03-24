package puppy

import "github.com/GoesToEleven/dog"

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof Woof"
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}