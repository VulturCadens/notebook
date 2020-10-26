# WebAssembly

```shell
GOOS=js GOARCH=wasm go build -o app.wasm

cp $(go env GOROOT)/misc/wasm/wasm_exec.js .

python3 -m http.server
```