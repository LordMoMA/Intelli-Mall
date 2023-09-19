install-tools:
	@echo installing tools
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install github.com/vektra/mockery/v2@latest
	@go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	@go install github.com/cucumber/godog/cmd/godog@latest
	@echo done

generate:
	@echo running code generation
	@go generate ./...
	@echo done

build: build-monolith build-services

rebuild: clean-monolith clean-services build

clean-monolith:
	docker image rm intellimall-monolith

clean-services:
	docker image rm intellimall-baskets intellimall-cosec intellimall-customers intellimall-depot intellimall-notifications intellimall-ordering intellimall-payments intellimall-search intellimall-stores

build-monolith:
	docker build -t intellimall-monolith --file docker/Dockerfile .

build-services:
	docker build -t intellimall-baskets --file docker/Dockerfile.microservices --build-arg=service=baskets .
	docker build -t intellimall-cosec --file docker/Dockerfile.microservices --build-arg=service=cosec .
	docker build -t intellimall-customers --file docker/Dockerfile.microservices --build-arg=service=customers .
	docker build -t intellimall-depot --file docker/Dockerfile.microservices --build-arg=service=depot .
	docker build -t intellimall-notifications --file docker/Dockerfile.microservices --build-arg=service=notifications .
	docker build -t intellimall-ordering --file docker/Dockerfile.microservices --build-arg=service=ordering .
	docker build -t intellimall-payments --file docker/Dockerfile.microservices --build-arg=service=payments .
	docker build -t intellimall-search --file docker/Dockerfile.microservices --build-arg=service=search .
	docker build -t intellimall-stores --file docker/Dockerfile.microservices --build-arg=service=stores .

