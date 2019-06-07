FROM techinno/service_b:1
WORKDIR /app
ADD goapp /app/goapp
ENV HOST="http://httpbin.org"
ENV PORT=""
ENV URI="/delay/1"

WORKDIR /app
EXPOSE 1323
CMD ["app"]
