# MultiTask

MultiTask is simple task manager. you can use it to multiple tasks and receive result from these tasks.

# Quick start

**create two task to count number.**

```go
mt := NewMultiTask()

mt.Do(func(ch chan<- interface{}) {
    for i := 0; i < 100; i++ {
        ch <- 1
    }
})
mt.Do(func(ch chan<- interface{}) {
    for i := 0; i < 100; i++ {
        ch <- 1
    }
})

cnt := 0
mt.Fetch(func(result interface{}) {
    cnt += 1
})

mt.Wait()
```

**create two task, not to recv.**

```go
mt := NewMultiTask()

mt.Do(func(ch chan<- interface{}) {
    for i := 0; i < 100; i++ {
        // do something
    }
})
mt.Do(func(ch chan<- interface{}) {
    for i := 0; i < 100; i++ {
        // do something
    }
})

mt.Wait()
```