FROM scratch
ADD main /root
ADD config.yml /root
EXPOSE 3000
WORKDIR /root
CMD ["./main"]