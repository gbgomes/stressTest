FROM golang:1.21.0 as builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o stresstest .

FROM scratch
COPY --from=builder /app/stresstest .

#ENTRYPOINT ["tail", "-f", "/dev/null"]

ENTRYPOINT ["./stresstest"]

