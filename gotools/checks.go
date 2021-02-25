/**********

#### checks 常用检查方法

```
* CheckEmail //验证邮箱是否正确
* CheckMobileNum //验证手机号是否正确
```

*****/
package gotools

import "regexp"

/**
 * 验证邮箱是否正确
 */
func CheckEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

/**
 * 验证手机号是否正确
 */
func CheckMobileNum(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}