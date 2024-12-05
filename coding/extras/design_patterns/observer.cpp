#include<iostream>
#include<bits/stdc++.h>

using namespace std;

class Observer{
    public:
    virtual ~Observer() = default;
    virtual void update() = 0;
};

class Subject{
    public:
        void attach(const shared_ptr<Observer>& ob){
            observers.push_back(ob);
        }

        void detach(const shared_ptr<Observer>& ob){
            observers.erase(remove(observers.begin() , observers.end(), ob),observers.end());
        }

        void notify(){
            for(auto o : observers){
                o->update();
            }
        }

    private:
        vector<shared_ptr<Observer>> observers;
};

class Stock: public Subject{
    public:
        void setPrice(float p){
            price = p;
            notify();
        }

        float getPrice(){
            return price;
        }
    private:
        float price = 0.0f;
};

class Investor : public Observer{
    public:
        Investor(string n , shared_ptr<Stock> s) : name(n) , stock(s){}

        void update(){
            cout<<" Investor : "<<name<<" notified for stock price  "<<stock->getPrice()<<endl;
        }
    private:
        string name;
        shared_ptr<Stock> stock;
};

int main(){
    auto st = make_shared<Stock>();

    auto in1 = make_shared<Investor>("Alice",st);
    auto in2 = make_shared<Investor>("Bob ",st);

    st->attach(in1);
    st->attach(in2);

    st->setPrice(101.12);
    st->setPrice(104.31);

    st->detach(in1);
    st->setPrice(150.9);
}