```shell
docker build -t dotfiles . --build-arg USERNAME=$(whoami)
```

```shell
docker run -it -v $(pwd):/home/$(whoami)/.dotfiles dotfiles /bin/bash
```

```shell
DOTPATH=$(pwd) go run -tags linux,amd64,server ./cmd/dotfiles/ deploy
DOTPATH=$(pwd) go run -tags linux,amd64,server ./cmd/dotfiles/ init
```

```shell
DOTPATH=$(pwd) go test -cover -tags linux,amd64,server -v ./... -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```
