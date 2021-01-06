FROM golang

WORKDIR /api

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 5000

CMD ["go", "run", "api"]