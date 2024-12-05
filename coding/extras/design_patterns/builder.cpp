#include<bits/stdc++.h>

using namespace std;

class Car{
    public:
        Car(){}
        void setEngine(std::string e){
            engine_ = e;
        }
        void setWheels(std::string w){
            wheels_ = w;
        }
         void showSpecifications() const {
            std::cout << "Car Specifications:\n"
                  << "Engine: " << engine_ << "\n"
                  << "Wheels: " << wheels_ << "\n";         
        }
    private:
        std::string wheels_;
        std::string engine_;
};

class CarBuilder{
    public:
        virtual void buildEngine() = 0;
        virtual void buildWheels() = 0;
        virtual std::shared_ptr<Car> getCar() = 0;
};

class SUVCarBuilder : public CarBuilder{
    public:
    SUVCarBuilder(){
        car_ = make_shared<Car>();
    }

    void buildEngine() override{
        car_->setEngine("ABC");
    }

    void buildWheels() override{
        car_->setWheels("XYZ");
    }

    std::shared_ptr<Car> getCar(){
        return car_;
    }
    private:
    std::shared_ptr<Car> car_;
};

class SportsCarBuilder : public CarBuilder{
    public:
    SportsCarBuilder(){
        car_ = make_shared<Car>();
    }

    void buildEngine() override{
        car_->setEngine("Sports SABCX");
    }

    void buildWheels() override{
        car_->setWheels("sPORTS XYZ");
    }

    std::shared_ptr<Car> getCar(){
        return car_;
    }
    private:
    std::shared_ptr<Car> car_;
};

class CarDirector{

    public:
        void setBuilder(shared_ptr<CarBuilder> builder){
            builder_ = builder;
        }

        void constructCar(){
            builder_->buildEngine();
            builder_->buildWheels();
        }
    private:
        std::shared_ptr<CarBuilder> builder_;
};
int main(){

    CarDirector cd;

    std::shared_ptr<SUVCarBuilder> scb = make_shared<SUVCarBuilder>();
    cd.setBuilder(scb);
    cd.constructCar();
    std::shared_ptr<Car> sCar = scb->getCar();
    sCar->showSpecifications();

    std::shared_ptr<CarBuilder> spCB = make_shared<SportsCarBuilder>();
    cd.setBuilder(spCB);
    cd.constructCar();
    std::shared_ptr<Car> sportsCar = spCB->getCar();
    sportsCar->showSpecifications();
}