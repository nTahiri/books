clean:
	rm -f books

build: clean
	go build -o books

test:
	go test -v ./go

run: build
	chmod +x books
	./books
docker_build: clean
	docker build --network=host -t books .

docker_run:
	docker run -p 8080:8080 --rm -it books