init: build up

build:
	cd ./docker && docker-compose build
up:
	cd ./docker && docker-compose up -d