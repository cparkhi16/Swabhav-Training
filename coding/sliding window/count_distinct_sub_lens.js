function countDistinctElements(arr, k) {
    let windowStart = 0;
    let freqMap = {};
    let count = [];

    for (let windowEnd = 0; windowEnd < arr.length; windowEnd++) {
        const endElement = arr[windowEnd];
        
        // Add the current element to the frequency map
        if (!(endElement in freqMap)) {
            freqMap[endElement] = 0;
        }
        freqMap[endElement]++;
        
        // If we've reached the size of the window
        if (windowEnd >= k - 1) {
            // Count distinct elements
            count.push(Object.keys(freqMap).length);
            
            // Remove the element going out of the window
            const startElement = arr[windowStart];
            freqMap[startElement]--;
            if (freqMap[startElement] === 0) {
                delete freqMap[startElement];
            }
            windowStart++; // Slide the window forward
        }
    }

    return count;
}

// Test Cases
console.log(countDistinctElements([1, 2, 3, 2, 1, 4, 5], 3)); // [3, 2, 3, 3, 3]
console.log(countDistinctElements([4, 1, 1, 1, 2, 3, 4], 4)); // [2, 3, 3, 4]
