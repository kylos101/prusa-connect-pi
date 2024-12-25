# prusa_connect

## Instructions

1. Download a release for your pi
2. Copy the token for your camera from Prusa Connect, and set it as a `CONNECT_TOKEN` environment variable on your pi
3. Run the binary (perhaps as a service)
4. Print something
5. Observe the first and last images via the Print History from Prusa Connect

## dependencies

A Raspberry Pi zero, with a connected camera, and `rpicam-still` available on the OS.

## development

### generate an API client using docker

[Reference](https://openapi-generator.tech/docs/installation#docker)

```bash
sudo docker run --rm   -v ${PWD}:/local openapitools/openapi-generator-cli generate   -i /local/specs/prusaconnect.0.22.0.yaml   -g go   -o /local/client
```


### replacing the go.mod entry with the generated client

```bash
go mod edit -replace github.com/kylos101/prusa-connect-pi/v2/pkg/openapi=./pkg/openapi
```

### creating a manual release

Prior to doing, you must have set a GITHUB_TOKEN environment variable.

```bash
sudo docker run --rm --privileged   -v $PWD:/go/src/github.com/kylos101/prusa-connect-pi -v /var/run/docker.sock:/var/run/docker.sock -w /go/src/github.com/kylos101/prusa-connect-pi -e GITHUB_TOKEN goreleaser/goreleaser release
```