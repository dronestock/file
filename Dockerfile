FROM storezhang/alpine:3.16.2


LABEL author="storezhang<华寅>" \
email="storezhang@gmail.com" \
qq="160290688" \
wechat="storezhang" \
# TODO 增加描述信息
description="Drone持续集成Docker插件，增加以下功能：1、xxx；2、xxx"


# 复制文件
COPY plugin /bin


RUN set -ex \
    \
    \
    \
    && apk update \
    && apk --no-cache add docker \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/plugin \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 执行命令
ENTRYPOINT /bin/plugin
