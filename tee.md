# tee
```
# 往zdf.log文件追加aaa和bbb，EOF是结束输入符号
tee -a /home/zdf.log <<'EOF'
> aaa
> bbb
> EOF
```