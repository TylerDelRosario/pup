package puppy

import (
	"github.com/TylerDelRosario/dog"
)

func Bark() string {
	return "Woof!!!"
}

func Barks() string {
	return "Woof!! Woof!! Woof!!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}