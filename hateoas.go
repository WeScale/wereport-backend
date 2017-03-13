package main

import (
	"fmt"
	"reflect"
)

func MarshalHateoas(subject interface{}) map[string]interface{} {
	links := make(map[string]string)
	out := make(map[string]interface{})
	subjectValue := reflect.Indirect(reflect.ValueOf(subject))
	subjectType := subjectValue.Type()
	for i := 0; i < subjectType.NumField(); i++ {
		field := subjectType.Field(i)
		name := subjectType.Field(i).Name
		out[field.Tag.Get("json")] = subjectValue.FieldByName(name).Interface()
	}
	switch s := subject.(type) {
	case Contrat:
		links["self"] = fmt.Sprintf("/contrats/%s", s.ID.String())
	}
	out["_links"] = links
	return out
}
