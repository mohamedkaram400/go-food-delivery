

<!-- ============================================ Business Objects ================================================= -->


- Nonus in the system

Customer

Address

Restaurant

Menu

Category

Menu Item

Cart

Cart Item

Order

Order Item

Driver

Delivery

Payment

Review


<!-- ============================================ Relationships ================================================= -->

Customer
has many
Addresses

Restaurant
has many
Menus

Menu
has many
Categories

Category
has many
Menu Items

Customer
has one
Cart

Cart
contains many
Menu Items

Customer
places many
Orders

Order
contains many
Order Items

Driver
has many
Deliveries

Order
has one
Delivery
