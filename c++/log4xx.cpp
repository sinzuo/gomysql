#include <log4cxx/logger.h>
#include <log4cxx/basicconfigurator.h>
#include <log4cxx/helpers/exception.h>
#include <log4cxx/propertyconfigurator.h>

using namespace log4cxx;
using namespace log4cxx::helpers;

LoggerPtr logger_file(Logger::getLogger("file"));//获取配置文件中file对应的句柄
LoggerPtr logger_file(Logger::getLogger("console"));//获取配置文件中console对应的句柄

int main(){
    PropertyConfigurator::configure("./log4cxx.properties");//加载配置文件，下面会细说
    LOG4CXX_INFO(logger_file,"This is a test");
  
    return 0;
}