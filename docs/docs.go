// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-08-22 12:50:02.1589442 +0700 ICT

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a documentation API.",
        "title": "Forex",
        "version": "1.0"
    },
    "basePath": "/forex",
	"schemes": [
		"http"
	],
	"consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/currency-exchange/:id": {
      "delete": {
        "description": "Delete Currency Exchange",
        "summary": "Delete Currency Exchange",
        "tags": [
          "Currency Exchange"
        ],
        "operationId": "CurrencyExchange1Delete2",
        "produces": [
          "application/json"
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "security": [],
        "x-unitTests": [
          {
            "request": {
              "method": "DELETE",
              "uri": "/currency-exchange/1",
              "headers": {
                "Content-Type": "application/x-www-form-urlencoded"
              },
              "body": "from=USD&to=IDR"
            },
            "expectedResponse": {
              "x-allowExtraHeaders": true,
              "x-bodyMatchMode": "NONE",
              "x-arrayOrderedMatching": false,
              "x-arrayCheckCount": false,
              "x-matchResponseSchema": true,
              "headers": {}
            },
            "x-testShouldPass": true,
            "x-testEnabled": true,
            "x-testName": "Delete Currency Exchange",
            "x-testDescription": "List Currency Exchange"
          }
        ],
        "x-operation-settings": {
          "CollectParameters": false,
          "AllowDynamicQueryParameters": false,
          "AllowDynamicFormParameters": false,
          "IsMultiContentStreaming": false
        }
      }
    },
    "/currency-exchange": {
      "get": {
        "description": "List Currency Exchange",
        "summary": "List Currency Exchange",
        "tags": [
          "Currency Exchange"
        ],
        "operationId": "CurrencyExchangeGet",
        "produces": [
          "application/json"
        ],
        "parameters": [],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "security": [],
        "x-unitTests": [
          {
            "request": {
              "method": "GET",
              "uri": "/currency-exchange"
            },
            "expectedResponse": {
              "x-allowExtraHeaders": true,
              "x-bodyMatchMode": "NONE",
              "x-arrayOrderedMatching": false,
              "x-arrayCheckCount": false,
              "x-matchResponseSchema": true,
              "headers": {}
            },
            "x-testShouldPass": true,
            "x-testEnabled": true,
            "x-testName": "List Currency Exchange",
            "x-testDescription": "List Currency Exchange"
          }
        ],
        "x-operation-settings": {
          "CollectParameters": false,
          "AllowDynamicQueryParameters": false,
          "AllowDynamicFormParameters": false,
          "IsMultiContentStreaming": false
        }
      },
      "post": {
        "description": "Create Currency Exchange",
        "summary": "Create Currency Exchange",
        "tags": [
          "Currency Exchange"
        ],
        "operationId": "CurrencyExchangePost",
        "produces": [
          "application/json"
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "parameters": [
          {
            "name": "from",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "to",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "security": [],
        "x-unitTests": [
          {
            "request": {
              "method": "POST",
              "uri": "/currency-exchange",
              "headers": {
                "Content-Type": "application/x-www-form-urlencoded"
              },
              "body": "from=USD&to=IDR"
            },
            "expectedResponse": {
              "x-allowExtraHeaders": true,
              "x-bodyMatchMode": "NONE",
              "x-arrayOrderedMatching": false,
              "x-arrayCheckCount": false,
              "x-matchResponseSchema": true,
              "headers": {}
            },
            "x-testShouldPass": true,
            "x-testEnabled": true,
            "x-testName": "Create Currency Exchange",
            "x-testDescription": "List Currency Exchange"
          }
        ],
        "x-operation-settings": {
          "CollectParameters": false,
          "AllowDynamicQueryParameters": false,
          "AllowDynamicFormParameters": false,
          "IsMultiContentStreaming": false
        }
      }
    },
    "/daily-currency-exchange/trends": {
      "get": {
        "description": "",
        "summary": "Get Currency Trend",
        "tags": [
          "Daily Currency Exchange"
        ],
        "operationId": "DailyCurrencyExchangeTrendsGet",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "from",
            "in": "query",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "to",
            "in": "query",
            "required": true,
            "type": "string",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "security": [],
        "x-unitTests": [
          {
            "request": {
              "method": "GET",
              "uri": "/daily-currency-exchange/trends?from=USD&to=IDR"
            },
            "expectedResponse": {
              "x-allowExtraHeaders": true,
              "x-bodyMatchMode": "NONE",
              "x-arrayOrderedMatching": false,
              "x-arrayCheckCount": false,
              "x-matchResponseSchema": true,
              "headers": {}
            },
            "x-testShouldPass": true,
            "x-testEnabled": true,
            "x-testName": "Get Currency Trend",
            "x-testDescription": ""
          }
        ],
        "x-operation-settings": {
          "CollectParameters": false,
          "AllowDynamicQueryParameters": false,
          "AllowDynamicFormParameters": false,
          "IsMultiContentStreaming": false
        }
      }
    },
    "/daily-currency-exchange/lists": {
      "get": {
        "description": "",
        "summary": "Get Currency By Date",
        "tags": [
          "Daily Currency Exchange"
        ],
        "operationId": "DailyCurrencyExchangeListsGet",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "date",
            "in": "query",
            "required": true,
            "type": "string",
            "format": "date-time",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "security": [],
        "x-unitTests": [
          {
            "request": {
              "method": "GET",
              "uri": "/daily-currency-exchange/lists?date=2018-07-01"
            },
            "expectedResponse": {
              "x-allowExtraHeaders": true,
              "x-bodyMatchMode": "NONE",
              "x-arrayOrderedMatching": false,
              "x-arrayCheckCount": false,
              "x-matchResponseSchema": true,
              "headers": {}
            },
            "x-testShouldPass": true,
            "x-testEnabled": true,
            "x-testName": "Get Currency By Date",
            "x-testDescription": ""
          }
        ],
        "x-operation-settings": {
          "CollectParameters": false,
          "AllowDynamicQueryParameters": false,
          "AllowDynamicFormParameters": false,
          "IsMultiContentStreaming": false
        }
      }
    },
    "/daily-currency-exchange": {
      "post": {
        "description": "",
        "summary": "Create Daily Currency Exchange",
        "tags": [
          "Daily Currency Exchange"
        ],
        "operationId": "DailyCurrencyExchangePost",
        "produces": [
          "application/json"
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "parameters": [
          {
            "name": "date",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "from",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "to",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "rate",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "created_by",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "security": [],
        "x-unitTests": [
          {
            "request": {
              "method": "POST",
              "uri": "/daily-currency-exchange",
              "headers": {
                "Content-Type": "application/x-www-form-urlencoded"
              },
              "body": "date=2018-07-01&from=USD&to=IDR&rate=0.75709&created_by=1"
            },
            "expectedResponse": {
              "x-allowExtraHeaders": true,
              "x-bodyMatchMode": "NONE",
              "x-arrayOrderedMatching": false,
              "x-arrayCheckCount": false,
              "x-matchResponseSchema": true,
              "headers": {}
            },
            "x-testShouldPass": true,
            "x-testEnabled": true,
            "x-testName": "Create Daily Currency Exchange",
            "x-testDescription": ""
          }
        ],
        "x-operation-settings": {
          "CollectParameters": false,
          "AllowDynamicQueryParameters": false,
          "AllowDynamicFormParameters": false,
          "IsMultiContentStreaming": false
        }
      }
    }
  }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
