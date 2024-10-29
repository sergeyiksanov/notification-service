PHONY: generate
generate:
	mkdir -p pkg
	protoc --go_out=pkg --go_opt=paths=source_relative \
		api/v1/api.proto
