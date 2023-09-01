all: build-image run-container
build-image:
	docker build -t ascii-art-web .
run-container:
	docker run --name cont -dp 4000:4000 ascii-art-web