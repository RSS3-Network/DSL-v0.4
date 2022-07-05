VERSION=$(shell git rev-parse --short HEAD)

initialize:
	go mod tidy

.PHONY: build
build:
	mkdir -p ./build
	go build \
    	-ldflags "-w -s -X github.com/naturalselectionlabs/pregod/common/constant.Version=$(VERSION)" \
    	-o ./build/pregod_hub ./service/hub/cmd/main.go
	go build \
    	-ldflags "-w -s -X github.com/naturalselectionlabs/pregod/common/constant.Version=$(VERSION)" \
		-o ./build/pregod_indexer ./service/indexer/cmd/main.go

# fail coverage report under xx percent
DIR_TEST = .test
FILE_COVERAGE = $(DIR_TEST)/coverage.txt
COVERAGE_FAIL_UNDER ?= 45

TEST_FLAGS := \
	-race -failfast -p=1 \
	-covermode=atomic \

$(DIR_TEST):
	mkdir -p $(DIR_TEST)

test: $(DIR_TEST)
	for package in $$(go list ./... | grep -v gen); do \
		coverfile="$(DIR_TEST)/$$(echo $$package | tr / -).cover"; \
		go test $(TEST_FLAGS) -coverprofile="$$coverfile" -coverpkg=./... "$$package" || exit 1; \
		done
	gocovmerge $(DIR_TEST)/*.cover > $(FILE_COVERAGE)
	# remove auto generate code
	sed -i '/\/worker\/token\/contract\/erc20\.go/d' $(FILE_COVERAGE)
	sed -i '/indexer\/internal\/worker\/crossbell\/contract/d' $(FILE_COVERAGE)
	sed -i '/internal\/datasource\/moralis\/mrc20/d' $(FILE_COVERAGE)

# Open up a browser to view coverage report
coverage: test
	go tool cover -html=$(FILE_COVERAGE)

coverage_fail: test
	go tool cover -func $(FILE_COVERAGE) | tee $(DIR_TEST)/cover.txt
	test `tail -1 $(DIR_TEST)/cover.txt | cut -d\) -f2 | tr -d '[:space:]|%' | cut -d. -f1` -ge $(COVERAGE_FAIL_UNDER)

# lint
lint:
	golangci-lint run --timeout=10m --skip-dirs=gen
