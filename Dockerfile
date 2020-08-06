# not suitble for docker
#FROM alpine:3.7
#
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
#	apk add tzdata ca-certificates && \
#	cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
#	echo "Asia/Shanghai" > /etc/timezone && apk add --update bash && apk add curl && \
#	apk add wget ansible openssh python && \
#	rm -rf /var/cache/apk/*
#ENV AngoBaseDir /tmp
#ENV ANSIBLE_HOST_KEY_CHECKING False
#ENV ANSIBLE_PYTHON_INTERPRETER /usr/bin/python2
#
#COPY ./dist/ango_linux_amd64/ango /usr/bin/ango
#
#ENTRYPOINT ["/usr/bin/ango"]
