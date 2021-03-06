package api

func init() {
	Swagger.Add("legacy_legacy", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/legacy/legacy.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v0/events/data-collector": {
      "get": {
        "summary": "This is used by chef-server, it requests a GET /data-collector/v0 to check\nAutomate's status.\nWe proxy /data-collector/v0 to /api/v0/events/data-collector, so this is\nwhere we need to respond.\nSince this is for legacy-support only, we don't bother much about having\ngoogle.protobuf.Empty as argument.",
        "operationId": "Status",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.legacy.StatusResponse"
            }
          }
        },
        "tags": [
          "LegacyDataCollector"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.legacy.StatusResponse": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    }
  }
}
`)
}
