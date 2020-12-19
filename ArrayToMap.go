package gomu

import (
	"fmt"
	"reflect"
)

func reflectString() reflect.Type {
	return reflect.TypeOf("r")
}

//ArrayToMap ...
func ArrayToMap(mf interface{}, sl interface{}) error {
	ptrMapField := reflect.ValueOf(mf)
	valElem := ptrMapField.Elem()

	ptrSlice := reflect.ValueOf(sl)
	valSlice := ptrSlice.Elem()

	if valSlice.Len() < 1 {
		return nil
	}

	reflectEntityType := reflect.TypeOf(valSlice.Index(0).Interface())

	virtualMapType := reflect.MapOf(reflectString(), reflectEntityType)
	virtualMap := reflect.MakeMapWithSize(virtualMapType, 0)

	for i := 0; i < valSlice.Len(); i++ {
		each := valSlice.Index(i)
		eachValue := reflect.ValueOf(each.Interface())
		uidString := each.FieldByName("ID").Interface().(fmt.Stringer).String()
		virtualMap.SetMapIndex(reflect.ValueOf(uidString), eachValue.Convert(reflectEntityType))
	}

	valSlice.Set(reflect.Zero(reflect.ValueOf(sl).Elem().Type()))
	valElem.Set(virtualMap)
	return nil
}
