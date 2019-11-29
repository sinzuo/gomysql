#include <iostream>

using namespace std;



template <typename T>
class Complex{
    
public:
    //构造函数
    Complex(T a, T b)
    {
        this->a = a;
        this->b = b;
    }
    
    //运算符重载
    Complex<T> operator+(Complex &c)
    {
        Complex<T> tmp(this->a+c.a, this->b+c.b);
        return tmp;
    }
        

    T a;
    T b;
};

int main()
{
    //对象的定义，必须声明模板类型，因为要分配内容
    Complex<int> a(10,20);  
    Complex<int> b(20,30);
    Complex<int> c = a + b;

    cout<< c.a <<endl;
    
    return 0;
}



