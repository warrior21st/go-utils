package commonutil

import (
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"github.com/shopspring/decimal"
)

//是否为null或空字符串
func IsNilOrWhiteSpace(s string) bool {
	return strings.Trim(s, " ") == ""
}

//忽略大小写比较两个字符串是一致
func EqualStringIgnoreCase(a, b string) bool {
	return strings.EqualFold(a, b)
}

//判断字符串是否包含数组中任意一个元素
func StringContainsAny(s string, arr *[]string) {
	b := false
	for i := range *arr {
		b = strings.Contains(s, (*arr)[i])
		if b {
			break
		}
	}
}

//判断字符串是否包含数组中任意一个元素，忽略大小写
func StringContainsAnyIgnoreCase(s string, arr *[]string) {
	ss := strings.ToLower(s)
	b := false
	for i := range *arr {
		b = strings.Contains(ss, strings.ToLower((*arr)[i]))
		if b {
			break
		}
	}
}

func BytesToStringNoCopy(bytes *[]byte) *string {
	return (*string)(unsafe.Pointer(bytes))
}

func StringToBytesNoCopy(s *string) *[]byte {
	return (*[]byte)(unsafe.Pointer(s))
}

func ParseInt(s string) int {
	var i, err = strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func ParseInt64(s string) int64 {
	var i, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func ParseInt32(s string) int32 {
	var i, err = strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}

	return int32(i)
}

func ParseDecimal(s string) decimal.Decimal {
	d, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}

	return d
}

func ParseFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return f
}

func ParseFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}

	return float32(f)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Float32ToString(f float32) string {
	return Float64ToString(float64(f))
}

func DecimalToString(d decimal.Decimal) string {
	return d.String()
}

func ParseBool(b string) bool {
	return b != "0"
}

func BoolToString(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func IntToBool(i int32) bool {
	return i != 0
}

//按指定的开始下标与长度截取字符串
func Substring(source string, start int, length int) string {
	var r = []rune(source)
	len := len(r)
	if start+length > len {
		length = len - start
	}

	return string(r[start : start+length])
}

func ReadFile(path string) string {
	return string(ReadFileBytes(path))
}

func ReadFileBytes(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	// 要记得关闭
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	return byteValue
}

//获取程序根目录
func GetProgramRootPath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path
}

//将uri形式的路径映射为物理路径
func MapPath(uri string) string {
	path := GetProgramRootPath()
	arr := strings.Split(uri, "/")
	for s := range arr {
		if !IsNilOrWhiteSpace(arr[s]) {
			path = CombinePath(path, arr[s])
		}
	}

	return path
}

//拼接path
func CombinePath(args ...string) string {
	return strings.Join(args, string(os.PathSeparator))
}

func IsExistPath(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

//十进制左移
func DecimalLeftShift(d decimal.Decimal, decimals int) decimal.Decimal {
	return decimal.NewFromBigInt(d, 0).DivRound(decimal.NewFromInt(10).Pow(decimal.NewFromInt(decimals)), decimals)
}

//十进制右移
func DecimalRightShift(d decimal.Decimal, decimals int) decimal.Decimal {
	return d.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(decimals)))
}

//decimal转[]byte
func DecimalToBytes(d decimal.Decimal) []byte {
	return d.BigInt().Bytes()
}

//[]byte转decimal
func BytesToDecimal(bytes []byte) decimal.Decimal {
	return decimal.NewFromBigInt(new(big.Int).SetBytes(bytes), 0)
}