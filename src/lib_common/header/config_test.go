package header

import (
	"testing"
	"strconv"
	"util/logger"
	"fmt"
)

func Test1(t *testing.T) {
	a := ""

	n, e := strconv.Atoi(a)
	if e == nil {
		fmt.Println(n)
	} else {
		logger.Error("转换失败：", e)
	}
}


func Test2(t *testing.T) {
	fmt.Println(OperationHeadByteMap[0])
	fmt.Println(OperationHeadByteMap[1])
	fmt.Println(OperationHeadByteMap[2])
	fmt.Println(OperationHeadByteMap[3])
}
