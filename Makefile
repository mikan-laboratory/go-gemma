ENVIRONMENT ?= production
TOKEN ?= 68850cc7becba8a8cca282354f

IMAGE_NAME ?= go-gemma-image
CONTAINER_NAME ?= go-gemma-container

.PHONY: build run all clean clean-image clean-all terminal

all: build run
clean-all: clean clean-image

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -d --name $(CONTAINER_NAME) -p 8081:8081 \
	-e ENVIRONMENT=$(ENVIRONMENT) \
	-e TOKEN=$(TOKEN) \
	$(IMAGE_NAME)

clean:
	docker stop $(CONTAINER_NAME) && docker rm $(CONTAINER_NAME)

clean-image:
	docker rmi $(IMAGE_NAME)

terminal:
	docker exec -it $(CONTAINER_NAME) /bin/bash
