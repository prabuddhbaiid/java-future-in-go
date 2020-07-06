package main

import "time"

type futureTask struct {
	result    result
	done      bool
	running   bool
	cancelled bool
	channel   chan result
}

func createFutureTask(f func() result) *futureTask {
	//future channel
	futureObjChannel := make(chan result, 1)

	//reference to future object
	futureObj := &futureTask{
		result:    result{},
		done:      false,
		running:   true,
		cancelled: false,
		channel:   futureObjChannel}

	//initilizing the future task in a new thread
	go func() {
		//channel waiting for a result
		futureObjChannel <- f()

		//close channel
		defer close(futureObjChannel)

		//if get() or getWithTimeout() has not been called on the futureObject
		//take the result from future channel
		//update attributes
		//close channel
		if len(futureObjChannel) > 0 {
			futureObj.result = <-futureObjChannel
			futureObj.running = false
			futureObj.done = true
		}
	}()

	return futureObj
}

func (futureTask *futureTask) isDone() bool {
	return futureTask.done
}

func (futureTask *futureTask) isCancelled() bool {
	return futureTask.cancelled
}

func (futureTask *futureTask) cancel(b bool) bool {
	switch {

	//if input to cancel is false, then don't cancel
	case !b:
		return false

	//if future is done/has returned an error, i.e. cancelled or timed out
	case futureTask.done:
		return false

	//if future is running, then cancel, send a nil result and an error message that it has been cancelled
	case futureTask.running:
		futureTask.cancelled = true
		futureTask.result = result{res: nil, myError: myError{"cancelled"}}
		futureTask.running = false
		futureTask.done = true
		futureTask.channel <- futureTask.result

		return true
	}
	return false
}

func (futureTask *futureTask) get() result {

	//if task is cancelled/done/has returned an error...return the result immediately
	if futureTask.done {
		return futureTask.result
	}

	//wait till task in future has been accomplished
	//update attributes
	select {
	case futureTask.result = <-futureTask.channel:
		futureTask.done = true
		futureTask.running = false

		return futureTask.result
	}
}

func (futureTask *futureTask) getWithTimeout(timeout int) result {

	if futureTask.done {
		return futureTask.result
	}

	//wait for race between task in future and timeout
	//update attributes
	select {
	case futureTask.result = <-futureTask.channel:

	case <-time.After(time.Duration(time.Second * time.Duration(timeout))):
		<-futureTask.channel
		futureTask.result = result{res: nil, myError: myError{"timeout"}}
	}

	futureTask.running = false
	futureTask.done = true
	return futureTask.result
}
