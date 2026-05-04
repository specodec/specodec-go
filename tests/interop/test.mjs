import { execSync } from 'child_process';
import { existsSync, mkdirSync, rmSync, readdirSync, writeFileSync, readFileSync } from 'fs';
import { join, dirname } from 'path';
import { fileURLToPath } from 'url';

const __dir = dirname(fileURLToPath(import.meta.url));
const CACHE = join(__dir, '.tests-cache');
const EMIT_GEN = join(__dir, 'emit', 'emit_gen');
const OUT_DIR = join(__dir, 'output');

function run(cmd) {
  console.log('  >', cmd);
  execSync(cmd, { stdio: 'inherit' });
}

console.log('\n=== Step 1: Install dependencies ===');
run(`cd ${__dir} && npm install`);

if (!existsSync(join(CACHE, 'vectors'))) { console.log('\n=== Step 2: Clone tests repo ==='); if (existsSync(CACHE)) rmSync(CACHE, { recursive: true }); run(`git clone --depth=1 https://github.com/specodec/tests ${CACHE}`) } else { console.log('\n=== Step 2: Using cached .tests-cache ===') }

console.log('\n=== Step 3: Generate vectors ===');
run(`cd ${CACHE} && npm install --frozen-lockfile`);
run(`cd ${CACHE} && node gen_types.mjs`);

const VEC_DIR = join(CACHE, 'vectors');

console.log('\n=== Step 4: Generate emit code ===');
if (existsSync(EMIT_GEN)) rmSync(EMIT_GEN, { recursive: true });
mkdirSync(EMIT_GEN, { recursive: true });

run(`cd ${__dir} && node_modules/.bin/tsp compile ${CACHE}/alltypes.tsp --emit=@specodec/typespec-emitter-golang \
  --option @specodec/typespec-emitter-golang.emitter-output-dir=${EMIT_GEN}`);

const goFiles = readdirSync(EMIT_GEN).filter(f => f.endsWith('.go'));
if (goFiles.length > 0) {
  console.log(`  ✓ Generated ${goFiles.join(', ')}`);
  
  // Move generated files into proper package structure
  const pkgDir = join(EMIT_GEN, 'specodec_all_types');
  mkdirSync(pkgDir, { recursive: true });
  for (const f of goFiles) {
    const src = join(EMIT_GEN, f);
    const dest = join(pkgDir, f);
    const content = readFileSync(src, 'utf-8');
    writeFileSync(dest, content);
    rmSync(src);
  }
  console.log(`  ✓ Moved to specodec_all_types package`);
  
  // Fix generated code to use correct types
  const genFile = join(pkgDir, 'all_types_types.go');
  let content = readFileSync(genFile, 'utf-8');
  content = content.replace(/specodec\.Writer/g, 'specodec.SpecWriter');
  content = content.replace(/specodec\.Reader/g, 'specodec.SpecReader');
  content = content.replace(/\*specodec\.SpecWriter/g, 'specodec.SpecWriter');
  content = content.replace(/\*specodec\.SpecReader/g, 'specodec.SpecReader');
  writeFileSync(genFile, content);
  console.log(`  ✓ Fixed SpecWriter/SpecReader references`);
  
  // Fix generated code to use pointer parameters
  content = content.replace(/func write(\w+)\(w specodec\.SpecWriter, obj (\w+)\)/g, 'func write$1(w specodec.SpecWriter, obj *$2)');
  content = content.replace(/\(obj, w\)/g, '(obj, w)');
  writeFileSync(genFile, content);
  console.log(`  ✓ Fixed pointer parameters for write functions`);
} else {
  console.error('  FAIL: No generated Go files');
  process.exit(1);
}

console.log('\n=== Step 5: Generate test runner ===');
mkdirSync(join(__dir, 'emit'), { recursive: true });
run(`cd ${__dir} && VEC_DIR=${VEC_DIR} node generate_emit_runner.mjs`);

console.log('\n=== Step 6: Run tests ===');
if (existsSync(OUT_DIR)) rmSync(OUT_DIR, { recursive: true });
mkdirSync(OUT_DIR, { recursive: true });

// Use local runtime source instead of remote
const runtimeDir = join(__dir, '..', '..');  // specodec-runtime-golang root
const goMod = `module emit_go

go 1.23

require github.com/specodec/specodec-runtime-golang v0.0.0

replace github.com/specodec/specodec-runtime-golang => ${runtimeDir}
`;
writeFileSync(join(__dir, 'emit', 'go.mod'), goMod);

// Update generated code imports to match new module path
const emitGenDir = join(__dir, 'emit', 'emit_gen', 'specodec_all_types');
const genFile = join(emitGenDir, 'all_types_types.go');
let genContent = readFileSync(genFile, 'utf-8');
genContent = genContent.replace(/github\.com\/specodec\/specodec-go/g, 'github.com/specodec/specodec-runtime-golang');
genContent = genContent.replace(/github\.com\/specodec\/specodec-runtime-go/g, 'github.com/specodec/specodec-runtime-golang');
genContent = genContent.replace(/github\.com\/specodec\/specodec-runtime-golanglang/g, 'github.com/specodec/specodec-runtime-golang');
writeFileSync(genFile, genContent);
console.log('  ✓ Updated runtime import paths');

try { run(`cd ${__dir}/emit && go mod tidy`); } catch (e) { console.log("Go mod tidy completed (some failures expected)"); }
try { run(`cd ${__dir}/emit && VEC_DIR=${VEC_DIR} OUT_DIR=${OUT_DIR} go run run_emit.go`); } catch (e) { console.log("Go tests completed (some failures expected)"); }

console.log('\n=== Step 7: Compare output ===');
const manifest = JSON.parse(readFileSync(join(VEC_DIR, 'manifest.json'), 'utf-8'));
let match = 0, mismatch = 0;

for (const [name] of Object.entries(manifest.scalars || {})) {
  const expected = join(VEC_DIR, 'scalars', `${name}.mp`);
  const actual = join(OUT_DIR, 'scalars', `${name}.mp`);
  if (!existsSync(actual)) { mismatch++; console.log(`MISSING: ${name}.mp`); continue; }
  if (readFileSync(expected).equals(readFileSync(actual))) match++;
  else { mismatch++; console.log(`MISMATCH: ${name}.mp`); }
}
for (const model of manifest.testModels || []) {
  for (const fmt of ['msgpack', 'json', 'gron', 'unformatted.json']) {
    const expected = join(VEC_DIR, `${model}.${fmt}`);
    const actual = join(OUT_DIR, `${model}.${fmt}`);
    if (!existsSync(expected)) continue;
    if (!existsSync(actual)) { mismatch++; console.log(`MISSING: ${model}.${fmt}`); continue; }
    if (readFileSync(expected).equals(readFileSync(actual))) match++;
    else { mismatch++; console.log(`MISMATCH: ${model}.${fmt}`); }
  }
}
const total = match + mismatch;
console.log(`${match}/${total} match, ${mismatch} mismatch`);
if (mismatch > 0) process.exit(1);

console.log('\n=== ALL PASSED ===');