# not  suitble for docker , ansible is written by python.
FROM liumiaocn/ansible
MAINTAINER "louisehong <louisehong4168@gmail.com>"
ENV AngoBaseDir /tmp
ENV ANSIBLE_HOST_KEY_CHECKING False

COPY ./dist/ango_linux_amd64/ango /usr/bin/ango

ENTRYPOINT ["/usr/bin/ango"]


## usage
# docker run --rm -v ~/.ssh:/root/.ssh -v /etc/ansible/hosts:/etc/ansible/hosts louisehong/ango \
#  deploy -f  http://www.fenghong.tech/ansible/test/test.yml -t v1.2.0