#include <iostream>
#include <vector>
#include <queue>
#include <cstdlib>
#include <ctime>

class Player {
    std::string name;
    int position;

public:
    Player(const std::string& name) : name(name), position(0) {}

    void setPosition(int pos) { position = pos; }
    int getPosition() const { return position; }
    std::string getName() const { return name; }
};

class Dice {
    int sides;

public:
    Dice(int sides = 6) : sides(sides) {}

    int roll() const {
        return 1 + (rand() % sides);
    }
};

struct Snake {
    int head;
    int tail;

    Snake(int h, int t) : head(h), tail(t) {}
};

struct Ladder {
    int start;
    int end;

    Ladder(int s, int e) : start(s), end(e) {}
};

class Board {
    int size;
    std::vector<Snake> snakes;
    std::vector<Ladder> ladders;

public:
    Board(int size) : size(size) {}

    void addSnake(const Snake& s) { snakes.push_back(s); }
    void addLadder(const Ladder& l) { ladders.push_back(l); }

    int getSize() const { return size; }

    int getNextPosition(int pos) const {
        for (const auto& s : snakes)
            if (s.head == pos) return s.tail;

        for (const auto& l : ladders)
            if (l.start == pos) return l.end;

        return pos;
    }
};

class Game {
    Board board;
    Dice dice;
    std::queue<Player> players;

public:
    Game(Board b, Dice d, const std::vector<Player>& pList)
        : board(b), dice(d) {
        for (auto& p : pList) players.push(p);
    }

    void play() {
        while (true) {
            Player current = players.front();
            players.pop();

            int roll = dice.roll();
            int newPos = current.getPosition() + roll;

            if (newPos > board.getSize()) {
                std::cout << current.getName() << " rolled a " << roll << " but stayed at " << current.getPosition() << " (overshoot)\n";
                players.push(current);
            } else {
                int finalPos = board.getNextPosition(newPos);
                std::cout << current.getName() << " rolled a " << roll
                          << " and moved from " << current.getPosition()
                          << " to " << finalPos << "\n";

                current.setPosition(finalPos);

                if (finalPos == board.getSize()) {
                    std::cout << current.getName() << " wins the game!\n";
                    break;
                }

                players.push(current);
            }
        }
    }
};

int main() {
    srand(static_cast<unsigned int>(time(0))); // Seed for randomness

    Board board(100);

    // Add snakes
    board.addSnake(Snake(99, 54));
    board.addSnake(Snake(70, 55));
    board.addSnake(Snake(52, 42));
    board.addSnake(Snake(25, 2));
    board.addSnake(Snake(95, 72));

    // Add ladders
    board.addLadder(Ladder(6, 25));
    board.addLadder(Ladder(11, 40));
    board.addLadder(Ladder(60, 85));
    board.addLadder(Ladder(46, 90));
    board.addLadder(Ladder(17, 69));

    // Create players
    std::vector<Player> players = {
        Player("Alice"),
        Player("Bob")
    };

    Dice dice;
    Game game(board, dice, players);
    game.play();

    return 0;
}
