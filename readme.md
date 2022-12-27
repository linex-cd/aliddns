#使用说明

1.把config.json-模板文件重命名为config.json

2.在阿里云控制台申请accesskey，并在config.json中填写对应的id和secret

3.在域名控制台中，在需要解析的域名的解析页，手动添加一个解析记录，IP随便写。然后打开Chrome的网络调试。刷新页面，在Network找dns/rr_api/list.json（版本更新可能会是其他的请求地址）。找到对应的解析记录的recordId，然后更新到config.json中。

4.启动exe即可。
