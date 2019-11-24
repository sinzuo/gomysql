#include "Poco/Thread.h"  
#include "Poco/Runnable.h"  
#include   
class HelloRunnable: public Poco::Runnable  
{  
       virtual void run()  
       {  
            std::cout << "Hello, bingzhe!" << std::endl;  
       }  
};  
int main(int argc, char** argv)  
{  
       HelloRunnable runnable;  
       Poco::Thread thread;  
       thread.start(runnable);  
       thread.join();  
       return 0;  
}  