package lazy

import (
	"reflect"
)

//Pointer Or Value TO Value (default value of the indirect type in case of nil value)
func PoVToV(pov interface{}) (v interface{}) {
	if reflect.TypeOf(pov).Kind() != reflect.Ptr {
		//It's not a pointer
		return pov
	}

	if !reflect.ValueOf(pov).IsNil() {
		v = reflect.ValueOf(pov).Elem().Interface()
	} else {
		//The pointer is nil, return the default value of the indirect type
		v = reflect.Zero(reflect.TypeOf(pov).Elem()).Interface()
	}

	return v
}
