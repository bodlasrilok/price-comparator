#!/bin/bash

export NOW=$(shell date "+%Y-%m-%d")
export PKGS=$(shell go list ./... | grep -v vendor/ | grep -v cmd/ | grep -v model/ | grep -v pkg/account | grep -v usecase$$ | grep -v repository\/\\w*$$)
export TEST_OPTS=-cover -race
export MIN_COVER=30 #minimum coverage percentage

# need to improve this in the future
gotest:
	@go test -v -cover -race ./...
	
# go build command
gobuildpricecomparator:
	@go build -v -o ./build/price-comparator ./cmd/price-comparator/*.go

# go run command
gorunpricecomparator:
	make gobuildpricecomparator
	@./build/price-comparator -log_info files/var/log/tokopedia/price-comparator.log -log_debug files/var/log/tokopedia/price-comparator.debug.log

lint:
	@golangci-lint run ./... --skip-dirs=vendor --disable=errcheck --enable=gosec --fast --exclude=G401,G501,G502,G503,G504,G505

vet:
	@go vet ${PKGS} 2>&1

test:
	@echo "${NOW} == TESTING..."
	@go test ${TEST_OPTS} ${PKGS} | tee .dev/test.out
	@.dev/test.sh .dev/test.out ${MIN_COVER}

configure:
	@echo "CONFIGURING YOUR MACHINE FOR DEVELOPMENT ⚙️ ⚙️ ⚙️ "
	@echo "SETUP GIT PRE-COMMIT HOOK ️️☝️ "
	@cp .dev/pre-commit .git/hooks/
	@chmod +x .git/hooks/pre-commit
	@echo "SETUP GIT COMMIT MESSAGE CONVENTION ✌️"
	@git config commit.template .dev/.gitmessage
	@echo "Done, You are ready to Go 💯 💯 💯 " 