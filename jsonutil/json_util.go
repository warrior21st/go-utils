package jsonutil

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/warrior21st/goutils/commonutil"
)

//读取json指定key的值，嵌套key使用":"分隔，如"AppSettings:DBConnectionString"
func ReadJsonValue(jsonByte []byte, keys string) string {
	var f interface{}
	err := json.Unmarshal(jsonByte, &f)
	if err != nil {
		panic(err)
	}

	return ReadJsonValFromDecodedBytes(f, keys)
}

//从json字符串中读取指定值，嵌套key使用":"分隔，如"AppSettings:DBConnectionString"
func ReadJsonValueFromJsonString(json string, keys string) string {
	return ReadJsonValue([]byte(json), keys)
}

//从已解组json的bytes数组中读取指定值，嵌套key使用":"分隔，如"AppSettings:DBConnectionString"
func ReadJsonValFromDecodedBytes(f interface{}, keys string) string {
	val := ""
	ftemp := f
	keysArr := strings.Split(keys, ":")
	l := len(keysArr)
	for i := 0; i < l; i++ {
		m := ftemp.(map[string]interface{})
		if i < l-1 {
			ftemp = m[keysArr[i]]
		} else {
			switch m[keysArr[i]].(type) {
			case bool:
				val = strconv.FormatBool(m[keysArr[i]].(bool))
				break
			case int:
				val = strconv.FormatInt(int64(m[keysArr[i]].(int)), 10)
				break
			case int32:
				val = strconv.FormatInt(int64(m[keysArr[i]].(int32)), 10)
				break
			case int64:
				val = strconv.FormatInt(m[keysArr[i]].(int64), 10)
				break
			case float32:
				val = strconv.FormatFloat(float64(m[keysArr[i]].(float32)), 'f', -1, 64)
				break
			case float64:
				val = strconv.FormatFloat(m[keysArr[i]].(float64), 'f', -1, 64)
				break
			case string:
				val = m[keysArr[i]].(string)
				break
			default:
				jbytes, err := json.Marshal(m[keysArr[i]])
				if err != nil {
					panic(err)
				}
				return string(jbytes)
				//fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
				//break
			}
		}
	}

	return string(val)
}

func SerializeJson(val *interface{}) *string {
	bytes, err := json.Marshal(val)
	if err == nil {
		panic(err)
	}

	return commonutil.BytesToStringNoCopy(&bytes)
}

// func DesJson(jsonStr string) map[string]string {
// 	var f interface{}
// 	err := json.Unmarshal([]byte(jsonStr), &f)
// 	if err != nil {
// 		panic(err)
// 	}
// 	m := f.(map[string]interface{})

// }
