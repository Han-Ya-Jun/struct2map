package struct2map

import (
	"errors"
	"reflect"
	"strings"
)

/*
* @Author: yajun.han
* @Date: 2020/7/12 1:59 上午
* @Name：struct2map
* @Description:
 */


var (
	ErrNotPtr       = errors.New("need a pointer")
	ErrNotValidElem = errors.New("pointer not point to struct")
	ErrNotValidTag  = errors.New("not valid tag")
	ErrNotValidKey  = errors.New("not valid key")
	ErrIgnore       = errors.New("ignore key")
	ErrOmitempty    = errors.New("omitempty key")
	ErrNeedTag      = errors.New("need struct2map tag")
	TagName         = "struct2map"
	TagIgnore       = "-"
	TagOmitempty    = "omitempty"
)

// get key
func getKey(tagStr string) (key string, err error) {
	for _, tag := range strings.Split(tagStr, ";") {
		tagList := strings.Split(tag, ",")
		if tag == TagIgnore {
			err = ErrIgnore
		}
		if len(tagList) == 2 && tagList[1] == TagOmitempty {
			err = ErrOmitempty
			key = tagList[0]
		} else {
			key = tagList[0]
		}
	}
	return key, err
}

// struct2map 支持嵌套
func Struct2map(st interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if IsNil(st) || st == nil {
		return m, nil
	}
	stType := reflect.TypeOf(st)
	if stType.Kind() != reflect.Ptr {
		return nil, ErrNotPtr
	}
	eleType := stType.Elem()
	if eleType.Kind() != reflect.Struct {
		return nil, ErrNotValidElem
	}
	stVal := reflect.Indirect(reflect.ValueOf(st))
	FillMap(m, stVal, eleType)
	return m, nil
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func FillMap(m map[string]interface{}, val interface{}, eleType reflect.Type) {
	for i := 0; i < eleType.NumField(); i++ {
		if eleType.Field(i).Type.Kind() == reflect.Struct {
			FillMap(m, val.(reflect.Value).Field(i), eleType.Field(i).Type)
			continue
		}
		tagStr, ok := eleType.Field(i).Tag.Lookup(TagName)
		if !ok {
			continue
		}
		key, err := getKey(tagStr)
		if err == ErrNotValidKey || err == ErrNotValidTag {
			continue
		}
		if err == ErrIgnore {
			continue
		}
		value := val.(reflect.Value).Field(i).Interface()
		if err == ErrOmitempty && (IsNil(value) || value == "") {
			continue
		}
		m[key] = value
	}
}
