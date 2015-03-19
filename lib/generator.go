package main

import (
	"fmt"
	"os"
	"reflect"
	"text/template"
	"time"
)

type ReflectedField struct {
	Name, Type string
	Default    interface{}
}

type HeaderData struct {
	RequestAliases, ResponseAliases map[FrameId]FrameAlias
	Version                         string
	TimeStr                         string
	HelloFields                     []ReflectedField
}

var generatedTypes = map[string]string{
	"int":    "int32_t",
	"string": "char *",
	"uint8":  "uint8_t",
}

func reflectStruct(f interface{}) []ReflectedField {
	val := reflect.ValueOf(f)
	if val.Kind() != reflect.Struct {
		panic("reflectStruct can only work on structs!")
	}
	fields := make([]ReflectedField, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag
		valueType := valueField.Type().Name()
		generatedName := ""
		generatedType, ok := generatedTypes[valueType]
		if !ok {
			panic(fmt.Errorf("Not sure how to convert %s Golang type to C. Teach me!", valueType))
		}

		if tag.Get("codec") == "" {
			generatedName = typeField.Name
		} else {
			generatedName = tag.Get("codec")
		}
		fields = append(fields, ReflectedField{
			Name:    generatedName,
			Type:    generatedType,
			Default: valueField.Interface(),
		})
	}
	return fields
}

func main() {
	var ts HelloResponseFrame
	generatedAt, _ := time.Now().MarshalText()
	tmpl, err := template.ParseGlob("*.tmpl")
	var headerData = HeaderData{
		RequestAliases:  FrameRequestAliases,
		ResponseAliases: FrameResponseAliases,
		Version:         Version,
		TimeStr:         string(generatedAt),
		HelloFields: reflectStruct(ts),
	}
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "framelib.h.tmpl", headerData)
	if err != nil {
		panic(err)
	}
}
