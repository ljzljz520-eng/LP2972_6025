# LP2972_6025

Go project for module `ticketverify`.

## Standard commands

```bash
go build ./...
go test -count=1 ./...
```

## Run

```bash
go run ./cmd/verifier
```

## Docker validation

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh my-go-task linux/arm64
./build_benzhi_docker.sh my-go-task linux/amd64
docker run -it my-go-task:latest
```
