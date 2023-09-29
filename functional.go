package functional

import "reflect"

func IsNothing(x any) bool {
	val := reflect.ValueOf(x)
	if KindOf(val) == reflect.Ptr {
		return val.IsNil()
	}
	return false
}

func IsPtr(x any) bool {
	return KindOf(x) == reflect.Ptr
}

func PtrOf[T any](x T) *T {
	return &x
}

func TypeOf(x any) reflect.Type {
	return reflect.TypeOf(x)
}

func KindOf(obj any) reflect.Kind {
	return reflect.ValueOf(obj).Kind()
}
