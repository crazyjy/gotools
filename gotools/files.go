/**********

#### files 常用文件方法

```
* CheckDir //检查文件夹是否存在
* CheckFile //检查文件是否存在
* CreateDir //新建文件夹
* CreateFile //新建空文件
* ReadFile //读取文件所有内容（快速）
* LoopReadFile //方式循环读取文件，每次 1M
* WriteNewFile //覆盖写入新文件（快速）
* AppendToFile //追加内容到文件中（快速）
* LoopCopyFile //循环方式拷贝文件 每次10M
* DelFile //删除文件
* GetFileList //获取文件夹下文件列表
* DelDir //删除文件夹和文件
```

*****/

package gotools

import (
	"io"
	"io/ioutil"
	"os"
)

/**
 * 检查文件夹是否存在
 */
func CheckDir(path string) bool {
	s,err:=os.Stat(path)
	if err!=nil{
		return false
	}
	return s.IsDir()
}

/**
 * 检查文件是否存在
 */
func CheckFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/**
 * 新建文件夹
 */
func CreateDir(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

/**
 * 新建空文件
 */
func CreateFile(path string) bool {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return false
	}
	return true
}

/**
 * 读取文件所有内容（快速）
 */
func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	info, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(info), nil
}

/**
 * 方式循环读取文件，每次 1M
 */
func LoopReadFile(path string) (data string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	var content []byte
	var tmp = make([]byte, 1024 * 1024 * 1)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			break
		}
		content = append(content, tmp[:n]...)
	}
	data = string(content)
	return
}

/**
 * 覆盖写入新文件（快速）
 */
func WriteNewFile(path string, info string) bool  {
	err := ioutil.WriteFile(path, []byte(info), 0666)
	if err != nil {
		return false
	}
	return true
}

/**
 * 追加内容到文件中（快速）
 */
func AppendToFile(path string, content string) bool {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return false
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		return false
	}
	return true
}

/**
 * 循环方式拷贝文件 每次10M
 */
func LoopCopyFile(dstFileName string, srcFileName string) (bool, error) {
	source, _ := os.Open(srcFileName)
	destination, _ := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	buf := make([]byte, 1024 * 1024 * 10)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return false, err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return false, err
		}
	}
	return true, nil
}

/**
 * 删除文件
 */
func DelFile(path string) bool {
	if CheckFile(path) {
		err := os.Remove(path)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

/**
 * 获取文件夹下文件列表
 */
func GetFileList(path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return files, nil
}

/**
 * 删除文件夹和文件
 */
func DelDir(path string) bool {
	err := os.RemoveAll(path)
	if err != nil {
		return false
	}
	return true
}