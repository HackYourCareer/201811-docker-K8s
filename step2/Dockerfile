FROM alpine:3.7

COPY ./maze /
COPY ./run.sh /

ENV WORK_DIR /generated

ENTRYPOINT ["./run.sh"]
CMD ["5", "5"]
