package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"strconv"

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

func TypeListToUnmarshalStringList(typeList basetypes.ListValue) (stringList []string) {
	for _, t := range typeList.Elements() {
		var n string
		json.Unmarshal([]byte(t.String()), &n)
		stringList = append(stringList, n)
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

func StringListToTypeInt64List(stringList []string) (typeList basetypes.ListValue) {
	var valueList []attr.Value
	for _, s := range stringList {
		sInt, _ := strconv.ParseInt(s, 10, 64)
		valueList = append(valueList, basetypes.NewInt64Value(sInt))
	}
	val, _ := basetypes.NewListValue(types.Int64Type, valueList)
	return val
}

func ToStringList(in []interface{}) []string {
	out := make([]string, 0, len(in))
	for _, val := range in {
		out = append(out, val.(string))
	}
	return out
}

func PrefixedUniqueId(prefix string) string {
	b := make([]byte, 16)
	rand.Read(b)
	return prefix + hex.EncodeToString(b)
}
