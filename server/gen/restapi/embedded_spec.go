// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "GPS Viewer",
    "version": "0.0.1"
  },
  "basePath": "/gps-viewer/v1",
  "paths": {
    "/gqs/{deviceId}": {
      "get": {
        "description": "Return device gps info",
        "produces": [
          "application/json; charset=utf-8"
        ],
        "tags": [
          "gps"
        ],
        "operationId": "getGpsByDeviceId",
        "parameters": [
          {
            "type": "string",
            "x-example": 99999,
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operationId",
            "schema": {
              "$ref": "#/definitions/DeviceGPS"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "当該データなし",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Register a device GPS by ID",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json; charset=utf-8"
        ],
        "tags": [
          "gps"
        ],
        "operationId": "registerGpsByDeviceId",
        "parameters": [
          {
            "type": "string",
            "x-example": 99999,
            "name": "deviceId",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/registerDeviceGPS"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/DeviceGPS"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "DeviceGPS": {
      "type": "object",
      "required": [
        "deviceId",
        "deviceType",
        "lat",
        "lng",
        "recordedAt"
      ],
      "properties": {
        "deviceId": {
          "type": "string",
          "example": 99999
        },
        "deviceType": {
          "type": "string",
          "example": "microcomputer"
        },
        "lat": {
          "type": "number",
          "format": "double",
          "example": 34.69139
        },
        "lng": {
          "type": "number",
          "format": "double",
          "example": 135.18306
        },
        "recordedAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "description": "エラーコード",
          "type": "integer"
        },
        "message": {
          "description": "エラーメッセージ",
          "type": "string"
        }
      }
    },
    "registerDeviceGPS": {
      "type": "object",
      "required": [
        "deviceId",
        "deviceType",
        "lat",
        "lng",
        "recordedAt"
      ],
      "properties": {
        "deviceId": {
          "type": "string",
          "example": 99999
        },
        "deviceType": {
          "type": "string",
          "example": "microcomputer"
        },
        "lat": {
          "type": "number",
          "format": "double",
          "example": 34.69139
        },
        "lng": {
          "type": "number",
          "format": "double",
          "example": 135.18306
        },
        "recordedAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    }
  },
  "tags": [
    {
      "description": "GPS情報",
      "name": "gps"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "GPS Viewer",
    "version": "0.0.1"
  },
  "basePath": "/gps-viewer/v1",
  "paths": {
    "/gqs/{deviceId}": {
      "get": {
        "description": "Return device gps info",
        "produces": [
          "application/json; charset=utf-8"
        ],
        "tags": [
          "gps"
        ],
        "operationId": "getGpsByDeviceId",
        "parameters": [
          {
            "type": "string",
            "x-example": 99999,
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operationId",
            "schema": {
              "$ref": "#/definitions/DeviceGPS"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "当該データなし",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Register a device GPS by ID",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json; charset=utf-8"
        ],
        "tags": [
          "gps"
        ],
        "operationId": "registerGpsByDeviceId",
        "parameters": [
          {
            "type": "string",
            "x-example": 99999,
            "name": "deviceId",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/registerDeviceGPS"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/DeviceGPS"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "DeviceGPS": {
      "type": "object",
      "required": [
        "deviceId",
        "deviceType",
        "lat",
        "lng",
        "recordedAt"
      ],
      "properties": {
        "deviceId": {
          "type": "string",
          "example": 99999
        },
        "deviceType": {
          "type": "string",
          "example": "microcomputer"
        },
        "lat": {
          "type": "number",
          "format": "double",
          "example": 34.69139
        },
        "lng": {
          "type": "number",
          "format": "double",
          "example": 135.18306
        },
        "recordedAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "description": "エラーコード",
          "type": "integer"
        },
        "message": {
          "description": "エラーメッセージ",
          "type": "string"
        }
      }
    },
    "registerDeviceGPS": {
      "type": "object",
      "required": [
        "deviceId",
        "deviceType",
        "lat",
        "lng",
        "recordedAt"
      ],
      "properties": {
        "deviceId": {
          "type": "string",
          "example": 99999
        },
        "deviceType": {
          "type": "string",
          "example": "microcomputer"
        },
        "lat": {
          "type": "number",
          "format": "double",
          "example": 34.69139
        },
        "lng": {
          "type": "number",
          "format": "double",
          "example": 135.18306
        },
        "recordedAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    }
  },
  "tags": [
    {
      "description": "GPS情報",
      "name": "gps"
    }
  ]
}`))
}
