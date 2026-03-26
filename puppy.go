package puppy

import "github.com/GoesToEleven/dog"

//added v1.0.1 minor changes

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof Woof"
}

func BigBarks() string {
	return "And when grown up it will " + dog.WhenGrownUp(Barks())
}

func FromLatestVersion() string {
	return "from v1.2.0"
}
