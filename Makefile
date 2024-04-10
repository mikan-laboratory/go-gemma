ENVIRONMENT ?= production
TOKEN ?= 8e712f41d6015bc34c27f8071b42e7542131d2d7d00923668a73866059ce0c6b

IMAGE_NAME ?= go-gemma-image
CONTAINER_NAME ?= go-gemma-container

.PHONY: build run all clean clean-image clean-all terminal

all: build run
clean-all: clean clean-image

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -d --name $(CONTAINER_NAME) -p 8080:8080 \
	-e ENVIRONMENT=$(ENVIRONMENT) \
	-e TOKEN=$(TOKEN) \
	$(IMAGE_NAME)

clean:
	docker stop $(CONTAINER_NAME) && docker rm $(CONTAINER_NAME)

clean-image:
	docker rmi $(IMAGE_NAME)

terminal:
	docker exec -it $(CONTAINER_NAME) /bin/bash
