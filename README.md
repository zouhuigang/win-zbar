### windows下golang二维码解码

1.首先安装 tdm64-gcc-5.1.0-2.exe，方便解析c++/c库，如果没加环境变量，添加进环境变量。path


2.安装windows版zbar软件zbar-0.10-setup.exe，识别二维码。

![https://raw.githubusercontent.com/zouhuigang/win-zbar/master/images/zbar.png](https://raw.githubusercontent.com/zouhuigang/win-zbar/master/images/zbar.png)


3.最重要的部分，添加cgo库和包进环境变量中，否则找不到包或库文件,导致报错

	CGO_LDFLAGS=-LD:\software\zbar\lib
	CGO_CFLAGS=-ID:\software\zbar\include

![https://raw.githubusercontent.com/zouhuigang/win-zbar/master/images/path.png](https://raw.githubusercontent.com/zouhuigang/win-zbar/master/images/path.png)


然后bgcode.go文件

	package decode
	
	// #cgo LDFLAGS: -lzbar
	// #include <zbar.h>
	import "C"


4.点击install.bat编译成功，生成src.exe,点击即可识别出结果。

![https://raw.githubusercontent.com/zouhuigang/win-zbar/master/images/result.png](https://raw.githubusercontent.com/zouhuigang/win-zbar/master/images/result.png)

大功告成，最重要的是第三步~！！！