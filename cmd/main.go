package main

import "fmt"

func returnError(value int) error {
	return fmt.Errorf("This is an erroe with value %v", value)
}

func main() {
	err := returnError(5)

}
