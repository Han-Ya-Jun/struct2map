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

// struct2map
func Struct2map(st interface{}) (map[string]interface{}, error) {
	stType := reflect.TypeOf(st)
	if stType.Kind() != reflect.Ptr {
		return nil, ErrNotPtr
	}
	eleType := stType.Elem()
	if eleType.Kind() != reflect.Struct {
		return nil, ErrNotValidElem
	}
	stVal := reflect.Indirect(reflect.ValueOf(st))
	m := make(map[string]interface{})
	for i := 0; i < eleType.NumField(); i++ {
		tagStr, ok := eleType.Field(i).Tag.Lookup(TagName)
		if !ok {
			return nil, ErrNeedTag
		}
		key, err := getKey(tagStr)
		if err == ErrNotValidKey || err == ErrNotValidTag {
			return nil, err
		}
		if err == ErrIgnore {
			continue
		}
		if err == ErrOmitempty && (IsNil(stVal.Field(i).Interface()) || stVal.Field(i).Interface() == "") {
			continue
		}
		m[key] = stVal.Field(i).Interface()
	}
	return m, nil
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
