# 说明
多线程爬虫，从网络上获取ssr，ss，v2ray的配置项。（暂没有代理可用性检测功能，之后会更新上。）

尴尬的是，国内好像搜不到可以用的https代理，我每天会push，大家每天pull吧。

# 该软件功能如下
1、 配置文件：初次使用，可直接修改配置文件。
    
    proxy：配置是否使用代理。（因为大多数的共享网站被墙了。）
        on：从网络上获取http代理列表。
        off：在使用代理的情况下，可以不用开启。
        example：
            proxy:on

    代理类型：列表 
    list：
        [ssr,ss,v2ray]
        example:
        list:[ssr,ss,v2ray]
    
    输出类型

2、 输出到文件。
    输出到表格文件和控制台。会自动进行对比，过滤掉已存在的，并标志更新。以保证最新获取的一直在文件顶部。（后期会加入自动测试是否可用功能）

3、 未来。
    针对某一些网站，采用验证码识别，甚至是机器学习的功能来获取数据。（哎，机器学习的机器好贵的，没钱啊）

# 我

博客： https://www.cnblogs.com/Leon-The-Professional/

邮箱： wongyinlong@yeah.net

最近在写自己的一个项目，代理原理研究的文档可能要放一放了。尽量五月底前写了吧。好慢啊我。还要学英语，有人教我英语吗？我叫你汉语呀。