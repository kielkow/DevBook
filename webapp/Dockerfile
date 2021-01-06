FROM golang

WORKDIR /webapp

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 3000

CMD ["go", "run", "webapp"]