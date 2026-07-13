FROM go1.26.4
WORKDIR /home/aaron/coding/morsecode-converter

COPY morsecode-converter ./go.mod
EXPOSE 8080

RUN useradd app 
USER app

CMD ["go run go.main"]
