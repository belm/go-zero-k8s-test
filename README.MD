# go-zero k8s test

```shell
goctl api go --dir . --api user.api --style go_zero
go mod init

goctl docker -go user.go
cp ./Dockerfile ../../

docker build  -t user-api:1.0 .

docker login  http://192.168.57.151:8077
docker tag  user-api:1.0  192.168.57.151:8077/go-zero-k8s-test/user-api:1.0
docker push 192.168.57.151:8077/go-zero-k8s-test/user-api:1.0

goctl kube deploy -secret docker-login -replicas 2 -nodePort 30888 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name user-api -namespace go-zero-k8s-test -image 192.168.57.151:8077/go-zero-k8s-test/user-api:1.0 -o user-api-deploy.yaml -port 8888 

```

vim /etc/docker/daemon.json
{"insecure-registries":["192.168.57.151:8077"] }

重启docker
systemctl restart  docker

harbor实际是docker-compose启动的， 到harbor安装目录可以通过docker-compose restart重启


## k8s拉取harbor私有仓库配置
harbor宿主机  登录harbor

```
docker login 192.168.57.151:8077
```
系统会在目录/root/.docker/config.json 生成授权凭证

cat /root/.docker/config.json
```
{
	"auths": {
		"192.168.57.151:8077": {
			"auth": "bHVtOkFiY2QxMjM0"
		}
	}
}
```

对秘钥文件进行base64加密
```
cat /root/.docker/config.json  | base64 -w 0
```

创建docker-secret.yaml
```
apiVersion: v1
kind: Secret
metadata:
  name: docker-login
type: kubernetes.io/dockerconfigjson
data:
  .dockerconfigjson: ewoJImF1dGhzIjogewoJCSIxOTIuMTY4LjU3LjE1MTo4MDc3IjogewoJCQkiYXV0aCI6ICJiSFZ0T2tGaVkyUXhNak0wIgoJCX0KCX0KfQ==
```

kubectl create -f docker-secret.yaml -n go-zero-looklook