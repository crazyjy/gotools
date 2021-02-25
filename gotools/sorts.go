
/**********

#### sorts 常用排序方法

```
* BubbleSortAscInt // 冒泡升序INT
* BubbleSortDescInt // 冒泡降序INT
* BubbleSortAscStr // 冒泡升序String
* BubbleSortDescStr // 冒泡降序String
* GoSortAscInt // 内置升序INT
* GoSortAscStr // 内置升序String
```

*****/

package gotools

import (
	"sort"
)

/**
 * 冒泡升序INT
 */
func BubbleSortAscInt(numSlice []int) []int {
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] > numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}
		}
	}
	return numSlice
}

/**
 * 冒泡降序INT
 */
func BubbleSortDescInt(numSlice []int) []int {
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] < numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}
		}
	}
	return numSlice
}

/**
 * 冒泡升序string
 */
func BubbleSortAscStr(strSlice []string) []string {
	for i := 0; i < len(strSlice); i++ {
		for j := i + 1; j < len(strSlice); j++ {
			if strSlice[i] > strSlice[j] {
				temp := strSlice[i]
				strSlice[i] = strSlice[j]
				strSlice[j] = temp
			}
		}
	}
	return strSlice
}

/**
 * 冒泡降序string
 */
func BubbleSortDescStr(strSlice []string) []string {
	for i := 0; i < len(strSlice); i++ {
		for j := i + 1; j < len(strSlice); j++ {
			if strSlice[i] < strSlice[j] {
				temp := strSlice[i]
				strSlice[i] = strSlice[j]
				strSlice[j] = temp
			}
		}
	}
	return strSlice
}

/**
 * 内置升序排序INT
 */
func GoSortAscInit(intSlice []int) []int {
	sort.Ints(intSlice)
	return intSlice
}

/**
 * 内置升序排序String
 */
func GoSortAscStr(strSlice []string) []string {
	sort.Strings(strSlice)
	return strSlice
}