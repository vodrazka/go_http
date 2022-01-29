FROM golang

COPY chat.go chat.go

ENTRYPOINT ["go", "run", "chat.go"]