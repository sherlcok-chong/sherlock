package main

import (
	"fmt"
	"time"
)

func main() {
	res1, err := getArea(4, -6)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("⾯积是:", res1)
	}
}

type MyError struct {
	When time.Time
	What string
}

//2、⾃定义错误类型实现error接⼝的⽅法 ：Error() string
func (e *MyError) Error() string {
	return fmt.Sprintf("%v : %v", e.When, e.What)
}
func getArea(width, length float64) (float64, error) {
	errorMsg := ""
	if width < 0 && length < 0 {
		errorMsg = fmt.Sprintf("⻓度:%v ，宽度:%v ，均为负数", length,
			width)
	} else if length < 0 {
		errorMsg = fmt.Sprintf("⻓度:%v ，出现负数", length)
	} else if width < 0 {
		errorMsg = fmt.Sprintf("宽度:%v ，出现负数", width)
	}
	if errorMsg != "" {
		return 0, &MyError{time.Now(), errorMsg}
	} else {
		return width * length, nil
	}
}
