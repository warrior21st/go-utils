package jsonrpcutil

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
)

var rpcRequestId int64 = 0

//使用数组作为参数请求rpc api
func RequestJsonRPCApi(url string, method string, paras []string) (string, error) {
	paraStr := ""
	if paras != nil && len(paras) > 0 {
		for i := range paras {
			paraStr += "\"" + paras[i] + "\"" + ","
		}
		paraStr = paraStr[0 : len(paraStr)-1]
	}
	requestId := GetNextRpcRequestId()
	var jsonContent = "{\"jsonrpc\":\"2.0\",\"method\":\"" + method + "\",\"params\":[" + paraStr + "],\"id\":" + strconv.FormatInt(requestId, 10) + "}"
	response, err := http.Post(url, "application/json", strings.NewReader(jsonContent))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", errors.New("StatusCode:" + strconv.Itoa(response.StatusCode) + "Content:" + string(body))
	}

	var resultJson interface{}
	err = json.Unmarshal(body, &resultJson)
	if err != nil {
		return "", err
	}
	m := resultJson.(map[string]interface{})
	if int64(m["id"].(float64)) != requestId {
		return "", errors.New("response id error,request id:" + strconv.FormatInt(rpcRequestId, 10) + "," + "response id:" + strconv.FormatInt(int64(m["id"].(float64)), 10))
	}

	resb, err := json.Marshal(m["result"])
	if err != nil {
		return "", err
	}
	return string(resb), nil
}

func GetNextRpcRequestId() int64 {
	atomic.CompareAndSwapInt64(&rpcRequestId, math.MaxInt64, 0)
	return atomic.AddInt64(&rpcRequestId, 1)
}
