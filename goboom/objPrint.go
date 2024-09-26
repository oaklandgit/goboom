package goboom

import (
	"fmt"
	"reflect"
)

func (obj *GameObject) Print() {
    v := reflect.ValueOf(obj).Elem()
    t := v.Type()

    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i).Interface()
        fmt.Printf("%s: %v\n", field.Name, value)
    }
}