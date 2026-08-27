

## Install gin framework
 go get -u github.com/gin-gonic/gin


go get google.golang.org/protobuf
go get google.golang.org/grpc 


brew install golang-migrate
migrate -version

migrate create -ext sql -dir migrations -seq create_users_table


protoc \
  --go_out=services/identity_service \
  --go_opt=paths=source_relative \
  --go-grpc_out=services/identity_service \
  --go-grpc_opt=paths=source_relative \
  proto/identity/identity.proto


go work use ./services/api_gateway_service ./services/identity_service ./services/restaurant_service ./services/order_service ./services/delivery_service ./services/payment_service ./services/notification_service