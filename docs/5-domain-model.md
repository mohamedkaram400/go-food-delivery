

## 1. Customer

A customer is a user who can browse restaurants, manage a cart,
place orders, track deliveries, and submit reviews.

### - Relationships

- A customer can have many addresses.
- A customer can have one active cart.
- A customer can place many orders.
- A customer can submit many reviews.


## 3. Restaurant

A restaurant represents a food provider on the platform.

### - Relationships
- A restaurant can have many menus.
- A restaurant can have many categories throguh its menu.
- A restaurant can have many menu items.
- A restaurant can have many menus.
- A restaurant can receive many orders.
- A restaurant can approved by an admin before appearing in search bar


## 4. Address

An address represents the customer's delivery location.

### - Relationships
- A address can belong to one customer.
- A customer can have many addresses.
- An order uses one delivery address.


## 5. Category

A category groups the related menu items.

### - Relationships
- A category belongs to one menu.
- A category contains many menu items.


## 6. Menu

A menu represents a collection of food items offered by a restaurant.

### - Relationships
* A menu has many menu items
- A menu contains many categories.
- A menu can be active or inactive.


## 7. MenuItem

A menu item represents a food or drink item (product)

### - Relationships
- A menu item belongs to one category.
- A menu item belongs to one restaurant.
- A menu item can appear in many cart items.

### Rules

- A restaurant can mark a menu item as unavailable.
- A menu item that belongs to an active order cannot be deleted.


## 8. Cart 

A cart represnts the cllection of menu items/cart items a customer intends to purchase.

### - Relationships


## 9. CartItem

A cart item represnts a menu item and it's selected qty insde the cart.

### - Relationships
- A cart item belongs to one cart.
- A cart item references one menu item.

### Example

Burger × 2
Fries × 1


## 10. Order

An order represents a customer's confirmed request to purchase food
from a restaurant.

### - Relationships

- An order belongs to one customer.
- An order belongs to one restaurant.
- An order contains many order items.
- An order has one payment.
- An order has one delivery.
- An order uses one delivery address.

### Rules

- An order must contain at least one item.
- All items in an order must belong to the same restaurant.
- Order total is calculated automatically.
- Customer cannot modify an order after restaurant acceptance.
- Customer can cancel an order only while it is pending.


## 11. OrderItem

An order item represents a menu item purchased as part of an order.

### - Relationships
- An order item belongs to one order.
- An order item references one menu item.


## 12. Driver

A driver is a user responsible for delivering orders.

### - Relationships
- A driver can have many deliveries.
- A driver can have at most one active delivery.

### Rules

- Driver must be online before accepting a delivery.
- Driver cannot mark a delivery as delivered before picking it up.
- Driver location is updated while actively delivering.


## 13. Delivery

A delivery represents the process of transporting an order
from a restaurant to a customer.

### - Relationships
- A delivery belongs to one order.
- A delivery is assigned to one driver.
- A driver can have many deliveries over time.

## 14. Payment

A payment represents the financial transaction associated with an order.

### - Relationships


- A payment belongs to one order.
- An order has one payment.

### Rules

- Every order has exactly one payment.
- Payment must succeed before the restaurant starts preparing the order.
- Failed payment causes the order to be cancelled.


## 15. Review

A review represents feedback submitted by a customer after completing an order.

### - Relationships

- A review belongs to one customer.
- A review is associated with one order.
- A review can target a restaurant.
- A review can target a driver.

### Rules

- A customer cannot review an order they did not receive.
- Reviews can only be submitted after the order is completed.



<!-- ================================================================================================ -->

Customer
 ├── has many → Addresses
 ├── has one → Active Cart
 ├── places many → Orders
 └── creates many → Reviews

Restaurant
 ├── has many → Menus
 └── receives many → Orders

Menu
 ├── belongs to → Restaurant
 └── has many → Categories

Category
 └── has many → Menu Items

Cart
 └── has many → Cart Items

Cart Item
 └── references → Menu Item

Order
 ├── belongs to → Customer
 ├── belongs to → Restaurant
 ├── has many → Order Items
 ├── has one → Payment
 ├── has one → Delivery
 └── uses one → Address

Order Item
 └── references → Menu Item

Delivery
 ├── belongs to → Order
 └── assigned to → Driver

Review
 ├── belongs to → Customer
 ├── belongs to → Order
 └── targets → Restaurant / Driver