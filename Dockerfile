FROM alpine

RUN mkdir -p /uploader/files -m 777
VOLUME /uploader/files/
WORKDIR /uploader/
EXPOSE 8888
COPY uploader /uploader/
CMD ["./uploader"]
