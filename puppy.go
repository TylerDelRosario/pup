package puppy

import (
	"github.com/TylerDelRosario/dog"
	"github.com/TylerDelRosario/pup/new"
)

func Bark() string {
	return new.Test()
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