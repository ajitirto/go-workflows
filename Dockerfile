FROM gcr.io/distroless/static-debian12
ARG binary

COPY ./$binary /app
ENTRYPOINT ["/app"]
