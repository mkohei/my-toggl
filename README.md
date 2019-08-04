# my-toggl

please make `conf.json`

```
{
  // for toggl
  "api_token" : "xxxxxxxxxxx",
  "email" : "test@mai.com",
  "workspace_name" : "Workspace"

  // for backlog
  "space_url" : "https://xxxxxxxxxx.backlog.jp",
  "apikey" : "xxxxxxxxxxx"

  // for check
  "backlog_project_keys" : [
    "PROJECT", "XXX"
  ],
  "toggl_projects_required_backlog_issue" : [
    "Togglプロジェクト", "YYY"
  ]
}
```

## check

```
$ go run main.go 2019-06
```