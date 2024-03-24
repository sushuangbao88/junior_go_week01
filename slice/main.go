package main

import (
	"errors"
	"fmt"
)

func main() {
	//TestV1()
	TestV2()
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
	s1, err := SliceDel(sInt, 3)
	if err != nil {
		fmt.Println("操作失败：", err.Error())
	}
	fmt.Printf("删除整型切片：%v \n", s1)

	sString := []string{"one", "two", "three", "four", "five"}
	s2, err := SliceDel(sString, 2)
	if err != nil {
		fmt.Println("操作失败：", err.Error())
	}
	fmt.Printf("删除字符串切片：%v \n", s2)
}

// 删除切片的指定下标元素操作的泛型版
func SliceDel[T any](s []T, index int) ([]T, error) {
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
