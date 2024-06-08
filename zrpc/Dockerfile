# Escolher a imagem base
FROM golang:1.22 as builder

# Definir o diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar os arquivos do projeto para o contêiner
COPY . .

# Construir o aplicativo
WORKDIR /app/zrpc
RUN CGO_ENABLED=0 GOOS=linux go build -o zrpc .

# Usar uma imagem Docker leve para a execução
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o executável
COPY --from=builder /app/zrpc .

# Copiar o arquivo de configuração
COPY --from=builder /app/zrpc/etc/zrpc.yaml /etc/zrpc.yaml


# Comando para executar o aplicativo
CMD ["./zrpc"]
