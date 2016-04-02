```sh
% cd $GOPATH/src/github.com/otiai10/gcpx/gaex/cloudsqlx
% goapp serve ./
% appcfg.py -E PROJECT_ID:{YOUR_PROJECT_ID} -E INSTANCE_NAME:{YOUR_CLOUDSQL_INSTANCE_NAME} -E DATABASE_NAME:{YOUR_DATABASE_NAME} -A {YOUR_APP_NAME} -V v1 update ./
# または、app.yamlに書くことも可能
```
