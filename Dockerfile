FROM scratch

COPY ango /usr/bin/ango

ENTRYPOINT ["/usr/bin/ango"]
