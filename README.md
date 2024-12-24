# prusa_connect

## development

### generate a client

```
sudo docker run --rm   -v ${PWD}:/local openapitools/openapi-generator-cli generate   -i /local/specs/prusaconnect.0.22.0.yaml   -g go   -o /local/client
```