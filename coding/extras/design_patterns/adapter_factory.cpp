#include <iostream>
#include <bits/stdc++.h>
using namespace std;

class OldReqHandler{
    public:
        void oldRequest(){
            cout<<" Handling req by old req handler "<<endl;
        }
};

class ModernInterface{
    public: 
        virtual void request() = 0;
};

class Adapter: public ModernInterface {
    public:
        Adapter (OldReqHandler o) : adaptee(o){}
        void request() override{
            adaptee.oldRequest();
            cout<<" succesfully sent the req to old handler "<<endl;
        }
    private:
        OldReqHandler adaptee;
};

class AdapterFactory{
    public:
        static unique_ptr<ModernInterface> createAdapter(OldReqHandler o){
            return std::make_unique<Adapter>(o);
        }
};

int main(){
    
    OldReqHandler o;
    auto a = AdapterFactory::createAdapter(o);
    a->request();
}