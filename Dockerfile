FROM go1.26.4
WORKDIR https://github.com/aaron03emnify/morsecode-converter.git

COPY src ./go.mod
EXPOSE 8080

RUN useradd app 
USER app

CMD ["go run go.main"]
