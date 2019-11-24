#include "stdafx.h"
#include "Poco/NotificationCenter.h"  
#include "Poco/Notification.h"  
#include "Poco/Observer.h"  
#include "Poco/NObserver.h"  
#include "Poco/AutoPtr.h"  
#include   
using Poco::NotificationCenter;  
using Poco::Notification;  
using Poco::Observer;  
using Poco::NObserver;  
using Poco::AutoPtr;  
class BaseNotification: public Notification  
{  
public: void dosome(){
        printf("fuck!");
    }
};  
class SubNotification: public BaseNotification  
{  

};  
  
  
class Target  
{  
public:  
    void handleBase(BaseNotification* pNf)  
    {  
        std::cout << "handleBase: " << pNf->name() << std::endl;  
        pNf->dosome();
        pNf->release(); // we got ownership, so we must release  
    }  
    void handleSub(const AutoPtr& pNf)  
    {  
        std::cout << "handleSub: " << pNf->name() << std::endl;  
    }  
};  
  
  
int main(int argc, char** argv)  
{  
    NotificationCenter nc;  
    Target target;  
    nc.addObserver(  
        Observer(target, &Target::handleBase)  
        );  
    nc.addObserver(  
        NObserver(target, &Target::handleSub)  
        );  
    nc.postNotification(new BaseNotification);  
    nc.postNotification(new SubNotification);  
    nc.removeObserver(  
        Observer(target, &Target::handleBase)  
        );  
    nc.removeObserver(  
        NObserver(target, &Target::handleSub)  
        );  
    return 0;  
}  