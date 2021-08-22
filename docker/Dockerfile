# Global Args, will be used these values in both build and deploy environemtns build-stages
ARG APP=xks-go
ARG WORK_DIR=/app/$APP

# Build environment
# -----------------

FROM golang:1.13-alpine as build-env
ARG APP
ARG WORK_DIR
WORKDIR $WORK_DIR

RUN apk update && apk add --no-cache gcc musl-dev git

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -a -o ./bin/$APP ./cmd/app
RUN ls $WORK_DIR/bin/

# Deployment environment
# ----------------------
FROM 287432581570.dkr.ecr.us-west-2.amazonaws.com/centos-base:7
ARG APP
ARG WORK_DIR
## Todo - move add user, /etc/xpc to base image
RUN /sbin/adduser -u 800 xplat && \
    mkdir -p -m755 $WORK_DIR && \
    mkdir -p /etc/xpc

RUN chown -R xplat:xplat /app /etc/xpc && \
    chmod -R u=rwX,g=rX /etc/xpc && \
    chmod -R 755 $WORK_DIR
RUN echo $WORK_DIR/bin/$APP
COPY config.yaml $WORK_DIR

COPY --from=build-env $WORK_DIR/bin/$APP $WORK_DIR
RUN chmod +x $WORK_DIR/$APP
EXPOSE 8080
CMD ["/app/xks-go/xks-go"]