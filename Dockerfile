FROM scratch
ENV PORT=2000
COPY docker/p-vip-gateway /
ENTRYPOINT ["./p-vip-gateway"]
