## Openresty-Best-Practices

《OpenResty 最佳实践》代码实践 

原书地址： [openresty-best-practices](https://moonbingbing.gitbooks.io/openresty-best-practices)

### 安装

从[官网](http://openresty.org/en/download.html)下载最新的源码包:

```bash
## 将压缩包解压
tar -xvf openresty-VERSION.tar.gz

## 切换到解压后的目录
cd openresty-VERSION/

## 根据需求指定cofigure选项(还可以--add-module=PATH添加自己的第三方Nginx C模块)
./configure --prefix=/opt/openresty --with-pcre-jit --without-http_redis2_module --with-http_iconv_module -j8

## 编译
make -j8

## 安装
sudo make install
```

修改`/etc/profile`，配置环境变量后保存，通过source使其生效:

```bash
OPENRESTY_HOME=/opt/openresty
LUAJIT_HOME=/opt/openresty/luajit
export PATH=$PATH:$OPENRESTY_HOME/bin:$LUAJIT_HOME/bin
```

### 推荐阅读

[agentzh的Nginx教程](https://openresty.org/download/agentzh-nginx-tutorials-zhcn.html)

[programming-openresty-zh](https://github.com/iresty/programming-openresty-zh)

[Nginx开发从入门到精通](http://tengine.taobao.org/book/index.html)