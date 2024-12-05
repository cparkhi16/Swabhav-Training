#include <iostream>
#include <memory>

// Forward declaration of Context
class TrafficLightContext;

// State Interface
class TrafficLightState {
public:
    virtual ~TrafficLightState() = default;
    virtual void handle(TrafficLightContext& context) = 0;
    virtual const char* getStateName() const = 0;
};

// Concrete State: Red Light
class RedLight : public TrafficLightState {
public:
    void handle(TrafficLightContext& context) override;
    const char* getStateName() const override { return "Red Light"; }
};

// Concrete State: Green Light
class GreenLight : public TrafficLightState {
public:
    void handle(TrafficLightContext& context) override;
    const char* getStateName() const override { return "Green Light"; }
};

// Concrete State: Yellow Light
class YellowLight : public TrafficLightState {
public:
    void handle(TrafficLightContext& context) override;
    const char* getStateName() const override { return "Yellow Light"; }
};

// Context Class
class TrafficLightContext {
private:
    std::shared_ptr<TrafficLightState> state;

public:
    TrafficLightContext(std::shared_ptr<TrafficLightState> initialState)
        : state(std::move(initialState)) {}

    void setState(std::shared_ptr<TrafficLightState> newState) {
        state = std::move(newState);
    }

    void request() {
        state->handle(*this);
    }

    const char* getStateName() const {
        return state->getStateName();
    }
};

// State Transitions
void RedLight::handle(TrafficLightContext& context) {
    std::cout << "Switching from Red to Green.\n";
    context.setState(std::make_shared<GreenLight>());
}

void GreenLight::handle(TrafficLightContext& context) {
    std::cout << "Switching from Green to Yellow.\n";
    context.setState(std::make_shared<YellowLight>());
}

void YellowLight::handle(TrafficLightContext& context) {
    std::cout << "Switching from Yellow to Red.\n";
    context.setState(std::make_shared<RedLight>());
}

// Main Function
int main() {
    auto redState = std::make_shared<RedLight>();
    TrafficLightContext trafficLight(redState);

    // Simulate state transitions
    for (int i = 0; i < 5; ++i) {
        std::cout << "Current State: " << trafficLight.getStateName() << "\n";
        trafficLight.request();
    }

    return 0;
}
