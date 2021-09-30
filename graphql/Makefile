pre-deploy:
	go mod vendor

post-deploy:
	rm -rf ./vendor

build-http:
	go build -tags httplambda -mod=vendor -o /asset-output/http

build-ws:
	go build -tags wslambda -mod=vendor -o /asset-output/ws

build-event-steam:
	go build -tags eventstreamlambda -mod=vendor -o /asset-output/eventstreamlambda
