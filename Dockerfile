# LATEST GO IMAGE
FROM golang:1.27

# MAKE SURE NANO EDITOR EXISTS
RUN apt update && apt install nano -y

# SAAVEDRA AS MAIN DIRECTORY
WORKDIR /usr/src/Saavedra

# EXCLUDED db, private, .env, .gitignore, LICENSE, README.md, testAPI.sh
COPY assets/ config/ env/ migration/ service/ utils/ entrypoint.sh go.mod go.sum main.go ./

# DOWNLOAD GO DEPENDENCIES
RUN go mod download

# COMPILE GO SAAVEDRA AS AN APP
RUN go build ./...

# START COONTAINER WITH NO PROCESS OF SAAVEDRA
ENTRYPOINT [ "tail", "-f", "/dev/null" ]
