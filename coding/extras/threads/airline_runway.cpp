#include <iostream>
#include <queue>
#include <vector>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <chrono>

using namespace std;

// Flight Structure
struct Flight {
    int id;
    int priority; // Higher priority flights land first
    Flight(int id, int priority) : id(id), priority(priority) {}

    // Overload operator to make priority queue a max heap (higher priority first)
    bool operator<(const Flight& f) const {
        return priority < f.priority;
    }
};

// Runway Class
class Runway {
public:
    int id;
    bool isAvailable;

    Runway(int id) : id(id), isAvailable(true) {}
};

// Air Traffic Control System
class AirTrafficControl {
private:
    bool stopProcessing = false;  
    priority_queue<Flight> flightQueue;
    vector<Runway> runways;
    mutex mtx;
    condition_variable cv;

public:
    AirTrafficControl(int numRunways) {
        for (int i = 0; i < numRunways; i++) {
            runways.emplace_back(i + 1);
        }
    }

    // Request landing for a flight
    void requestLanding(Flight flight) {
        unique_lock<mutex> lock(mtx);
        cout << "Flight " << flight.id << " requests landing (Priority: " << flight.priority << ")\n";
        flightQueue.push(flight);
        cv.notify_one();
    }

    // Allocate a runway when available
    void allocateRunway() {
        cout<<" here "<<endl;
        while (true) {
            unique_lock<mutex> lock(mtx);
            if (stopProcessing && flightQueue.empty()) break;
            cv.wait(lock, [this]() { return !flightQueue.empty(); });
             
            for (auto &runway : runways) {
                if (runway.isAvailable) {
                    Flight flight = flightQueue.top();
                    flightQueue.pop();
                    runway.isAvailable = false;
                    
                    cout << "Allocating Runway " << runway.id << " to Flight " << flight.id << "\n";

                    // Simulate landing time
                    lock.unlock();
                    this_thread::sleep_for(chrono::seconds(3));
                    lock.lock();

                    runway.isAvailable = true;
                    cout << "Runway " << runway.id << " is now free\n";
                    break;
                }
            }
        }
    }
     void stop() {
        lock_guard<mutex> lock(mtx);
        stopProcessing = true;
        cv.notify_one();  // Wake up `allocateRunway()` to allow exit
    }
};

int main() {
     AirTrafficControl atc(2);

    thread runwayAllocator(&AirTrafficControl::allocateRunway, &atc);

    atc.requestLanding(Flight(101, 2));
    this_thread::sleep_for(chrono::seconds(1));
    atc.requestLanding(Flight(102, 1));
    this_thread::sleep_for(chrono::seconds(1));
    atc.requestLanding(Flight(103, 3));

    this_thread::sleep_for(chrono::seconds(5));
    atc.stop();  // Tell the thread to exit
    runwayAllocator.join();  // Wait for the thread to finish

    cout << "All flights processed. Exiting...\n";
    return 0;
    
}
