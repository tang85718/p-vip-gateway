FROM scratch
ENV PORT=2000
COPY p-vip-gateway /
COPY entrypoint.sh /
ENTRYPOINT ["./entrypoint.sh"]
