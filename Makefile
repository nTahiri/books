clean:
	rm -f books

build: clean
	go build -o books

run: build
	chmod +x books
	./books
docker_build: build
	docker build -t books .

docker_run:
	docker run -it books --port 8080