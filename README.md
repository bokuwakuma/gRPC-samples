#タスク管理マイクロサービス

## 動作環境
- Go 1.12
- grpc-go 1.19.1
- protobuf 3.7.0
- protoc-gen-go 1.3.1
- Docker
  - Docker Desktop 2.0.0.3
  - Docker CE 18.09.3 ＋ Docker Compose 1.23.2

## 起動方法

### .protoファイルをコンパイルする

```
$ protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/activity/activity.proto
$ protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/task/task.proto
$ protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/user/user.proto
$ protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/project/project.proto
```

### 各サービスをビルドする

```
$ docker-compose build
```

### 各サービスを起動する

```
$ docker-compose up
```