#!/bin/sh

function getIpAddr(){
	# 获取IP命令
	ipaddr=`ifconfig -a|grep inet|grep -v 127.0.0.1|grep -v inet6|awk '{print $2}'|tr -d "addr:"​`
	array=(`echo $ipaddr | tr '\n' ' '` ) 	# IP地址分割，区分是否多网卡
	#array=(172.20.32.214 192.168.1.10);
	num=${#array[@]}  						#获取数组元素的个数

	# 选择安装的IP地址
	if [ $num -eq 1 ]; then
		#echo "*单网卡"
		local_ip=${array[*]}
	elif [ $num -gt 1 ];then
		echo -e "\033[035m******************************\033[0m"
		echo -e "\033[036m*    请选择安装的IP地址		\033[0m"
		echo -e "\033[032m*      1 : ${array[0]}		\033[0m"
		echo -e "\033[034m*      2 : ${array[1]} 		\033[0m"
		echo -e "\033[035m******************************\033[0m"
		#选择需要安装的服务类型
		input=""
		while :
		do
			read -r -p "*请选择安装的IP地址(序号): " input
			case $input in
				1)
					local_ip=${array[0]}
					#echo "选择网段1的IP为：${local_ip}"
					break
					;;
				2)
					local_ip=${array[1]}
					#echo "选择网段2的IP为：${local_ip}"
					break
					;;
				*)
				echo "*请输入有效的数字:"
					;;
			esac
		done
	else
		echo -e "\033[31m*未设置网卡IP，请检查服务器环境！ \033[0m"
		exit 1
	fi
}

# 校验IP地址合法性
function isValidIp() {
	local ip=$1
	local ret=1

	if [[ $ip =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ ]]; then
		ip=(${ip//\./ }) # 按.分割，转成数组，方便下面的判断
		[[ ${ip[0]} -le 255 && ${ip[1]} -le 255 && ${ip[2]} -le 255 && ${ip[3]} -le 255 ]]
		ret=$?
	fi
	return $ret
}

local_ip=''
getIpAddr	#自动获取IP
isValidIp ${local_ip}	# IP校验
if [ $? -ne 0 ]; then
	echo -e "\033[31m*自动获取的IP地址无效，请重试！ \033[0m"
	exit 1
fi
echo "*选择安装的IP地址为：${local_ip}"

for port in $(seq 6379 6384);
do 
mkdir -p ./node-${port}/conf
touch ./node-${port}/conf/redis.conf
cat  << EOF > ./node-${port}/conf/redis.conf
port ${port}
requirepass 1234
bind 0.0.0.0
protected-mode no
daemonize no
appendonly yes
cluster-enabled yes 
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip ${local_ip}
cluster-announce-port ${port}
cluster-announce-bus-port 1${port}
EOF
done
