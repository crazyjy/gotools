---

#### strings 常用字符串方法

```
* RandomString //获取任意长度随机字符串
* UniqueString32 //获取32位唯一id
```

---


#### curls 常用网络方法

```
* HttpGet //发送普通Get请求
* HttpParamsGet //发送普通Get带参数请求
* HttpHeadGet //带请求头get请求
* HttpPost //发送post请求
* HttpPostJson //发送JSON数据POST方法一
* HttpPostJson2 //发送JSON数据POST方法二
```

---


#### times 常用时间方法

```
* NowDataTime //获取东八区当前时间 Y-m-d H:i:s
* NowDateTimeMap //获取东八区当前时间map Y m d H i s
* NowDateTimeArr //获取东八区当前时间 时，分，秒 ["23","05","40"]
* MonthRange //获取当月日期范围 ["2021-02-01", "2021-02-28"]
* LostMonth //获取前几个月日期
* FutureMonth //获取后几个月日期
* LostDay //获取前几天日期
* FutureDay //获取后几天日期
* LostTimeS //获取前几秒时间
* FutureTimeS //获取后几秒时间
```

---


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

---


#### checks 常用检查方法

```
* CheckEmail //验证邮箱是否正确
* CheckMobileNum //验证手机号是否正确
```

---


#### sorts 常用排序方法

```
* BubbleSortAscInt // 冒泡升序INT
* BubbleSortDescInt // 冒泡降序INT
* BubbleSortAscStr // 冒泡升序String
* BubbleSortDescStr // 冒泡降序String
* GoSortAscInt // 内置升序INT
* GoSortAscStr // 内置升序String
```

---


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

---
