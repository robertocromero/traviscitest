package traviscitest

import "errors"
import "fmt"

func IsHelloWorld(word string) (bool, error) {
	switch word {
	case "hello":
		fmt.Println("This is Hello Wooorld!")
		return true, nil
	case "notHello":
		err := errors.New("What's this?")
		fmt.Println("This is not Hello Wooooorld!")
		return false, err
	}
	return false, nil
}

