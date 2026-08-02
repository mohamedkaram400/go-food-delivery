


<!-- ============================================ Business Rules ================================================= -->

- Restaurant Rules
* A restaurant cannot accept an order that has already been cancelled.
* Restaurant cannot receive orders when closed.
* Restaurant cannot delete a menu item that belongs to an active order.
* Restaurant can temporarily mark menu items as unavailable.
* Restaurant must be approved by Admin before appearing in search.


- Customer Rules
* Customer must verify their email before placing an order.
* Customer cannot place an empty order.
* Customer can save up to 10 addresses.
* Customer can have only one active cart.
* Customer cannot review an order they never received.


- Order Rules
* Order must contine at least one itme.
* All items in an order must belong to the same restaurant
* Order total is calculated automatically.
* Customer cannot modify an order after the restaurant accepte it .
* Customer can cancel only while order is still pending.
* An order can have one payment.


- Driver Rules
* Driver must be online before accepting deliveries.
* Driver can only have one active delivery.
* Driver cannot mark an order as delivered before picking it up.
* Driver location updates every few seconds while delivering.


- Payment Rules
* Every order has exactly one payment
* Payment must succeed before the restaurant starts preparing the food.
* Failed payments automatically cancel the order.