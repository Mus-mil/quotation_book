TARGET = quotation_book
TEST_SOURSE =

all: clean build run

build:
	go build -o $(TARGET) cmd/quotation_book/main.go

run:
	./$(TARGET)

test:
	go test ./...

clean:
	rm -rf $(TARGET) *.o