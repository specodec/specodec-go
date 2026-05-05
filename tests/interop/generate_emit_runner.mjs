import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";

const __dir = path.dirname(fileURLToPath(import.meta.url));
const VEC_DIR = process.env.VEC_DIR || path.join(__dir, ".tests-cache", "vectors");

const manifestPath = path.join(VEC_DIR, "manifest.json");
const manifest = JSON.parse(fs.readFileSync(manifestPath, "utf-8"));

const models = manifest.testModels;
const scalars = manifest.scalars;
const modelNamespaces = manifest.modelNamespaces || {};

const emitGenDir = path.join(__dir, "emit", "emit_gen");

// ── Scan generated Go packages to map model→package ──
const packageMap = {};
const modelPackage = {};

function scanGoDir(dir, base) {
  const entries = fs.readdirSync(dir, { withFileTypes: true });
  for (const entry of entries) {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      scanGoDir(full, base);
    } else if (entry.name.endsWith(".go")) {
      const content = fs.readFileSync(full, "utf-8");
      const pm = content.match(/^package\s+(\S+)/m);
      if (!pm) continue;
      const pkg = pm[1];
      if (!packageMap[pkg]) {
        packageMap[pkg] = path.relative(base, dir);
      }
      for (const model of models) {
        if (content.includes(model + "Codec")) {
          modelPackage[model] = pkg;
        }
      }
    }
  }
}

if (fs.existsSync(emitGenDir)) {
  scanGoDir(emitGenDir, emitGenDir);
}

const packages = Object.keys(packageMap);
if (packages.length === 0) packages.push("specodec_all_types");
if (!packageMap[packages[0]]) packageMap[packages[0]] = packages[0];

for (const model of models) {
  if (!modelPackage[model]) modelPackage[model] = packages[0];
}

function getReadMethod(type) {
  const map = {
    "int32": "ReadInt32", "int64": "ReadInt64",
    "uint32": "ReadUint32", "uint64": "ReadUint64",
    "float32": "ReadFloat32", "float64": "ReadFloat64",
    "string": "ReadString", "bytes": "ReadBytes",
    "bool": "ReadBool",
  };
  return map[type] || "ReadInt32";
}

function getWriteMethod(type) {
  const map = {
    "int32": "WriteInt32", "int64": "WriteInt64",
    "uint32": "WriteUint32", "uint64": "WriteUint64",
    "float32": "WriteFloat32", "float64": "WriteFloat64",
    "string": "WriteString", "bytes": "WriteBytes",
    "bool": "WriteBool",
  };
  return map[type] || "WriteInt32";
}

function toSnakeCase(name) {
  let snake = name.replace(/([A-Z])/g, "_$1").toLowerCase().replace(/^_/, "");
  snake = snake.replace(/\./g, "_").replace(/-/g, "_");
  return snake;
}

function nsToSnake(ns) {
  return ns.split('.').map(p => toSnakeCase(p)).join('_');
}

function toPascalCase(s) {
  return s.split('_').map(p => p.charAt(0).toUpperCase() + p.slice(1)).join('');
}

// ── Group models: by namespace if available, else by generated package ──
let testGroups = {};
const hasModelNs = Object.keys(modelNamespaces).length > 0;

if (hasModelNs) {
  for (const model of models) {
    const ns = modelNamespaces[model] || "AllTypes";
    if (!testGroups[ns]) testGroups[ns] = [];
    testGroups[ns].push(model);
  }
} else {
  for (const model of models) {
    const pkg = modelPackage[model];
    if (!testGroups[pkg]) testGroups[pkg] = [];
    testGroups[pkg].push(model);
  }
}

// ── Test module name from group key ──
function groupKeyToSnake(key) {
  if (hasModelNs) {
    return nsToSnake(key);
  }
  // Package name → simple snake: "specodec_all_types_nested" → "all_types_nested"
  return key.replace(/^specodec_/, "").replace(/_types$/, "");
}

// ── Import helper ──
function genPkgImport(pkg) {
  const dir = packageMap[pkg] || pkg;
  return `\t${pkg} "emit_go/emit_gen/${dir}"`;
}

// ── Scalar function ──
function genScalarFunc(name, info) {
  const snake = toSnakeCase(name);
  const pascal = toPascalCase(snake);
  const funcName = `runScalar_${pascal}`;
  return `
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
}

// ── Model function ──
function genModelFunc(model) {
  const snake = toSnakeCase(model);
  const pascal = toPascalCase(snake);
  const pkgName = modelPackage[model];
  const funcName = `runModel_${pascal}`;
  return `
func ${funcName}(vecDir, outDir string) (passed, failed int) {
\tp, f := tryTest("${model} mp", func() {
\t\tdata, err := os.ReadFile(filepath.Join(vecDir, "${model}.msgpack"))
\t\tif err != nil { panic(err) }
\t\tr := specodec.NewMsgPackReader(data)
\t\tobj := ${pkgName}.${model}Codec.Decode(r)
\t\tw := specodec.NewMsgPackWriter()
\t\t${pkgName}.${model}Codec.Encode(w, obj)
\t\terr = os.WriteFile(filepath.Join(outDir, "${model}.msgpack"), w.ToBytes(), 0644)
\t\tif err != nil { panic(err) }
\t})
\tpassed += p; failed += f

\tp, f = tryTest("${model} json", func() {
\t\tdata, err := os.ReadFile(filepath.Join(vecDir, "${model}.json"))
\t\tif err != nil { panic(err) }
\t\tr := specodec.NewJsonReader(data)
\t\tobj := ${pkgName}.${model}Codec.Decode(r)
\t\tw := specodec.NewJsonWriter()
\t\t${pkgName}.${model}Codec.Encode(w, obj)
\t\terr = os.WriteFile(filepath.Join(outDir, "${model}.json"), w.ToBytes(), 0644)
\t\tif err != nil { panic(err) }
\t})
\tpassed += p; failed += f

\tp, f = tryTest("${model} unformatted", func() {
\t\tdata, err := os.ReadFile(filepath.Join(vecDir, "${model}.unformatted.json"))
\t\tif err != nil { panic(err) }
\t\tr := specodec.NewJsonReader(data)
\t\tobj := ${pkgName}.${model}Codec.Decode(r)
\t\tw := specodec.NewJsonWriter()
\t\t${pkgName}.${model}Codec.Encode(w, obj)
\t\terr = os.WriteFile(filepath.Join(outDir, "${model}.unformatted.json"), w.ToBytes(), 0644)
\t\tif err != nil { panic(err) }
\t})
\tpassed += p; failed += f

\tp, f = tryTest("${model} gron", func() {
\t\tdata, err := os.ReadFile(filepath.Join(vecDir, "${model}.gron"))
\t\tif err != nil { panic(err) }
\t\tr := specodec.NewGronReader(data)
\t\tobj := ${pkgName}.${model}Codec.Decode(r)
\t\tw := specodec.NewGronWriter()
\t\t${pkgName}.${model}Codec.Encode(w, obj)
\t\terr = os.WriteFile(filepath.Join(outDir, "${model}.gron"), w.ToBytes(), 0644)
\t\tif err != nil { panic(err) }
\t})
\tpassed += p; failed += f
\treturn
}
`;
}

// ── Generate test files ──
const srcDir = path.join(__dir, "emit");
fs.mkdirSync(srcDir, { recursive: true });

const testFiles = [];

// 1. Scalars
if (Object.keys(scalars).length > 0) {
  let funcs = '';
  let calls = '';
  for (const [name, info] of Object.entries(scalars)) {
    funcs += genScalarFunc(name, info);
    const snake = toSnakeCase(name);
    const pascal = toPascalCase(snake);
    calls += `\tp, f = runScalar_${pascal}(vecDir, outDir); passed += p; failed += f\n`;
  }

  const code = `package main

import (
\t"os"
\t"path/filepath"
\tspecodec "github.com/specodec/specodec-runtime-golang"
)

${funcs}

func runScalars(vecDir, outDir string) (passed, failed int) {
\tvar p, f int
${calls}
\treturn
}
`;
  const fname = "test_scalars.go";
  fs.writeFileSync(path.join(srcDir, fname), code);
  console.log("Generated emit/" + fname + " (" + Object.keys(scalars).length + " scalars)");
  testFiles.push({ name: "runScalars", file: fname });
}

// 2. Model test files per group
for (const [key, groupModels] of Object.entries(testGroups)) {
  let funcs = '';
  let calls = '';
  const groupPkgs = new Set();
  for (const model of groupModels) {
    funcs += genModelFunc(model);
    const snake = toSnakeCase(model);
    const pascal = toPascalCase(snake);
    calls += `\tp, f = runModel_${pascal}(vecDir, outDir); passed += p; failed += f\n`;
    groupPkgs.add(modelPackage[model]);
  }

  const pkgImports = [...groupPkgs].map(p => genPkgImport(p)).join("\n");

  const baseSnake = groupKeyToSnake(key);
  const runFuncName = "run" + toPascalCase(baseSnake);

  const code = `package main

import (
\t"os"
\t"path/filepath"
\tspecodec "github.com/specodec/specodec-runtime-golang"
${pkgImports}
)

${funcs}

func ${runFuncName}(vecDir, outDir string) (passed, failed int) {
\tvar p, f int
${calls}
\treturn
}
`;
  const fname = "test_" + baseSnake + ".go";
  fs.writeFileSync(path.join(srcDir, fname), code);
  console.log("Generated emit/" + fname + " (" + groupModels.length + " models)");
  testFiles.push({ name: runFuncName, file: fname });
}

// 3. Generate main.go
const mainCalls = testFiles.map(t => `\tp, f = ${t.name}(vecDir, outDir); passed += p; failed += f`).join("\n");

const mainCode = `package main

import (
\t"fmt"
\t"os"
\t"path/filepath"
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

func main() {
\tvecDir := os.Getenv("VEC_DIR")
\toutDir := os.Getenv("OUT_DIR")
\tos.MkdirAll(filepath.Join(outDir, "scalars"), 0755)

\tpassed := 0
\tfailed := 0
\tvar p, f int

${mainCalls}

\tfmt.Printf("emit-go: %d passed, %d failed\\n", passed, failed)
\tif failed > 0 { os.Exit(1) }
}
`;

fs.writeFileSync(path.join(srcDir, "main.go"), mainCode);
console.log("Generated emit/main.go with " + testFiles.length + " test modules");
