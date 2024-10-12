FROM alpine:latest

RUN apk add --no-cache \
    curl \ 
    bash \
    xz \
    sudo 

RUN addgroup nixbld && \
    adduser -D -G nixbld nixuser 

RUN curl -L https://nixos.org/nix/install | sh

COPY . .

CMD ["./nix/var/nix/profiles/default/bin/nix-shell"]