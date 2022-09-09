```shell
docker build -t dotfiles . --build-arg USERNAME=$(whoami)
```

```shell
docker run -it -v $(pwd):/home/$(whoami)/.dotfiles dotfiles /bin/bash
```

```shell
DOTPATH=$(pwd) go run -tags linux,amd64,server ./cmd/dotfiles/
```
