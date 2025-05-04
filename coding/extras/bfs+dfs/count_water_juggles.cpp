#include <iostream>
#include <vector>
#include <queue>
#include <set>
#include <tuple>
#include <map>
#include <algorithm>

using namespace std;

// Structure to represent the state of the jars
struct State {
    int a, b, c;
    int moves;

    bool operator<(const State& other) const {
        if (a != other.a) return a < other.a;
        if (b != other.b) return b < other.b;
        return c < other.c;
    }
};

int solve(int capacityA, int capacityB, int capacityC, int target) {
    if (target > capacityA && target > capacityB && target > capacityC) {
        return -1; // Target cannot be reached
    }

    queue<State> q;
    set<tuple<int, int, int>> visited;

    q.push({0, 0, capacityC, 0});
    visited.insert({0, 0, capacityC});

    while (!q.empty()) {
        State current = q.front();
        q.pop();

        if (current.a == target || current.b == target || current.c == target) {
            return current.moves;
        }

        if (current.moves >= 6) {
            continue;
        }

        // Possible moves:
        // 1. Pour from C to A
        int pourCA = min(current.c, capacityA - current.a);
        State nextCA = {current.a + pourCA, current.b, current.c - pourCA, current.moves + 1};
        if (visited.find({nextCA.a, nextCA.b, nextCA.c}) == visited.end()) {
            visited.insert({nextCA.a, nextCA.b, nextCA.c});
            q.push(nextCA);
        }

        // 2. Pour from C to B
        int pourCB = min(current.c, capacityB - current.b);
        State nextCB = {current.a, current.b + pourCB, current.c - pourCB, current.moves + 1};
        if (visited.find({nextCB.a, nextCB.b, nextCB.c}) == visited.end()) {
            visited.insert({nextCB.a, nextCB.b, nextCB.c});
            q.push(nextCB);
        }

        // 3. Pour from B to A
        int pourBA = min(current.b, capacityA - current.a);
        State nextBA = {current.a + pourBA, current.b - pourBA, current.c, current.moves + 1};
        if (visited.find({nextBA.a, nextBA.b, nextBA.c}) == visited.end()) {
            visited.insert({nextBA.a, nextBA.b, nextBA.c});
            q.push(nextBA);
        }

        // 4. Pour from B to C
        int pourBC = min(current.b, capacityC - current.c);
        State nextBC = {current.a, current.b - pourBC, current.c + pourBC, current.moves + 1};
        if (visited.find({nextBC.a, nextBC.b, nextBC.c}) == visited.end()) {
            visited.insert({nextBC.a, nextBC.b, nextBC.c});
            q.push(nextBC);
        }

        // 5. Pour from A to B
        int pourAB = min(current.a, capacityB - current.b);
        State nextAB = {current.a - pourAB, current.b + pourAB, current.c, current.moves + 1};
        if (visited.find({nextAB.a, nextAB.b, nextAB.c}) == visited.end()) {
            visited.insert({nextAB.a, nextAB.b, nextAB.c});
            q.push(nextAB);
        }

        // 6. Pour from A to C
        int pourAC = min(current.a, capacityC - current.c);
        State nextAC = {current.a - pourAC, current.b, current.c + pourAC, current.moves + 1};
        if (visited.find({nextAC.a, nextAC.b, nextAC.c}) == visited.end()) {
            visited.insert({nextAC.a, nextAC.b, nextAC.c});
            q.push(nextAC);
        }
    }

    return -1; // Target not reached within 6 moves
}

int main() {
    int capacityA, capacityB, capacityC, target;

    cout << "Enter the capacity of jar A: ";
    cin >> capacityA;

    cout << "Enter the capacity of jar B: ";
    cin >> capacityB;

    cout << "Enter the capacity of jar C (initial amount): ";
    cin >> capacityC;

    cout << "Enter the target capacity: ";
    cin >> target;

    int minMoves = solve(capacityA, capacityB, capacityC, target);

    if (minMoves != -1) {
        cout << "Minimum juggles to measure " << target << " units of water: " << minMoves << endl;
    } else {
        cout << "Target capacity cannot be measured within 6 juggles." << endl;
    }

    return 0;
}