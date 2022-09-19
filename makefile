.PHONY: gen_mocks
gen_mocks:
	sh scripts/gen_mocks.sh

.PHONY: test_all
test_all:
	ginkgo -r --race -cover --randomize-all --randomize-suites --trace --progress -skip-package=vendor
