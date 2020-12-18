package gomu

import (
	"reflect"
)

//ToMapAllFields ...
func ToMapAllFields(elem interface{}) error {
	ptrElem := reflect.ValueOf(elem)
	valElem := ptrElem.Elem()
	typeOfS := valElem.Type()

	for i := 0; i < valElem.NumField(); i++ {
		mapF := typeOfS.Field(i).Name
		arrF := typeOfS.Field(i).Tag.Get("gomu")
		if arrF != "" {
			mapVal := valElem.FieldByName(mapF).Addr().Interface()
			arrVal := valElem.FieldByName(arrF).Addr().Interface()

			ToMap(mapVal, arrVal)

		}
	}

	return nil
}
