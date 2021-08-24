generate-swagger:
	rm -rf ./server/gen/models
	rm -rf ./server/gen/restapi/gpsviewer
	rm -rf ./client/gen/models
	rm -rf ./client/gen/client
	swagger generate server -a gpsviewer -A gpsviewer --exclude-main --strict-additional-properties -t server/gen -f ./swagger.yml
	swagger generate client -a gpsviewer -A gpsviewer --strict-additional-properties -t client/gen -f ./swagger.yml
