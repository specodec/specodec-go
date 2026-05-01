package specodec

import "strings"

type SpecCodec[T any] struct {
	Encode func(w SpecWriter, obj *T)
	Decode func(r SpecReader) *T
}

// ---------------------------------------------------------------------------
// FormatEntry: a reader/writer factory pair for one MIME type
// ---------------------------------------------------------------------------
type FormatEntry struct {
	ContentType string
	NewWriter   func() SpecWriter
	NewReader   func(body []byte) SpecReader
}

// ---------------------------------------------------------------------------
// FormatRegistry: maps content-type substrings to format entries
// ---------------------------------------------------------------------------
type FormatRegistry struct {
	entries []FormatEntry
}

func NewFormatRegistry() *FormatRegistry { return &FormatRegistry{} }

func (r *FormatRegistry) Register(e FormatEntry) *FormatRegistry {
	r.entries = append(r.entries, e)
	return r
}

func (r *FormatRegistry) Match(contentType string) FormatEntry {
	for _, e := range r.entries {
		parts := strings.SplitN(e.ContentType, "/", 2)
		if len(parts) == 2 && strings.Contains(contentType, parts[1]) {
			return e
		}
	}
	return r.entries[0] // default: first registered (JSON)
}

// ---------------------------------------------------------------------------
// Default registry: json + msgpack + gron
// ---------------------------------------------------------------------------
var DefaultRegistry = NewFormatRegistry().
	Register(FormatEntry{
		ContentType: "application/json",
		NewWriter:   func() SpecWriter { return NewJsonWriter() },
		NewReader:   func(body []byte) SpecReader { return NewJsonReader(body) },
	}).
	Register(FormatEntry{
		ContentType: "application/msgpack",
		NewWriter:   func() SpecWriter { return NewMsgPackWriter() },
		NewReader:   func(body []byte) SpecReader { return NewMsgPackReader(body) },
	}).
	Register(FormatEntry{
		ContentType: "application/gron",
		NewWriter:   func() SpecWriter { return NewGronWriter() },
		NewReader:   func(body []byte) SpecReader { return NewGronReader(body) },
	})

// ---------------------------------------------------------------------------
// Dispatch / Respond
// ---------------------------------------------------------------------------
func Dispatch[T any](codec SpecCodec[T], body []byte, contentType string) *T {
	return DispatchWith(codec, body, contentType, DefaultRegistry)
}

func DispatchWith[T any](codec SpecCodec[T], body []byte, contentType string, registry *FormatRegistry) *T {
	fmt := registry.Match(contentType)
	r := fmt.NewReader(body)
	return codec.Decode(r)
}

type RespondResult struct {
	Body        []byte
	ContentType string
}

func Respond[T any](codec SpecCodec[T], obj *T, accept string) RespondResult {
	return RespondWith(codec, obj, accept, DefaultRegistry)
}

func RespondWith[T any](codec SpecCodec[T], obj *T, accept string, registry *FormatRegistry) RespondResult {
	fmt := registry.Match(accept)
	w := fmt.NewWriter()
	codec.Encode(w, obj)
	return RespondResult{Body: w.ToBytes(), ContentType: fmt.ContentType}
}
