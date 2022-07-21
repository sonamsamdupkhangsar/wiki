# Wiki 
This Wiki go project is based on the (Wiki)[https://go.dev/doc/articles/wiki/].

## build
```go biuld wiki.go```

## run
```./wiki```

## Docker image
To build a docker image for wiki:

```docker build -t wiki .```

To run the docker image:

```docker run -it --rm -p 8081:8080 wiki```

### File structure
Html templates stored in tmpl.  Data files stored in data folder.