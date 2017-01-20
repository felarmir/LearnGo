package lesson3

import (
	"errors"
	"fmt"
)

func OwnErrorMessage() {
	err := errors.New("My error message")
	fmt.Println(err.Error())
}
