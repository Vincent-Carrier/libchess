.PHONY: test wasm tiny_wasm

test:
	go test | rg FAIL --context 20

wasm:
	GOOS=js GOARCH=wasm go build -o ../vcar.dev/static/chess/main.wasm

tiny_wasm:
	tinygo build -o ../vcar.dev/static/chess/main.min.wasm \
							 -no-debug \
							 ./cmd/wasm/main.go
