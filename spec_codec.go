package specodec

import "strings"

type SpecCodec[T any] struct {
	Encode func(w SpecWriter, obj *T)
	Decode func(r SpecReader) *T
}

// ---------------------------------------------------------------------------
// FormatEntry: a reader/writer factory pair for one format
// ---------------------------------------------------------------------------
type FormatEntry struct {
	Name      string // e.g. "json", "msgpack", "gron"
	NewWriter func() SpecWriter
	NewReader func(body []byte) SpecReader
}

// ---------------------------------------------------------------------------
// FormatRegistry: maps format name substrings to format entries
// ---------------------------------------------------------------------------
type FormatRegistry struct {
	entries []FormatEntry
}

func NewFormatRegistry() *FormatRegistry { return &FormatRegistry{} }

func (r *FormatRegistry) Register(e FormatEntry) *FormatRegistry {
	r.entries = append(r.entries, e)
	return r
}

func (r *FormatRegistry) Match(format string) FormatEntry {
	for _, e := range r.entries {
		if strings.Contains(format, e.Name) {
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
		Name:      "json",
		NewWriter: func() SpecWriter { return NewJsonWriter() },
		NewReader: func(body []byte) SpecReader { return NewJsonReader(body) },
	}).
	Register(FormatEntry{
		Name:      "msgpack",
		NewWriter: func() SpecWriter { return NewMsgPackWriter() },
		NewReader: func(body []byte) SpecReader { return NewMsgPackReader(body) },
	}).
	Register(FormatEntry{
		Name:      "gron",
		NewWriter: func() SpecWriter { return NewGronWriter() },
		NewReader: func(body []byte) SpecReader { return NewGronReader(body) },
	})

// ---------------------------------------------------------------------------
// Dispatch / Respond
// ---------------------------------------------------------------------------
func Dispatch[T any](codec SpecCodec[T], body []byte, format string) *T {
	return DispatchWith(codec, body, format, DefaultRegistry)
}

func DispatchWith[T any](codec SpecCodec[T], body []byte, format string, registry *FormatRegistry) *T {
	fmt := registry.Match(format)
	r := fmt.NewReader(body)
	return codec.Decode(r)
}

type RespondResult struct {
	Body []byte
	Name string // format name: "json" | "msgpack" | "gron"
}

func Respond[T any](codec SpecCodec[T], obj *T, format string) RespondResult {
	return RespondWith(codec, obj, format, DefaultRegistry)
}

func RespondWith[T any](codec SpecCodec[T], obj *T, format string, registry *FormatRegistry) RespondResult {
	fmt := registry.Match(format)
	w := fmt.NewWriter()
	codec.Encode(w, obj)
	return RespondResult{Body: w.ToBytes(), Name: fmt.Name}
}
