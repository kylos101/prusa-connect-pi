# prusa_connect

## development

### generate an API client using docker

[Reference](https://openapi-generator.tech/docs/installation#docker)

```bash
sudo docker run --rm   -v ${PWD}:/local openapitools/openapi-generator-cli generate   -i /local/specs/prusaconnect.0.22.0.yaml   -g go   -o /local/client
```