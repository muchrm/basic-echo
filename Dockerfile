FROM scratch

WORKDIR /root
ADD main /root
ADD config.yml /root

EXPOSE 3000
CMD ["./main"]