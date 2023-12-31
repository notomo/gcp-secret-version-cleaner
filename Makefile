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
destroy_dry_run:
	$(MAKE) _execute CLEANER_ARGS="destroy --dry-run"
destroy_run:
	$(MAKE) _execute CLEANER_ARGS="destroy"
disable_dry_run:
	$(MAKE) _execute CLEANER_ARGS="disable --dry-run"
disable_run:
	$(MAKE) _execute CLEANER_ARGS="disable"
