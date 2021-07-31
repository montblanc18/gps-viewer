generate-swagger:
	swagger generate server -a gpsviewer -A gpsviewer --exclude-main --strict-additional-properties -t server/gen -f ./swagger.yaml
	swagger generate client -a gpsviewer -A gpsviewer --strict-additional-properties -t client/gen -f ./swagger.yaml
