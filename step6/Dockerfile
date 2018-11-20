FROM tsmsap/mazegen:latest

COPY ./work.sh /
COPY ./maze-worker /maze-worker

ENV MAZEGEN_CMD /work.sh
ENV WORKER_DIR /workers

ENTRYPOINT ["/maze-worker"]
