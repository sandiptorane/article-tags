# test and code coverage
GO_TEST_PKG=
GO_COVER_PKG=
# atomic: Concurrent access to the same coverage counters is guaranteed to be
# executed one at a time, avoiding race conditions.
COVER_MODE=atomic # other options [count, set]
COVER_PROFILE=coverage.txt # coverage profile, write out file


## Generate mocks for all the interfaces
mock:
	rm -rf mocks
	mockery --all --keeptree --case underscore --with-expecter --exported

## Run all the tests by excluding auto-generated/conditional packages
unit-test:
	go test ./... -mod=readonly -cover -covermode=${COVER_MODE} -coverprofile=${COVER_PROFILE} -coverpkg=${GO_COVER_PKG}
