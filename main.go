package main

import (
	"fmt"
	"time"
)

func main() {

	//Test cases

	fmt.Println("Case 1 : Get")
	//case 1 - simple get(), blocks execution of code below, call get any number of times
	fut := createFutureTask(func() result {
		time.Sleep(10 * time.Second)
		r := result{res: 15}
		return r
	})

	//expected output
	/*
		Result: {15 {}}
		Done: true
		Cancelled: false
		{15 {}}  printing Result for the second time
	*/

	fut.get()
	fmt.Println("Result:", fut.result)
	fmt.Println("Done:", fut.isDone())
	fmt.Println("Cancelled:", fut.isCancelled())
	fmt.Println(fut.get(), " printing Result for the second time")

	fmt.Println("Case 2 : Cancel a Future")
	//case 2 - cancel a future in between, you can call get any number of times..will return "cancelled" as error
	fut2 := createFutureTask(func() result {
		time.Sleep(10 * time.Second)
		r := result{res: 15}
		return r
	})

	//expected output
	/*
		Cancelled: true
		Done: false
		Result: {<nil> {cancelled}}
	*/

	fut2.cancel(true)
	fmt.Println("Cancelled:", fut2.isCancelled())
	fmt.Println("Done:", fut2.isDone())
	fmt.Println("Result:", fut2.get())

	fmt.Println("Case 3 : Get With Timeout")
	//case 3 - timeout is more than function time, so that no timeout occurs
	fut3 := createFutureTask(func() result {
		time.Sleep(10 * time.Second)
		r := result{res: 15}
		return r
	})

	//expected output
	/*
		Done: true
		Result: {15 {}}
	*/

	// so that no timeout occurs
	fut3.getWithTimeout(11)
	fmt.Println("Done:", fut3.isDone())
	fmt.Println("Result:", fut3.result)

	fmt.Println("Case 4 : Timeout!")
	//case 4 - timeout is less than function time, so that timeout occurs
	fut4 := createFutureTask(func() result {
		time.Sleep(10 * time.Second)
		r := result{res: 15}
		return r
	})

	//expected output
	/*
		Done: false
		Result: {<nil> {timeout}}
	*/

	fut4.getWithTimeout(5)
	fmt.Println("Done:", fut4.isDone())
	fmt.Println("Result:", fut4.result)

	fmt.Println("Case 5 : Failed Future")
	//case 5 - future fails
	fut5 := createFutureTask(func() result {
		time.Sleep(4 * time.Second)
		r := result{myError: myError{"future failed"}}
		return r
	})

	//expected output
	/*
		Done: false
		Result: {<nil> {future failed}}
	*/

	fut5.get()
	fmt.Println("Done:", fut5.isDone())
	fmt.Println("Result:", fut5.result)

	fmt.Println("Case 6 : Wait for result")
	//case 6 - Dont get future, just wait for result
	fut6 := createFutureTask(func() result {
		time.Sleep(4 * time.Second)
		r := result{res: 15}
		return r
	})

	//expected output
	/*
		Done: true
		Result: {15 {}}
	*/

	time.Sleep(5 * time.Second)
	fmt.Println("Done:", fut6.isDone())
	fmt.Println("Result:", fut6.result)

}
