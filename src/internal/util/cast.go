package util

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Slice(slice interface{}, newSliceType reflect.Type) interface{} {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Slice called with non-slice value of type %T", slice))
	}
	if newSliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Slice called with non-slice type of type %T", newSliceType))
	}
	newSlice := reflect.New(newSliceType)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hdr.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
}
