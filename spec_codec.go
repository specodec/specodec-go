package specodec

import "strings"

type SpecCodec[T any] struct {
	EncodeJson    func(obj *T) []byte
	EncodeMsgPack func(obj *T) []byte
	Decode        func(r SpecReader) *T
}

func Dispatch[T any](codec SpecCodec[T], body []byte, contentType string) *T {
	if strings.Contains(contentType, "msgpack") {
		return codec.Decode(NewMsgPackReader(body))
	}
	return codec.Decode(NewJsonReader(body))
}

func Respond[T any](codec SpecCodec[T], obj *T, accept string) []byte {
	if strings.Contains(accept, "msgpack") {
		return codec.EncodeMsgPack(obj)
	}
	return codec.EncodeJson(obj)
}
