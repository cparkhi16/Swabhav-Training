#include <iostream>
#include <vector>
#include <map>

using namespace std;

bool carPooling(vector<vector<int>>& trips, int capacity) {
    map<int, int> passengerChanges;

    // Record the changes at each location
    for (const auto& trip : trips) {
        int numPassengers = trip[0];
        int from = trip[1];
        int to = trip[2];

        passengerChanges[from] += numPassengers; // pick up
        passengerChanges[to] -= numPassengers;   // drop off
    }

    int currentPassengers = 0;
    for (const auto& [location, change] : passengerChanges) {
        currentPassengers += change;
        if (currentPassengers > capacity) {
            return false;
        }
    }

    return true;
}

int main() {
    vector<vector<int>> trips1 = {{2,1,5}, {3,3,7}};
    int capacity1 = 4;
    cout << boolalpha << carPooling(trips1, capacity1) << endl; // Output: false

    vector<vector<int>> trips2 = {{2,1,5}, {3,3,7}};
    int capacity2 = 5;
    cout << boolalpha << carPooling(trips2, capacity2) << endl; // Output: true

    return 0;
}
