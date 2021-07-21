package modle

import (
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	fmt.Println("测试book中的所有数据")
	t.Run("????",testGetBooks)
}

func testGetBooks(t *testing.T)  {
	boos:= GetBook(3)
	fmt.Println(boos)
}