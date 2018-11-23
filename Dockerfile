FROM golang

MAINTAINER poplark

# -------- 镜像操作指令 --------
ENV APP_DIR /go/src/bee-web
RUN mkdir -p $APP_DIR
ADD . $APP_DIR

# 解决无法修改/etc/hosts的问题
#RUN cp /etc/hosts /tmp/hosts #路径长度最好保持一致
#RUN mkdir -p -- /lib-override && cp /lib/x86_64-linux-gnu/libnss_files.so.2 /lib-override
#RUN sed -i 's:/etc/hosts:/tmp/hosts:g' /lib-override/libnss_files.so.2
#ENV LD_LIBRARY_PATH /lib-override
#RUN echo "192.30.253.112 github.com" >> /tmp/hosts
#RUN echo "151.101.88.249 github.global.ssl.fastly.net" >> /tmp/hosts

WORKDIR $APP_DIR
# RUN go build 相关命令
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 3001

# -------- 容器启动时执行指令 --------
# Set the entrypoint
ENTRYPOINT ./bee-web
# #CMD ["bee-web"] 有 entrypoint，不需要这个，否则重复了

