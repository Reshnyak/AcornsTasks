package main

//go:generate oapi-codegen --config=oapi_server.config.yaml docs/swagger.yaml
//go:generate oapi-codegen --config=oapi_models.config.yaml docs/swagger.yaml

////go:generate node widdershins --search true --language_tabs 'go:Go' -c --summary internal/gateway/swagger.yaml -o swagger.html
