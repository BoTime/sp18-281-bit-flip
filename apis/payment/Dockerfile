FROM golang:latest
EXPOSE 80
RUN mkdir /app
ADD ./src/cmpe281/payment /app/src/cmpe281/payment
ENV GOPATH /app
WORKDIR /app/
RUN (cd src/cmpe281/payment; go get; go build)
ENV IP 0.0.0.0
ENV PORT 80
CMD ["sh", "-c", "/app/bin/payment --dbuser=$DBUSER --dbpass=$DBPASS --ip=$IP --port=$PORT"]
