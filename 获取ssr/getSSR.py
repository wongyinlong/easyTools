import requests
import base64
import time

"""
 该文件可以独立运行，获取某个网站的信息， 功能单一，正在更改。
"""
from io import BytesIO
from bs4 import BeautifulSoup
from PIL import Image

url = "https://www.youneed.win/free-ssr"
#url = "http://www.baidu.com"
ssrDir = "C:\myProgramFiles\ShadowsocksR-4.7.0"  # ssr 目录，windows下
# data:image/jpeg;base64,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


def get_captcha(html):
    # 获取验证码
    soup = BeautifulSoup(html, "lxml")
    div_img = soup.find()
    img = div_img.find("img").get("src")
    img_data = img.split(",")[-1]
    binary_data = base64.b64decode(img_data)
    file = BytesIO(binary_data)
    image = Image.open(file)
    image.show()


def getContent():
    # 获取ssr链接list
    s = requests.session()
    resp = s.get(url)
    try:
        if resp.status_code == 200:
            get_captcha(resp.content)
            captcha = input("enter code ")
            resp = s.post(url, data={"passster_captcha": captcha})
            return resp.content
    except Exception as e:
        print(e)
        print(resp.status_code)
        return None


def getList(content):
    # 处理网页
    soup = BeautifulSoup(content, "lxml")
    table_tb = soup.find_all(attrs={'class': 'v2ray'})
    if len(table_tb) == 0 :
        print("验证码输入错误，或者没有获取到有效的网页内容。")
        quit()
    ssrs = []
    for item in table_tb:
        ssr = item.find("a").get("data")
        print(ssr)
        ssrs.append(ssr)
    return ssrs


def writToFile(ssrs):
    # 写到文件
    t = time.localtime()
    
    file_name = "./"+str(t.tm_year)+str(t.tm_mon)+str(t.tm_mday)+str(t.tm_hour)+str(t.tm_min)+str(t.tm_sec)+".txt"
    f = open(file_name, "w+") 
    for i in ssrs:
        f.write(i) 
        f.write("\n")
    f.close()


# def SSR(listSSR):
#     pass


# def CheckSSR():
#     pass


if __name__ == "__main__":
    '''
    先检查代理列表里可用的代理。
    必须先开启代理，然后再使用。因为爬去的网站不使用代理是访问不了的。
    只爬取ssr连接。
    功能包括爬取，去重，检查是否可用。每日点开一次自动更新ssr列表。
    '''
    content = getContent()
    ssr_list = getList(content)
    writToFile(ssr_list)
"""
3-05可以用了，白天把写文件写咯，其他功能再说吧。
"""


# -TODO: 添加功能，从其他网站获取ssr 和ss 以及v2ray.并输出xls表格(在控制台输出)初次使用需要进行配置，
# 配置选项为两个，1是否开启代理（自己获取网上的免费代理做），2获取代理类型，有ssr，ss，v2ray
# 使用python多线程做。多网站进行map reduce操作。