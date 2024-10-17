FROM ubuntu:22.04

RUN apt update && apt install -y openssh-server net-tools
RUN echo 'root:Pa33w0rd' | chpasswd
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN service ssh start

EXPOSE 22

CMD [ "/usr/sbin/sshd", "-D" ]