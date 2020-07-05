package main

type future interface {
	cancel(bool) bool
	get() result
	getWithTimeout(int) result
	isCancelled() bool
	isDone() bool
}
