FROM go1.26.4
WORKDIR /home/aaron/coding/morsecode-converter

COPY /home/aaron/coding ./home/aaron/coding/morsecode-converter
EXPOSE 8080

RUN useradd app 
USER app

CMD ["go run go.main"]
