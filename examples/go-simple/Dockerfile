FROM alpine AS go-builder

RUN apk add --no-cache go

WORKDIR /src/projects
ADD . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o ./app *.go

FROM alpine
COPY --from=go-builder /src/projects/app /bin/app

ENTRYPOINT ["/bin/app"]
