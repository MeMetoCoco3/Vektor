build: 
	go build -o bin/vektor ./

run:
	./bin/vektor

test:
	go test ./Vektor -v
