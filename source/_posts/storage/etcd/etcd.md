---
title: etcd的安装与配置
date: 2018-11-13 15:14:11
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/59736948-53c5e400-928e-11e9-9a25-c7b7d94b861f.png
tags: 
    - etcd
    - database 
categories:
    - storage
    - etcd
---

## 说明
我们将在  10.1.19.214:2380,  10.1.15.229:2380,    10.1.13.165:2380 三个机器上分别创建三个etcd节点以组成集群

## 创建用户
安装比较简单， 但为了安全性起见， 推荐为etcd创建单独的用户

```bash
sudo groupadd etcd
sudo useradd -r -g etcd etcd
```

## 下载安装

### 下载
```bash
cd /opt
wget https://github.com/etcd-io/etcd/releases/download/v3.3.9/etcd-v3.3.9-linux-amd64.tar.gz
```


###安装
执行如下命令下载加压并创建好制定的文件与目录，设置好软链接等

```bash

cd /opt
tar xpf etcd-v3.3.9-linux-amd64.tar.gz
mv etcd-v3.3.9-linux-amd64 etcd-v3.3.9
ln -sf /opt/etcd-v3.3.9 /opt/etcd
cd etcd
mkdir data
mkdir wal
mkdir ssl
touch conf.yml

```

## 配置
执行如下命令

```bash
 vim conf.yml
```
修改内容 , 如下加粗的内容都要改


```cnf
# This is the configuration file for the etcd server.

# Human-readable name for this member.
name: 'etcd1'

# Path to the data directory.
data-dir: /opt/etcd/data

# Path to the dedicated wal directory.
wal-dir: /opt/etcd/wal

# Number of committed transactions to trigger a snapshot to disk.
snapshot-count: 10000

# Time (in milliseconds) of a heartbeat interval.
heartbeat-interval: 500

# Time (in milliseconds) for an election to timeout.
election-timeout: 1000

# Raise alarms when backend size exceeds the given quota. 0 means use the
# default quota.
quota-backend-bytes: 0

# List of comma separated URLs to listen on for peer traffic.

# List of comma separated URLs to listen on for client traffic.
listen-client-urls: http://10.111.8.11:2379,http://127.0.0.1:2379
listen-peer-urls: http://10.111.8.11:2380

# Maximum number of snapshot files to retain (0 is unlimited).
max-snapshots: 5

# Maximum number of wal files to retain (0 is unlimited).
max-wals: 5

# Comma-separated white list of origins for CORS (cross-origin resource sharing).
cors:

# List of this member's peer URLs to advertise to the rest of the cluster.
# The URLs needed to be a comma-separated list.
initial-advertise-peer-urls: http://10.111.8.11:2380

# List of this member's client URLs to advertise to the public.
# The URLs needed to be a comma-separated list.
advertise-client-urls: http://10.111.8.11:2379

# Discovery URL used to bootstrap the cluster.
discovery:

# Valid values include 'exit', 'proxy'
discovery-fallback: 'proxy'

# HTTP proxy to use for traffic to discovery service.
discovery-proxy:

# DNS domain used to bootstrap initial cluster.
discovery-srv:

# Initial cluster configuration for bootstrapping.
initial-cluster: etcd0=http://10.111.10.151:2380,etcd1=http://10.111.8.11:2380,etcd2=http://10.111.9.145:2380,etcd3=http://10.111.10.152:2380,etcd4=http://10.111.8.12:2380

# Initial cluster token for the etcd cluster during bootstrap.
initial-cluster-token: 'etcd-cluster'

# Initial cluster state ('new' or 'existing').
initial-cluster-state: 'new'

# Reject reconfiguration requests that would cause quorum loss.
strict-reconfig-check: false

# Accept etcd V2 client requests
enable-v2: true

# Enable runtime profiling data via HTTP server
enable-pprof: true

# Valid values include 'on', 'readonly', 'off'
proxy: 'off'

# Time (in milliseconds) an endpoint will be held in a failed state.
proxy-failure-wait: 5000

# Time (in milliseconds) of the endpoints refresh interval.
proxy-refresh-interval: 30000

# Time (in milliseconds) for a dial to timeout.
proxy-dial-timeout: 1000

# Time (in milliseconds) for a write to timeout.
proxy-write-timeout: 5000

# Time (in milliseconds) for a read to timeout.
proxy-read-timeout: 0

#client-transport-security:
 # Path to the client server TLS cert file.
# cert-file: 
 #/opt/etcd/ssl/etcd.pem

# Path to the client server TLS key file.
# key-file:
 # /opt/etcd/ssl/etcd-key.pem

# Enable client cert authentication.
# client-cert-auth: false

# Path to the client server TLS trusted CA cert file.
# trusted-ca-file:
 # /opt/etcd/ssl/etcd-root-ca.pem

# Client TLS using generated certificates
 # auto-tls:
 # false

# peer-transport-security:
 # Path to the peer server TLS cert file.
# cert-file:
 # /opt/etcd/ssl/etcd.pem

# Path to the peer server TLS key file.
# key-file:
 # /opt/etcd/ssl/etcd-key.pem

# Enable peer client cert authentication.
# client-cert-auth: false

# Path to the peer server TLS trusted CA cert file.
# trusted-ca-file:
 #/opt/etcd/ssl/etcd-root-ca.pem

# Peer TLS using generated certificates.
 # auto-tls:
 # false

# Enable debug-level logging for etcd.
debug: false

logger: zap

# Specify 'stdout' or 'stderr' to skip journald logging even when running under systemd.
log-outputs: [stderr]

# Force to create a new one member cluster.
force-new-cluster: false
auto-compaction-mode: periodic
auto-compaction-retention: "1"


```

注意： 里面关于集群的一些点， 和本机配置的一些点， 如果要配置tls/https 需要自行准备证书， 然后配置被如上注释了， 根据经验，如果配置了https， 再转到http需要清空数据才可以转， 否则是不能转的

## 最后的步骤

更改权限使得使用etcd用户来启动etcd节点

```bash
cd /opt/etcd
chown -R etcd:etcd /opt/etcd

```
配置systemd 的 unit文件并加入开机启动

vim /usr/lib/systemd/system/etcd.service  并贴入如下内容
```ini
[Unit]
Description=Etcd Server
After=network.target
After=network-online.target
Wants=network-online.target
Documentation=https://github.com/coreos
[Service]
User=etcd
Type=notify
WorkingDirectory=/opt/etcd/
ExecStart=/opt/etcd/etcd --config-file=/opt/etcd/conf.yml
Restart=on-failure
RestartSec=5
LimitNOFILE=65536
[Install]
WantedBy=multi-user.target
```

## 验证
为了方便etcdctl的使用

```
ln -sf /opt/etcd/etcdctl /usr/bin/etcdctl
```

etcd 多版本支持， 支持的有v2的API， 与V3 的API， 但V2与V3 是完全隔离的， V2 更像ZK， V3 更像KV 

### V2 验证集群健康状况
```
etcdctl cluster-health
```

如果健康则返回如下

```
member 42137edb694602d3 is healthy: got healthy result from http://10.1.15.229:2379
member b813a1f117f7f288 is healthy: got healthy result from http://10.1.13.165:2379
member e6992029de967f70 is healthy: got healthy result from http://10.1.19.214:2379

```

### V3 验证集群监控

在此之前需要先设置环境变量
```
export ETCDCTL_API=3
```
验证命令如下
```
etcdctl --endpoints=http://10.1.13.165:2379,http://10.1.15.229:2379,http://10.1.19.214:2379 endpoint health
```
如果健康则返回如下
```
http://10.1.13.165:2379 is healthy: successfully committed proposal: took = 1.605107ms
http://10.1.19.214:2379 is healthy: successfully committed proposal: took = 2.097534ms
http://10.1.15.229:2379 is healthy: successfully committed proposal: took = 2.983391ms
```


## 配置推荐
### CPU
etcd是可以利用多核的性能的

轻度使用	CPU
```
QPS < 2000	2-4core
```
中度使用
```
QPS > 5000	8-16core
```
### 内存
轻度使用	Mem
```
QPS < 2000	4-8G
```
中度使用
```
QPS > 5000	16-64G
```
### 硬盘
不需要做raid，raid0 就可以，  因为etcd就是高可用的， 推荐使用固态硬盘

### 网络
需要稳定可靠的网络， 网络不行的话，容易导致可用性比较低

如果要多数据中心部署， 尽量离近一些。



## Etcd的一些限制
### 请求大小限制
etcd多用于处理小的key-value对的元信息，大的请求虽然也可以工作， 但会增加延时，目前默认最大支持1MB的请求， 这个限制可以通过配置来更改

### 存储大小限制
默认的存储限制是 2GB, 可以用 --quota-backend-bytes flag 来配置;最大支持 8GB.

## https 
https 与http互转是有坑的， 可能需要用代码来数据迁移， data下面的与 wal 里面的清空， 通过调整配置可以互转
怀疑是这两个地方的一些文件里记录了之前一些tls，的一些设置， 导致互转不成。