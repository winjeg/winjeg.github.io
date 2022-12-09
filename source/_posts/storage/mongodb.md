---
title: MongoDb 草记
date: 2016-07-13 15:14:11
toc: true
# img: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags: 
    - mongodb
    - database 
categories:
    - storage
    - mongodb
---

备份
```bash
mongodump -h dbhost -d dbname -o dbdirectory
```
恢复
```bash
sudo mongorestore -h host --db database --dir mongodata/database
```