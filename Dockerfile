FROM golang:1.22.2

LABEL Name="asciiartwebdockerize" maintainer="amosjoel91@gmail.com" version="latest"

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 9000

CMD [ "go", "run", "main.go", "9000" ]