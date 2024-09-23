FROM golang:latest

WORKDIR /api/

COPY . /api/.

RUN go mod tidy 
EXPOSE 8000:8000
CMD [ "go", "run", "main.go" ]