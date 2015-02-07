tester-on-docer
===============

轻量级高隔离性跨平台golang 单元测试框架

## 说明
利用docker，隔离执行golang的单元测试 

## 示例


[windows_demo](doc/readme_windows.md)

[linux_demo](doc/readme_windows.md)


### 配置说明

	[ssh] 		# 可省略,不可为空,连接boot2docker的ssh配置;linux和mac下可省略
	user	=
	passwd	=
	ip		=
	port	=

	[global]	
	debug	=  # 是否开启debug
	sudo 	=  # 是否使用sudo

	[image] 	# 测试使用的docker image及image信息,os和arch填写参照golang交叉编译的GOOS和GOARCH变量
	os		= 
	arch	= 
	name	= #格式 REPOSITORY:TAG
	rebuild = # 值为true/false ,!如果值设置为true,将删除相关container

	[path] 		# 如果宿主机是Windows和OS X,则两个值分别为宿主机和boot2docker文件夹映射路径;如果宿主机是Linux,则两个路径要求一致	
	host	= # 测试执行路径和配置文件必须在该路径或该路径的子目录下
	boot2docker	= 

## 要求
- docker:
	windows 和OS X下docker 1.3+ 
- go 1.3+
- 宿主机可交叉编译
