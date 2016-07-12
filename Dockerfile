FROM alpine

VOLUME /uploader/files/
WORKDIR /uploader/
COPY uploader /uploader/
EXPOSE 8888
CMD ["./uploader"]
