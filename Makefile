export GO111MODULE := on

.PHONY: test
test: vet
	go test -race -cover ./...

.PHONY: vet
vet: lint
	go vet ./...

.PHONY: lint
lint: fmt
	golint -set_exit_status ./dcos/...

.PHONY: fmt
fmt:
	gofmt -w -s .

.PHONY: generate
generate: generate-client fmt

define run_generator
docker run -u $(shell id -u):$(shell id -g) \
	-v $(CURDIR):/local -w /local \
	openapitools/openapi-generator-cli:v4.0.0-beta2 \
	generate -i openapi/dcos.yaml -g go -o dcos \
	-t templates \
	-DpackageName=dcos -DwithGoCodegenComment=true -Dmodels -Dapis \
	$(1)
endef

.PHONY: generate-client
generate-client:
	@$(RM) -rf $(shell grep -l 'Code generated by OpenAPI Generator' dcos/*.go) \
		dcos/api dcos/docs
	$(call run_generator,-DsupportingFiles=client.go)
	$(call run_generator,-DsupportingFiles=response.go)
	$(call run_generator,-DsupportingFiles=configuration.go)
	$(call run_generator,-DsupportingFiles=README.md)
	cp templates/model_edgelb_v2_frontend.go.tmpl dcos/model_edgelb_v2_frontend.go

.PHONY: validate
validate: generate
	git diff
	@export CHANGED_FILES="$$(git status --porcelain)" && \
		[ -z "$${CHANGED_FILES}" ] || \
		([ -n "$${CHANGED_FILES}" ] && printf "\nValidation failed: the following files have changed when running generate:\n\n$${CHANGED_FILES}\n\n" && exit 1)
