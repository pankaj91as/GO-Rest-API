FROM cytopia/ansible:latest

RUN apk update
RUN apk add openssh sshpass net-tools

COPY ansible/inventory.ini /data
COPY ansible/pinghost.yml /data

