package specodec

type SpecReader interface {
	BeginObject()
	HasNextField() bool
	ReadFieldName() string
	EndObject()
	BeginArray()
	HasNextElement() bool
	EndArray()
	ReadString() string
	ReadBool() bool
	ReadInt32() int32
	ReadInt64() int64
	ReadUint32() uint32
	ReadUint64() uint64
	ReadFloat32() float32
	ReadFloat64() float64
	ReadNull()
	ReadBytes() []byte
	ReadEnum() string
	IsNull() bool
	Skip()
}
