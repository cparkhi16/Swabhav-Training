#include <iostream>
#include <vector>
#include <queue>

using namespace std;

// BFS function definition
void bfs(vector<vector<int>>& image, int sr, int sc, int originalColor, int newColor) {
    queue<pair<int, int>> q;
    q.push({sr, sc});
    image[sr][sc] = newColor;

    vector<pair<int, int>> directions = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};

    while (!q.empty()) {
        auto [x, y] = q.front();
        q.pop();

        for (auto [dx, dy] : directions) {
            int newX = x + dx;
            int newY = y + dy;

            if (newX >= 0 && newX < image.size() && newY >= 0 && newY < image[0].size() && image[newX][newY] == originalColor) {
                image[newX][newY] = newColor;
                q.push({newX, newY});
            }
        }
    }
}
// Flood fill function using BFS
vector<vector<int>> floodFillBFS(vector<vector<int>>& image, int sr, int sc, int newColor) {
    int originalColor = image[sr][sc];
    if (originalColor != newColor) {
        bfs(image, sr, sc, originalColor, newColor);
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

    floodFillBFS(image, sr, sc, newColor);

    cout << "Image after flood fill (BFS):" << endl;
    printImage(image);

    return 0;
}
