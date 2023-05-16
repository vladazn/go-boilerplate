
proto:
	protoc --proto_path=api/proto \
	--go_out=api/client --go_opt=paths=source_relative \
	--go-grpc_out=api/client --go-grpc_opt=paths=source_relative api/proto/*.proto


.PHONY: gateway
gateway:
	protoc --proto_path=api/proto \
      --grpc-gateway_out ./api/gateway \
      --grpc-gateway_opt logtostderr=true \
      --grpc-gateway_opt paths=source_relative \
      --grpc-gateway_opt standalone=true \
      --grpc-gateway_opt grpc_api_configuration=api/yaml/core.yaml \
      api/proto/*.proto


.PHONY: swagger
swagger:
	protoc -I=./api/proto --proto_path=api/proto \
      --openapiv2_out ./api/swagger \
      --openapiv2_opt grpc_api_configuration=api/yaml/core.yaml \
      --openapiv2_opt openapi_configuration=api/yaml/core.swagger.yaml \
      --openapiv2_opt allow_merge=true \
      --openapiv2_opt enums_as_ints=true \
      api/proto/*.proto

.PHONY: infra
infra:
	docker-compose -f infra/docker-compose.yaml -p boilerplate up -d