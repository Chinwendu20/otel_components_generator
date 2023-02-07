GOCMD?= go
GOTEST=$(GOCMD) test
GO_ACC=go-acc
LINT=golangci-lint
IMPI=impi

GOOS := $(shell $(GOCMD) env GOOS)
GOARCH := $(shell $(GOCMD) env GOARCH)
GH := $(shell which gh)

.PHONY: test
test:
	cd exporters
	$(GOTEST) .
	cd extensions
	$(GOTEST) .
	cd processors
	$(GOTEST) .
	cd receivers
	$(GOTEST) .
	cd ..
	$(GOTEST) .

.PHONY: fmt
fmt:
	cd extensions
	gofmt -w -s .
	goimports -w  -local github.com/Chinwendu20/otel_components_generator/extensions .
	cd receivers
	gofmt -w -s .
	goimports -w  -local github.com/Chinwendu20/otel_components_generator/receivers .
	cd exporters
	gofmt -w -s .
	goimports -w  -local github.com/Chinwendu20/otel_components_generator/exporters .
	cd processors
	gofmt -w -s .
	goimports -w  -local github.com/Chinwendu20/otel_components_generator/processors .
	cd ..
	gofmt -w -s .
	goimports -w  -local github.com/Chinwendu20/otel_components_generator .

.PHONY: tidy
tidy:
	cd exporters
	rm -fr go.sum
	$(GOCMD) mod tidy -compat=1.19
	cd extensions
	rm -fr go.sum
	$(GOCMD) mod tidy -compat=1.19
	cd processors
	rm -fr go.sum
	$(GOCMD) mod tidy -compat=1.19
	cd receivers
	rm -fr go.sum
	$(GOCMD) mod tidy -compat=1.19
	cd ..
	rm -fr go.sum
	$(GOCMD) mod tidy -compat=1.19

.PHONY: lint
lint:
	cd exporters
	$(LINT) run
	cd extensions
	$(LINT) run
	cd processors
	$(LINT) run
	cd receivers
	$(LINT) run
	cd ..
	$(LINT) run
.PHONY: moddownload
moddownload:
	cd exporters
	$(GOCMD) mod download
	cd extensions
	$(GOCMD) mod download
	cd processors
	$(GOCMD) mod download
	cd receivers
	$(GOCMD) mod download
	cd ..
	$(GOCMD) mod download
