package main

import (
	"fmt"
	"sync"
)

func main() {
	c := &sync.WaitGroup{}
	ss := []string{"A", "B", "C"}
	ch := make([]chan struct{}, 3)
	for i := 0; i < len(ch); i++ {
		ch[i] = make(chan struct{})
	}
	for it, vv := range ss {
		cc := it
		v := vv
		c.Add(1)
		go func() {
			// 收到消息后可发送
			for i := 0; i < 5; i++ {
				//fmt.Println(cc)
				<-ch[cc]
				fmt.Println(v)
				ch[(cc+1)%3] <- struct{}{}
			}
			if cc == 0 {
				// 把三号的最后一次信号接受了，防止阻塞
				<-ch[0]
			}
			c.Done()
		}()
	}
	ch[0] <- struct{}{}
	c.Wait()
}
