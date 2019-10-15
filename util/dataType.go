package util

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"unsafe"
)

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Json2Map(jsonStr string) map[string]interface{} {
	if jsonStr == "" {
		return nil
	}

	var mapRt map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapRt)
	if err != nil {
		logrus.Info(err)
	}
	return mapRt
}

func Map2Json(maps map[string]interface{}) string {
	if maps == nil {
		return ""
	}

	jsonBytes, err := json.Marshal(maps)
	if err != nil {
		logrus.Info(err)
	}
	return Bytes2Str(jsonBytes)
}