FROM ubuntu

WORKDIR /opt
RUN mkdir volume
COPY cloud .

CMD ["./cloud", "-json", "volume/sample.json"]
