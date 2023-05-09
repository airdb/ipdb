FROM airdb/builder:alpine-go1.20 as builder

WORKDIR /build

ADD ./ /build

RUN cd /build/ && \
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o ipdb


FROM airdb/base:latest

WORKDIR /app/ipdb

COPY --from=builder /build/ipdb /app/ipdb/

EXPOSE 8080

#ENTRYPOINT ["sleep", "3600"]
ENTRYPOINT ["/app/ipdb/ipdb"]
