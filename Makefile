launch_args=
test_args=-coverprofile cover.out && go tool cover -func cover.out
cover_args=-cover -coverprofile=cover.out `go list ./...` && go tool cover -html=cover.out

SERVICE_NAME=nexus-service
VERSION?= $(shell git describe --match 'v[0-9]*' --tags --always)
DOCKER_IMAGE_NAME=krobus00/${SERVICE_NAME}
CONFIG?=./config.yml
NAMESPACE?=default
PACKAGE_NAME=github.com/krobus00/${SERVICE_NAME}

# make tidy
tidy:
	go mod tidy

generate:
	go get github.com/99designs/gqlgen@v0.17.26
	go generate ./...

# make lint
lint:
	golangci-lint run

# make run dev server
# make run dev worker
# make run server
# make run worker
# make run migration
# make run migration MIGRATION_ACTION=up
# make run migration MIGRATION_ACTION=create MIGRATION_NAME=create_table_products
# make run migration MIGRATION_ACTION=up MIGRATION_STEP=1
run:
ifeq (dev server, $(filter dev server,$(MAKECMDGOALS)))
	$(eval launch_args=server $(launch_args))
	air --build.cmd 'go build -ldflags "-s -w -X $(PACKAGE_NAME)/internal/config.serviceVersion=$(VERSION) -X $(PACKAGE_NAME)/internal/config.serviceName=$(SERVICE_NAME)" -o bin/nexus-service main.go' --build.bin "./bin/nexus-service $(launch_args)"
else ifeq (dev worker, $(filter dev worker,$(MAKECMDGOALS)))
	$(eval launch_args=worker $(launch_args))
	air --build.cmd 'go build -ldflags "-s -w -X $(PACKAGE_NAME)/internal/config.serviceVersion=$(VERSION) -X $(PACKAGE_NAME)/internal/config.serviceName=$(SERVICE_NAME)" -o bin/nexus-service main.go' --build.bin "./bin/nexus-service $(launch_args)"
else ifeq (worker, $(filter worker,$(MAKECMDGOALS)))
	$(eval launch_args=worker $(launch_args))
	$(shell if test -s ./bin/nexus-service; then ./bin/nexus-service $(launch_args); else echo nexus binary not found; fi)
else ifeq (server, $(filter server,$(MAKECMDGOALS)))
	$(eval launch_args=server $(launch_args))
	$(shell if test -s ./bin/nexus-service; then ./bin/nexus-service $(launch_args); else echo nexus binary not found; fi)
endif

# make build
build:
	# build binary file
	go build -ldflags "-s -w -X $(PACKAGE_NAME)/internal/config.serviceVersion=$(VERSION) -X $(PACKAGE_NAME)/internal/config.serviceName=$(SERVICE_NAME)" -o ./bin/nexus-service ./main.go
ifeq (, $(shell which upx))
	$(warning "upx not installed")
else
	# compress binary file if upx command exist
	upx -9 ./bin/nexus-service
endif

# make image VERSION="vx.x.x"
image:
	docker build -t ${DOCKER_IMAGE_NAME}:${VERSION} . -f ./deployments/Dockerfile

# make deploy VERSION="vx.x.x"
# make deploy VERSION="vx.x.x" NAMESPACE="staging"
# make deploy VERSION="vx.x.x" NAMESPACE="staging" CONFIG="./config-staging.yml"
deploy:
	helm upgrade --install nexus-service ./deployments/helm/server-nexus-service --set-file appConfig="${CONFIG}" --set app.container.version="${VERSION}" -n ${NAMESPACE}

# make test
test:
ifeq (, $(shell which richgo))
	go test ./... $(test_args)
else
	richgo test ./... $(test_args)
endif

# make cover
cover: test
ifeq (, $(shell which richgo))
	go test $(cover_args)
else
	richgo test $(cover_args)
endif

# make changelog VERSION=vx.x.x
changelog: tidy generate lint
	git-chglog -o CHANGELOG.md --next-tag $(VERSION)

%:
	@:
