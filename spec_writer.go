package specodec

type SpecWriter interface {
	WriteString(value string)
	WriteBool(value bool)
	WriteInt32(value int32)
	WriteInt64(value int64)
	WriteUint32(value uint32)
	WriteUint64(value uint64)
	WriteFloat32(value float32)
	WriteFloat64(value float64)
	WriteNull()
	WriteBytes(value []byte)
	WriteEnum(value string)
	BeginObject(fieldCount int)
	WriteField(name string)
	EndObject()
	BeginArray(elementCount int)
	NextElement()
	EndArray()
	ToBytes() []byte
}
