# GPS Viewer

1. 

# Setup

## awscli

```bash
aws configure set aws_access_key_id dummy --profile local
aws configure set aws_secret_access_key dummy --profile local
aws configure set region ap-northeast-1 --profile local
```

## go-swagger

```bash
go get -u github.com/go-swagger/go-swagger/cmd/swagger@v0.27.0
```

### update swagger code

```bash
make generate-swagger
## server side
# swagger generate server -a gpsviewer -A gpsviewer --exclude-main --strict-additional-properties -t server/gen -f ./swagger.yml
## client side
# swagger generate client -a gpsviewer -A gpsviewer --strict-additional-properties -t client/gen -f ./swagger.yml
```

## docker 

```bash
docker-compose -f docker-compose.yml up -d
```
