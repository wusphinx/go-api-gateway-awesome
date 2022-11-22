# [easegress](https://megaease.com/zh/easegress/)

## start

```
$ docker exec  easegress_gateway_1  /init.sh
```

### 测试
```
$ curl  http://127.0.0.1:10080/one/pine
{"message":"pong"}
```