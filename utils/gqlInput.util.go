package utils

import (
	"reflect"
)

func ConvertUpdateInputToDBColumnNames(m, g interface{}) []string {
	var columns []string
	mrt := reflect.TypeOf(m)              // sqlboilerのリフレクション取得
	grt := reflect.TypeOf(g)              // gqlgenのリフレクション取得
	for i := 0; i < mrt.NumField(); i++ { // sqlboilerのfieldをループ
		mf := mrt.Field(i)
		if mf.Name == "UpdatedAt" { // updated_atがあれば更新対象に含める
			columns = append(columns, "updated_at")
		}
		for i := 0; i < grt.NumField(); i++ { // gqlgenのfieldをループ
			gf := grt.Field(i)
			if mf.Name == gf.Name { // gqlgenのfieldにあるカラムのみ更新対象に含める
				columns = append(columns, mf.Tag.Get(`boil`))
				break
			}
		}
	}
	return columns
}