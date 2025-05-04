#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <memory>
#include <algorithm>
#include <ctime>

// 1. Enums
enum class BookStatus {
    AVAILABLE,
    BORROWED,
    RESERVED,
    LOST,
    DAMAGED
};

// 2. Book (metadata)
class Book {
public:
    std::string title;
    std::string author;
    std::string subject;
    int publicationYear;

    Book(std::string t, std::string a, std::string s, int y)
        : title(std::move(t)), author(std::move(a)), subject(std::move(s)), publicationYear(y) {}
};

// 3. BookItem (copy)
class BookItem {
    int barcodeId;
    Book* bookInfo;
    BookStatus status;
    int borrowerId;
    std::time_t borrowTime;

public:
    BookItem(int barcode, Book* info)
        : barcodeId(barcode), bookInfo(info), status(BookStatus::AVAILABLE),
          borrowerId(-1), borrowTime(0) {}

    int getBarcode() const { return barcodeId; }
    Book* getBookInfo() const { return bookInfo; }
    BookStatus getStatus() const { return status; }
    int getBorrowerId() const { return borrowerId; }

    bool isAvailable() const { return status == BookStatus::AVAILABLE; }

    void borrow(int userId) {
        status = BookStatus::BORROWED;
        borrowerId = userId;
        borrowTime = std::time(nullptr);
    }

    void returnBook() {
        status = BookStatus::AVAILABLE;
        borrowerId = -1;
        borrowTime = 0;
    }
};

// 4. IUser and its extensions
class IUser {
public:
    virtual std::string getName() const = 0;
    virtual int getId() const = 0;
    virtual ~IUser() = default;
};

class Customer : public IUser {
    int id;
    std::string name;
    std::vector<int> borrowedItems; // barcodes

public:
    Customer(int id, std::string name) : id(id), name(std::move(name)) {}

    std::string getName() const override { return name; }
    int getId() const override { return id; }

    void borrowBook(int barcodeId) { borrowedItems.push_back(barcodeId); }
    void returnBook(int barcodeId) {
        borrowedItems.erase(std::remove(borrowedItems.begin(), borrowedItems.end(), barcodeId), borrowedItems.end());
    }

    int getBorrowedBookCount() const { return borrowedItems.size(); }
};

class Librarian : public IUser {
    int id;
    std::string name;

public:
    Librarian(int id, std::string name) : id(id), name(std::move(name)) {}

    std::string getName() const override { return name; }
    int getId() const override { return id; }
};

// 5. Library System
class Library {
    int nextBarcode = 1000;
    std::map<int, std::unique_ptr<BookItem>> bookItems;   // barcode -> BookItem
    std::map<std::string, std::unique_ptr<Book>> books;   // title -> Book
    std::map<int, std::unique_ptr<IUser>> users;          // userId -> user

    const int maxBorrowLimit = 5;

public:
    void addUser(std::unique_ptr<IUser> user) {
        users[user->getId()] = std::move(user);
    }

    void addBook(const std::string& title, const std::string& author,
                 const std::string& subject, int year, int copies = 1) {

        if (books.find(title) == books.end()) {
            books[title] = std::make_unique<Book>(title, author, subject, year);
        }

        for (int i = 0; i < copies; ++i) {
            int barcode = nextBarcode++;
            bookItems[barcode] = std::make_unique<BookItem>(barcode, books[title].get());
            std::cout << "[Library] Added copy barcode: " << barcode << " for book: " << title << "\n";
        }
    }

    void borrowBook(int barcodeId, int userId) {
        if (bookItems.find(barcodeId) == bookItems.end() || users.find(userId) == users.end()) {
            std::cout << "[Error] Invalid barcode/user\n";
            return;
        }

        BookItem* item = bookItems[barcodeId].get();
        auto* customer = dynamic_cast<Customer*>(users[userId].get());

        if (!customer) {
            std::cout << "[Error] Only customers can borrow books.\n";
            return;
        }

        if (!item->isAvailable()) {
            std::cout << "[Error] Book not available\n";
            return;
        }

        if (customer->getBorrowedBookCount() >= maxBorrowLimit) {
            std::cout << "[Error] Borrow limit exceeded.\n";
            return;
        }

        item->borrow(userId);
        customer->borrowBook(barcodeId);

        std::cout << "[Borrowed] Book: " << item->getBookInfo()->title
                  << ", Barcode: " << barcodeId
                  << ", Borrower: " << customer->getName() << "\n";
    }

    void returnBook(int barcodeId, int userId) {
        if (bookItems.find(barcodeId) == bookItems.end() || users.find(userId) == users.end()) {
            std::cout << "[Error] Invalid barcode/user\n";
            return;
        }

        BookItem* item = bookItems[barcodeId].get();
        auto* customer = dynamic_cast<Customer*>(users[userId].get());

        if (!customer) {
            std::cout << "[Error] Only customers can return books.\n";
            return;
        }

        if (item->getBorrowerId() != userId) {
            std::cout << "[Error] This book was not borrowed by you.\n";
            return;
        }

        item->returnBook();
        customer->returnBook(barcodeId);

        std::cout << "[Returned] Book: " << item->getBookInfo()->title
                  << ", Barcode: " << barcodeId
                  << ", by: " << customer->getName() << "\n";
    }

    void showAllBooks() {
        std::cout << "----------------------------------------\n";
        std::cout << "Current Books in Library:\n";
        for (const auto& [barcode, item] : bookItems) {
            std::cout << "Barcode: " << barcode
                      << ", Title: " << item->getBookInfo()->title
                      << ", Status: " << (item->isAvailable() ? "Available" : "Borrowed") << "\n";
        }
        std::cout << "----------------------------------------\n";
    }
};

// 6. Main driver
int main() {
    Library lib;

    // Add users
    auto librarian = std::make_unique<Librarian>(1, "Mr. Smith");
    auto customer = std::make_unique<Customer>(2, "Alice");

    lib.addUser(std::move(librarian));
    lib.addUser(std::move(customer));

    // Librarian adds 2 copies of the same book
    lib.addBook("The Great Gatsby", "F. Scott Fitzgerald", "Fiction", 1925, 2);

    lib.showAllBooks();

    // Customer borrows two copies
    lib.borrowBook(1000, 2); // Alice borrows first copy
    lib.borrowBook(1001, 2); // Alice borrows second copy
    lib.borrowBook(1001, 2);

    lib.showAllBooks();

    // Customer returns one copy
    lib.returnBook(1000, 2);

    lib.showAllBooks();

    return 0;
}
