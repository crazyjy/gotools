
/**********

#### formats 常用转换方法

```
* IntToString //数字格式化成字符串
* Int64ToString //64数字格式化成字符串
* StringToInt //字符串格式化成数字
* StringToInt64 //字符串格式化成数字64
* ByteToString //字节格式化成字符串
* StringToByte //字符串格式化成字节
* RuneToString //UTF8-字节格式化成字符串
* StringToRune //字符串格式化成UTF8-字节
* InitString //初始化输出字符串
```

*****/

package gotools

import (
	"fmt"
	"strconv"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToInt(s string) int {
	int, _ := strconv.Atoi(s)
	return int
}

func StringToInt64(s string) int64 {
	int64, _ := strconv.ParseInt(s, 10, 64)
	return int64
}

func ByteToString(b []byte) string {
	return string(b)
}

func StringToByte(s string) []byte {
	return []byte(s)
}

func RuneToString(b []rune) string {
	return string(b)
}

func StringToRune(s string) []rune {
	return []rune(s)
}

func InitString(s interface{}) string {
	return fmt.Sprintf("%v", s)
}

