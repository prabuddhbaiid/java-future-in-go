package main

import "fmt"

type result struct {
	res     interface{}
	myError myError
}

type myError struct {
	message string
}

func (e myError) Error() string {
	return fmt.Sprintf("%v", e)
}
