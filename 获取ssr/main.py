# _*_ coding:utf-8 _*_
import configparser
import os
import time


def createConfigFile(c):

    ProxySwitch = "on"
    ProxyList = ["ssr","v2ray","ss"]


    print("是否从网络上自动获取http代理爬虫。 过程可能会比较长，处理速度和代理有关系。")
    print("1 不需要，我有代理。")
    print("2 需要哦。")
    print("回车键默认配置。")
    while True:
        proxyChoose = input()
        if proxyChoose == "1":
            ProxySwitch = "off"
            break
        if proxyChoose == "2":
            break
        if proxyChoose == "":
            break
        print("请重新输入，1 自动获取代理， 2 不获取代理")

    print("请选择需要获取的代理类型。可多选")
    print("1 SSR")
    print("2 V2ray")
    print("3 SS")
    print("回车键默认配置。")
    while True:
        typeChoose = input()
        typeChoose = list(set(typeChoose))
        tempType = []
        if len(typeChoose) == 0:
            break
        if len(typeChoose) > 3:
            print("请重新输入,")
            continue
        if 3>=len(typeChoose) >0 :
             for i in typeChoose:
                 if i =="1":
                        tempType.append("ssr")
                 if i =="2":
                        tempType.append("v2ray")
                 if i =="3":
                        tempType.append("ss")
        if len(tempType) == 0:
            print("请重新输入")
            continue
        ProxyList = tempType
        break
 

    print("选择结果是，代理列表:",ProxyList)
    print("选择结果是，代理开启:",ProxySwitch)
  
            

def readConfig(c):
    pass



if __name__ == "__main__":
    '''
    检查配置文件
    生成配置文件
    程序入口
    :( 我傻叉了，刚才定义变量写成golang 的方式了。Orz  
    '''
    conf = configparser.ConfigParser()
    if not os.path.exists("./config/conf.ini"):
        print("初次使用，生成配置文件")
        time.sleep(0.5)
        createConfigFile(conf)
    readConfig(conf)
