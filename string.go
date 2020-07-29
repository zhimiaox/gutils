package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/axgle/mahonia"
)

// MustUtf8 强制字符串必须为utf8
func MustUtf8(s string) string {
	if !utf8.ValidString(s) {
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		s = string(v)
	}
	return s
}

// ConvertToString 转换字符串编码
// src 字符串 srccode 源编码 tagCode 目标编码
func ConvertToString(src, srcCode, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// ParseArrString 解析字符串数组
// str 待解析字符串 sep 分割符号 dto 期望解析切片(指针)
func ParseArrString(str, sep string, dto interface{}) error {
	arr := strings.Split(str, sep)
	typ := reflect.TypeOf(dto)
	if typ.Kind() != reflect.Ptr {
		return fmt.Errorf("dto is not ptr")
	}
	etyp := typ.Elem()
	if etyp.Kind() != reflect.Slice {
		return fmt.Errorf("dto type is not slice")
	}
	eetyp := etyp.Elem()
	rval := make([]reflect.Value, 0)
	switch eetyp.Kind() {
	case reflect.Bool:
		for _, s := range arr {
			if v, err := strconv.ParseBool(s); err == nil {
				rval = append(rval, reflect.ValueOf(v))
			}
		}
	case reflect.Int:
		for _, s := range arr {
			if v, err := strconv.Atoi(s); err == nil {
				rval = append(rval, reflect.ValueOf(v))
			}
		}
	case reflect.Int8:
		for _, s := range arr {
			if v, err := strconv.ParseInt(s, 0, 8); err == nil {
				rval = append(rval, reflect.ValueOf(int8(v)))
			}
		}
	case reflect.Int16:
		for _, s := range arr {
			if v, err := strconv.ParseInt(s, 0, 16); err == nil {
				rval = append(rval, reflect.ValueOf(int16(v)))
			}
		}
	case reflect.Int32:
		for _, s := range arr {
			if v, err := strconv.ParseInt(s, 0, 32); err == nil {
				rval = append(rval, reflect.ValueOf(int32(v)))
			}
		}
	case reflect.Int64:
		for _, s := range arr {
			if v, err := strconv.ParseInt(s, 0, 64); err == nil {
				rval = append(rval, reflect.ValueOf(v))
			}
		}
	case reflect.Uint:
		for _, s := range arr {
			if v, err := strconv.ParseUint(s, 0, 0); err == nil {
				rval = append(rval, reflect.ValueOf(uint(v)))
			}
		}
	case reflect.Uint8:
		for _, s := range arr {
			if v, err := strconv.ParseUint(s, 0, 8); err == nil {
				rval = append(rval, reflect.ValueOf(uint8(v)))
			}
		}
	case reflect.Uint16:
		for _, s := range arr {
			if v, err := strconv.ParseUint(s, 0, 16); err == nil {
				rval = append(rval, reflect.ValueOf(uint16(v)))
			}
		}
	case reflect.Uint32:
		for _, s := range arr {
			if v, err := strconv.ParseUint(s, 0, 32); err == nil {
				rval = append(rval, reflect.ValueOf(uint32(v)))
			}
		}
	case reflect.Uint64:
		for _, s := range arr {
			if v, err := strconv.ParseUint(s, 0, 64); err == nil {
				rval = append(rval, reflect.ValueOf(v))
			}
		}
	case reflect.Float64:
		for _, s := range arr {
			if v, err := strconv.ParseFloat(s, 64); err == nil {
				rval = append(rval, reflect.ValueOf(v))
			}
		}
	case reflect.Float32:
		for _, s := range arr {
			if v, err := strconv.ParseFloat(s, 32); err == nil {
				rval = append(rval, reflect.ValueOf(float32(v)))
			}
		}
	case reflect.String:
		for _, s := range arr {
			if s != "" {
				rval = append(rval, reflect.ValueOf(s))
			}
		}
	default:
		return fmt.Errorf("dto type Unknown")
	}
	val := reflect.ValueOf(dto).Elem()
	if val.CanSet() {
		val.Set(reflect.Append(val, rval...))
	}
	return nil
}