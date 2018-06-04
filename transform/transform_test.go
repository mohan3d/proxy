package transform

import (
	"testing"
)

func transformerExists(e Transformer, lst []Transformer) bool {
	for _, a := range lst {
		if e == a {
			return true
		}
	}
	return false
}

func TestAvailableTransformrsHTML(t *testing.T) {
	expected := []Transformer{
		HTMLTransformer{},
		CSSTransformer{},
	}
	actual := AvailableTransformrs(HTML)

	for _, e := range expected {
		if !transformerExists(e, actual) {
			t.Errorf("expected %T exists, not found", e)
		}
	}
}

func TestAvailableTransformrsCSS(t *testing.T) {
	expected := []Transformer{
		CSSTransformer{},
	}
	actual := AvailableTransformrs(CSS)

	for _, e := range expected {
		if !transformerExists(e, actual) {
			t.Errorf("expected %T exists, not found", e)
		}
	}
}
