# gcp-secret-version-cleaner

command to delete gcp secret manager's secret version with filter

## Install

```
go install github.com/notomo/gcp-secret-version-cleaner@latest
```

## Example

- destroy except latest
```
gcp-secret-version-cleaner --project=${PROJECT_NAME} --secret-name=${SECRET_NAME} destroy --keep-recent-count=1 --filter=''
```
- filter by field (ref. [https://cloud.google.com/secret-manager/docs/filtering](https://cloud.google.com/secret-manager/docs/filtering) )
```
gcp-secret-version-cleaner --project=${PROJECT_NAME} --secret-name=${SECRET_NAME} destroy --filter='createTime<2020-01-01'
```
- dry run destroy
```
gcp-secret-version-cleaner --project=${PROJECT_NAME} --secret-name=${SECRET_NAME} destroy --dry-run
```
- dry run disable
```
gcp-secret-version-cleaner --project=${PROJECT_NAME} --secret-name=${SECRET_NAME} disable --dry-run
```
