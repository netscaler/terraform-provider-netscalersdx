package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TypeListToStringList(typeList basetypes.ListValue) (stringList []string) {
	for _, t := range typeList.Elements() {
		stringList = append(stringList, t.String())
	}
	return
}

func StringListToTypeList(stringList []string) (typeList basetypes.ListValue) {
	var valueList []attr.Value
	for _, s := range stringList {
		valueList = append(valueList, basetypes.NewStringValue(s))
	}
	val, _ := basetypes.NewListValue(types.StringType, valueList)
	return val
}

func ToStringList(in []interface{}) []string {
	out := make([]string, 0, len(in))
	for _, val := range in {
		out = append(out, val.(string))
	}
	return out
}
