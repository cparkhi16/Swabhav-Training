#include <stdio.h>
#include <stdlib.h>

#define N 5  // Size of the array
#define M 10 // Number of integers to allocate space for

int main() {
    // Declare x as a pointer to an array of N integers
    int (*x)[N];

    // Dynamically allocate memory for M integers (treated as an array of N integers)
    x = (int(*)[N])malloc(M * sizeof(*x));  // Allocation for M integers

    // Check if memory allocation is successful
    if (x == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }

    // Print the sizes
    printf("Size of x: %zu\n", sizeof(x));       // Size of the pointer x itself
    printf("Size of *x: %zu\n", sizeof(*x));     // Size of what *x points to (array of N integers)

    // Free the allocated memory
    free(x);

    return 0;
}
