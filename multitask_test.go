package multitask

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	executor := NewMultiTask()
	now := time.Now()
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 10; i++ {
			ch <- 1
		}
	})

	cnt := 0
	executor.Fetch(func(result interface{}) {
		cnt += 1
	})

	executor.Wait()
	t.Logf("all task cost: %dms\n", time.Since(now).Milliseconds())
	t.Logf("task count: %d\n", cnt)
	assert.Equal(t, 310, cnt)
}

func TestTaskWithEmpty(t *testing.T) {
	executor := NewMultiTask()
	now := time.Now()
	executor.Do(func(ch chan<- interface{}) {
	})
	executor.Do(func(ch chan<- interface{}) {
	})
	executor.Do(func(ch chan<- interface{}) {
	})
	executor.Do(func(ch chan<- interface{}) {
	})

	cnt := 0
	executor.Fetch(func(result interface{}) {
		cnt += 1
	})

	executor.Wait()
	t.Logf("all task cost: %dms\n", time.Since(now).Milliseconds())
	t.Logf("task count: %d\n", cnt)
	assert.Equal(t, 0, cnt)
}

func TestTaskWithNoRecv(t *testing.T) {
	executor := NewMultiTask()
	now := time.Now()
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 10; i++ {
			ch <- 1
		}
	})

	executor.Wait()
	t.Logf("all task cost: %dms\n", time.Since(now).Milliseconds())
}

func TestTaskWithWorkSleep(t *testing.T) {
	executor := NewMultiTask(WithQueueSize(5))
	now := time.Now()
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
			time.Sleep(time.Millisecond * 2)
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
			time.Sleep(time.Millisecond * 1)
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
			time.Sleep(time.Millisecond * 5)
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 10; i++ {
			ch <- 1
		}
	})

	cnt := 0
	executor.Fetch(func(result interface{}) {
		cnt += 1
		time.Sleep(time.Millisecond * 2)
	})

	executor.Wait()
	t.Logf("all task cost: %dms\n", time.Since(now).Milliseconds())
	t.Logf("task count: %d\n", cnt)
	assert.Equal(t, 310, cnt)
}

func TestTaskWithRecvSleep(t *testing.T) {
	executor := NewMultiTask()
	now := time.Now()
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 10; i++ {
			ch <- 1
		}
	})

	cnt := 0
	executor.Fetch(func(result interface{}) {
		cnt += 1
		time.Sleep(time.Millisecond * 2)
	})

	executor.Wait()
	t.Logf("all task cost: %dms\n", time.Since(now).Milliseconds())
	t.Logf("task count: %d\n", cnt)
	assert.Equal(t, 310, cnt)
}

func TestTaskWithSmallBuffer(t *testing.T) {
	executor := NewMultiTask(WithQueueSize(2))
	now := time.Now()
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 10; i++ {
			ch <- 1
		}
	})

	cnt := 0
	executor.Fetch(func(result interface{}) {
		cnt += 1
	})

	executor.Wait()
	t.Logf("all task cost: %dms\n", time.Since(now).Milliseconds())
	t.Logf("task count: %d\n", cnt)
	assert.Equal(t, 310, cnt)
}

func TestTaskWithLargeBuffer(t *testing.T) {
	executor := NewMultiTask(WithQueueSize(5000))
	now := time.Now()
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
	})
	executor.Do(func(ch chan<- interface{}) {
		for i := 0; i < 10; i++ {
			ch <- 1
		}
	})

	cnt := 0
	executor.Fetch(func(result interface{}) {
		cnt += 1
	})

	executor.Wait()
	t.Logf("all task cost: %dms\n", time.Since(now).Milliseconds())
	t.Logf("task count: %d\n", cnt)
	assert.Equal(t, 310, cnt)
}
