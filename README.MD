# 1.启动Docker服务（守护进程）
```bash
# Linux
sudo systemctl start docekr         # 启动服务
sudo systemctl enable docker        # 设置开机自动启动
sudo systemctl status docekr        # 查看服务状态

# Windows/macOS
Docker Desktop                      # 手动从App启动
```

***

# 2.启动一个已经存在的容器
如果要启动一个已经停止的**Docker容器**(需要知道容器名称或者ID):
```bash
docker start [容器名或ID]           # 启动容器
docker start -i [容器名或ID]        # 启动并进入交互模式（类似docker attach）
```

***

# 3.创建并启动新容器
如果要从镜像创建并启动一个**新容器**(常用`docker run`):
```bash
docker run [选项] [镜像名]          # 例如：docker run -it ubuntu /bin/bash
常用选项：
  -d      后台运行（守护模式）
  -it     交互模式（进入容器终端）
  --name  指定容器名称
  -p      端口映射（如 -p 80:80）
  -v      挂载数据卷
```

# 4.重启容器
重启正在运行或者已经停止的容器
```bash
docker restart [容器名或ID]
```
