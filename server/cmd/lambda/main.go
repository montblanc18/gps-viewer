package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-openapi/loads"
	"github.com/montblanc18/gps-viewer/server/gen/restapi"
	"github.com/montblanc18/gps-viewer/server/gen/restapi/gpsviewer"
)

var httpAdapter *httpadapter.HandlerAdapter

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if httpAdapter == nil {
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			log.Fatalln(err)
		}

		api := gpsviewer.NewGpsviewerAPI(swaggerSpec)
		server := restapi.NewServer(api)
		server.ConfigureAPI()

		// see https://github.com/go-swagger/go-swagger/issues/962#issuecomment-478382896
		httpAdapter = httpadapter.New(server.GetHandler())
	}
	return httpAdapter.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
