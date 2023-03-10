include ./Makefile.Common
GOCMD?= go
GOTEST=$(GOCMD) test
GO_ACC=go-acc
LINT=golangci-lint
IMPI=impi

GOOS := $(shell $(GOCMD) env GOOS)
GOARCH := $(shell $(GOCMD) env GOARCH)
GH := $(shell which gh)

ALL_MODULES := $(shell find . -type f -name "go.mod" -exec dirname {} \; | sort | egrep  '^./' )

# All source code and documents. Used in spell check.
ALL_DOC := $(shell find . \( -name "*.md" -o -name "*.yaml" \) \
                                -type f | sort)

# Define a delegation target for each module
GOMODULES = $(ALL_MODULES)

.PHONY: $(GOMODULES)
$(GOMODULES):
	@echo "Running target '$(TARGET)' in module '$@'"
	$(MAKE) -C $@ $(TARGET)

# Triggers each module's delegation target
.PHONY: for-all-target
for-all-target: $(GOMODULES)

.PHONY: gotest
gotest:
	@$(MAKE) for-all-target TARGET="test"

.PHONY: gomoddownload
gomoddownload:
	@$(MAKE) for-all-target TARGET="moddownload"

.PHONY: golint
golint:
	@$(MAKE) for-all-target TARGET="lint"

.PHONY: goimpi
goimpi:
	@$(MAKE) for-all-target TARGET="impi"

.PHONY: gofmt
gofmt:
	@$(MAKE) for-all-target TARGET="fmt"

.PHONY: misspell
misspell: $(MISSPELL)
	$(MISSPELL) -error $(ALL_DOC)

.PHONY: misspell-correction
misspell-correction: $(MISSPELL)
	$(MISSPELL) -w $(ALL_DOC)

.PHONY: crosslink
crosslink: $(CROSSLINK)
	@echo "Executing crosslink"
	$(CROSSLINK) --root=$(shell pwd) --prune

DEPENDABOT_PATH=".github/dependabot.yml"
.PHONY: internal-gendependabot
internal-gendependabot:
	@echo "Add rule for \"${PACKAGE}\" in \"${DIR}\"";
	@echo "  - package-ecosystem: \"${PACKAGE}\"" >> ${DEPENDABOT_PATH};
	@echo "    directory: \"${DIR}\"" >> ${DEPENDABOT_PATH};
	@echo "    schedule:" >> ${DEPENDABOT_PATH};
	@echo "      interval: \"weekly\"" >> ${DEPENDABOT_PATH};

# This target should run on /bin/bash since the syntax DIR=$${dir:1} is not supported by /bin/sh.
.PHONY: gendependabot
gendependabot: $(eval SHELL:=/bin/bash)
	@echo "Recreating ${DEPENDABOT_PATH} file"
	@echo "# File generated by \"make gendependabot\"; DO NOT EDIT." > ${DEPENDABOT_PATH}
	@echo "" >> ${DEPENDABOT_PATH}
	@echo "version: 2" >> ${DEPENDABOT_PATH}
	@echo "updates:" >> ${DEPENDABOT_PATH}
	$(MAKE) internal-gendependabot DIR="/" PACKAGE="github-actions"
	$(MAKE) internal-gendependabot DIR="/" PACKAGE="gomod"
	@set -e; for dir in $(ALL_MODULES); do \
		$(MAKE) internal-gendependabot DIR=$${dir:1} PACKAGE="gomod"; \
	done

.PHONY: gotest-with-cover
gotest-with-cover: $(GOCOVMERGE)
	@$(MAKE) for-all-target TARGET="test-with-cover"
	$(GOCOVMERGE) $$(find . -name coverage.out) > coverage.txt

# Build the ocg executable.
.PHONY: ocg
ocg:
	# pushd main/ && GO111MODULE=auto $(GOCMD) build -trimpath -o ../../bin/ocg_$(GOOS)_$(GOARCH) . && popd
	pushd main && $(GOCMD) build -trimpath -o ../../bin/ocg_$(GOOS)_$(GOARCH) && popd

VERSIONYAML_PATH="./versions.yaml"
.PHONY: internal-genversionyml
internal-genversionyml:
	@echo    "          - github/Chinwendu20/otel_components_generator${DIR}" >> ${VERSIONYAML_PATH};

# This target should run on /bin/bash since the syntax DIR=$${dir:1} is not supported by /bin/sh.
INCLUDE_MODULES := $(filter-out ./internal/tools, $(ALL_MODULES))
VERSION?="dev"
.PHONY: genversionyml
genversionyml: $(eval SHELL:=/bin/bash)
	@echo "Recreating ${VERSIONYAML_PATH} file"
	@echo "# File generated by \"make genversionyml\"; DO NOT EDIT." > ${VERSIONYAML_PATH}
	@echo "" >> ${VERSIONYAML_PATH}
	@echo "module-sets:" >> ${VERSIONYAML_PATH}
	@echo "  stable:" >> ${VERSIONYAML_PATH}
	@echo "    version:${VERSION}:" >> ${VERSIONYAML_PATH}
	@echo "    modules:" >> ${VERSIONYAML_PATH}
	@set -e; for dir in $(INCLUDE_MODULES); do \
		$(MAKE) internal-genversionyml DIR=$${dir:1}; \
	done
	@echo "">>${VERSIONYAML_PATH}
	@echo "excluded-modules:" >>${VERSIONYAML_PATH}
	@echo "          - github/Chinwendu20/otel_components_generator/internal/tools" >>${VERSIONYAML_PATH};
.PHONY: multimod-verify
multimod-verify: $(MULTIMOD)
	@echo "Validating versions.yaml"
	$(MULTIMOD) verify

MODSET?=stable
.PHONY: multimod-prerelease
multimod-prerelease: $(MULTIMOD)
	$(MULTIMOD) prerelease -s=true -b=false -v ./versions.yaml -m ${MODSET}
	$(MAKE) gotidy

FILENAME?=$(shell git branch --show-current).yaml
.PHONY: chlog-new
chlog-new: $(CHLOG)
	$(CHLOG) new --filename $(FILENAME)

.PHONY: chlog-validate
chlog-validate: $(CHLOG)
	$(CHLOG) validate

.PHONY: chlog-preview
chlog-preview: $(CHLOG)
	$(CHLOG) update --dry

.PHONY: chlog-update
chlog-update: $(CHLOG)
	$(CHLOG) update --version $(VERSION)

.PHONY: gotidy
gotidy:
	@$(MAKE) for-all-target TARGET="tidy"
