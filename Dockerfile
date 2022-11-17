FROM golang:1.19.3-bullseye

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /home/travel-accommodation/server

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

EXPOSE 8080

CMD air