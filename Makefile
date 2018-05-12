METALINTER_CONF_FILE := "/go/config/gometalinter/config.json"

env:
	@docker run -it --rm \
		-v $(shell pwd):/go/src/github.com/uniplaces/logfairy:consistent \
		-v ~/.ssh/github:/root/.ssh/github \
		-w /go/src/github.com/uniplaces/logfairy \
		uniplaces/uniplaces-go:1.10.0 /bin/sh
.PHONY: env # Enter the environment

test:
	@docker run -it --rm \
		-v $(shell pwd):/go/src/github.com/uniplaces/logfairy:consistent \
		-w /go/src/github.com/uniplaces/logfairy \
		uniplaces/uniplaces-go:1.10.0 go test ./... -covermode=atomic
.PHONY: test # Run all tests
