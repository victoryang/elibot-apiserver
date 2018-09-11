# apidoc
## How to generate apidoc website
### get apidoc source
	git clone https://github.com/apidoc/apidoc.git

### build docker image
	cd apidoc
	docker build -t apidoc/apidoc .

### cp directory input and template to /apidoc
	cp -a /path/to/input .
	cp -a /path/to/template .

### run docker image, generate website
	docker run --rm -v ${PWD}:/apidoc -it apidoc/apidoc --input ./input --template ./template --output ./output -v

## test api
	curl -H "Content-Type: application/json" -X PUT -d '{"index":"1", "data":"1"}'

	-H   --header
	-X   --request
	-d   --data