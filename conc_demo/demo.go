package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/pool"
	"github.com/sourcegraph/conc/stream"
	"go.uber.org/atomic"
	"log"
	"time"
)

func main() {
	streamDemo()
}

// waitGroupDemo 自动recover
// 并发的处理，并且等待消费完成，适用于不用保证顺序并发消费场景
func waitGroupDemo() {
	var count atomic.Int64

	var wg conc.WaitGroup
	// 开启10个goroutine并发执行 count.Add(1)
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			if i == 7 {
				panic("bad thing")
			}
			count.Add(1)
		})
	}
	// 等待10个goroutine都执行完，并且recover panic
	wg.WaitAndRecover()

	fmt.Println(count.Load())
}

// poolDemo goroutine池示例
// 执行结果也是无序的，使用线程池进行无序的任务消费
func poolDemo() {
	// 创建一个最大数量为3的goroutine池
	p := pool.New().WithMaxGoroutines(3)
	// 使用p.Go()提交5个任务
	for i := 0; i < 5; i++ {
		idx := i
		p.Go(func() {
			fmt.Println("pluto ", idx)
		})
	}
	p.Wait()
}

// poolDemo 支持context的池
// 在发生err或panic的时候，结束所有goroutine
func poolDemoCancelOnError() {
	p := pool.New().
		WithMaxGoroutines(4). // 设定线程池的goroutine数量
		WithContext(context.Background()).
		WithCancelOnError() // 出错时取消所有goroutine
	// 提交3个任务
	for i := 0; i < 3; i++ {
		idx := i
		p.Go(func(ctx context.Context) error {
			if idx == 2 {
				return errors.New("cancel all other tasks")
			}
			log.Println(idx)
			<-ctx.Done()
			return nil
		})
	}
	err := p.Wait()
	fmt.Println(err)
}

// streamDemo 并发的流式任务示例
// 保证消费的顺序性
func streamDemo() {
	times := []int{20, 52, 16, 45, 4, 80}

	s := stream.New()
	for _, millis := range times {
		dur := time.Duration(millis) * time.Millisecond
		// 提交任务
		s.Go(func() stream.Callback {
			// 虽然进行了休眠，但还是先入先出，保证了消费的顺序性
			time.Sleep(dur)
			return func() { fmt.Println(dur) }
		})
	}
	s.Wait()
}
