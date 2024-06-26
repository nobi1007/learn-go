# Learning Golang

Completed [https://go.dev/doc/tutorial/create-module](https://go.dev/doc/tutorial/create-module) tutorial

### Publishing a new version:

1.

```bash
cd greetings

git commit -m "mymodule: changes for v0.0.1"
git tag v0.0.1

git push origin v0.0.1

GOPROXY=proxy.golang.org go list -m github.com/nobi1007/learn-go@v0.0.1
```

2. cd into `hello` and delete old version of imported module from go.mod and then run `go mod tidy`
