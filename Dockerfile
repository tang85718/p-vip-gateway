FROM scratch
ENV PORT=2000
COPY docker/p-vip-gateway /
COPY docker/entrypoint.sh /
ENTRYPOINT ["./entrypoint.sh"]
