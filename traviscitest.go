package traviscitest

import "errors"
import "fmt"

func IsHelloWorld(word string) (bool, error) {
	switch word {
	case "hello":
		fmt.Println("This is Hello World!")
		return true, nil
	case "notHello":
		err := errors.New("What's this?")
		fmt.Println("This is not Hello World!")
		return false, err
	}
	return false, nil
}

