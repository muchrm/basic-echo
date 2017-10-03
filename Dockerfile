FROM scratch

RUN mkdir /root
COPY config.yml /root
COPY main /root

WORKDIR /root
CMD ["./main"]