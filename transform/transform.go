package transform

const (
	// HTML represents html content.
	HTML byte = iota + 1
	// CSS represents css content.
	CSS
)

// Transformer interface for all transformers (css, html)
type Transformer interface {
	Transform([]byte) []byte
}

type HTMLTransformer struct{}

func (t HTMLTransformer) Transform(html []byte) []byte {
	return []byte("")
}

type CSSTransformer struct{}

func (t CSSTransformer) Transform(html []byte) []byte {
	return []byte("")
}

// AvailableTransformrs returns a list of all required transformrs
// HTML requires (HTMLTransformer, CSSTransformer, JavaScriptTransformer)
// CSS requires CSSTransformer
// JavaScript requires JavaScriptTransformer
func AvailableTransformrs(contentType byte) []Transformer {
	var transformrs []Transformer
	return transformrs
}

// Transform applies available transformers on content
// returns transformed content.
func Transform(content []byte, contentType byte) ([]byte, error) {
	transformers := AvailableTransformrs(contentType)

	for _, transformer := range transformers {
		content = transformer.Transform(content)
	}
	return content, nil
}
