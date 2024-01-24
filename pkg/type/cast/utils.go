package cast

import "reflect"

func FillSameField(source interface{}, dest interface{}) {
	valSource := reflect.ValueOf(source)
	valDest := reflect.ValueOf(dest)

	if valSource.Kind() == reflect.Ptr {
		valSource = valSource.Elem()
	}
	if valDest.Kind() == reflect.Ptr {
		valDest = valDest.Elem()
	}

	for i := 0; i < valSource.NumField(); i++ {
		fieldSource := valSource.Field(i)
		fieldDest := valDest.FieldByName(valSource.Type().Field(i).Name)

		if fieldDest.IsValid() && fieldDest.CanSet() && fieldSource.Kind() == fieldDest.Kind() {
			fieldDest.Set(fieldSource)
		}
	}
}
