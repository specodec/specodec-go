package main

import "os"

func main() {
    vec := envOr("VEC_DIR", "/app/vectors")
    out := envOr("OUT_DIR", "/app/output_emit_go")
    os.MkdirAll(out, 0755)
    runEmit(vec, out)
}

func envOr(k, d string) string {
    if v := os.Getenv(k); v != "" {
        return v
    }
    return d
}
