FROM scratch

WORKDIR /root
COPY config.yml /root
COPY main /root

CMD ["./main"]