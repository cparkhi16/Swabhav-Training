#include <iostream>
#include <map>
#include <vector>
using namespace std;

class MyCalendarThree {
    map<int, int> timeline;
    int maxBooking = 0;

public:
    MyCalendarThree() {}

    int book(int startTime, int endTime) {
        timeline[startTime]++;
        timeline[endTime]--;

        int active = 0;
        for (const auto& [time, count] : timeline) {
            active += count;
            maxBooking = max(maxBooking, active);
        }

        return maxBooking;
    }
};


int main() {
    MyCalendarThree myCalendar;

    vector<pair<int, int>> bookings = {
        {10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55}
    };

    for (auto& [start, end] : bookings) {
        int k = myCalendar.book(start, end);
        cout << "Booking [" << start << ", " << end << ") => k = " << k << endl;
    }

    return 0;
}
