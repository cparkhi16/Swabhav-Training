#include <iostream>
#include <vector>
#include <string>

enum class Symbol { Empty, X, O };

class Player {
    std::string name;
    Symbol symbol;
public:
    Player(const std::string& name, Symbol symbol) : name(name), symbol(symbol) {}
    Symbol getSymbol() const { return symbol; }
    std::string getName() const { return name; }
};

class Board {
    std::vector<std::vector<Symbol>> grid;
public:
    Board() : grid(3, std::vector<Symbol>(3, Symbol::Empty)) {}

    bool makeMove(int row, int col, Symbol symbol) {
        if (row < 0 || row >= 3 || col < 0 || col >= 3 || grid[row][col] != Symbol::Empty)
            return false;
        grid[row][col] = symbol;
        return true;
    }

    void display() const {
        for (const auto& row : grid) {
            for (const auto& cell : row) {
                char c = cell == Symbol::X ? 'X' : (cell == Symbol::O ? 'O' : '.');
                std::cout << c << " ";
            }
            std::cout << "\n";
        }
    }

    bool checkWin(Symbol symbol) const {
        // Rows and Columns
        for (int i = 0; i < 3; ++i) {
            if ((grid[i][0] == symbol && grid[i][1] == symbol && grid[i][2] == symbol) ||
                (grid[0][i] == symbol && grid[1][i] == symbol && grid[2][i] == symbol))
                return true;
        }
        // Diagonals
        if ((grid[0][0] == symbol && grid[1][1] == symbol && grid[2][2] == symbol) ||
            (grid[0][2] == symbol && grid[1][1] == symbol && grid[2][0] == symbol))
            return true;
        return false;
    }

    bool isFull() const {
        for (const auto& row : grid)
            for (auto cell : row)
                if (cell == Symbol::Empty)
                    return false;
        return true;
    }
};

class Game {
    Board board;
    Player player1, player2;
    Player* currentPlayer;
public:
    Game(const Player& p1, const Player& p2) : player1(p1), player2(p2), currentPlayer(nullptr) {}

    void start() {
        currentPlayer = &player1;
        while (true) {
            board.display();
            std::cout << currentPlayer->getName() << "'s turn (" << (currentPlayer->getSymbol() == Symbol::X ? 'X' : 'O') << ")\n";
            int row, col;
            std::cout << "Enter row and column (0-based): ";
            std::cin >> row >> col;

            if (!board.makeMove(row, col, currentPlayer->getSymbol())) {
                std::cout << "Invalid move. Try again.\n";
                continue;
            }

            if (board.checkWin(currentPlayer->getSymbol())) {
                board.display();
                std::cout << currentPlayer->getName() << " wins!\n";
                break;
            }

            if (board.isFull()) {
                board.display();
                std::cout << "It's a draw!\n";
                break;
            }

            switchPlayer();
        }
    }

private:
    void switchPlayer() {
        currentPlayer = (currentPlayer == &player1) ? &player2 : &player1;
    }
};

int main() {
    Player p1("Player 1", Symbol::X);
    Player p2("Player 2", Symbol::O);
    Game game(p1, p2);
    game.start();
    return 0;
}
