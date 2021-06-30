run:
	# gin --all -i run main.go
	air

mod:
	go mod vendor

build:
	go build -mod=vendor

init:
	go get github.com/labstack/echo
	go get gopkg.in/go-playground/validator.v9
	# go get github.com/ilyakaznacheev/cleanenv
	go get github.com/caarlos0/env
	go mod vendor
	# go mod tidy
	# go get -u github.com/go-sql-driver/mysql

## docker 
up:
	docker-compose up -d

up-b:
	docker-compose up -d --build

kill:
	docker-compose kill

reload: kill up

## docker-clean: docker remove all containers in stack
clean:
	docker-compose rm -fv --all
	docker-compose down --rmi local --remove-orphans

net:
	docker network create my-network


## mysql: workspace container bash
db-bash:
	docker-compose exec sample-db bash
