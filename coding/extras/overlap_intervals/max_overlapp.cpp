// Program to find maximum guest at any time in a party
#include <bits/stdc++.h>
using namespace std;

vector<int> findMaxGuests(vector<int> &Entry, vector<int> &Exit)
{
    int n = Entry.size();
    
    // Sort arrival and Exit arrays
    sort(Entry.begin(), Entry.end());
    sort(Exit.begin(), Exit.end());

    // guests_in indicates number of guests at a time
    int guests_in = 1, max_guests = 1, time = Entry[0];
    int i = 1, j = 0;

    // Similar to merge in merge sort to process
    // all events in sorted order
    while (i < n && j < n)
    {
        // If next event in sorted order is arrival,
        // increment count of guests
        if (Entry[i] <= Exit[j])
        {
            guests_in++;

            // Update max_guests if needed
            if (guests_in > max_guests)
            {
                max_guests = guests_in;
                time = Entry[i];
            }
            i++; // increment index of arrival array
        }
        else // If event is Exit, decrement count
        {    // of guests.
            guests_in--;
            j++;
        }
    }

    return {max_guests, time};
}

// Driver program to test above function
int main()
{
    vector<int> Entry = {1, 2, 10, 5, 5};
    vector<int> Exit = {4, 5, 12, 9, 12};
    vector<int> ans = findMaxGuests(Entry, Exit);
    cout << ans[0] << " " << ans[1];
    return 0;
}
