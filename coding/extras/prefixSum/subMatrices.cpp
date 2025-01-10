#include <iostream>
#include <vector>
#include <unordered_map>
using namespace std;

int numSubmatrixSumTarget(vector<vector<int>>& matrix, int target) {
    int rows = matrix.size();
    int cols = matrix[0].size();
    int result = 0;

    // Iterate over all possible row ranges
    for (int rowStart = 0; rowStart < rows; rowStart++) {
        vector<int> collapsed(cols, 0); // Collapsed 1D array
        cout<<" rowstart "<<rowStart<<endl;
        for (int rowEnd = rowStart; rowEnd < rows; rowEnd++) {
            // Collapse rows between rowStart and rowEnd into a single array
            cout<<"rowEnd is "<<rowEnd<<endl;
            for (int col = 0; col < cols; col++) {
                collapsed[col] += matrix[rowEnd][col];
            }
            for(int i = 0 ; i < collapsed.size() ; i++){
                cout<<" collapsed val "<<collapsed[i]<<" ";
            }
            cout<<endl;
            // Use prefix sum + hashmap to count subarrays with sum = target
            unordered_map<int, int> prefixCount;
            prefixCount[0] = 1; // Base case
            int prefixSum = 0;

            for (int col = 0; col < cols; col++) {
                prefixSum += collapsed[col];

                // Check if there's a subarray with the required sum
                if (prefixCount.find(prefixSum - target) != prefixCount.end()) {
                    result += prefixCount[prefixSum - target];
                }

                // Update the prefix count
                prefixCount[prefixSum]++;
            }
        }
    }

    return result;
}

int main() {
    vector<vector<int>> matrix = {{0, 1, 0}, {1, 1, 1}, {0, 1, 0}};
    int target = 0;
    cout << "Number of submatrices with sum " << target << ": " << numSubmatrixSumTarget(matrix, target) << endl;
    return 0;
}
