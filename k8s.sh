#!/bin/bash
# ConfigMap必须在Pod使用它之前创建
# 使用envFrom时，将会自动忽略无效的键
# Pod只能使用同一个命名空间的ConfigMap
# 默认namespace是default

# 本地测试redis
# kubectl port-forward --namespace default svc/redis-master 6379:6379

# 删除
# kubectl delete configmap bmk-demo 

# 上传除了redis.toml的配置文件
kubectl create configmap bmk-demo --from-file=config/dev/ -n dev

# 查看
# kubectl get configmap bmk-demo -o yaml -n dev

# redis.toml 由运维人员创建secret
kubectl create secret generic redis-toml --from-file=secret/dev/redis.toml -n dev

# 测试环境
test()
{   
    namespace="test"
    kubectl delete configmap bmk-demo -n $namespace
    kubectl create configmap bmk-demo --from-file=config/test/ -n $namespace
    kubectl create secret generic redis-toml --from-file=secret/test/redis.toml
}

# deployment.yaml是模板，先修改, 因为要把redis这个secret的redis.toml挂载为/app/config/redis.toml,其他配置文件不得不一个个单独挂载
kubectl apply -f deployment.yaml
