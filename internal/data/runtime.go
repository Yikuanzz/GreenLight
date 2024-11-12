package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	// 生成目标字符串格式
	jsonValue := fmt.Sprintf("%d mins", r)
	// 给字符串加上双引号
	quotedJSONValue := strconv.Quote(jsonValue)
	// 以字节切片返回结果
	return []byte(quotedJSONValue), nil
}
