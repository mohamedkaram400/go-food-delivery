Pass 1 — Tables
Identity Service
- users
- customers
- addresses
- refresh_tokens
- roles
- permissions
- role_has_permission


Restaurant Service
- restaurants
- menus
- categories
- menu_items
- restaurant_working_hours (fueatre)


Order Service
- carts
- cart_items
- orders
- order_items


Payment Service
- payments


Delivery Service
- drivers
- deliveries
- driver_locations


Notification Service
- notifications



Pass 2 — Columns

* users
- id
- name
- email
- password
- role_id
- status
- email_verified_at
- created_at
- updated_at

* customers
- id
- phone
- user_id
- created_at
- updated_at

* addresses
- id
- customer_id
- label
- address_line
- city
- latitude
- longitude
- is_default
- created_at
- updated_at


* refresh_tokens
- id
- token
- expires_at
- created_at
- revoked_at
- user_id

* roles
- id
- name

permissions
- id
- name

* role_has_permission
- role_id
- permission_id

* restaurants
- id
- name
- description
- phone
- email
- status
- is_open
- owner_id
- created_at
- updated_at

* menus
- id
- restaurant_id
- name
- description
- is_active
- created_at
- updated_at


* categories
- id
- name
- description
- created_at
- updated_at
- menu_id


* menu_items
- id
- restaurant_id
- category_id
- name
- description
- price
- image_url
- is_available
- created_at
- updated_at

* restaurant_working_hours
- id
- restaurant_id
- day_of_week
- open_time
- close_time
- is_closed
- payment_id
- placed_at
- created_at
- updated_at

* carts
- id
- customer_id
- restaurant_id
- status
- created_at
- updated_at


* cart_items
- id
- cart_id
- menu_item_id
- quantity
- unit_price
- created_at
- updated_at

* orders
- id
- status
- subtotal
- delivery_fee
- total
- customer_id
- restaurant_id
- address_id
- payment_id


* order_items
- id
- quantity
- unit_price
- total_price
- order_id
- menu_item_id
- created_at
- updated_at

* payments
- id
- order_id
- amount
- status (pending, paid, failed, refunded)
- payment_method
- transaction_id
- paid_at
- created_at
- updated_at

* drivers
- id
- user_id
- status
- is_online
- created_at
- updated_at

* deliveries
- id
- order_id
- driver_id
- status
- assigned_at
- accepted_at
- picked_up_at
- delivered_at
- created_at
- updated_at

* driver_locations
- id
- driver_id
- latitude
- longitude
- recorded_at

* notifications
- id
- user_id
- type
- title
- message
- status
- read_at
- created_at

Pass 3 — Relationships

Identity Service
----------------

User
1
↓
1
Customer

User
1
↓
N
Refresh Tokens

User
N
↓
1
Role

Role
N
↕
N
Permissions

Customer
1
↓
N
Addresses


Restaurant Service
------------------

Restaurant
1
↓
N
Menus

Restaurant
1
↓
N
Working Hours

Menu
1
↓
N
Categories

Category
1
↓
N
Menu Items

Restaurant
1
↓
N
Orders


Order Service
-------------

Customer
1
↓
N
Orders

Customer
1
↓
1
Active Cart

Cart
1
↓
N
Cart Items

Cart Item
N
↓
1
Menu Item

Order
1
↓
N
Order Items

Order Item
N
↓
1
Menu Item

Order
N
↓
1
Restaurant

Order
N
↓
1
Address

Order
1
↓
1
Payment


Delivery Service
----------------

Driver
1
↓
N
Deliveries

Driver
1
↓
N
Driver Locations

Delivery
N
↓
1
Driver

Delivery
1
↓
1
Order


Notification Service
--------------------

User
1
↓
N
Notifications


Pass 4 — Constraints



Pass 5 — Indexes