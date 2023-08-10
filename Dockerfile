FROM storezhang/alpine:3.18.3 AS builder

COPY docker /docker
COPY file /docker/usr/local/bin


FROM storezhang/alpine:3.18.3


LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Drone持续集成Ftp插件，提供如下功能：1、文件上传功能；2、支持主流文件上传方式；3、支持Ftp；4、支持Webdav；5、支持Scp"

COPY --from=builder /docker /

RUN set -ex \
    \
    \
    \
    && apk update \
    && apk --no-cache add openssh-client sshpass curl \
    # 增加执行权限 \
    && chmod +x /usr/local/bin/scpx \
    && chmod +x /usr/local/bin/file \
    \
    \
    \
    && rm -rf /var/cache/apk/*

ENTRYPOINT /usr/local/bin/file
