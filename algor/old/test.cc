//
// Created by nekoimi on 2021/6/24.
//
#include <ctime>
#include <iostream>
#include <exception>
#include <vector>
#include "test.h"

using namespace std;

#define DEBUG 1

#ifdef DEBUG

#endif

class Box {
public:
    virtual double getV() = 0;

protected:
    int width;
    int height;

    int getW() {
        return width;
    }

    int getH() {
        return height;
    }
};

class SmallBox : public Box {
public:
    double getV() {
        return (getW() + getH()) * 2;
    }

    SmallBox(int w, int h) {
        this->width = w;
        this->height = h;
        cout << "构造函数" << endl;
    }

    ~SmallBox() {
        cout << "析构函数" << endl;
    }
};


class MyException : public exception {
    const char *what() const noexcept {
        return "MyException";
    }
};

namespace FirstSpace {
    void showSomeThing() {
        cout << "show some thing" << endl;
    }
}

template<typename T>
inline T const &Max(T const &a, T const &b) {
    return a < b ? a : b;
}

template<class T>
class TemplateClass {
};

// 泛型 Stack<T>

template<class T>
class Stack {
private:
    vector<T> items;
public:
    void push(T const &);

    T pop();

    bool empty() const {
        return items.empty();
    }
};

template<class T>
void Stack<T>::push(T const &item) {
    items.push_back(item);
}

template<class T>
T Stack<T>::pop() {
    if (items.empty()) {
        throw out_of_range("Stack<T>::pop Stack<T> is empty!");
    }

    return items.back();
}

int _main() {
    //    time_t t = time(0);
//    tm *_tm = localtime(&t);
//    cout << 1900 + _tm->tm_year << endl;
//
//    SmallBox *smallBox = new SmallBox(100, 100);
//    try {
//        smallBox->getV();
//    } catch (out_of_range) {
//
//    }
//    delete smallBox;
//
//    double *privateValue = NULL;
//    privateValue = new double;
//    *privateValue = 3.14;
//
//
//    FirstSpace::showSomeThing();
//
//    cout << "Max(1, 2)" << Max(1, 2) << endl;
//    cout << "Max(1.5, 2.5)" << Max(1.5, 2.5) << endl;
//    cout << "Max('hello', 'world')" << Max("Hello", "World") << endl;
//
//    Stack<int> intStack;
//    Stack<string> stringStack;
//    intStack.push(3);
//    intStack.push(2);
//    intStack.push(1);
//
//    stringStack.push("Nationality");
//    stringStack.push("国际");
//    stringStack.push("Land");
//    stringStack.push("陆地");
//    stringStack.push("HomeLand");
//    stringStack.push("祖国；家乡");
//
//    intStack.pop();
//    stringStack.pop();
    return 0;
}
