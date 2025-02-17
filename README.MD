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

# 2.查看容器相关的命令
## 2.1 列出正在运行的容器
```bash
docker ps
```
输出示例:
```bash
CONTAINER ID   IMAGE          COMMAND       CREATED         STATUS         PORTS     NAMES
a1b2c3d4e5f6   nginx:latest   "nginx -g ..." 2 minutes ago   Up 2 minutes   80/tcp    webserver
```

## 2.2 列出所有容器
```bash
docker ps -a
# 或
docker ps --all
```
输出示例:
```bash
CONTAINER ID   IMAGE          COMMAND       CREATED         STATUS                      PORTS     NAMES
a1b2c3d4e5f6   nginx:latest   "nginx -g ..." 2 minutes ago   Exited (0) 1 minute ago              webserver
x9y8z7q6p5o4   ubuntu:20.04   "/bin/bash"    5 days ago      Created                               test-container
```
## 2.3 常用选项
* `-q`: 仅显示容器ID（适用于脚本操作）
```bash
docker ps -aq
```
* `--filter`: 按条件过滤（列入按状态、名称等）
```bash
docker ps -a --filter "status=exited"       # 列出已停止的容器
docker ps -a --filter "name=web"            # 列出名称包含 "web" 的容器
```

## 2.4 显示容器的详细信息
如果要查看容器的详细信息（如 IP 地址、挂载卷等），可以使用：
``` bash
docker inspect [容器名或ID]
```

## 2.5 组合其他命令
* **删除所有已停止的容器:**
``` bash
docker rm $(docker ps -aq)
```
* **格式话输出(例如仅显示容器名和状态):**
``` bash
docker ps -a --format "table {{.Names}}\t{{.Status}}"
```

## 注意事项
* 如果没有任何容器，命令会返回空列表
* 容器ID是唯一标识符，通常只需要输入前四位即可完成操作（如 `docker start a1b2`）。

***

# 3.启动一个已经存在的容器
如果要启动一个已经停止的**Docker容器**(需要知道容器名称或者ID):
```bash
docker start [容器名或ID]           # 启动容器
docker start -i [容器名或ID]        # 启动并进入交互模式（类似docker attach）
```

***

# 4.创建并启动新容器
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

# 5.重启容器
重启正在运行或者已经停止的容器
```bash
docker restart [容器名或ID]
```
