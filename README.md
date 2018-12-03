# goecho-postgre-sample

## requirement

- docker,docker-compose
- golang
- nodejs,npm

## backend

```
$ go get github.com/s-take/goecho-postgre-sample
$ cd $GOPATH/src/github.com/s-take/goecho-postgre-sample/
$ go get -u github.com/golang/dep/cmd/dep
$ dep ensure
$ docker-compose build
$ docker-compose up -d
```

## frontend

```
$ cd $GOPATH/src/github.com/s-take/goecho-postgre-sample/frontend
$ npm install
$ npm run serve
```
