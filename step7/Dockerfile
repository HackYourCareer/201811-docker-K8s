FROM alpine:3.7

COPY ./maze-controller /

ENV WORKER_DIR /workers
ENV WORKER_HOST host.docker.internal:8081
ENV REDIS_ADDR host.docker.internal:6379

ENTRYPOINT ["/maze-controller"]
