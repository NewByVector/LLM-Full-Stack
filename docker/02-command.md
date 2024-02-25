- `docker ps` 查看正在运行的容器
- `docker system prune` 清理所有停止的容器
- `docker logs container_id` 查看容器日志
- `docker run image_name` 创建并运行一个容器
- `docker create image_name` 创建一个容器
- `docker start container_id` 启动一个容器
- `docker stop container_id` 停止一个容器
- `docker kill container_id` 强制停止一个容器
- `docker exec -it container_id bash` 在容器内执行命令
- `docker build -t image_name -f docker_file .` 创建镜像文件(image_name=dockerID/镜像名:版本号)
- `docker commit container_id image_name` 保存容器为镜像
- `docker compose up --build` 重新构建镜像并启动容器
- `docker compose down` 停止容器