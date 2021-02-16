# business
cd cmd/business
rm -f business
go build -o business main.go
echo "打包business成功"
pkill business
echo "停止business服务"
nohup ./business &
echo "启动business服务"

# logic
cd ../logic
rm -f logic
go build -o logic main.go
echo "打包logic成功"
pkill logic
echo "停止logic服务"
nohup ./logic &
echo "启动logic服务"

# tcp_conn
cd ../tcp_conn
rm -f tcp_conn
go build -o tcp_conn main.go
echo "打包tcp_conn成功"
pkill tcp_conn
echo "停止tcp_conn服务"
nohup ./tcp_conn &
echo "启动tcp_conn服务"

