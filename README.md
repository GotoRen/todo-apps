# todo-apps: Sample application for Kubernetes
- React.js+Golang API sample.

## After git cloning you should run
```
### api
$ cp api/.env.docker{.sample,}
$ cp api/.envrc{.sample,}

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
REPOSITORY      TAG       IMAGE ID       CREATED         SIZE
todo-apps-ui    latest    ef1d01f978b3   3 minutes ago   124MB
todo-apps-api   latest    6ecb4f9cf926   5 minutes ago   17.8MB
todo-apps-db    latest    3a5c0bb06443   11 days ago     538MB

### イメージ名とTAG名を変更
$ docker tag ef1d01f978b3 ren1007/todo-apps-ui:v1.0
$ docker tag 6ecb4f9cf926 ren1007/todo-apps-api:v1.0

### 再度 Dockerイメージを確認
$ docker images | grep todo-apps- | grep v1.0
REPOSITORY                           TAG       IMAGE ID       CREATED         SIZE
ren1007/todo-apps-ui                 v1.0      ef1d01f978b3   5 minutes ago   124MB
ren1007/todo-apps-api                v1.0      6ecb4f9cf926   7 minutes ago   17.8MB

### DockerHub にログイン
$ docker login

### イメージをアップロード
$ docker push ren1007/todo-apps-ui:v1.0
$ docker push ren1007/todo-apps-api:v1.0
```

## Compose v2
- https://zenn.dev/miroha/articles/whats-docker-compose-v2
- https://github.com/compose-spec/compose-spec/blob/master/spec.md#version-top-level-element
