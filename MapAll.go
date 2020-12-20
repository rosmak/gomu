package gomu

import (
	"reflect"
)

//MapAll ...
func MapAll(obj interface{}, rootObj interface{}) error {
	//typeOfObj should be "struct" or "slice"
	typeOfObj := (reflect.ValueOf(reflect.Indirect(reflect.ValueOf(obj)).Interface()).Kind()).String()

	switch typeOfObj {
	case "struct":
		MapAllFromStruct(obj)
		break
	case "slice":
		refSliceVal := reflect.Indirect(reflect.ValueOf(obj))
		for i := 0; i < refSliceVal.Len(); i++ {
			eachElem := refSliceVal.Index(i).Addr().Interface()
			MapAllFromStruct(eachElem)
		}
		if rootObj != nil {
			a := reflect.ValueOf(rootObj).Interface()
			ArrayToMap(a, obj)
		}
	}

	return nil
}
