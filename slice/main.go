package main

import (
	"errors"
	"fmt"
	"runtime"
)

func main() {
	//TestV1()
	//TestV2()
	TestV3()
}

func TestV1() {
	s := []int{1, 2, 3}
	fmt.Printf("%v \n", s)
	var err error

	s1, err := SliceDelInt(s, 0)
	if err != nil {
		fmt.Println("操作失败：", err.Error())
	}
	fmt.Printf("删除第一个元素：%v \n", s1)

	s2, err := SliceDelInt(s, 1)
	if err != nil {
		fmt.Println("操作失败：", err.Error())
	}
	fmt.Printf("删除第二个元素：%v \n", s2)

	s3, err := SliceDelInt(s, 2)
	if err != nil {
		fmt.Println("操作失败：", err.Error())
	}
	fmt.Printf("删除最后一个元素：%v \n", s3)

	_, err = SliceDelInt(s, 3)
	if err != nil {
		fmt.Println("操作失败：", err.Error())
	}
}

// 删除整型切片中指定下标元素
func SliceDelInt(s []int, index int) ([]int, error) {
	if index < 0 || index >= len(s) {
		return s, errors.New("非法下标值，不能进行删除操纵")
	}
	//创建新切片
	newS := make([]int, 0, len(s)-1)

	//以index为界限，拼接index左右两边的两个小切片
	if index > 0 {
		for _, v := range s[:index] {
			newS = append(newS, v)
		}
	}

	if index < (len(s) - 1) {
		for _, v := range s[index+1:] {
			newS = append(newS, v)
		}
	}

	return newS, nil
}

func TestV2() {
	sInt := []int{1, 2, 3, 4, 5, 6}
	s1, err := SliceDelV1(sInt, 3)
	if err != nil {
		fmt.Println("操作失败：", err.Error())
	}
	fmt.Printf("删除整型切片：%v \n", s1)

	sString := []string{"one", "two", "three", "four", "five"}
	s2, err := SliceDelV1(sString, 2)
	if err != nil {
		fmt.Println("操作失败：", err.Error())
	}
	fmt.Printf("删除字符串切片：%v \n", s2)
}

// 删除切片的指定下标元素操作的泛型版
func SliceDelV1[T any](s []T, index int) ([]T, error) {
	if index < 0 || index >= len(s) {
		return s, errors.New("非法下标值，不能进行删除操纵")
	}
	//创建新切片
	newS := make([]T, 0, len(s)-1)

	//以index为界限，拼接index左右两边的两个小切片
	if index > 0 {
		for _, v := range s[:index] {
			newS = append(newS, v)
		}
	}

	if index < (len(s) - 1) {
		for _, v := range s[index+1:] {
			newS = append(newS, v)
		}
	}

	return newS, nil
}

func TestV3() {
	src1, _, err := SliceDelV2(make([]int, 10, 50), 0)
	if err != nil {
		fmt.Println("操作失败：", err.Error())
	}
	fmt.Printf("新容量：%d \n", cap(src1))

	src2, _, err := SliceDelV2(make([]int, 250, 1000), 0)
	if err != nil {
		fmt.Println("操作失败：", err.Error())
	}
	fmt.Printf("新容量：%d \n", cap(src2))

	src3, _, err := SliceDelV2(make([]int, 500, 1000), 0)
	if err != nil {
		fmt.Println("操作失败：", err.Error())
	}
	fmt.Printf("新容量：%d \n", cap(src3))

}

// (泛型)切片删除操作优化版V2，相对V1版本，优化点：
// 1、删除过程不进行内存重新分配
// 2、返回值添加被删除的元素
// 3、将缩容操作独立，与删除操作进行解耦
func SliceDelV2[T any](s []T, index int) ([]T, T, error) {
	length := len(s)
	var zeroT T
	if index < 0 || index >= length {
		return s, zeroT, fmt.Errorf("下标：%d的，超出了0～%d的范围", index, length-1)
	}

	res := s[index] //提前取出要删除的元素
	for i := index; i < length-1; i++ {
		s[i] = s[i+1] //从index位置开始，后一个元素覆盖前一个元素
	}

	s = s[:length-1] //去除最后一个元素
	s = Shrink(s)    //缩容

	return s, res, nil
}

// 对切片进行缩容操作
func Shrink[T any](s []T) []T {
	c := cap(s)
	l := len(s)

	newCap, ok := calCapacity(c, l)
	if !ok {
		return s
	}

	newS := make([]T, 0, newCap)
	newS = append(newS, s...)

	return newS
}

// 判断是否要缩容，以及返回新容量的大小
func calCapacity(c, l int) (int, bool) {
	if c <= 64 {
		return c, false //原本的容量就很小了，不需要所容
	}
	demarc := 256
	if runtime.Version() < "go1.18" {
		demarc = 1024 //go1.18之前的版本临界值是1024，之后就变成256了
	}

	if l <= demarc && c/l > 2 {
		return l * 2, true
	}

	if l > demarc && float64(c)/float64(l) > 1.25 {
		return int(float64(l) * 1.25), true
	}

	return c, false
}
