# 使用说明

1.把config.json-模板文件重命名为config.json。

2.在阿里云控制台申请accesskey，并在config.json中填写对应的id和secret。

3.在域名控制台中，在需要解析的域名的解析页，手动添加一个解析记录，IP随便写。然后打开Chrome的网络调试。刷新页面，在Network找dns/rr_api/list.json（版本更新可能会是其他的请求地址）。找到对应的解析记录的recordId，然后更新到config.json中。

4.如果是windows系统，请先安装curl。如果是mac/linux，可能已经有curl了。请手动确认是否安装。

5.启动aliddns（.exe）即可。

# 编译说明

1. go get 安装依赖。如果连不上服务器，先go env -w GOPROXY=https://goproxy.cn使用代理。

2. go run aliddns.go 测试程序。

3. go build aliddns.go 编译得到二进制程序。