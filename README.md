# gcp-secret-version-cleaner

command to delete gcp secret manager's secret version with filter

## Example

- destroy except latest
```
go run github.com/notomo/gcp-secret-version-cleaner@latest --project=${PROJECT_NAME} --secret-name=${SECRET_NAME} destroy --keep-recent-count=1 --filter=''
```
- filter by field (ref. [https://cloud.google.com/secret-manager/docs/filtering](https://cloud.google.com/secret-manager/docs/filtering) )
```
go run github.com/notomo/gcp-secret-version-cleaner@latest --project=${PROJECT_NAME} --secret-name=${SECRET_NAME} destroy --filter='createTime<2020-01-01'
```
- dry run destroy
```
go run github.com/notomo/gcp-secret-version-cleaner@latest --project=${PROJECT_NAME} --secret-name=${SECRET_NAME} destroy --dry-run
```
- dry run disable
```
go run github.com/notomo/gcp-secret-version-cleaner@latest --project=${PROJECT_NAME} --secret-name=${SECRET_NAME} disable --dry-run
```
