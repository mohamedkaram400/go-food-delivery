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

* customers
- id
- phone
- user_id

* addresses
- id
- label
- address_line
- city
- latitude
- longitude
- is_default
- customer_id

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
- name
- description
- is_active
- restaurant_id
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
- name
- description
- price
- image_url
- is_avilable
- restaurant_id
- category_id
- created_at
- updated_at

* restaurant_working_hours
- id
- restaurant_id
- day_of_week
- open_time
- close_time
- is_closed

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


* payments
- id

* drivers
- id

* deliveries
- id

* driver_locations
- id

* notifications
- id


Pass 1 — Relationships

Restaurant

1

↓

N

Menus

<!-- ====== -->

Menu

1

↓

N

Categories

<!-- ====== -->
Category

1

↓

N

Menu Items

<!-- ====== -->

Customer

1

↓

N

Orders

<!-- ====== -->
Order

1

↓

N

Order Items

<!-- ====== -->
Cart

1

↓

N

Cart Items