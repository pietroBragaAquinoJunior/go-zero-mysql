# go-zero-mysql

Projeto backend em go-zero com endpoint e comunicação com banco de dados MySQL.

canal: https://www.youtube.com/@pietrobraga523/videos

Primeiro você precisará ter o docker e o docker-compose instalados:
```bash
docker --version
docker-compose --version
```

Navegue até a pasta do projeto e execute o comando:
```bash
docker-compose up -d
```

Comandos utilizados para gerar o código do projeto:
```bash
goctl rpc protoc ./protos/zrpc.proto --go_out=./common/pb --go-grpc_out=./common/pb --zrpc_out=./zrpc
goctl model mysql ddl -c -src zrpc/go_zero_mysql_init_table.sql -dir ./zrpc/internal/models
```
![Screenshot from 2024-08-01 17-04-30](https://github.com/user-attachments/assets/b551a3e1-09f7-4ada-a341-2a67e9a65a69)
