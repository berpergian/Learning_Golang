package model

import (
	"reflect"
	"strings"
)

type Document interface {
	SetDocType()
	GetDocType() string
}

type BaseDoc struct {
	DocType string `bson:"docType"`
}

func (b *BaseDoc) SetDocTypeFrom(v interface{}) {
	t := reflect.TypeOf(v)

	// if pointer, get element
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	b.DocType = strings.ToLower(t.Name())
}

func (b *BaseDoc) GetDocType() string {
	return b.DocType
}
