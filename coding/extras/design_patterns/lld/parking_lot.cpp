#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <memory>
#include <chrono>

enum class VehicleCategory { TwoWheeler, Hatchback, Sedan, SUV, Bus };
enum class ParkingSlotType { TwoWheeler, Compact, Medium, Large };

// Forward declarations
class Vehicle;
class ParkingSlot;
class Ticket;

// ------------------- Address -------------------
struct Address {
    std::string street, city, state, zip;
};

// ------------------- Vehicle -------------------
class Vehicle {
protected:
    std::string licensePlate;
    VehicleCategory category;

public:
    Vehicle(std::string plate, VehicleCategory cat) : licensePlate(plate), category(cat) {}
    VehicleCategory getCategory() const { return category; }
    std::string getPlate() const { return licensePlate; }
};

class Bike : public Vehicle {
public:
    Bike(std::string plate) : Vehicle(plate, VehicleCategory::TwoWheeler) {}
};

class Car : public Vehicle {
public:
    Car(std::string plate) : Vehicle(plate, VehicleCategory::Sedan) {}
};

// ------------------- ParkingSlot -------------------
class ParkingSlot {
    std::string name;
    bool available = true;
    std::shared_ptr<Vehicle> vehicle;
    ParkingSlotType slotType;

public:
    ParkingSlot(std::string name, ParkingSlotType type) : name(name), slotType(type) {}

    bool isAvailable() const { return available; }
    ParkingSlotType getSlotType() const { return slotType; }

    void addVehicle(std::shared_ptr<Vehicle> v) {
        vehicle = v;
        available = false;
    }

    void removeVehicle(std::shared_ptr<Vehicle> v) {
        vehicle = nullptr;
        available = true;
    }

    std::shared_ptr<Vehicle> getVehicle() const { return vehicle; }
};

// ------------------- Ticket -------------------
class Ticket {
    long long startTime;
    std::shared_ptr<Vehicle> vehicle;
    std::shared_ptr<ParkingSlot> slot;

public:
    Ticket(std::shared_ptr<Vehicle> v, std::shared_ptr<ParkingSlot> s)
        : vehicle(v), slot(s) {
        startTime = std::chrono::system_clock::now().time_since_epoch().count();
    }

    static std::shared_ptr<Ticket> createTicket(std::shared_ptr<Vehicle> v, std::shared_ptr<ParkingSlot> s) {
        return std::make_shared<Ticket>(v, s);
    }

    long long getStartTime() const { return startTime; }
    std::shared_ptr<Vehicle> getVehicle() const { return vehicle; }
    std::shared_ptr<ParkingSlot> getParkingSlot() const { return slot; }
};

// ------------------- ParkingFloor -------------------
class ParkingFloor {
    std::string name;
    std::unordered_map<ParkingSlotType, std::unordered_map<std::string, std::shared_ptr<ParkingSlot>>> parkingSlots;

    ParkingSlotType pickSlotType(VehicleCategory category) {
        switch (category) {
            case VehicleCategory::TwoWheeler: return ParkingSlotType::TwoWheeler;
            case VehicleCategory::Hatchback:
            case VehicleCategory::Sedan: return ParkingSlotType::Compact;
            case VehicleCategory::SUV: return ParkingSlotType::Medium;
            case VehicleCategory::Bus: return ParkingSlotType::Large;
        }
        return ParkingSlotType::Compact;
    }

public:
    ParkingFloor(std::string name,
                 std::unordered_map<ParkingSlotType, std::unordered_map<std::string, std::shared_ptr<ParkingSlot>>> slots)
        : name(name), parkingSlots(slots) {}

    std::shared_ptr<ParkingSlot> getRelevantSlotForVehicleAndPark(std::shared_ptr<Vehicle> vehicle) {
        ParkingSlotType type = pickSlotType(vehicle->getCategory());
        auto& slotMap = parkingSlots[type];
        for (auto& pair : slotMap) {
            if (pair.second->isAvailable()) {
                pair.second->addVehicle(vehicle);
                return pair.second;
            }
        }
        return nullptr;
    }
};

// ------------------- ParkingLot (Singleton) -------------------
class ParkingLot {
    std::string name;
    Address address;
    std::vector<std::shared_ptr<ParkingFloor>> floors;

    static std::shared_ptr<ParkingLot> instance;

    ParkingLot(std::string name, Address addr, std::vector<std::shared_ptr<ParkingFloor>> flrs)
        : name(name), address(addr), floors(flrs) {}

public:
    static std::shared_ptr<ParkingLot> getInstance(const std::string& name, Address addr,
                                                   std::vector<std::shared_ptr<ParkingFloor>> flrs) {
        if (!instance) {
            instance = std::shared_ptr<ParkingLot>(new ParkingLot(name, addr, flrs));
        }
        return instance;
    }

    void addFloor(const std::string& name,
                  std::unordered_map<ParkingSlotType, std::unordered_map<std::string, std::shared_ptr<ParkingSlot>>> slots) {
        floors.push_back(std::make_shared<ParkingFloor>(name, slots));
    }

    std::shared_ptr<Ticket> assignTicket(std::shared_ptr<Vehicle> vehicle) {
        for (auto& floor : floors) {
            auto slot = floor->getRelevantSlotForVehicleAndPark(vehicle);
            if (slot) return Ticket::createTicket(vehicle, slot);
        }
        return nullptr;
    }

    double scanAndPay(std::shared_ptr<Ticket> ticket) {
        long long endTime = std::chrono::system_clock::now().time_since_epoch().count();
        ticket->getParkingSlot()->removeVehicle(ticket->getVehicle());

        int duration = static_cast<int>((endTime - ticket->getStartTime()) / 1000);
        double price = duration * 0.05; // dummy pricing
        return price;
    }
};

std::shared_ptr<ParkingLot> ParkingLot::instance = nullptr;

// ------------------- Main -------------------
int main() {
    Address addr = {"1st Street", "City", "State", "00000"};

    auto slot1 = std::make_shared<ParkingSlot>("S1", ParkingSlotType::Compact);
    auto slot2 = std::make_shared<ParkingSlot>("S2", ParkingSlotType::Compact);

    std::unordered_map<std::string, std::shared_ptr<ParkingSlot>> compactSlots = {
        {"S1", slot1}, {"S2", slot2}
    };
    std::unordered_map<ParkingSlotType, std::unordered_map<std::string, std::shared_ptr<ParkingSlot>>> slots = {
        {ParkingSlotType::Compact, compactSlots}
    };

    auto floor = std::make_shared<ParkingFloor>("F1", slots);
    auto lot = ParkingLot::getInstance("MyLot", addr, {floor});

    auto car = std::make_shared<Car>("ABC123");
    auto ticket = lot->assignTicket(car);

    if (ticket) {
        std::cout << "Ticket issued for car: " << car->getPlate() << "\n";
        double price = lot->scanAndPay(ticket);
        std::cout << "Total price: " << price << "\n";
    } else {
        std::cout << "No slot available.\n";
    }

    return 0;
}
