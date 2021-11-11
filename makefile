

generate:
	 protoc --go_out=./gen --go_opt=paths=source_relative \
        --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
        proto/service.proto
build:
	docker build  --target=release -t watermap:latest  .

run:
	docker run \
		-e GIN_MODE=debug \
		-e grpc_port=8080 \
		-p 8080:8080 \
		--expose 8080 \
		-e http_port=8081 \
		-p 8081:8081 \
		--expose 8081 \
		watermap:latest
		

