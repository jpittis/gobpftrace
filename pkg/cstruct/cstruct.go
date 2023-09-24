package cstruct

import (
	"fmt"
	"reflect"
	"strings"
)

// Given a Go struct, generate the code representing its associated C struct.
// Designed for use with bpftrace. Supported types are not comprehensive.
func FromGoStruct(t reflect.Type, skipProtobufPrivate bool) (string, error) {
	result := fmt.Sprintf("struct %s {\n", t.Name())
	for _, field := range reflect.VisibleFields(t) {
		if skipProtobufPrivate && strings.HasPrefix(field.Name, "XXX_") {
			continue
		}
		cfields, err := convertStructField(field)
		if err != nil {
			return "", err
		}
		for _, cfield := range cfields {
			result += fmt.Sprintf("    %s %s;\n", cfield.t, cfield.name)
		}
	}
	return result + "}", nil
}

type cStructField struct {
	t    string
	name string
}

func convertStructField(field reflect.StructField) ([]cStructField, error) {
	switch field.Type.Kind() {
	case reflect.Slice:
		ctype, err := convertType(field.Type.Elem())
		if err != nil {
			return nil, err
		}
		return []cStructField{
			{fmt.Sprintf("%s*", ctype), field.Name},
			{"uint64_t", fmt.Sprintf("%sLen", field.Name)},
			{"uint64_t", fmt.Sprintf("%sCap", field.Name)},
		}, nil
	default:
		ctype, err := convertType(field.Type)
		if err != nil {
			return nil, err
		}
		return []cStructField{{ctype, field.Name}}, nil
	}
}

func convertType(t reflect.Type) (string, error) {
	switch t.Kind() {
	case reflect.Int64:
		return "int64_t", nil
	case reflect.Int32:
		return "int32_t", nil
	case reflect.Bool:
		return "uint8_t", nil
	case reflect.Uint8:
		return "uint8_t", nil
	default:
		// Support for more types can be added above.
		return "", fmt.Errorf("Unknown type kind %v", t.Kind())
	}
}
