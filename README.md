```shell
docker build -t dotfiles . --build-arg USERNAME=$(whoami)
```

```shell
docker run -it -v $(pwd):/home/$(whoami)/.dotfiles dotfiles /bin/bash
```
