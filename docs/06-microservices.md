# Microservices Design

## Architecture Overview
The system follows a microservice architecture.

Each service:

- Owns its business logic.
- Owns its database.
- Can be deployed independently.
- Communicates with other services using gRPC.
- Publishes events when important business actions occur.

## Service Communication
External Clients

↓

REST API

↓

API Gateway

↓

gRPC

↓

Microservices

Asynchronous communication between services uses events.

## API Gateway

## Identity Service
Purpose

Responsible for user authentication and identity management.

Owns

- Customer
- Driver
- Restaurant User
- Admin
- Roles
- Refresh Tokens

Responsibilities

- Register users
- Login users
- Generate JWT tokens
- Refresh tokens
- Manage user profiles

Exposes

- Login
- Register
- Get User

Database

Identity Database

## Restaurant Service
Purpose
Responsible for restaurant management.

Owns
- Restaurant
- Working Hours
- Restaurant Profile

Responsibilities
- Approve Restaurant
- Update Restaurant
- Restaurant Availability
- Working Hours
- Restaurant Profile
- Menu
- Categories
- Menu Items

## Order Service
Purpose
Manage customer orders.

Owns
- Cart
- Cart Item
- Order
- Order Item

Responsibilities
- Create Cart
- Update Cart
- Checkout
- Create Order
- Cancel Order
- Order History
- Track Order Status

## Delivery Service
Purpose
Manage deliveries.

Owns
- Driver
- Delivery

Responsibilities
- Assign Driver
- Update Driver Location
- Pickup
- Delivered
- Delivery Tracking

## Payment Service
Purpose
Handle payments.

Owns
- Payment

Responsibilities
- Create Payment
- Verify Payment
- Refund
- Payment Status

## Notification Service
Purpose
Notify users.

Owns
- Notifications

Responsibilities
- Email
- SMS
- Push Notification

## Service Dependencies
Gateway

↓

Identity

↓

Order

↓

Restaurant

↓

Payment

↓

Delivery

↓

Notification

## Database Ownership
Identity Service

↓

Identity Database

Restaurant Service

↓

Restaurant Database

Catalog Service

↓

Catalog Database

Order Service

↓

Order Database

Delivery Service

↓

Delivery Database

Payment Service

↓

Payment Database

Notification Service

↓

Notification Database