FROM golang:1.26.2-alpine3.23

WORKDIR /app

# Copiando apenas os módulos primeiro para aproveitar o cache
COPY go.mod go.sum ./
RUN go mod download

# Copiamos o resto dos arquivos
COPY . .

RUN go build -o ./app .

# O Air vai gerenciar o build e a execução
CMD ["go", "tool", "air", "-c", ".air.toml"]