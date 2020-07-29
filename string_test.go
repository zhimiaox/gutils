package utils

import "testing"


func TestParseArrString(t *testing.T) {
	tint := make([]int, 0)
	err := ParseArrString("1,,4,o,5,8888,-2", ",", &tint)
	if err != nil || len(tint) != 5 {
		t.Fail()
	}
	tint8 := make([]int8, 0)
	err = ParseArrString("1,,4,o,5,8888,-2", ",", &tint8)
	if err != nil || len(tint8) != 4 {
		t.Fail()
	}
	tint16 := make([]int16, 0)
	err = ParseArrString("1,,4,o,5,8888,-2", ",", &tint16)
	if err != nil || len(tint16) != 5 {
		t.Fail()
	}
	tint32 := make([]int32, 0)
	err = ParseArrString("1,,4,o,5,8888,-2", ",", &tint32)
	if err != nil || len(tint32) != 5 {
		t.Fail()
	}
	tint64 := make([]int64, 0)
	err = ParseArrString("1,,4,o,5,8888,-2", ",", &tint64)
	if err != nil || len(tint64) != 5 {
		t.Fail()
	}
	tuint := make([]uint, 0)
	err = ParseArrString("1,,4,o,5,8888,-2", ",", &tuint)
	if err != nil || len(tuint) != 4 {
		t.Fail()
	}
	tuint8 := make([]uint8, 0)
	err = ParseArrString("1,,4,o,5,8888,-2", ",", &tuint8)
	if err != nil || len(tuint8) != 3 {
		t.Fail()
	}
	tuint16 := make([]uint16, 0)
	err = ParseArrString("1,,4,o,5,8888,-2", ",", &tuint16)
	if err != nil || len(tuint16) != 4 {
		t.Fail()
	}
	tuint32 := make([]uint32, 0)
	err = ParseArrString("1,,4,o,5,8888,-2", ",", &tuint32)
	if err != nil || len(tuint32) != 4 {
		t.Fail()
	}
	tuint64 := make([]uint64, 0)
	err = ParseArrString("1,,4,o,5,8888,-2", ",", &tuint64)
	if err != nil || len(tuint64) != 4 {
		t.Fail()
	}
	tfloat32 := make([]float32, 0)
	err = ParseArrString("1,,4,o,5,88.88,-2.9", ",", &tfloat32)
	if err != nil || len(tfloat32) != 5 {
		t.Fail()
	}
	tfloat64 := make([]float64, 0)
	err = ParseArrString("1,,4,o,5,88.88,-2.9", ",", &tfloat64)
	if err != nil || len(tfloat64) != 5 {
		t.Fail()
	}
	tstring := make([]string, 0)
	err = ParseArrString("1,,4,o,5,88.88,-2.9", ",", &tstring)
	if err != nil || len(tstring) != 6 {
		t.Fail()
	}
}

