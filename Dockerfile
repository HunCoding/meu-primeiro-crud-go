FROM golang:1.19 AS BUILDER


WORKDIR /app
COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
 GOOS=linux go build -o meuprimeirocrudgo .

FROM golang:1.19-alpine3.15 as runner

RUN adduser -D huncoding

COPY --from=BUILDER /app/meuprimeirocrudgo /app/meuprimeirocrudgo

RUN chown -R huncoding:huncoding /app
RUN chmod +x /app/meuprimeirocrudgo

EXPOSE 8080

USER huncoding

CMD ["./meuprimeirocrudgo"]