plugin:
	cd cmd/protoc-gen-service-catalog && go build

snapshot:
	protoc \
		--plugin=cmd/protoc-gen-service-catalog/protoc-gen-service-catalog \
		--service-catalog_out=. \
		--service-catalog_opt=paths=source_relative \
		--go_out=. \
		--go_opt=paths=source_relative \
		--proto_path=. \
		example/example.proto example/another_example.proto
