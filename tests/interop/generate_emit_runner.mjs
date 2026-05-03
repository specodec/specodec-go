import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";

const __dir = path.dirname(fileURLToPath(import.meta.url));
const VEC_DIR = process.env.VEC_DIR || path.join(__dir, ".tests-cache", "vectors");

const manifestPath = path.join(VEC_DIR, "manifest.json");
const manifest = JSON.parse(fs.readFileSync(manifestPath, "utf-8"));

const models = manifest.testModels;
const scalars = manifest.scalars;

function getReadMethod(type) {
  const map = {
    "int32": "ReadInt32",
    "int64": "ReadInt64",
    "uint32": "ReadUint32",
    "uint64": "ReadUint64",
    "float32": "ReadFloat32",
    "float64": "ReadFloat64",
    "string": "ReadString",
    "bytes": "ReadBytes",
    "bool": "ReadBool",
  };
  return map[type] || "ReadInt32";
}

function getWriteMethod(type) {
  const map = {
    "int32": "WriteInt32",
    "int64": "WriteInt64",
    "uint32": "WriteUint32",
    "uint64": "WriteUint64",
    "float32": "WriteFloat32",
    "float64": "WriteFloat64",
    "string": "WriteString",
    "bytes": "WriteBytes",
    "bool": "WriteBool",
  };
  return map[type] || "WriteInt32";
}

function toSnakeCase(name) {
  let snake = name.replace(/([A-Z])/g, "_$1").toLowerCase().replace(/^_/, "");
  snake = snake.replace(/\./g, "_").replace(/-/g, "_");
  return snake;
}

let scalarFuncs = "";
let scalarCalls = "";
for (const [name, info] of Object.entries(scalars)) {
  const snake = toSnakeCase(name);
  const funcName = `testScalar_${snake}`;
  scalarFuncs += `
func ${funcName}(vecDir, outDir string) (passed, failed int) {
\treturn tryTest("${name} mp", func() {
\t\tdata, err := os.ReadFile(filepath.Join(vecDir, "scalars/${name}.mp"))
\t\tif err != nil { panic(err) }
\t\tr := specodec.NewMsgPackReader(data)
\t\tval := r.${getReadMethod(info.type)}()
\t\tw := specodec.NewMsgPackWriter()
\t\tw.${getWriteMethod(info.type)}(val)
\t\terr = os.WriteFile(filepath.Join(outDir, "scalars/${name}.mp"), w.ToBytes(), 0644)
\t\tif err != nil { panic(err) }
\t})
}
`;
  scalarCalls += `\tp, f = ${funcName}(vecDir, outDir); passed += p; failed += f\n`;
}

let modelFuncs = "";
let modelCalls = "";
for (const model of models) {
  const snake = toSnakeCase(model);
  const funcName = `testModel_${snake}`;
  modelFuncs += `
func ${funcName}(vecDir, outDir string) (passed, failed int) {
\tp, f := tryTest("${model} mp", func() {
\t\tdata, err := os.ReadFile(filepath.Join(vecDir, "${model}.msgpack"))
\t\tif err != nil { panic(err) }
\t\tr := specodec.NewMsgPackReader(data)
\t\tobj := specodec_all_types.${model}Codec.Decode(r)
\t\tw := specodec.NewMsgPackWriter()
\t\tspecodec_all_types.${model}Codec.Encode(w, obj)
\t\terr = os.WriteFile(filepath.Join(outDir, "${model}.msgpack"), w.ToBytes(), 0644)
\t\tif err != nil { panic(err) }
\t})
\tpassed += p; failed += f

\tp, f = tryTest("${model} json", func() {
\t\tdata, err := os.ReadFile(filepath.Join(vecDir, "${model}.json"))
\t\tif err != nil { panic(err) }
\t\tr := specodec.NewJsonReader(data)
\t\tobj := specodec_all_types.${model}Codec.Decode(r)
\t\tw := specodec.NewJsonWriter()
\t\tspecodec_all_types.${model}Codec.Encode(w, obj)
\t\terr = os.WriteFile(filepath.Join(outDir, "${model}.json"), w.ToBytes(), 0644)
\t\tif err != nil { panic(err) }
\t})
\tpassed += p; failed += f

\tp, f = tryTest("${model} unformatted", func() {
\t\tdata, err := os.ReadFile(filepath.Join(vecDir, "${model}.unformatted.json"))
\t\tif err != nil { panic(err) }
\t\tr := specodec.NewJsonReader(data)
\t\tobj := specodec_all_types.${model}Codec.Decode(r)
\t\tw := specodec.NewJsonWriter()
\t\tspecodec_all_types.${model}Codec.Encode(w, obj)
\t\terr = os.WriteFile(filepath.Join(outDir, "${model}.unformatted.json"), w.ToBytes(), 0644)
\t\tif err != nil { panic(err) }
\t})
\tpassed += p; failed += f

\tp, f = tryTest("${model} gron", func() {
\t\tdata, err := os.ReadFile(filepath.Join(vecDir, "${model}.gron"))
\t\tif err != nil { panic(err) }
\t\tr := specodec.NewGronReader(data)
\t\tobj := specodec_all_types.${model}Codec.Decode(r)
\t\tw := specodec.NewGronWriter()
\t\tspecodec_all_types.${model}Codec.Encode(w, obj)
\t\terr = os.WriteFile(filepath.Join(outDir, "${model}.gron"), w.ToBytes(), 0644)
\t\tif err != nil { panic(err) }
\t})
\tpassed += p; failed += f
\treturn
}
`;
  modelCalls += `\tp, f = ${funcName}(vecDir, outDir); passed += p; failed += f\n`;
}

const code = `package main

import (
\t"fmt"
\t"os"
\t"path/filepath"
\tspecodec "github.com/specodec/specodec-go"
\tspecodec_all_types "emit_go/emit_gen/specodec_all_types"
)

func tryTest(name string, fn func()) (passed, failed int) {
\tdefer func() {
\t\tif r := recover(); r != nil {
\t\t\tfmt.Println("FAIL " + name + ": " + fmt.Sprint(r))
\t\t\tfailed = 1
\t\t}
\t}()
\tfn()
\tpassed = 1
\treturn
}

func toSnakeCase(s string) string {
\tresult := ""
\tfor i, c := range s {
\t\tif c >= 65 && c <= 90 {
\t\t\tif i > 0 { result += "_" }
\t\t\tresult += string(c + 32)
\t\t} else {
\t\t\tresult += string(c)
\t\t}
\t}
\treturn result
}

${scalarFuncs}
${modelFuncs}

func main() {
\tvecDir := os.Getenv("VEC_DIR")
\toutDir := os.Getenv("OUT_DIR")
\tos.MkdirAll(filepath.Join(outDir, "scalars"), 0755)

\tpassed := 0
\tfailed := 0
\tvar p, f int

${scalarCalls}
${modelCalls}

\tfmt.Printf("emit-go: %d passed, %d failed\\n", passed, failed)
\tif failed > 0 { os.Exit(1) }
}
`;

const srcDir = path.join(__dir, "emit");
fs.mkdirSync(srcDir, { recursive: true });
const outFile = path.join(srcDir, "run_emit.go");
fs.writeFileSync(outFile, code);
console.log("Generated emit/run_emit.go");
