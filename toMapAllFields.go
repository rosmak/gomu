package gomu

import (
	"reflect"
)

//MapAllFromStruct ...
func MapAllFromStruct(elem interface{}) error {
	ptrElem := reflect.ValueOf(elem)
	valElem := ptrElem.Elem()
	typeOfS := valElem.Type()

	for i := 0; i < valElem.NumField(); i++ {
		mapF := typeOfS.Field(i).Name
		arrF := typeOfS.Field(i).Tag.Get("gomu")
		if arrF != "" {
			mapVal := valElem.FieldByName(mapF).Addr().Interface()
			arrVal := valElem.FieldByName(arrF).Addr().Interface()
			arrFromField := valElem.FieldByName(arrF)
			for j := 0; j < arrFromField.Len(); j++ {
				eachElem := arrFromField.Index(j).Addr().Interface()
				MapAllFromStruct(eachElem)
			}

			ToMap(mapVal, arrVal)

		}
	}

	return nil
}
