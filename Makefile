depend:
	dep ensure

go_run:
	go run cmd/watermark/main.go

go_build:
	go build -o build/watermark cmd/watermark/main.go

run:
	./build/watermark