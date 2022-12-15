# todo-apps: Sample application for Kubernetes
- React.js+Golang API sample.

## After git cloning you should run
```
### api
$ cp api/.env{.sample,}
$ cp api/.env.docker{.sample,}

### ui
$ cp ui/.env{.sample,}

### db
$ cp db/.env{.sample,}
```

## Boot command
```
$ make
```

## Upload Docker image
```
### Dockerイメージを確認
$ docker images | grep todo-apps-
REPOSITORY                           TAG       IMAGE ID       CREATED         SIZE
todo-apps-ui                         latest    ab4704be53e8   6 minutes ago   123MB
todo-apps-api                        latest    fd738cde1aa6   8 minutes ago   18.3MB
todo-apps-db                         latest    50b222f7759d   42 hours ago    538MB

### イメージ名とTAG名を変更
$ docker tag ab4704be53e8 ren1007/todo-apps-ui:v1
$ docker tag fd738cde1aa6 ren1007/todo-apps-api:v1

### 再度 Dockerイメージを確認
$ docker images | grep todo-apps- | grep v1
REPOSITORY                           TAG       IMAGE ID       CREATED          SIZE
ren1007/todo-apps-ui                 v1        ab4704be53e8   12 minutes ago   123MB
ren1007/todo-apps-api                v1        fd738cde1aa6   14 minutes ago   18.3MB

### DockerHub にログイン
$ docker login

### イメージをアップロード
$ docker push ren1007/todo-apps-ui:v1
$ docker push ren1007/todo-apps-api:v1
```

## Compose v2
- https://zenn.dev/miroha/articles/whats-docker-compose-v2
- https://github.com/compose-spec/compose-spec/blob/master/spec.md#version-top-level-element
