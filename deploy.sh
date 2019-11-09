namespace=$1
version=$2
name=bmk-demo
image=192.168.41.34/devops/bmk-demo
config_dir=dev

if [ "$namespace" != "dev" ]; then
  config_dir='test'
fi
sed -ie "s/VERSION/$version/g" deployment.yaml
sed -ie "s/NAMESPACE/$namespace/g" deployment.yaml
sed -ie "s/NAMESPACE/$namespace/g" service.yaml
# cat deployment.yaml
kubectl delete configmap $name -n $namespace >/dev/null 2>&1 || true
kubectl create configmap $name --from-file=config/$config_dir/ -n $namespace
kubectl get services/$name -n $namespace || kubectl apply -f service.yaml
kubectl get deployments/$name -n $namespace || kubectl apply -f deployment.yaml
kubectl set image deployments/$name $name="$image:$version" -n $namespace
