FROM alpine
RUN apk add --no-cache ca-certificates bash git curl
COPY motivationapp /motivationapp
COPY images /images
ADD images /images
EXPOSE 443
ENTRYPOINT ["/motivationapp"]