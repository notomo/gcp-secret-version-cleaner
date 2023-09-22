test:
	go test -v ./...

lint:
	go vet ./...

LOG_DIR:=/tmp/gcp-secret-version-cleaner
PROJECT:=example
SECRET_NAME:=example
_execute:
	rm -rf ${LOG_DIR}
	go run main.go --project=${PROJECT} --secret-name=${SECRET_NAME} --log-dir=${LOG_DIR} ${CLEANER_ARGS}
destory_dry_run:
	$(MAKE) _execute CLEANER_ARGS="destory --dry-run"
destory_run:
	$(MAKE) _execute CLEANER_ARGS="destory"
