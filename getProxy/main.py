# _*_ coding:utf-8 _*_
import configparser
import os
import time

proxyListDir = "./proxy.list"
aimWebSit = "https://twitter.com"
proxyWebSits = [
    "https://www.website1.com"
    "https://www.website1.com"
]

total = 20  # 代理数量
ping = 100   # 代理延迟低于100ms

def getProxyFormWebSite():
    """
    从网站获取代理，并检测是否可用
    """
    pass

def writeProxyToFile():
    """
    写代理到文件
    """
    pass

def readProxyFromFile():
    """
    读代理从文件
    """
    pass
                
def checkProxy():
    """
    检查代理
    """
    pass



if __name__ == "__main__":
    '''
    检查本地文件 proxy.list
    如果没有则从网上获取httpsProxy，socks5Proxy。
    获取20个写入文件
    :( 我傻叉了，刚才定义变量写成golang 的方式了。Orz  
    '''
    proxyListFileExist =  os.path.exists(proxyListDir)
    if proxyListFileExist:
        # 读文件到代理列表
        pass
    # 开始检查代理列表，并计数
    # 代理列表长度是否大于等于20(total)
    # 若不足够，则从网页上获取
    # 写入代理到文件
    # 开始多线程


