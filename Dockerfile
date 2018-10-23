FROM alpine
RUN apk add --no-cache ca-certificates bash git curl
COPY chikkin-server /chikkin-server
# COPY images /images
# ADD images /images
EXPOSE 443
ENTRYPOINT ["/chikkin-server"]