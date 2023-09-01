# syntax=docker/dockerfile:1
# Specify the base image for the app.
FROM golang:1.19-alpine AS builder
# Specify that we now need to execute any commands in this directory.
WORKDIR /app
# Copy everything from this project into the filesystem of the contaner.
COPY . .
# Compile the binary exe for our app.
RUN go build -o main .

FROM alpine
WORKDIR /app
COPY --from=builder /app .

# LABELS
LABEL version="1.0"
LABEL name="ASCII-ART-WEB-DOCKERIZE"
LABEL git-repo="git@git.01.alem.school:smustafi/ascii-art-web-dockerize.git"
LABEL authors="smustafi, vtarasso"
LABEL release-date="15/02/2023"

# For localhost:4000
EXPOSE 4000
# Start the application
CMD ["./main"]
