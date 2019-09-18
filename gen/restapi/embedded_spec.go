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
    "description": "A PoC for future use",
    "title": "A Property management application",
    "version": "1.0.0"
  },
  "paths": {
    "/properties": {
      "get": {
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the property addresses",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/property"
              }
            }
          },
          "500": {
            "description": "server error"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/properties/{id}": {
      "get": {
        "summary": "Gets a property by id.",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "property ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/property"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "property": {
      "type": "object",
      "properties": {
        "bathrooms": {
          "type": "number",
          "format": "float"
        },
        "bedrooms": {
          "type": "number",
          "format": "float"
        },
        "city": {
          "type": "string",
          "minLength": 1
        },
        "price": {
          "type": "string",
          "minLength": 1
        },
        "property_id": {
          "type": "string",
          "minLength": 1
        },
        "property_type": {
          "type": "string",
          "minLength": 1
        },
        "real_estate_type": {
          "type": "string",
          "minLength": 1
        },
        "state": {
          "type": "string",
          "minLength": 1
        },
        "street": {
          "type": "string",
          "minLength": 1
        },
        "street_number": {
          "type": "string",
          "minLength": 1
        },
        "street_suffix": {
          "type": "string",
          "minLength": 1
        },
        "zip": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
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
    "description": "A PoC for future use",
    "title": "A Property management application",
    "version": "1.0.0"
  },
  "paths": {
    "/properties": {
      "get": {
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the property addresses",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/property"
              }
            }
          },
          "500": {
            "description": "server error"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/properties/{id}": {
      "get": {
        "summary": "Gets a property by id.",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "property ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/property"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "property": {
      "type": "object",
      "properties": {
        "bathrooms": {
          "type": "number",
          "format": "float"
        },
        "bedrooms": {
          "type": "number",
          "format": "float"
        },
        "city": {
          "type": "string",
          "minLength": 1
        },
        "price": {
          "type": "string",
          "minLength": 1
        },
        "property_id": {
          "type": "string",
          "minLength": 1
        },
        "property_type": {
          "type": "string",
          "minLength": 1
        },
        "real_estate_type": {
          "type": "string",
          "minLength": 1
        },
        "state": {
          "type": "string",
          "minLength": 1
        },
        "street": {
          "type": "string",
          "minLength": 1
        },
        "street_number": {
          "type": "string",
          "minLength": 1
        },
        "street_suffix": {
          "type": "string",
          "minLength": 1
        },
        "zip": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
}
