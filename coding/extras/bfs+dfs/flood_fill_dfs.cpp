#include <iostream>
#include <vector>

using namespace std;

// DFS function definition
void dfs(vector<vector<int>>& image, int x, int y, int originalColor, int newColor) {
    if (x < 0 || x >= image.size() || y < 0 || y >= image[0].size() || image[x][y] != originalColor)
        return;

    image[x][y] = newColor;

    dfs(image, x + 1, y, originalColor, newColor); // Down
    dfs(image, x - 1, y, originalColor, newColor); // Up
    dfs(image, x, y + 1, originalColor, newColor); // Right
    dfs(image, x, y - 1, originalColor, newColor); // Left
}

// Flood fill function using DFS
vector<vector<int>> floodFillDFS(vector<vector<int>>& image, int sr, int sc, int newColor) {
    int originalColor = image[sr][sc];
    if (originalColor != newColor) {
        dfs(image, sr, sc, originalColor, newColor);
    }
    return image;
}

// Function to print the image matrix
void printImage(const vector<vector<int>>& image) {
    for (const auto& row : image) {
        for (int pixel : row) {
            cout << pixel << ' ';
        }
        cout << endl;
    }
}

int main() {
    vector<vector<int>> image = {
        {1, 1, 1},
        {1, 1, 0},
        {1, 0, 1}
    };
    int sr = 1, sc = 1, newColor = 2;

    cout << "Original image:" << endl;
    printImage(image);

    floodFillDFS(image, sr, sc, newColor);

    cout << "Image after flood fill (DFS):" << endl;
    printImage(image);

    return 0;
}
