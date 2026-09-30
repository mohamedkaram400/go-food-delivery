

## Install gin framework
 go get -u github.com/gin-gonic/gin


go get google.golang.org/protobuf
go get google.golang.org/grpc 

brew install golang-migrate

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

go get -u gorm.io/gorm

go get github.com/joho/godotenv

github.com/golang-jwt/jwt/v5

go get github.com/go-playground/validator/v10


go get "google.golang.org/grpc/status"
go get "google.golang.org/grpc/codes"


migrate create -ext sql -dir migrations -seq create_refrash_tokens_table

migrate -path migrations -database "mysql://root:qazwsx123@tcp(127.0.0.1:3306)/identity_service" up
migrate -path migrations -database "mysql://root:qazwsx123@tcp(127.0.0.1:3306)/identity_service" force 4


<!-- mysql://root:qazwsx123@127.0.0.1/keme_dev?statusColor=686B6F&env=development&name=Keme&tLSMode=0&usePrivateKey=false&safeModeLevel=0&advancedSafeModeLevel=0&driverVersion=0&lazyload=false -->

protoc \
  --go_out=services/identity_service \
  --go_opt=paths=source_relative \
  --go-grpc_out=services/identity_service \
  --go-grpc_opt=paths=source_relative \
  proto/identity/identity.proto


go work use ./services/api_gateway_service ./services/identity_service ./services/restaurant_service ./services/order_service ./services/delivery_service ./services/payment_service ./services/notification_service

