FROM golang:1.16.3-buster
RUN apt update && \
    apt upgrade -y
RUN apt install \
    gcc \
    curl \
    git
RUN curl https://nim-lang.org/choosenim/init.sh -sSf | bash -s -- -y
ENV PATH /root/.nimble/bin:$PATH
