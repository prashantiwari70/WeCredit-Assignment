
FROM golang:alpine


WORKDIR /app


COPY . .


RUN go mod init go-log-service
RUN go mod tidy
RUN go build -o logservice .


EXPOSE 8080

CMD ["./logservice"]
