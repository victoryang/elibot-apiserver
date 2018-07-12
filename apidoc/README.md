# apidoc
## run apidoc example with docker
### get apidoc source
	git clone https://github.com/apidoc/apidoc.git

### build docker image
	cd apidoc
	docker build -t apidoc/apidoc .

### run docker image, replace $PWD with actual directory. This will generate an example website
	docker run --rm -v '$(PWD):/apidoc' -it apidoc/apidoc --input ./example --output ./docker-example -v

## test api
curl -H "Content-Type: application/json" -X PUT -d '{"index":"1", "data":"1"}'

-H   --header
-X   --request
-d   --data