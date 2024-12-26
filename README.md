# prusa-connect-pi-camera

## dependencies

A Raspberry Pi with a camera, where `rpicam-still -n -o test.jpg` just works.

## Instructions

1. Add a 3rd party camer in [Prusa Connect](https://connect.prusa3d.com/), copy the token, and set it as a `CONNECT_TOKEN` environment variable on your pi
2. Download and run the binary (see below instructions to setup a service)
3. Print something, pictures will be taken on an interval (5m is the default, set `CONNECT_INTERVAL` to override)
4. Observe the first and last images via the Print History from Prusa Connect

### Run the binary as a service

Setup a service in `/etc/systemd/system/prusa-connect-pi-camera.service` like so:

```
[Unit]
Description=Prusa Connect Pi Camera Service
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=your_user
WorkingDirectory=/home/your_user
ExecStart=/usr/local/bin/prusa-connect-pi-camera
Restart=on-failure
RestartSec=10
StandardOutput=syslog
StandardError=syslog
EnvironmentFile=/etc/prusa-connect-pi-camera/config.env

[Install]
WantedBy=multi-user.target
```

Create a file containing environment variables in `/etc/prusa-connect-pi-camera/config.env`:

```
CONNECT_TOKEN=your_token
#CONNECT_URL=some_other_url # Set to override the default url
#ENABLE_PPROF=true # Enable pprof for debugging if needed
#PPROF_PORT=6060 # Set port if you want other than 6060 for pprof
#CONNECT_DEBUG=true # Set debug to troubleshoot API calls
```

Setup and inspect the service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable prusa-connect-pi-camera.service
sudo systemctl start prusa-connect-pi-camera.service
sudo systemctl status prusa-connect-pi-camera.service
sudo journalctl -u prusa-connect-pi-camera.service
```


## development

### generate an API client using docker

[Reference](https://openapi-generator.tech/docs/installation#docker)

```bash
sudo docker run --rm   -v ${PWD}:/local openapitools/openapi-generator-cli generate   -i /local/specs/prusaconnect.0.22.0.yaml   -g go   -o /local/pkg/openapi
```


### replacing the go.mod entry with the generated client

```bash
go mod edit -replace github.com/kylos101/prusa-connect-pi-camera/v2/pkg/openapi=./pkg/openapi
```

### creating a manual release

Prior to doing, you must have set a GITHUB_TOKEN environment variable.

```bash
sudo docker run --rm --privileged   -v $PWD:/go/src/github.com/kylos101/prusa-connect-pi-camera -v /var/run/docker.sock:/var/run/docker.sock -w /go/src/github.com/kylos101/prusa-connect-pi-camera -e GITHUB_TOKEN goreleaser/goreleaser release
```