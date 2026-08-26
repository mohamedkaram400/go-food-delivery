                    Identity Service
                         ↑
                         │
                        gRPC
                         │
                         │
Restaurant Service ←── Order Service ──→ Payment Service
                         │
                         │
                        gRPC
                         │
                         ▼
                  Delivery Service


                    ┌─────────────────┐
                    │   API Gateway   │
                    │     REST        │
                    └────────┬────────┘
                             │ gRPC
        ┌────────────────────┼────────────────────┐
        │                    │                    │
        ▼                    ▼                    ▼
   Identity              Restaurant             Order
   Service                Service              Service
        │                    │                    │
        ▼                    ▼                    ▼
   Identity DB          Restaurant DB         Order DB

                             │
                             │ events
                             ▼
                       Message Broker
                       (Kafka/RabbitMQ)
                      /       |       \
                     /        |        \
                    ▼         ▼         ▼
              Payment    Delivery   Notification
              Service     Service      Service

                  
service RestaurantService {

    rpc GetRestaurant(...)

    rpc ValidateMenuItem(...)

    rpc UpdateAvailability(...)
}


service OrderService {

    rpc CreateOrder(...)

    rpc CancelOrder(...)

    rpc GetOrder(...)
}