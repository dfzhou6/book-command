# top
动态查看资源占用情况

大写M：按照内存倒序

大写P：按照CPU倒序

# ps aux
静态查看资源占用情况

ps aux | sort -k4nr | head -n10

按照进程资源占用信息中的第4列的数值倒序排序，查看前10位

# sort -k4nr
-k：表示第几列

-n：表示使用数值

-r：表示使用倒序

# lsof -i:3306
查看端口3306的进程信息

# ll /proc/[pid]/status
查看进程pid的进程信息，比如RSS物理内存

# top -p [pid]
查看进程[pid]的动态实时资源信息

# ps aux | grep [pid]
查看进程[pid]的静态资源信息


