# go-zero k8s test

```shell

goctl api go --api user.api --dir .

goctl rpc protoc user.proto --go_out=./ --go-grpc_out=./  --zrpc_out=./

go mod tidy

cd user/api
goctl docker -go user.go
mv Dockerfile ../../Dockerfile-user-api

cd user/rpc
goctl docker -go user.go
mv Dockerfile ../../Dockerfile-user-rpc

```
这里的Dockerfile稍微修改了一些配置

## coding 
* 导入github仓库
* 分别新建user-rpc和user-api制品仓库
* 分别新建构建计划  构建代码到制品仓库
* 制品仓库 配置访问令牌

## 安装 KubeSphere

### 1. 登录节点
```shell
ssh root@x.x.x.x
```

### 2.安装
```shell

#安装KubeSphere
wget https://nimoc.gitee.io/be/ops/k8s_file/ks/v3.1.1/kubesphere-installer.yaml
kubectl apply -f kubesphere-installer.yaml

## 下载集群配置 
wget https://nimoc.gitee.io/be/ops/k8s_file/ks/v3.1.1/cluster-configuration.yaml
kubectl apply -f cluster-configuration.yaml

# 查看安装情况
kubectl logs -n kubesphere-system $(kubectl get pod -n kubesphere-system -l app=ks-install -o jsonpath='{.items[0].metadata.name}') -f

# 访问管理页面
http://{你的节点公网IP}:30880
# 用户名 admin 密码 P@88w0rd
```
这个过程有点长，大概需要几分钟

## 安装etcd
 由于项目需要依赖etcd ，需要提前安装， 为了测试方便， 直接安装在k8s集群里面
 这里也可以安装redis, mysql等，也可以直接使用云厂商提供的redis和mysql, 配置里面修改对应的ip即可

我这里直接使用KubeSphere启动redis，并且新建服务， redis-svc 监听6379端口

etcd配置环境变量  ALLOW_NONE_AUTHENTICATION=yes
