gobserved
=========

```
$ docker run --name gobserved --rm -v$PWD/src:/usr/local/src levonet/golang:go2go bash -c "cd /usr/local/src/go2/observed && go tool go2go test"
$ docker run --name gobserved --rm -v$PWD/src:/usr/local/src levonet/golang:go2go bash -c "cd /usr/local/src/go/observed && go generate && go test"
```