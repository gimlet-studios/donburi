package ecs

import (
	"fmt"
	stdreflect "reflect"
	"testing"
	"unique"

	"github.com/goccy/go-reflect"
)

func TestLayer(t *testing.T) {
	l := getLayer(0)
	ll := getLayer(1)

	if l.id != 0 {
		t.Errorf("layer id is not 0")
	}
	if ll.id != 1 {
		t.Errorf("layer id is not 1")
	}
	if l.tag == ll.tag {
		t.Errorf("layer tag is same")
	}
	if l.tag.Name() != "Layer0" {
		t.Errorf("layer tag name is not Layer0")
	}
}

func BenchmarkStdKeyForType(b *testing.B) {
	var layer Layer
	for i := 0; i < b.N; i++ {
		s := stdKeyForType(stdreflect.TypeOf(layer))
		_ = s
	}
}

func BenchmarkUniqueHandleMap(b *testing.B) {
	var layer Layer
	renderers := make(map[unique.Handle[string]][]any)
	s := keyForType(reflect.TypeOf(layer))
	renderers[s] = []any{"1", "2", "3"}

	for i := 0; i < b.N; i++ {
		key := keyForType(reflect.TypeOf(layer))
		found := renderers[key]
		_ = found
	}
}

func BenchmarkUniqueKeyForType(b *testing.B) {
	var layer Layer
	for i := 0; i < b.N; i++ {
		s := keyForType(reflect.TypeOf(layer))
		_ = s
	}
}

func stdKeyForType(typ stdreflect.Type) string {
	return fmt.Sprintf("%s/%s", typ.PkgPath(), typ.Name())
}
