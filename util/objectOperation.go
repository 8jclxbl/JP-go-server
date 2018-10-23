package util

import "reflect"

func GetFieldName(instance interface{}) []string {
	ref := reflect.TypeOf(instance)
	count := ref.NumField()
	result := make([]string,0,count)
	for i:=0;i<count;i++ {
		result = append(result,ref.Field(i).Name)
	}
	return result
}
