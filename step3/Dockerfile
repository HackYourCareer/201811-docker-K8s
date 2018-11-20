FROM tomek:latest

RUN apk --update add imagemagick
RUN apk --no-cache add ttf-freefont

COPY ./convert.sh /
COPY ./mazegen.sh /

ENTRYPOINT ["./mazegen.sh"]
CMD ["hello kuba"]
