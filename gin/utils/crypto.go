package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func MD5(str []byte) string {
	h := md5.New()
	h.Write(str)
	hex := hex.EncodeToString(h.Sum(nil))
	a := hex[8:24]
	fmt.Printf("32位 = %v  16位 = %v", hex, a)
	return a
}

// uuid 加盐salt
func GenerateSalt() string {
	u := uuid.NewV4()
	res := fmt.Sprintf("%s", u)
	return res
}

// 字符串 转 数值型数组
func ArrToStr(arr []int) string {
	var str string
	var tmp []string
	for _, v := range arr {
		tmp = append(tmp, strconv.Itoa(v))
	}
	str = strings.Join(tmp, ",")
	return str
}

func StrToArr(str string) []int {
	var arr []int
	var tmp []string
	tmp = strings.Split(str, ",")

	for _, v := range tmp {
		i, _ := strconv.Atoi(v)
		arr = append(arr, i)
	}
	return arr
}

// 字符串json 转 结构体
