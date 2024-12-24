# prusa_connect

## development

### generate an API client using docker

[Reference](https://openapi-generator.tech/docs/installation#docker)

```bash
sudo docker run --rm   -v ${PWD}:/local openapitools/openapi-generator-cli generate   -i /local/specs/prusaconnect.0.22.0.yaml   -g go   -o /local/client
```


### replacing the go.mod entry with the generated client

```bash
go mod edit -replace github.com/kylos101/prusa-connect/v2/pkg/openapi=./pkg/openapi
```

### creating a manual release

Prior to doing, you must have set a GITHUB_TOKEN environment variable.

```bash
sudo docker run --rm --privileged   -v $PWD:/go/src/github.com/kylos101/prusa-connect -v /var/run/docker.sock:/var/run/docker.sock -w /go/src/github.com/kylos101/prusa-connect -e GITHUB_TOKEN goreleaser/goreleaser release
```