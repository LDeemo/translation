FROM debian:stretch as runner
#ARG WORKDIR
ARG BIN
#ENV ep=$WORKDIR/bin/$BIN

#WORKDIR $WORKDIR

RUN echo "" > /etc/apt/sources.list \
 && echo "deb http://mirrors.cloud.tencent.com/debian/ stretch main non-free contrib" >> /etc/apt/sources.list \
 && apt-get update \
 && apt-get install bash -y\
 && echo "alias ll='ls -l'" >> /root/.bashrc \
 && echo "alias hh='hostname'" >> /root/.bashrc \
 && echo 'export PS1="[\u@alpine \W]\\$ "' >> /root/.bashrc \
 && echo "export COLUMNS=207" >> /root/.bashrc \
 && apt-get install tzdata -y\
 && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
 && apt-get remove tzdata -y\
# && mkdir -p /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 \
 && apt-get install ca-certificates tcpdump -y \
 && apt-get install mysql-client -y \
 && apt-get install busybox -y \
 && apt-get -y autoremove \
 && apt-get clean

COPY --from=fullstorydev/grpcurl:v1.3.1 /bin/grpcurl /usr/bin
#RUN echo "#!/bin/bash \n $ep \$@" > ./entrypoint.sh
#RUN chmod +x ./entrypoint.sh

COPY build/build_service/ /usr/local/services

#ENTRYPOINT ["./entrypoint.sh"]
