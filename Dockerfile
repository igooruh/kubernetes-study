FROM --platform=linux/amd64 golang:1.24

COPY . .

RUN go build -o server .

CMD ["./server"]