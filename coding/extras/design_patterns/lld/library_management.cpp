#include <iostream>
#include <string>
#include <map>
#include <memory>

// ---------------- Interface for Book ----------------
class IBook {
public:
    virtual std::string getTitle() const = 0;
    virtual int getId() const = 0;
    virtual bool isAvailable() const = 0;
    virtual void borrowBook() = 0;
    virtual void returnBook() = 0;
    virtual ~IBook() = default;
};

// ---------------- Book Implementation ----------------
class Book : public IBook {
    int id;
    std::string title;
    bool available;

public:
    Book(int id, std::string title) : id(id), title(std::move(title)), available(true) {}
    std::string getTitle() const override { return title; }
    int getId() const override { return id; }
    bool isAvailable() const override { return available; }
    void borrowBook() override { available = false; }
    void returnBook() override { available = true; }
};

// ---------------- Interface for User ----------------
class IUser {
public:
    virtual std::string getName() const = 0;
    virtual int getId() const = 0;
    virtual ~IUser() = default;
};

// ---------------- User Implementation ----------------
class User : public IUser {
    int id;
    std::string name;

public:
    User(int id, std::string name) : id(id), name(std::move(name)) {}
    std::string getName() const override { return name; }
    int getId() const override { return id; }
};

// ---------------- Strategy Pattern for Operations ----------------
class ILibraryOperation {
public:
    virtual void execute(int bookId, int userId, std::map<int, std::unique_ptr<IBook>>& books) = 0;
    virtual ~ILibraryOperation() = default;
};

class BorrowBookOperation : public ILibraryOperation {
public:
    void execute(int bookId, int userId, std::map<int, std::unique_ptr<IBook>>& books) override {
        if (books.find(bookId) != books.end()) {
            IBook* book = books[bookId].get();
            if (book->isAvailable()) {
                book->borrowBook();
                std::cout << "User " << userId << " borrowed book: " << book->getTitle() << "\n";
            } else {
                std::cout << "Book already borrowed.\n";
            }
        } else {
            std::cout << "Book ID not found.\n";
        }
    }
};

class ReturnBookOperation : public ILibraryOperation {
public:
    void execute(int bookId, int userId, std::map<int, std::unique_ptr<IBook>>& books) override {
        if (books.find(bookId) != books.end()) {
            IBook* book = books[bookId].get();
            if (!book->isAvailable()) {
                book->returnBook();
                std::cout << "User " << userId << " returned book: " << book->getTitle() << "\n";
            } else {
                std::cout << "Book wasn't borrowed.\n";
            }
        } else {
            std::cout << "Book ID not found.\n";
        }
    }
};

// ---------------- Factory Pattern ----------------
class LibraryOperationFactory {
public:
    static std::unique_ptr<ILibraryOperation> createOperation(const std::string& type) {
        if (type == "borrow")
            return std::make_unique<BorrowBookOperation>();
        else if (type == "return")
            return std::make_unique<ReturnBookOperation>();
        else
            return nullptr;
    }
};

// ---------------- Library System ----------------
class Library {
    std::map<int, std::unique_ptr<IBook>> books;
    std::map<int, std::unique_ptr<IUser>> users;

public:
    void addBook(int id, const std::string& title) {
        books[id] = std::make_unique<Book>(id, title);
        std::cout << "Added book: " << title << "\n";
    }

    void addUser(int id, const std::string& name) {
        users[id] = std::make_unique<User>(id, name);
        std::cout << "Added user: " << name << "\n";
    }

    void performOperation(const std::string& operation, int bookId, int userId) {
        auto op = LibraryOperationFactory::createOperation(operation);
        if (op) {
            op->execute(bookId, userId, books);
        } else {
            std::cout << "Invalid operation.\n";
        }
    }

    void showBooks() const {
        for (const auto& [id, book] : books) {
            std::cout << "BookID: " << id << ", Title: " << book->getTitle()
                      << ", Available: " << (book->isAvailable() ? "Yes" : "No") << "\n";
        }
    }
};
int main() {
    Library lib;
    lib.addUser(1, "Alice");
    lib.addBook(101, "The Great Gatsby");
    lib.addBook(102, "1984");

    lib.showBooks();

    lib.performOperation("borrow", 101, 1);
    lib.performOperation("borrow", 101, 1); // Already borrowed
    lib.performOperation("return", 101, 1);
    lib.performOperation("return", 101, 1); // Already returned

    lib.showBooks();
    return 0;
}

