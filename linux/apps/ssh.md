# ssh 相关的知识
## SSH KEY 相互转换的方法
1. 确保你有公钥私钥对，确保安装了ssh 
如果没有可以用下面的命令生成 
ssh-keygen 
如果有私钥可以用下面的命令生成公钥 
```
ssh-keygen -y -f id_rsa > id_rsa.pub 
```
2. open ssh与  windows ppk相互转换
关键工具 `puttygen.exe`

## 免密码ssh登录的设置
1. 要把自己的公钥添加至目标机的
```
.ssh/authorized_keys 
```
文件中去，`authorized_keys` 的权限是 `600` 

2. ssh 连接 linux 自动补全需要bash_completion
还需要在 .ssh/config 文件中记录 host user ip三个字段

`ssh-copy-id` 会自动加免密到目标机器