
package main

import (
	"fmt"
	"time"
)

func main() {
	// 返回一个通道：chan，存储的是d时间见后的当前时间。
	//func After(d Duration) <-chan Time
	ch1 := time.After(3 * time.Second) // 16:19:00
	fmt.Printf("%T\n", ch1) // <-chan time.Time
	fmt.Println(time.Now())
	time2 := <-ch1
	fmt.Println(time2)

	//2.time.After(Duration),同time.NewTimer(Duration).C
	timer:=time.NewTimer(5*time.Second)
	fmt.Printf("%T\n",timer) //<-chan time.Time
	ch2:=timer.C //<-chan time.Time
	fmt.Println(<-ch2)

}

