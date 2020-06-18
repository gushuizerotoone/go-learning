package v_4_concurrent_slice_append

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// method 1, use lock to process concurrent append
func TestLock(t *testing.T) {
	beginTime := time.Now()
	defer func() {
		fmt.Printf("cost: %v", time.Since(beginTime))
	}()

	slc := make([]int, 0, 10000000)
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()
			// use lock
			lock.Lock()
			defer lock.Unlock()
			slc = append(slc, a)
		}(i)

	}
	wg.Wait()
	fmt.Println(len(slc))
}

// method 2, use channel to concurrent append,
// if the concurrent count is not very high, performance: method 1  > method 2
// if the concurrent count is high, performance: method 2 > method 1
func TestChannel(t *testing.T) {
	beginTime := time.Now()
	defer func() {
		fmt.Printf("cost: %v", time.Since(beginTime))
	}()
	var (
		wg sync.WaitGroup
		n  = 20
	)
	c := make(chan struct{})

	// new 了这个 job 后，该 job 就开始准备从 channel 接收数据了
	s := NewScheduleJob(n, func() { c <- struct{}{} }) // 这里用了一个channel在等range遍历完

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(v int) {
			defer wg.Done()
			s.AddData(v)

		}(i)
	}

	wg.Wait() // 发完了才能执行下面的close
	s.Close()
	<-c // 这里用了一个channel在等range遍历完

	fmt.Println(len(s.data))
}

type ServiceData struct {
	ch   chan int // 用来 同步的channel
	data []int    // 存储数据的slice
}

func (s *ServiceData) Schedule() {
	// 从 channel 接收数据
	for i := range s.ch {
		fmt.Println("append..")
		s.data = append(s.data, i)
	}
	fmt.Println("finish schedule..")
}

func (s *ServiceData) Close() {
	// 最后关闭 channel
	fmt.Println("close..")
	close(s.ch)
}

func (s *ServiceData) AddData(v int) {
	fmt.Println("add data..")
	s.ch <- v // 发送数据到 channel
}

func NewScheduleJob(size int, done func()) *ServiceData {
	s := &ServiceData{
		ch:   make(chan int, size),
		data: make([]int, 0),
	}

	go func() {
		// 并发地 append 数据到 slice
		s.Schedule()
		fmt.Println("before done...")
		done()
	}()

	return s
}

