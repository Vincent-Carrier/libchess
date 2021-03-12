.PHONY: test

test:
	go test | rg FAIL --context 20

main.wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm
