
all:	coladad wasm

coladad: cmd/coladad/*.go cmd/coladad/config/*.go cmd/coladad/handlers/*.go
	go build -o ./cmd/coladad/coladad cmd/coladad/main.go
	
wasm: src/*.rs
	wasm-pack build