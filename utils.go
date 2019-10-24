/*
example:

func main() {
	p := &Person{
		Name: "xx",
		Age: 12,
	}
	m := StructToMapViaReflect(p)
	fmt.Println(m)
}
*/
package yuqueg

import (
	"reflect"
)

//StructToMapStr struct to map
func StructToMapStr(obj interface{}) map[string]string {
	data := make(map[string]string)

	objV := reflect.ValueOf(obj)
	v := objV.Elem()
	typeOfType := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tField := typeOfType.Field(i)
		tFieldTag := string(tField.Tag)
		if len(tFieldTag) > 0 {
			data[tFieldTag] = field.String()
		} else {
			data[tField.Name] = field.String()
		}
	}
	return data
}
