FROM alpine
RUN apk add tzdata
WORKDIR /khumu
COPY khumu-comment /khumu/khumu-comment
ENV KHUMU_HOME /khumu
ENV KHUMU_ENVIRONMENT DEV
CMD ["./khumu/khumu-comment"]