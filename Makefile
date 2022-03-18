

race:
	go run -race cmd/main.go

docker:
	# docker-compose up --build
	docker build -t sysmetric . && docker run -i -t --network host sysmetric