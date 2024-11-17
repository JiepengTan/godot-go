.DEFAULT_GOAL := pc

CURRENT_PATH=$(shell pwd)
.PHONY: engine init initweb fmt gen test

test:
	cd ./test/ && go run . && cd ../

fmt:
	go fmt ./... 

pc:
	./cmd/scripts/init.sh

web: 
	./cmd/scripts/init_web.sh 

gen:
	cd ./cmd/codegen && go run . && cd ../../ && \
	$(MAKE) fmt 