

## Install gin framework
 go get -u github.com/gin-gonic/gin


protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  proto/post.proto


go work use ./services/api_gateway_service ./services/identity_service ./services/restaurant_service ./services/order_service ./services/delivery_service ./services/payment_service ./services/notification_service