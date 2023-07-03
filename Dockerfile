FROM golang:1.20-alpine
LABEL maintainer="zotov.artem.2019@gmail.com"
WORKDIR /GOproj/currencyConverter
COPY . ./
EXPOSE 8081

CMD ["./main"]


