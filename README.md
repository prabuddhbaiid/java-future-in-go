# java-future-in-go
Repository is an implementation of Java 8 Future interface in Golang

To run the files clone repo, open directory and use the command "go run ."

General Behaviour of FutureTask:

When a futureTask is created
* A new goroutine is created
* Other tasks can be executed while this goroutine executes
* It contains a result 
* Result contains res (value of the result) and myError
* MyError contains a message (indicating the error)
* Operations
  * We can block execution by calling get() method and waiting for FutureTask
  * We can block execution for a given time by calling getWithTimeOut(seconds) method
  * We can cancel a future by calling cancel(true) method
  * We can check if a future is completed by calling isDone() method
  * We can check if a future has been cancelled by calling isCancelled() method
  
The following cases have been executed:  
In each case (except 5), a simple int of 15 has been passed as the result of every computation (futureTask)  

case 1 - simple get(), blocks execution of code below, call get any number of times  
Expected Output -  
Result: {15 {}}  
Done: true  
Cancelled: false  
{15 {}}  printing Result for the second time  

case 2 - cancel a future in between, returns "cancelled" as error  
Expected Output -  
Cancelled: true  
Done: false  
Result: {<nil> {cancelled}}  

case 3 - simple getWithTimeOut()  
Expected Output -   
Done: true  
Result: {15 {}}  

case 4 - timeout error  
Expected Output -  
Done: false  
Result: {<nil> {timeout}}  

case 5 - future returns a failure  
Expected Output -   
Done: false  
Result: {<nil> {future failed}}  

case 6 - wait for result  
Expected Output -  
Done: true  
Result: {15 {}}  
