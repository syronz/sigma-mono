FROM ubuntu

WORKDIR /opt
RUN mkdir volume
COPY nodeapp .

CMD ["./nodeapp", "-json", "volume/sample.json"]
