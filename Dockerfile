FROM alpine:3.12.1
RUN apk add --no-cache dumb-init openssh
RUN adduser -D -s /sbin/nologin bastion
RUN passwd -u -d bastion
RUN ssh-keygen -A
RUN (echo 'PasswordAuthentication no'; \
    echo 'port 2222'; \
    echo 'PubkeyAuthentication yes'; \
    echo 'StrictModes no'; ) > /etc/ssh/sshd_config_custom
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/usr/sbin/sshd", "-D", "-e", "-f", "/etc/ssh/sshd_config_custom"]
