# GPS Viewer

1. 

# Setup

## go-swagger

```bash
go get -u github.com/go-swagger/go-swagger/cmd/swagger@v0.27.0
```

### update swagger

```bash
# server side
swagger generate server -a gpsviewer -A gpsviewer --exclude-main --strict-additional-properties -t server/gen -f ./swagger.yaml
# client side
swagger generate client -a gpsviewer -A gpsviewer --strict-additional-properties -t client/gen -f ./swagger.yaml
```