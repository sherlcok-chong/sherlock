package main
import "fmt"
func main() {
	/*
	 panic：词义"恐慌"，
	 recover："恢复"
	 go语⾔利⽤panic()，recover()，实现程序中的极特殊的异常的处理
	 panic(),让当前的程序进⼊恐慌，中断程序的执⾏
	 recover(),让程序恢复，必须在defer函数中执⾏
	*/
	funA()
	funB()
	funC()
	fmt.Println("main...over....") }
func funA() {
	fmt.Println("我是函数funA()...") }
func funB() { //外围函数
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "恢复啦。。。")
		}
	}()
	fmt.Println("我是函数funB()...")
	for i := 1; i <= 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			//让程序中断
			panic("funB函数，恐慌啦。。。") //打断程序的执⾏。。
		}
	}
	//当外围函数中的代码引发运⾏恐慌时，只有其中所有的延迟函数都执⾏完毕后，该运 ⾏时恐慌才会真正被扩展⾄调⽤函数。
}
func funC() {
	defer func() {
		fmt.Println("func的延迟函数。。。")
		if msg := recover(); msg != nil {
		 fmt.Println(msg, "恢复啦。。。")
		}
		fmt.Println("recover执⾏了" , recover())
	}()
	fmt.Println("我是函数funC()。。")
	panic("funC恐慌啦。。") }
