package sqlutils
import "errors"
import "reflect"

// struct pointer 
func CheckRequired(val interface{}) (error) {
	t       := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()

	for i := 0; i < t.NumField(); i++ {
		var tag        reflect.StructTag = typeOfT.Field(i).Tag
		var field      reflect.Value = t.Field(i)
		var fieldType  reflect.Type = field.Type()
		var attributes map[string]bool = GetColumnAttributesFromTag(&tag)
		if _ , ok := attributes["required"] ; ok {
			// check the column value
			if fieldType.String() == "string" && field.Interface().(string) == "" {
				return errors.New( fieldType.Name() + " field required.")
			} else if fieldType.String() == "int" && field.Interface().(int) == 0 {
				return errors.New( fieldType.Name() + " field required.")
			}
		}
	}
	return nil
}
