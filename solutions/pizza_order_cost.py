# Run this file: python pizza_order_cost.py — answers are typed at the prompts.


def pizzaOrderCost(pizza_size: str, num_toppings: int, delivery_distance: float) -> None:
    pizza_size = pizza_size.lower().strip()  # so "Small" and " small " both work
    if pizza_size == "small":
        cost = 8
    elif pizza_size == "large":
        cost = 12
    else:
        raise ValueError("pizza_size must be 'small' or 'large'")
    if num_toppings < 0 or delivery_distance < 0:
        raise ValueError("toppings and distance must be non-negative")
    cost += num_toppings  # $1 each
    # 0 mi = pickup (no fee). 1–5 mi = flat $5. Beyond 5 mi, pay only for miles past 5.
    if delivery_distance == 0:
        pass
    elif delivery_distance <= 5:
        cost += 5
    else:
        cost += delivery_distance - 5
    print(f"Total: ${cost:.2f}")  # .2f = two decimal places


pizzaOrderCost(
    input("What size pizza would you like? (small or large) "),
    int(input("How many toppings would you like? ")),  # input is always a string; need int for math
    float(input("How far are you from the pizza store? (in miles) ")),
)
