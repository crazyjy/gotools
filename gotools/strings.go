/**********

#### strings 常用字符串方法

```
* RandomString //获取任意长度随机字符串
* UniqueString32 //获取32位唯一id
```

*****/

package gotools

import (
	"crypto/md5"
	cryRand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	mathRand "math/rand"
	"time"
)

/**
 * 获取任意长度随机字符串
 */
func RandomString(long int) string  {
	str := "0123456789AabBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
	bytes := []byte(str)
	result := []byte{}
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	for i := 0; i < long; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/**
 * 获取32位唯一id
 */
func UniqueString32() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(cryRand.Reader, b); err != nil {
		return ""
	}
	//生成32位md5字串
	h := md5.New()
	h.Write([]byte(base64.URLEncoding.EncodeToString(b)))
	//生成Guid字串
	return hex.EncodeToString(h.Sum(nil))
}