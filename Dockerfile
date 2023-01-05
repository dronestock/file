FROM storezhang/alpine:3.16.2


LABEL author="storezhang<华寅>" \
email="storezhang@gmail.com" \
qq="160290688" \
wechat="storezhang" \
description="Drone持续集成Ftp插件，提供如下功能：1、文件上传功能"


# 复制文件
COPY file /bin


RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/file \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 执行命令
ENTRYPOINT /bin/file
