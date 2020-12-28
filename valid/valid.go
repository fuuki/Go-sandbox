package valid

import (
	"reflect"
)

// Prop is property
type Prop interface {
	IsValid() bool
}

func IsValidCheck(value interface{}) bool {
	v := reflect.ValueOf(value)
	t := v.Type()
	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			ft := t.Field(i)
			fv := v.FieldByName(ft.Name)
			if !IsValidCheck(fv.Interface()) {
				return false
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			e := v.Index(i)
			if !IsValidCheck(e.Interface()) {
				return false
			}
		}
	case reflect.Ptr:
		if v.IsNil() {
			// nil ポインタはパス
			return true
		}
		if !IsValidCheck(v.Elem().Interface()) {
			return false
		}
	default:
		// Other Kinds: Int, Bool, String, ......
		return isValidCheckElement(v)
	}
	return true
}

func isValidCheckElement(v reflect.Value) bool {
	if prop, ok := v.Interface().(Prop); ok {
		return prop.IsValid()
	}
	return true
}
