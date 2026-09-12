FROM golang:1.27

WORKDIR /usr/src/Saavedra

COPY ./entrypoint.sh /usr/src/Saavedra/entrypoint.sh
RUN chmod +x /usr/src/Saavedra/entrypoint.sh

# Establecemos el entrypoint
ENTRYPOINT ["tail", "-f", "/dev/null"]
