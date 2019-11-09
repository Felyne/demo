
image=192.168.41.34/pub/kubectl
ver=v1
docker build -t $image:$ver .
docker push $image:$ver
docker run --rm 192.168.41.34/pub/kubectl:$ver kubectl get nodes
