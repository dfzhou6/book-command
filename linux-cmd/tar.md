# tar

```
# 压缩日志文件zdf.log
tar -zcvf log.tar.gz zdf.log 
```

```
# 解压缩文件到指定目录
tar -zxvf log.tar.gz -C /tmp/zdf
```
备注：-z 是指gzip，-c是指压缩，-x是指解压缩，-v是指显示过程，-f是指定压缩后的文件名