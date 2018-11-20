FROM centos

MAINTAINER poplark

# -------- 镜像操作指令 --------
ENV APP_DIR /app/bee-web
RUN mkdir -p $APP_DIR

# RUN go build 相关命令

ADD . $APP_DIR

EXPOSE 3001

# -------- 容器启动时执行指令 --------
WORKDIR $APP_DIR
# Set the entrypoint
ENTRYPOINT ./bee-web
# #CMD ["bee-web"] 有 entrypoint，不需要这个，否则重复了

