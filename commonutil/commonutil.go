package commonutil

import (
	"encoding/binary"
	"encoding/hex"
	"math/big"

	// "math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/google/uuid"
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

func TryParseInt(s string) int {
	var i, _ = strconv.Atoi(s)
	return i
}

func ParseInt64(s string) int64 {
	var i, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func TryParseInt64(s string) int64 {
	var i, _ = strconv.ParseInt(s, 10, 64)
	return i
}

func ParseInt32(s string) int32 {
	var i, err = strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}

	return int32(i)
}

func TryParseInt32(s string) int32 {
	var i, _ = strconv.ParseInt(s, 10, 32)

	return int32(i)
}

func ParseDecimal(s string) decimal.Decimal {
	d, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}

	return d
}

func TryParseDecimal(s string) decimal.Decimal {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Zero
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

func TryParseFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)

	return f
}

func ParseFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}

	return float32(f)
}

func TryParseFloat32(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)

	return float32(f)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func FormatInt(i int) string {
	return strconv.Itoa(i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func FormatInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func FormatFloat64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Float32ToString(f float32) string {
	return Float64ToString(float64(f))
}

func DecimalToString(d decimal.Decimal) string {
	return d.String()
}

func FormatDecimal(d decimal.Decimal) string {
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

func FormatBool(b bool) string {
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
	len := len(source)
	if start+length > len {
		length = len - start
	}

	return string(source[start : start+length])
}

func ReadFile(path string) string {
	return string(ReadFileBytes(path))
}

func ReadFileBytes(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return b
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
		if IsNilOrWhiteSpace(arr[s]) || arr[s] == "." {
			continue
		}
		if arr[s] == ".." {
			path = filepath.Dir(path)
		} else {
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
func DecimalLeftShift(d decimal.Decimal, decimals int32) decimal.Decimal {
	return d.DivRound(decimal.NewFromInt(10).Pow(decimal.NewFromInt32(decimals)), decimals)
}

//十进制右移
func DecimalRightShift(d decimal.Decimal, decimals int32) decimal.Decimal {
	return d.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt32(decimals)))
}

//decimal转[]byte
func DecimalToBytes(d decimal.Decimal) []byte {
	return d.BigInt().Bytes()
}

//[]byte转decimal
func BytesToDecimal(bytes []byte) decimal.Decimal {
	return decimal.NewFromBigInt(new(big.Int).SetBytes(bytes), 0)
}

func Int64ToBytes(i int64) []byte {
	s1 := make([]byte, 8)
	binary.BigEndian.PutUint64(s1, uint64(i))
	return s1
}

func Int32ToBytes(i int64) []byte {
	s1 := make([]byte, 4)
	binary.BigEndian.PutUint32(s1, uint32(i))
	return s1
}

func IntToBytes(i int) []byte {
	s1 := make([]byte, 8)
	binary.BigEndian.PutUint64(s1, uint64(i))
	return s1
}

func Bytes2Hex(buffer []byte) string {
	return hex.EncodeToString(buffer)
}

func Bytes2HexWith0x(buffer []byte) string {
	return "0x" + hex.EncodeToString(buffer)
}

func Hex2Bytes(s string) []byte {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}

// has0xPrefix validates str begins with '0x' or '0X'.
func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func AppendToFile(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

//时间转字符串
func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

//时间转字符串
func TimeToDateString(t time.Time) string {
	return t.UTC().Format("2006-01-02")
}

//字符串转时间
func StringToTime(s string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.UTC)
	if err != nil {
		panic(err)
	}

	return t
}

func ParseTime(s string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.UTC)
	if err != nil {
		panic(err)
	}

	return t
}

func TryParseTime(s string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.UTC)
	if err != nil {
		return time.Unix(0, 0)
	}

	return t
}

func GenRandom(max int64) int64 {
	if max < 0 {
		max = max * -1
	}
	r := [16]byte(uuid.New())
	rBig := big.NewInt(0)
	rBig.Abs(big.NewInt(0).SetBytes(r[:]))
	rBig.Mod(rBig, big.NewInt(max+1))

	return rBig.Int64()
}

func TrimUselessZero(decimalStr string) string {
	result := decimalStr
	for strings.Contains(result, ".") && (result[len(result)-1] == '0' || result[len(result)-1] == '.') {
		result = result[:len(result)-1]
	}
	return result
}

func Concat(texts ...string) string {
	return strings.Join(texts, "")
}
