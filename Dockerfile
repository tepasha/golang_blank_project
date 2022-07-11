FROM golang:1.17.2-bullseye

RUN go version

COPY ./ ./

RUN go mod download
RUN go build -o pmc_api_base ./main.go

CMD ["./pmc_api_base"]