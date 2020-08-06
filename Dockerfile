FROM alpine:3.7

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
	apk add tzdata ca-certificates && \
	cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
	echo "Asia/Shanghai" > /etc/timezone && apk add --update bash && apk add curl && \
	apk add wget && \
	rm -rf /var/cache/apk/*

COPY ango /usr/bin/ango

ENTRYPOINT ["/usr/bin/ango"]
