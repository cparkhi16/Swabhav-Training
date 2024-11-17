#include <stdio.h>

int main(int argc, char *argv[]) {
    // Check if there are enough arguments
    if (argc > 1) {
        // Print the first character of the second command-line argument
        printf("%c\n", **++argv); 
    } else {
        // If no arguments are passed, print a message
        printf("No command-line arguments passed.\n");
    }

    return 0;
}
