#include "stdafx.h"
#include "Poco/ThreadPool.h"
#include "Poco/Runnable.h"
#include 
class HelloRunnable: public Poco::Runnable
{
    virtual void run()
    {
        std::cout << "Hello, bingzhe" << std::endl;
    }
};
int main(int argc, char** argv)
{
    HelloRunnable runnable;
    Poco::ThreadPool::defaultPool().start(runnable);
    Poco::ThreadPool::defaultPool().joinAll();
    return 0;
}