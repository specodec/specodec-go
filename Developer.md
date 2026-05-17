# Go Runtime — Developer Reference

> **Emitter**: `/home/ytr/Specodec/typespec-emitter-golang/src/index.ts`

---

## 1. Type Mapping Table

| TypeSpec Type | Go Type | Notes |
|---|---|---|
| `string` | `string` | |
| `boolean` | `bool` | |
| `int8` | `int8` | Explicit sized types |
| `int16` | `int16` | |
| `int32` | `int32` | |
| `int64` | `int64` | |
| `uint8` | `uint8` | |
| `uint16` | `uint16` | |
| `uint32` | `uint32` | |
| `uint64` | `uint64` | |
| `float32` | `float32` | |
| `float64`, `float`, `decimal` | `float64` | |
| `bytes` | `[]byte` | |
| `integer` | `int32` | |
| Enum | `string` (no native Go enum) | read/written as `w.WriteString()` |
| Array `<T>` | `[]T` (slice) | |
| Record `<V>` | `map[string]V` | |
| Model | `struct` with named type | |
| Union | interface + per-variant structs | |

---

## 2. Model Representation

Models are Go structs:

```go
type MyModel struct {
    Name string
    Age  int32
    Tags []string
}
```

Fields are exported (uppercase first letter). Codec functions named `WriteMyModel` and `ReadMyModel` are emitted.

---

## 3. Optional / Nullable

- Optional fields use **pointer types**: `*T` (e.g., `*string`, `*int32`).
- Required fields are value types; optional fields are pointers.
- The emitter generates `nil` checks in encode and `*T` allocation in decode.

---

## 4. Union Representation

Discriminated unions use Go **interfaces** with per-variant structs:

```go
type MyUnion interface { isMyUnion() }

type VariantA struct { Value int32 }
func (v VariantA) isMyUnion() {}

type Undefined struct{}
func (v Undefined) isMyUnion() {}
var UndefinedInst = Undefined{}
```

Union encode/dispatch uses a **type switch**:

```go
func WriteMyUnion(w SpecWriter, obj MyUnion) {
    switch v := obj.(type) {
    case VariantA:
        w.BeginObject(2)
        w.WriteField("_tag"); w.WriteString("variantA")
        w.WriteField("value"); w.WriteInt32(v.Value)
        w.EndObject()
    case Undefined:
        w.WriteNull()
    }
}
```

---

## 5. Enum Representation

Go represents enums as **`string`** — there is no native Go enum. Values are read/written via `WriteString`/`ReadEnum`. No type-safety is enforced at the Go level beyond `string`.

---

## 6. Ryu Implementation

- **Bit extraction**: `math.Float32bits(f)` and `math.Float64bits(d)`.
- **`umul128`**: Full manual 128-bit multiplication using 32-bit halves:
  - Splits `a` and `b` into `aLo`, `aHi`, `bLo`, `bHi` (each 32-bit).
  - Computes `ll = aLo * bLo`, `lh = aLo * bHi`, `hl = aHi * bLo`, `hh = aHi * bHi`.
  - Combines with carry tracking: `mid = lh + hl`, `overflow` check, `hi = hh + (mid >> 32) + overflow`, `lo = ll + ((mid & 0xFFFFFFFF) << 32)` with carry.
- **`mulShift64`**: Uses `umul128` twice, then combines hi/lo with carry and shift logic. Handles `shiftAmount >= 128`, `>= 64`, `== 0`, and intermediate cases.
- **All math uses `uint64`** natively (no BigInt), so arithmetic is true 64-bit with overflow wraparound that matches C.
- **Tables**: `[]uint64` arrays. f32: `FLOAT_POW5_INV_SPLIT` (57 entries), `FLOAT_POW5_SPLIT` (48 entries). f64: two-dimensional `[][]uint64` arrays.
- **`decimalLength9`**: Uses `uint32`; **`decimalLength17`**: Uses `uint64`. Both are unsigned comparisons.
- **`multipleOfPowerOf5_32`**: Iterative `pow5 *= 5` loop (not `math.Pow`).

---

## 7. MsgPack Reader/Writer

**Reader** (`MsgPackReader`):
- Accumulates over `[]byte` with `pos` cursor, reads via `binary.BigEndian.Uint16/Uint32/Uint64`.
- `readFloat32`: reads as float64 via `math.Float64frombits(binary.BigEndian.Uint64(...))`, then casts to `float32`.
- `readFloat64`: reads as explicit `float64`.
- `readInt64`: manually computes signed 64-bit from two unsigned 32-bit halves.
- `containerCount []int` tracks nesting for map/array iteration.
- NaN/Infinity: native float values from msgpack encoding.

**Writer** (`MsgPackWriter`):
- Accumulates via `[]byte` + `append`.
- `writeU64`: splits into two `uint32` writes via `uint32(v >> 32)` and `uint32(v)`.
- `writeFloat32`: uses `math.Float32bits(f)` → big-endian bytes.
- `writeFloat64`: uses `math.Float64bits(d)` → big-endian bytes.
- Initial buf capacity: 256 bytes (`make([]byte, 0, 256)`).

---

## 8. JSON Reader/Writer

**Reader** (`JsonReader`):
- Works on `string` (via `string(data)` from `[]byte`).
- Parses `\uXXXX` with **surrogate pair support** (same algorithm as baseline).
- `int64`/`uint64`: supports **quoted string parsing** (`json_int64` pattern) for values exceeding safe integer range.
- NaN: `math.NaN()`; Infinity: `math.Inf(1)` / `math.Inf(-1)`.
- `readFloat32`: parses as `float64` via `strconv.ParseFloat(raw, 32)` then casts to `float32`.
- Error handling: uses `panic()` with descriptive messages (recovered by caller via `defer/recover`).

**Writer** (`JsonWriter`):
- Accumulates via `strings.Builder`.
- Escape: handles standard JSON escapes + `\u00XX` for control chars `< 0x20`.
- `int64`/`uint64`: emitted as quoted strings (`strconv.FormatInt`/`FormatUint` inside quotes).
- NaN/Infinity: `"NaN"`, `"Infinity"`, `"-Infinity"` (quoted).
- Uses `formatFloat32`/`formatFloat64` (Ryu) from `float_fmt.go`.

---

## 9. Gron Reader/Writer

**Reader** (`GronReader`):
- Parses `path = value;` lines (semicolons stripped).
- Context stack: `[]gronCtx` structs with `prefix`, `typeStr`, `index` fields.
- `gronUnescape`: handles `\uXXXX` via `strconv.ParseInt(..., 16, 32)` → `rune` — **no surrogate pair support**.
- `readInt64`: unescapes (removes quotes) then `strconv.ParseInt(s, 10, 64)`.
- `readUint64`: unescapes then `strconv.ParseUint(s, 10, 64)`.
- `readFloat32`: checks for quoted `"NaN"`, `"Infinity"`, `"-Infinity"` then `strconv.ParseFloat(v, 32)`.

**Writer** (`GronWriter`):
- Accumulates `[]string` lines, path starting with `"json"`.
- `int64`/`uint64`: emitted as quoted decimal strings.
- NaN/Infinity: quoted `"NaN"`, `"Infinity"`, `"-Infinity"`.
- Uses `formatFloat32`/`formatFloat64`.

---

## 10. State Management

- **Mutable** struct-based state via pointer receivers (`func (r *JsonReader)`).
- Error handling via `panic`/`recover` pattern — callers wrap with `defer`.
- The `SpecCodec[T]` is a struct containing function fields (`Encode`, `Decode`).

---

## 11. SpecReader / SpecWriter Interfaces

### SpecReader

```go
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
```

### SpecWriter

```go
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
```

---

## 12. Emitter Generation Pattern

### Model encode
```go
func WriteMyModel(w SpecWriter, obj *MyModel) {
    w.BeginObject(2)
    w.WriteField("name")
    w.WriteString(obj.Name)
    w.WriteField("age")
    w.WriteInt32(obj.Age)
    w.EndObject()
}
```

### Model decode
```go
func ReadMyModel(r SpecReader) *MyModel {
    r.BeginObject()
    obj := &MyModel{}
    for r.HasNextField() {
        switch r.ReadFieldName() {
        case "name": obj.Name = r.ReadString()
        case "age": obj.Age = r.ReadInt32()
        default: r.Skip()
        }
    }
    r.EndObject()
    return obj
}
```

Pointer return for all models (`*MyModel`). The `SpecCodec` generic uses `func(w SpecWriter, obj *T)` and `func(r SpecReader) *T`.

---

## 13. Known Quirks / Bugs

- **Error handling**: All errors use `panic()` rather than returning error values. The test infrastructure uses `tryTest()` with `defer/recover()` to catch panics. Production callers must wrap with recover.
- **Gron unescape**: No surrogate pair handling.
- **`umul128`**: Manual implementation with overflow tracking — matches C semantics but is verbose compared to Rust's `u128` native type.
- **`SpecUndefined`** is a struct with a package-level `var Undefined` instance (not a singleton pattern with private constructor).
- Go generics (`SpecCodec[T any]`) require go 1.18+.

---

## 14. DevContainer

- **Base image**: `dev:all`
- **Go proxy**: `GOPROXY=https://goproxy.cn,direct` with `GONOSUMCHECK=*`, `GONOSUMDB=*`
- **Build**: `go mod download`, then `go vet ./... && go build ./...`
- **Output** (`FROM scratch`): copies `/app/go.mod` to `/out/go.mod`
