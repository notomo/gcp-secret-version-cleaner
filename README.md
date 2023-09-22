# gcp-secret-version-cleaner

command to delete gcp secret manager's secret version with filter

## Example

- destory except latest
```
go run github.com/notomo/gcp-secret-version-cleaner@latest --project=${PROJECT_NAME} --secret-name=${SECRET_NAME} --keep-recent-count=1 --filter=''
```
- filter by field (ref. [https://cloud.google.com/secret-manager/docs/filtering](https://cloud.google.com/secret-manager/docs/filtering) )
```
go run github.com/notomo/gcp-secret-version-cleaner@latest --project=${PROJECT_NAME} --secret-name=${SECRET_NAME} --filter='createTime<2020-01-01'
```
- dry run destory
```
go run github.com/notomo/gcp-secret-version-cleaner@latest --project=${PROJECT_NAME} --secret-name=${SECRET_NAME} --dry-run
```
