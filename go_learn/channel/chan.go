package channel_

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func Test_loopChan() {

	c := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		i := 0
		for {
			select {
			case <-ctx.Done():
				time.Sleep(100 * time.Millisecond)
				fmt.Println("timeout!")
				close(c)
				return
			default:
				time.Sleep(1 * time.Second)
				c <- "runs of " + strconv.Itoa(i)
			}
			i++
		}
	}(ctx)

	fmt.Println("reciving message..., press CTRL & C to stop or waiting for 10 seconds")
	for msg := range c {
		fmt.Println(msg)
	}
	fmt.Println("channel is closed")
}
