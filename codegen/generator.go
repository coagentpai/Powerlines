package main

import (
	"os"
	"reflect"
	"text/template"
	"time"
	"encoding/hex"
	"crypto/sha256"
	"github.com/karlbloedorn/Powerlines/protocol"
	"fmt"
)

type Direction uint8

const (
	Response Direction = iota
	Request
)

type ReflectedField struct {
	Name, Type string
}

type ReflectedStruct struct {
	Name string
	Id protocol.ContainerId
	Fields []ReflectedField
	FrameDirection Direction
}


type HeaderData struct {
	RequestAliases, ResponseAliases map[protocol.ContainerId]protocol.ContainerAlias
	Version                         string
	TimeStr                         string
	ReflectedStructs                []ReflectedStruct
}

var generatedTypes = map[string]string{
	"string": "std::string",
	"int32":    "int32_t",
	"int16":    "int16_t",
	"int8":    "int8_t",
	"uint8":  "uint8_t",
	"uint16": "uint16_t",
	"uint32": "uint32_t",
	"slice":  "std::vec",
	"map": "std::map",
}

func (rs *ReflectedStruct) IsResponse() bool {
	return rs.FrameDirection == Response
}


func (rs *ReflectedStruct) IsRequest() bool {
	return rs.FrameDirection == Request
}


func mapKind(lookup reflect.Kind) string {
	protocolType, ok := generatedTypes[lookup.String()]
	if !ok {
		panic(fmt.Errorf("No protocol type given for %s Golang type", lookup.String()))
	}
	return protocolType
}

func findProtocolTypeContainer(field reflect.Type, kindChan *chan reflect.Kind) string {
	elem := field.Elem()
	switch field.Kind() {
	case reflect.Slice, reflect.Array:
		return fmt.Sprintf(
			"%s<%s>",
			mapKind(reflect.Slice),
			findProtocolType(elem, kindChan),
		)
	case reflect.Map:
		return fmt.Sprintf(
			"%s<%s,%s>",
			mapKind(reflect.Map),
			findProtocolType(field.Key(), kindChan),
			findProtocolType(elem, kindChan),
		)
	}
	return ""
}

func findProtocolType(field reflect.Type, kindChan *chan reflect.Kind) string {
	*kindChan <- field.Kind()
	switch field.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array:
		return findProtocolTypeContainer(field, kindChan)
	default:
		return mapKind(field.Kind())
	}
}

func reflectStruct(f interface{},  kindChan *chan reflect.Kind) ([]ReflectedField) {
	structType := reflect.TypeOf(f)
	if structType.Kind() != reflect.Struct {
		panic("reflectStruct can only work on structs!")
	}
	fields := make([]ReflectedField, 0, structType.NumField())

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag
		generatedName := ""
		generatedType := findProtocolType(field.Type, kindChan)

		if tag.Get("frame") == "" {
			generatedName = field.Name
		} else {
			generatedName = tag.Get("frame")
		}
		fields = append(fields, ReflectedField{
			Name:    generatedName,
			Type:    generatedType,
		})
	}
	return fields
}

func main() {
	kindChan := make(chan reflect.Kind, 10)
	protocalSig := make(chan string)
	go func() {
		sha256hash := sha256.New()
		for {
			kind, more := <-kindChan
			if more {
				fmt.Fprintln(sha256hash, kind)
			} else {
				var hashValue []byte
				protocalSig <- hex.EncodeToString(sha256hash.Sum(hashValue))
				return
			}
		}
	}()
	generatedAt, _ := time.Now().MarshalText()
	reflectedStructs := make([]ReflectedStruct, 0, 10)
	for _,id := range protocol.AllContainerIds {
		if container, ok := protocol.ContainerStructMap[id]; ok {
			fields := reflectStruct(container, &kindChan)
			var structName string
			var direction Direction
			if alias, ok := protocol.ContainerRequestAliases[id]; ok {
				structName = alias.Short
				direction = Request
			}
			if alias, ok := protocol.ContainerResponseAliases[id]; ok {
				structName = alias.Short
				direction = Response
			}
			if structName == "" {
				panic(fmt.Errorf("No name set for struct: %d", id ))
			}
			reflectedStructs = append(reflectedStructs, ReflectedStruct{
				Fields: fields,
				Name: structName,
				Id: id,
				FrameDirection: direction,
			})
		}
	}
	tmpl, err := template.ParseGlob("*.tmpl")
	var headerData = HeaderData{
		RequestAliases:  protocol.ContainerRequestAliases,
		ResponseAliases: protocol.ContainerResponseAliases,
		TimeStr:         string(generatedAt),
	    ReflectedStructs: reflectedStructs,
	}
	if err != nil {
		panic(err)
	}
	close(kindChan)
	headerData.Version = <-protocalSig
	err = tmpl.ExecuteTemplate(os.Stdout, "containerlib.cpp.tmpl", headerData)
	if err != nil {
		panic(err)
	}
}
