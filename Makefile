build:
	docker build -t dev_backend_summer19 .

dev:
	- @make build
	- docker run --rm -it -v ${PWD}:/go/src -p 8080:80 dev_backend_summer19 bash
