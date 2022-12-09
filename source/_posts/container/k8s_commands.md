---
title: k8s 常用命令笔记
date: 2018-04-13 15:14:11
toc: true
# img: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - k8s
  - container
categories:
  - container
  - k8s
---


### k8s 常用命令
```
kubectl version
kubectl cluster-info
kubectl get nodes

kubectl run kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1 --port=8080
# 当完成run的时候就相当于部署了一个应用
kubectl get deployments
# expose 后相当于  service
kubectl expose deployment hello-node --type=LoadBalancer
kubectl get services

kubectl describe pods
kubectl get pods -o wide

kubectl logs $POD_NAME
kubectl exec -ti $POD_NAME bash
kubectl exec $POD_NAME env
kubectl describe services/kubernetes-bootcamp
export NODE_PORT=$(kubectl get services/kubernetes-bootcamp -o go-template='{{(index .spec.ports 0).nodePort}}')
echo NODE_PORT=$NODE_PORT
kubectl scale deployments/kubernetes-bootcamp --replicas=2
kubectl set image deployments/kubernetes-bootcamp kubernetes-bootcamp=jocatalin/kubernetes-bootcamp:v2

# 看是否更新完
kubectl rollout status deployments/kubernetes-bootcamp
kubectl rollout undo deployments/kubernetes-bootcamp
kubectl get events
kubectl config view
kubectl delete service hello-node
kubectl delete deployment hello-node

kubectl create -f https://k8s.io/docs/user-guide/walkthrough/pod-nginx-with-label.yaml
kubectl apply -f https://k8s.io/docs/user-guide/walkthrough//deployment-update.yaml

kubectl get pods -l app=nginx
kubectl delete pod -l app=nginx


kubectl create configmap example-redis-config --from-file=https://k8s.io/docs/tutorials/configuration/configmap/redis/redis-config
kubectl get configmap example-redis-config -o yaml

```

### 在一个k8s集群内部各个pod默认跟其他服务与pod是可见的

一个Pod可以包含多个应用， 但一个pod里面的应用应该是强绑定的，共享存储网络，跟一些容器信息
