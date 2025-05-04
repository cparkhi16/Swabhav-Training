#include <iostream>
#include <string>
#include <memory>

// Base Component
class Pizza {
public:
    virtual std::string getDescription() const = 0;
    virtual double getCost() const = 0;
    virtual ~Pizza() = default;
};

// Concrete Component
class PlainPizza : public Pizza {
public:
    std::string getDescription() const override {
        return "Plain Pizza";
    }

    double getCost() const override {
        return 5.0;
    }
};

// Decorator Base Class
class ToppingDecorator : public Pizza {
protected:
    std::unique_ptr<Pizza> pizza_;
public:
    ToppingDecorator(std::unique_ptr<Pizza> pizza) : pizza_(std::move(pizza)) {}
};

// Concrete Decorators
class Cheese : public ToppingDecorator {
public:
    Cheese(std::unique_ptr<Pizza> pizza) : ToppingDecorator(std::move(pizza)) {}

    std::string getDescription() const override {
        return pizza_->getDescription() + ", Cheese";
    }

    double getCost() const override {
        return pizza_->getCost() + 1.5;
    }
};

class Olives : public ToppingDecorator {
public:
    Olives(std::unique_ptr<Pizza> pizza) : ToppingDecorator(std::move(pizza)) {}

    std::string getDescription() const override {
        return pizza_->getDescription() + ", Olives";
    }

    double getCost() const override {
        return pizza_->getCost() + 1.0;
    }
};

class Pepperoni : public ToppingDecorator {
public:
    Pepperoni(std::unique_ptr<Pizza> pizza) : ToppingDecorator(std::move(pizza)) {}

    std::string getDescription() const override {
        return pizza_->getDescription() + ", Pepperoni";
    }

    double getCost() const override {
        return pizza_->getCost() + 2.0;
    }
};

// Client Code
int main() {
    // Start with a Plain Pizza
    std::unique_ptr<Pizza> pizza = std::make_unique<PlainPizza>();

    // Add Cheese
    pizza = std::make_unique<Cheese>(std::move(pizza));

    // Add Olives
    pizza = std::make_unique<Olives>(std::move(pizza));

    // Add Pepperoni
    pizza = std::make_unique<Pepperoni>(std::move(pizza));

    // Final Description and Cost
    std::cout << "Pizza Description: " << pizza->getDescription() << "\n";
    std::cout << "Total Cost: $" << pizza->getCost() << "\n";

    return 0;
}
