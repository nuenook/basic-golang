# chanel
 use arrow `<-`

 ## Sending 
 assign calculated value to chanel `c`
  ```go

  func foo(c chan int, someValue int) {
    c <- someValue * 5
  }
  ```

## Reciept
create `fooVal` channel with make be integer

start woking by routines `go foo(fooVal, 3)`

reciept with v1 , v2  
```go
    fooVal := make(chan int)
    go foo(fooVal, 5)
    go foo(fooVal, 3)

    v1 ,v2 := <-fooVal, <- fooVal
```

## Buffering
buffer 10 on `fooVal` channel
```go
fooVal := make(chane int, 10)
```