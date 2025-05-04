// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

class IUser{
    public:
    virtual int getUserId() = 0;
    virtual string getUserName() = 0;
};

class User: public IUser{
    public:
    User (string name, int id ): _name(name), _id(id){};
    
    int getUserId() override{
        return _id;
    }
    
    string getUserName() override {
        return _name;
    }
    
    private:
    string _name;
    int _id;
};

class IBook{
    public:
    virtual int getBookId() = 0;
    virtual string getBookName() = 0;
    virtual bool isAvailable() = 0;
    virtual void borrowBook() = 0;
    virtual void returnBook() = 0;
};

class Book : public IBook{
    public:
     Book(string name , int id ): _name(name), _id(id), available(true){}
     
     int getBookId() override {
         return _id;
     }
     string getBookName() override{
         return _name;
     }
     bool isAvailable() override{
         return available;
     }
     void borrowBook() override{
         available = false;
     }
     void returnBook() override{
         available = true;
     }
    private:
    string _name;
    int _id;
    bool available;
};

class ILibraryOperation{
    public:
    virtual void execute(int bookId, int userId, unordered_map<int , unique_ptr<IBook>>& books ) = 0;
    virtual ~ILibraryOperation() = default;
};

class BorrowOperation : public ILibraryOperation{
    public:
    BorrowOperation() {}
    virtual void execute(int bookId, int userId, unordered_map<int , unique_ptr<IBook>>& books ) override {
        if(books.find(bookId) != books.end()){
            auto b = books[bookId].get();
            if(b->isAvailable()){
                b->borrowBook();
                cout<<" User "<<userId<<" borrowed book "<<bookId<<endl;
            }else{
                cout<<" Book is not available "<<endl;
            }
        }else{
            cout<<" Book with Id "<<bookId<<" not found "<<endl;
        }
    }
};

class ReturnOperation : public ILibraryOperation{
    public:
    ReturnOperation() {}
    virtual void execute(int bookId, int userId, unordered_map<int , unique_ptr<IBook>>& books ) override {
        if(books.find(bookId) != books.end()){
            auto b = books[bookId].get();
            if(!b->isAvailable()){
                b->returnBook();
                cout<<" User "<<userId<<" returned book "<<bookId<<endl;
            }else{
                cout<<" Book wasn't borrowed "<<endl;
            }
        }else{
            cout<<" Book with Id "<<bookId<<" not found "<<endl;
        }
    }
};

class LibraryOpFactory{
    public:
    static unique_ptr<ILibraryOperation> createLibraryOperation(string s){
        if (s== "return"){
            return make_unique<ReturnOperation>();
        }else if(s== "borrow"){
            return make_unique<BorrowOperation>();
        }
        return nullptr;
    }
};

class Library{
    private:
    unordered_map<int ,unique_ptr<IBook>> books;
    unordered_map<int , unique_ptr<IUser>> users;
    public:
    
    Library(){}
    
    void addUser(int userId, string name){
        users[userId] = make_unique<User>(name,userId );
    }
    void addBook(int bookId, string name){
        books[bookId] = make_unique<Book>(name ,bookId );
    }
    
    void performOperation(string op , int userId , int bookId){
        auto opPtr = LibraryOpFactory::createLibraryOperation(op);
        
        if(opPtr){
            opPtr->execute(bookId, userId , books);
        }
    }
    
    void showBooks(){
        for(auto const& b : books){
            cout<<" Book id "<<b.first<<" book name "<<b.second->getBookName()<<endl;
        }
    }
};
int main() {
    Library lib;
    
    lib.addUser(1, "Chinmay");
    lib.addUser(2,"Keyur");
    
    lib.addBook(101, "ABCD");
    lib.addBook(102, "DEF");
    
    lib.performOperation("borrow", 1, 101);
    lib.performOperation("borrow", 2, 101);
    lib.performOperation("return", 1, 101);
    lib.performOperation("return", 1, 101);
    lib.performOperation("borrow", 2, 101);
    lib.showBooks();
    return 0;
}