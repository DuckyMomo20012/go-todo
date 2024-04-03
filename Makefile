.PHONY: openapi
openapi: openapi_http

.PHONY: openapi_http
openapi_http:
	oapi-codegen -generate fiber -o internal/tasks/ports/openapi_api.gen.go -package ports api/openapi/tasks.yaml
	oapi-codegen -generate types -o internal/tasks/ports/openapi_types.gen.go -package ports api/openapi/tasks.yaml
