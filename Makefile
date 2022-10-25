GITHUB_REPO=sentinelone-sdk-go
GITHUB_USER=rtk-logan-fillo

install:
	go mod tidy 

validate:
	npx @openapitools/openapi-generator-cli validate -i ./schema/sentinelone.yaml

generate:
	npx @openapitools/openapi-generator-cli generate \
		-i ./schema/sentinelone.yaml \
		-g go \
		-o ./sentinelone\
		--git-host github.com \
		--git-repo-id ${GITHUB_REPO} \
		--git-user-id ${GITHUB_USER} \
		--additional-properties=packageName=sentinelone,isGoSubmodule=true