package constants

// USAGE help说明文本
const USAGE = `
基本:
	img2base64 <filepath>
	转换文件为文本文件, 在当前目录下输出文件, 格式为 filename.ext.txt

用法:
	img2base64 [--dataurl] [--print] <filepath> [target-filepath]
		--dataurl 输出dataurl
		--print 打印到终端

其他命令:
	-r, --reverse [-path|-text] <base64-string|base64filepath> [target-filepath]
		反转base64字符串为文件
		--path 默认值, 从读取路径的文件, 输出去掉后缀的同名文件
		--text 读取命令行输入的文本, 输出名为'output'的文件
		注意命令行是有长度限制的, 不同终端限制长度不同

	-v, --version
		打印版本号
`

// VERSION 版本号
const VERSION = "0.8"
