.PHONY: test wasm tiny_wasm

test:
	go test -json | tparse -all

wasm:
	GOOS=js GOARCH=wasm go build -o ../vcar.dev/static/chess/main.wasm ./cmd/wasm/

tiny_wasm:
	tinygo build -o ../vcar.dev/static/chess/main.min.wasm \
							 -no-debug ./cmd/wasm/
