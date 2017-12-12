FROM scratch
ENV PORT=2000
COPY p-vip-gateway /
CMD [ "./p-vip-gateway", "-port=$PORT"]
