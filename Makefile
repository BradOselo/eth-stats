proto:
	sh third_party/protoc_gen.sh

build: api/proto
	go build -o server cmd/server/server.go

docker: api/proto 
	docker build  -t cardenasrjl/eth-stats:latest .

docker-push:	
	docker image push cardenasrjl/eth-stats:latest

up: build
	docker-compose up --build 
	
down:
	docker-compose down

restart: down up 

clean-up: down  
	docker system prune -f
	make docker
	make up