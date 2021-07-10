FROM golang

RUN mkdir /app
ADD . /app
WORKDIR /app
## pull in dependencies
RUN go mod download
## project build with the necessary go libraries included.
RUN go build -o main .
## run binary executable
CMD ["/app/main"]