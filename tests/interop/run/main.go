package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/specodec/specodec-go"
)

var vecDir, outDir string
var schema map[string]json.RawMessage
var schemaMap map[string]*ModelSchema

type FieldSchema struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Optional bool   `json:"optional,omitempty"`
	IsArray  bool   `json:"isArray,omitempty"`
	IsModel  bool   `json:"isModel,omitempty"`
}

type ModelSchema struct {
	Fields    []FieldSchema `json:"fields"`
	Recursive bool          `json:"recursive,omitempty"`
}

func main() {
	vecDir = envOr("VEC_DIR", "../vectors")
	outDir = envOr("OUT_DIR", "../output_go")
	os.MkdirAll(filepath.Join(outDir, "scalars"), 0755)

	manifestData, _ := os.ReadFile(filepath.Join(vecDir, "manifest.json"))
	var manifest map[string]json.RawMessage
	json.Unmarshal(manifestData, &manifest)

	schemaData, _ := os.ReadFile(filepath.Join(vecDir, "typeschema.json"))
	json.Unmarshal(schemaData, &schema)

	schemaMap = make(map[string]*ModelSchema)
	for name, raw := range schema {
		var ms ModelSchema
		json.Unmarshal(raw, &ms)
		schemaMap[name] = &ms
	}

	var scalars map[string]json.RawMessage
	json.Unmarshal(manifest["scalars"], &scalars)

	var scalarResults []map[string]interface{}
	var objectResults []map[string]interface{}

	fmt.Println("Go: processing scalars...")
	for name, raw := range scalars {
		var spec map[string]interface{}
		json.Unmarshal(raw, &spec)
		stype := spec["type"].(string)
		ref, _ := os.ReadFile(filepath.Join(vecDir, "scalars", name+".mp"))
		w := &specodec.MsgPackWriter{}
		r := specodec.NewMsgPackReader(ref)

		var err error
		switch stype {
		case "int32":
			v := r.ReadInt32(); w.WriteInt32(v)
		case "int64":
			v := r.ReadInt64(); w.WriteInt64(v)
		case "uint32":
			v := r.ReadUint32(); w.WriteUint32(v)
		case "uint64":
			v := r.ReadUint64(); w.WriteUint64(v)
		case "float32":
			v := r.ReadFloat32(); w.WriteFloat32(v)
		case "float64":
			v := r.ReadFloat64(); w.WriteFloat64(v)
		case "string":
			v := r.ReadString(); w.WriteString(v)
		case "bytes":
			v := r.ReadBytes(); w.WriteBytes(v)
		case "bool":
			v := r.ReadBool(); w.WriteBool(v)
		default:
			err = fmt.Errorf("unknown type: %s", stype)
		}

		if err != nil {
			scalarResults = append(scalarResults, map[string]interface{}{"name": name, "pass": false})
		} else {
			os.WriteFile(filepath.Join(outDir, "scalars", name+".mp"), w.ToBytes(), 0644)
			scalarResults = append(scalarResults, map[string]interface{}{"name": name, "pass": true})
		}
	}

	fmt.Println("Go: processing objects...")
	var testModels []string
	json.Unmarshal(manifest["testModels"], &testModels)

	for _, name := range testModels {
		mpOk, jsonOk, gronOk := processObject(name)
		objectResults = append(objectResults, map[string]interface{}{"name": name, "mp": mpOk, "json": jsonOk, "gron": gronOk})
	}

	scalarMap := map[string]map[string]bool{}
	for _, r := range scalarResults {
		scalarMap[r["name"].(string)] = map[string]bool{"pass": r["pass"].(bool)}
	}
	objectMap := map[string]map[string]bool{}
	for _, r := range objectResults {
		objectMap[r["name"].(string)] = map[string]bool{"mp": r["mp"].(bool), "json": r["json"].(bool), "gron": r["gron"].(bool)}
	}
	results := map[string]interface{}{"scalars": scalarMap, "objects": objectMap}
	rj, _ := json.Marshal(results)
	os.WriteFile(filepath.Join(outDir, "results.json"), rj, 0644)

	fail := countScalarFail(scalarResults) + countObjectFail(objectResults)
	pass := len(scalarResults) + len(objectResults) - fail
	fmt.Printf("Go done: %d passed, %d failed\n", pass, fail)
	if fail > 0 {
		os.Exit(1)
	}
}

// ═══════════════════════════════════
// Generic schema-driven decode/encode
// ═══════════════════════════════════

func readScalar(r specodec.SpecReader, typ string) interface{} {
	switch typ {
	case "string":
		return r.ReadString()
	case "boolean":
		return r.ReadBool()
	case "int8", "int16", "int32":
		return r.ReadInt32()
	case "int64":
		return r.ReadInt64()
	case "uint8", "uint16", "uint32":
		return r.ReadUint32()
	case "uint64":
		return r.ReadUint64()
	case "float32":
		return r.ReadFloat32()
	case "float64":
		return r.ReadFloat64()
	case "bytes":
		return r.ReadBytes()
	default:
		panic(fmt.Sprintf("unknown scalar: %s", typ))
	}
}

func writeScalarMP(w *specodec.MsgPackWriter, val interface{}, typ string) {
	switch typ {
	case "string":
		w.WriteString(val.(string))
	case "boolean":
		w.WriteBool(val.(bool))
	case "int8", "int16", "int32":
		w.WriteInt32(val.(int32))
	case "int64":
		w.WriteInt64(val.(int64))
	case "uint8", "uint16", "uint32":
		w.WriteUint32(val.(uint32))
	case "uint64":
		w.WriteUint64(val.(uint64))
	case "float32":
		w.WriteFloat32(val.(float32))
	case "float64":
		w.WriteFloat64(val.(float64))
	case "bytes":
		w.WriteBytes(val.([]byte))
	}
}

func writeScalarJSON(w *specodec.JsonWriter, val interface{}, typ string) {
	switch typ {
	case "string":
		w.WriteString(val.(string))
	case "boolean":
		w.WriteBool(val.(bool))
	case "int8", "int16", "int32":
		w.WriteInt32(val.(int32))
	case "int64":
		w.WriteInt64(val.(int64))
	case "uint8", "uint16", "uint32":
		w.WriteUint32(val.(uint32))
	case "uint64":
		w.WriteUint64(val.(uint64))
	case "float32":
		w.WriteFloat32(val.(float32))
	case "float64":
		w.WriteFloat64(val.(float64))
	case "bytes":
		w.WriteBytes(val.([]byte))
	}
}

func writeScalarGRON(w *specodec.GronWriter, val interface{}, typ string) {
	switch typ {
	case "string":
		w.WriteString(val.(string))
	case "boolean":
		w.WriteBool(val.(bool))
	case "int8", "int16", "int32":
		w.WriteInt32(val.(int32))
	case "int64":
		w.WriteInt64(val.(int64))
	case "uint8", "uint16", "uint32":
		w.WriteUint32(val.(uint32))
	case "uint64":
		w.WriteUint64(val.(uint64))
	case "float32":
		w.WriteFloat32(val.(float32))
	case "float64":
		w.WriteFloat64(val.(float64))
	case "bytes":
		w.WriteBytes(val.([]byte))
	}
}

type GenericObj map[string]interface{}

func decodeField(r specodec.SpecReader, field FieldSchema) interface{} {
	if field.IsArray {
		var arr []interface{}
		r.BeginArray()
		for r.HasNextElement() {
			if field.IsModel {
				arr = append(arr, decodeModel(r, field.Type))
			} else {
				arr = append(arr, readScalar(r, field.Type))
			}
		}
		r.EndArray()
		return arr
	}
	if field.IsModel {
		return decodeModel(r, field.Type)
	}
	return readScalar(r, field.Type)
}

func decodeModel(r specodec.SpecReader, modelName string) GenericObj {
	ms := schemaMap[modelName]
	o := GenericObj{}
	r.BeginObject()
	for r.HasNextField() {
		k := r.ReadFieldName()
		for _, f := range ms.Fields {
			if f.Name == k {
				o[k] = decodeField(r, f)
				break
			}
		}
	}
	r.EndObject()
	return o
}

func readScalarGron(r *specodec.GronReader, typ string) interface{} {
	switch typ {
	case "string":
		return r.ReadString()
	case "boolean":
		return r.ReadBool()
	case "int8", "int16", "int32":
		return r.ReadInt32()
	case "int64":
		return r.ReadInt64()
	case "uint8", "uint16", "uint32":
		return r.ReadUint32()
	case "uint64":
		return r.ReadUint64()
	case "float32":
		return r.ReadFloat32()
	case "float64":
		return r.ReadFloat64()
	case "bytes":
		return r.ReadBytes()
	default:
		panic(fmt.Sprintf("unknown scalar: %s", typ))
	}
}

func decodeFieldGron(r *specodec.GronReader, field FieldSchema) interface{} {
	if field.IsArray {
		var arr []interface{}
		r.BeginArray()
		for r.HasNextElement() {
			r.NextElement()
			if field.IsModel {
				arr = append(arr, decodeModelGron(r, field.Type))
			} else {
				arr = append(arr, readScalarGron(r, field.Type))
			}
		}
		r.EndArray()
		return arr
	}
	if field.IsModel {
		return decodeModelGron(r, field.Type)
	}
	return readScalarGron(r, field.Type)
}

func decodeModelGron(r *specodec.GronReader, modelName string) GenericObj {
	ms := schemaMap[modelName]
	o := GenericObj{}
	r.BeginObject()
	for r.HasNextField() {
		k := r.ReadFieldName()
		for _, f := range ms.Fields {
			if f.Name == k {
				o[k] = decodeFieldGron(r, f)
				break
			}
		}
	}
	r.EndObject()
	return o
}

func encodeFieldMP(w *specodec.MsgPackWriter, val interface{}, field FieldSchema) {
	if field.IsArray {
		arr := val.([]interface{})
		w.BeginArray(len(arr))
		for _, item := range arr {
			if field.IsModel {
				encodeModelInlineMP(w, item.(GenericObj), field.Type)
			} else {
				writeScalarMP(w, item, field.Type)
			}
		}
		w.EndArray()
		return
	}
	if field.IsModel {
		encodeModelInlineMP(w, val.(GenericObj), field.Type)
		return
	}
	writeScalarMP(w, val, field.Type)
}

func encodeModelMP(o GenericObj, modelName string) []byte {
	w := &specodec.MsgPackWriter{}
	encodeModelInlineMP(w, o, modelName)
	return w.ToBytes()
}

func encodeModelInlineMP(w *specodec.MsgPackWriter, o GenericObj, modelName string) {
	ms := schemaMap[modelName]
	count := 0
	for _, f := range ms.Fields {
		if f.Optional {
			if _, ok := o[f.Name]; !ok {
				continue
			}
		}
		count++
	}
	w.BeginObject(count)
	for _, f := range ms.Fields {
		if f.Optional {
			if _, ok := o[f.Name]; !ok {
				continue
			}
		}
		w.WriteField(f.Name)
		encodeFieldMP(w, o[f.Name], f)
	}
	w.EndObject()
}

func encodeFieldJSON(w *specodec.JsonWriter, val interface{}, field FieldSchema) {
	if field.IsArray {
		arr := val.([]interface{})
		w.BeginArray(len(arr))
		for _, item := range arr {
			w.NextElement()
			if field.IsModel {
				encodeModelInlineJSON(w, item.(GenericObj), field.Type)
			} else {
				writeScalarJSON(w, item, field.Type)
			}
		}
		w.EndArray()
		return
	}
	if field.IsModel {
		encodeModelInlineJSON(w, val.(GenericObj), field.Type)
		return
	}
	writeScalarJSON(w, val, field.Type)
}

func encodeModelJSON(o GenericObj, modelName string) []byte {
	w := &specodec.JsonWriter{}
	encodeModelInlineJSON(w, o, modelName)
	return w.ToBytes()
}

func encodeModelInlineJSON(w *specodec.JsonWriter, o GenericObj, modelName string) {
	ms := schemaMap[modelName]
	count := 0
	for _, f := range ms.Fields {
		if f.Optional {
			if _, ok := o[f.Name]; !ok { continue }
		}
		count++
	}
	w.BeginObject(count)
	for _, f := range ms.Fields {
		if f.Optional {
			if _, ok := o[f.Name]; !ok {
				continue
			}
		}
		w.WriteField(f.Name)
		encodeFieldJSON(w, o[f.Name], f)
	}
	w.EndObject()
}

func encodeFieldGRON(w *specodec.GronWriter, val interface{}, field FieldSchema) {
	if field.IsArray {
		arr := val.([]interface{})
		w.BeginArray(len(arr))
		for _, item := range arr {
			w.NextElement()
			if field.IsModel {
				encodeModelInlineGRON(w, item.(GenericObj), field.Type)
			} else {
				writeScalarGRON(w, item, field.Type)
			}
		}
		w.EndArray()
		return
	}
	if field.IsModel {
		encodeModelInlineGRON(w, val.(GenericObj), field.Type)
		return
	}
	writeScalarGRON(w, val, field.Type)
}

func encodeModelGRON(o GenericObj, modelName string) []byte {
	w := specodec.NewGronWriter()
	encodeModelInlineGRON(w, o, modelName)
	return w.ToBytes()
}

func encodeModelInlineGRON(w *specodec.GronWriter, o GenericObj, modelName string) {
	ms := schemaMap[modelName]
	count := 0
	for _, f := range ms.Fields {
		if f.Optional {
			if _, ok := o[f.Name]; !ok {
				continue
			}
		}
		count++
	}
	w.BeginObject(count)
	for _, f := range ms.Fields {
		if f.Optional {
			if _, ok := o[f.Name]; !ok {
				continue
			}
		}
		w.WriteField(f.Name)
		encodeFieldGRON(w, o[f.Name], f)
	}
	w.EndObject()
}

func processObject(name string) (bool, bool, bool) {
	mpOk := false
	jsonOk := false
	gronOk := false

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  FAIL %s.msgpack: %v\n", name, r)
			}
		}()
		mpBuf, err := os.ReadFile(filepath.Join(vecDir, name+".msgpack"))
		if err != nil {
			fmt.Printf("  FAIL %s.msgpack: %v\n", name, err)
			return
		}
		r := specodec.NewMsgPackReader(mpBuf)
		o := decodeModel(r, name)
		mpOut := encodeModelMP(o, name)
		os.WriteFile(filepath.Join(outDir, name+".msgpack"), mpOut, 0644)
		mpOk = true
	}()

	var compactOut []byte
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  FAIL %s.json: %v\n", name, r)
			}
		}()
		jsonBuf, err := os.ReadFile(filepath.Join(vecDir, name+".json"))
		if err != nil {
			fmt.Printf("  FAIL %s.json: %v\n", name, err)
			return
		}
		r2 := specodec.NewJsonReader(jsonBuf)
		o2 := decodeModel(r2, name)
		compactOut = encodeModelJSON(o2, name)
		os.WriteFile(filepath.Join(outDir, name+".json"), compactOut, 0644)
		jsonOk = true
	}()

	if compactOut != nil {
		prettyPath := filepath.Join(vecDir, name+".pretty.json")
		if data, err := os.ReadFile(prettyPath); err == nil {
			func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Printf("  FAIL %s.pretty.json: %v\n", name, r)
						jsonOk = false
					}
				}()
				r3 := specodec.NewJsonReader(data)
				o3 := decodeModel(r3, name)
				prettyOut := encodeModelJSON(o3, name)
				if !bytes.Equal(prettyOut, compactOut) {
					fmt.Printf("  FAIL %s.pretty.json: re-encoded bytes differ\n", name)
					jsonOk = false
				}
			}()
		}
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  FAIL %s.gron: %v\n", name, r)
			}
		}()
		gronBuf, err := os.ReadFile(filepath.Join(vecDir, name+".gron"))
		if err != nil {
			fmt.Printf("  FAIL %s.gron: %v\n", name, err)
			return
		}
		r4 := specodec.NewGronReader(gronBuf)
		o4 := decodeModelGron(r4, name)
		gronOut := encodeModelGRON(o4, name)
		os.WriteFile(filepath.Join(outDir, name+".gron"), gronOut, 0644)
		gronOk = true
	}()

	return mpOk, jsonOk, gronOk
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" { return v }
	return def
}

func countScalarFail(results []map[string]interface{}) int {
	n := 0
	for _, r := range results {
		if !r["pass"].(bool) { n++ }
	}
	return n
}

func countObjectFail(results []map[string]interface{}) int {
	n := 0
	for _, r := range results {
		if !r["mp"].(bool) || !r["json"].(bool) || !r["gron"].(bool) { n++ }
	}
	return n
}
