如果放在lbm\solution\src目录，则可以不用配置WorkDir参数
参数说明：
	1、WorkDir:工作目录，需要配置为账户产品代码solution的绝对路径(如：I:\\kbss_ums\\src\\trunk\\lbm\\solution)，如果此程序放在solution\src\目录，则可不设置WorkDir参数
	2、VCDir：存放vsvarsall.bat的绝对路径
	3、Platform:构建的目标平台，取值:[x86|x64]，其中x86对应32位平台，x64对应64位平台
	4、OutDir：构建成功时，用于存放lbm动态库的目录；如果没有设置，则输出到makefile_template文件中配置的KCBP_DIR的目录的kbsslbm目录
	5、CompileAll:是否全量构建，取值为：true或false
	6、LibDir:指定链接的路径，如：kcbp\lib目录。如果没有指定则会使用makefile_template文件中配置的KCBP_DIR的目录的lib子目录