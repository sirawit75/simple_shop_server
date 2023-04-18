protogen:
	protoc  --proto_path=api --go_out=pkg  \
	--go-grpc_out=pkg  \
	--grpc-gateway_out=pkg \
	api/*.proto


userdb:
	docker run --name userdb  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createuserdb:
	docker exec -it userdb createdb --username=root --owner=root user

dropuserdb:
	docker exec -it userdb dropdb user


productdb:
	docker run --name productdb  -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createproductdb:
	docker exec -it productdb createdb --username=root --owner=root product

dropproductdb:
	docker exec -it  productdb dropdb product

loggerdb:
	docker run --name loggerdb -p 27017:27017 -d mongo:6.0-jammy


test:
	go test -short ./...


compose:
	docker-compose up -d