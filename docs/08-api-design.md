
Identity Service
===================

### Auth
POST   /auth/register
POST   /auth/login
POST   /auth/refresh
POST   /auth/logout

GET    /users/me
PATCH  /users/me

GET    /users/:id
PATCH  /users/:id


Restaurant Service
===================

### Restaurant
POST /restaurants
GET /restaurants
GET /restaurants/:id
PATCH /restaurants/:id
DELETE /restaurants/:id


### Menu
GET    /restaurants/:id/menu

# Categories

POST   /menus/:menu_id/categories
GET    /menus/:menu_id/categories
GET    /categories/:id
PATCH  /categories/:id
DELETE /categories/:id

POST   /restaurants/:id/menu/categories
POST   /restaurants/:id/menu/items
GET    /restaurants/:id/menu/items
PATCH  /menu-items/:id
DELETE  /menu-items/:id

### Menu Items

POST   /menus/:menu_id/items
GET    /menus/:menu_id/items
GET    /menu-items/:id
PATCH  /menu-items/:id
DELETE /menu-items/:id

### Availability
PATCH  /menu-items/:id/availability


Order Service
==============
### Cart
GET    /cart
POST   /cart/items
POST /cart/items/:id
DELETE /cart/items/:id
DELETE /cart

### Orders
POST /orders
GET /orders
GET /orders/:id
GET /orders/:id/cancel

Restaurant Order Management
==========================
POST /orders/:id/accept
POST /orders/:id/reject
POST /orders/:id/prepare
POST /orders/:id/ready


Delivery Service
=================
GET /deliveries
GET /deliveries/:id
POST /deliveries/:id/accept
POST /deliveries/:id/pickup
POST /deliveries/:id/deliver

POST /drivers/location
GET  /drivers/me/location
PATCH /drivers/me/status


Payment Service
================
GET /payments
GET /payments/:id
POST /payments/:id/pay


Notification Service
====================
GET /notifications
GET /notifications/:id
PATCH /notifications/:id/read
DELETE /notifications/:id




                    ┌──────────────────┐
                    │ Identity Service │
                    │                  │
                    │ identity_db      │
                    └────────┬─────────┘
                             │
                            gRPC
                             │
              ┌──────────────┴──────────────┐
              │                             │
              ▼                             ▼
     ┌─────────────────┐          ┌─────────────────┐
     │ Restaurant      │          │ Order           │
     │ Service         │◄──gRPC──►│ Service         │
     │                 │          │                 │
     │ restaurant_db   │          │ order_db        │
     └─────────────────┘          └────────┬────────┘
                                           │
                                          gRPC
                                           │
                                           ▼
                                  ┌─────────────────┐
                                  │ Payment         │
                                  │ Service         │
                                  │                 │
                                  │ payment_db      │
                                  └─────────────────┘


                    Events / Message Broker
                           │
              ┌────────────┼────────────┐
              ▼            ▼            ▼
          Delivery    Notification   Other services
           Service       Service
              │            │
              ▼            ▼
        delivery_db    notification_db