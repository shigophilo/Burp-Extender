接收到的数据包转为base64

排除的host	excludes.txt
-----------------------------------------------------------------
获取header	将请求头中的所有字段都保留下来 header.txt
接收url		保存在url命名的文件夹下 根据提交方法命名 uris_get.txt uris_POST.txt	注意去重 lock
参数		保存在url命名的文件夹下 parameters.txt	注意去重 lock	json xml 普通
------------------------------------------------------------------此处应该写成拓展,yaml文件
判断Accept	json格式的加逗号判断是不是fastjson
cookie		判断是不是shiro

危险利用提示可以弹窗

对接urlfinfer
