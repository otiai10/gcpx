```sh
% cd $GOPATH/src/github.com/otiai10/gcpx/gaex/cloudsqlx
% goapp serve ./
% appcfg.py -E DATABASE_URI:"root@cloudsql({YOUR_PROJECT_ID}:{YOUR_CLOUDSQL_INSTANCE_NAME})/{YOUR_DATABASE_NAME}" -A {YOUR_APP_NAME} -V v1 update ./
# または、app.yamlに書くことも可能
```
