goctl rpc protoc ./protos/zrpc.proto --go_out=./common/pb --go-grpc_out=./common/pb --zrpc_out=./zrpc

goctl model mysql ddl -c -src zrpc/go_zero_mysql_init_table.sql -dir ./zrpc/internal/models
