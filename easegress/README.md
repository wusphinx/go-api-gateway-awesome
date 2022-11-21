# [easegress](https://megaease.com/zh/easegress/)

## start

```
$ echo '
kind: HTTPServer
name: server-demo
port: 10080
keepAlive: true
https: false
rules:
  - paths:
    - pathPrefix: /ping
      backend: pipeline-demo' | egctl object update
```

```
$ echo '
name: pipeline-demo
kind: Pipeline
flow:
  - filter: proxy
filters:
  - name: proxy
    kind: Proxy
    pools:
    - servers:
      - url: http://one:8080
      loadBalance:
        policy: roundRobin' | egctl object create
```

### 测试
```
$ curl  http://127.0.0.1:10080/ping
{"message":"pong"}
```