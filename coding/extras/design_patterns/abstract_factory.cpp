#include<iostream>
#include<bits/stdc++.h>
using namespace std;

class Button {
    public:
        virtual void render() = 0;
        virtual ~Button() = default;
};

class Checkbox{
    public:
        virtual void toggle() = 0;
        virtual ~Checkbox() = default;
};

class WindowsButton : public Button{
    public:
        void render() override{
            cout<<" Rendering windows btn "<<endl;
        }
};

class MacOsButton : public Button {
    public:
        void render() override{
            cout<<" Rendering mac btn "<<endl;
        }
};

class WindowsCheckbox : public Checkbox{
    public:
        void toggle() override{
            cout<<" Toggling windows checkbox "<<endl;
        }
};

class MacOsCheckbox : public Checkbox{
    public:
        void toggle() override{
            cout<<" Toggling macos checkbox "<<endl;
        }
};

class GUIFactory{
    public:
        virtual std::unique_ptr<Button> createButton() = 0;
        virtual std::unique_ptr<Checkbox> createCheckbox() = 0;
};

class WindowsFactory : public GUIFactory{
    public:
        std::unique_ptr<Button> createButton() override{
            return make_unique<WindowsButton>();
        }
        std::unique_ptr<Checkbox> createCheckbox() override{
            return make_unique<WindowsCheckbox>();
        }
};

class MacOsFactory : public GUIFactory{
    public:
        std::unique_ptr<Button> createButton() override{
            return make_unique<MacOsButton>();
        }
        std::unique_ptr<Checkbox> createCheckbox() override{
            return make_unique<MacOsCheckbox>();
        }
};

int main(){

    std::unique_ptr<GUIFactory> factory;

    #ifdef _WIN32
        factory = make_unique<WindowsFactory>();
    #else
        factory = make_unique<MacOsFactory>();
    #endif

    auto btn = factory->createButton();
    auto checkbox = factory->createCheckbox();

    btn->render();
    checkbox->toggle();
    cout<<"end";
}