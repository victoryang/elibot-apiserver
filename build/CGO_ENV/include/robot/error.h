//file name is Error.h
//this file is define the error and err number
#ifndef _ERRORH_
#define  _ERRORH_
//错误等级
#define ERROR_LEVEL_0	2	//普通提示
#define ERROR_LEVEL_1	0	//警告
#define ERROR_LEVEL_2	0	//报警
//通用报警，不分模块
#define COMMON_ERROR_NUM0	0x01000//动态库打开失败 
#define COMMON_ERROR_NUM0_SUB0	0
#define COMMON_ERROR_NUM0_SUB1	1
#define COMMON_ERROR_NUM0_SUB2	2
#define COMMON_ERROR_NUM0_SUB3	3
#define COMMON_ERROR_NUM0_SUB4	4
#define COMMON_ERROR_NUM0_SUB5	5
#define COMMON_ERROR_NUM0_SUB6	6
#define COMMON_ERROR_NUM0_SUB7	7

#define COMMON_ERROR_NUM1	0x01001//动态库函数打开失败
#define COMMON_ERROR_NUM1_SUB0	0
#define COMMON_ERROR_NUM1_SUB1	1
#define COMMON_ERROR_NUM1_SUB2	2
#define COMMON_ERROR_NUM1_SUB3	3
#define COMMON_ERROR_NUM1_SUB4	4
#define COMMON_ERROR_NUM1_SUB5	5
#define COMMON_ERROR_NUM1_SUB6	6
#define COMMON_ERROR_NUM1_SUB7	7
#define COMMON_ERROR_NUM1_SUB8	8
#define COMMON_ERROR_NUM1_SUB9	9
#define COMMON_ERROR_NUM1_SUB10	10
#define COMMON_ERROR_NUM1_SUB11	11
#define COMMON_ERROR_NUM1_SUB12	12
#define COMMON_ERROR_NUM1_SUB13	13
#define COMMON_ERROR_NUM1_SUB14	14
#define COMMON_ERROR_NUM1_SUB15	15
#define COMMON_ERROR_NUM1_SUB16	16
#define COMMON_ERROR_NUM1_SUB17	17
#define COMMON_ERROR_NUM1_SUB18	18
#define COMMON_ERROR_NUM1_SUB19	19
#define COMMON_ERROR_NUM1_SUB20	20
#define COMMON_ERROR_NUM1_SUB21	21

#define COMMON_ERROR_NUM2	0x01002//机器人类型不符
#define COMMON_ERROR_NUM2_SUB0	0 
#define COMMON_ERROR_NUM2_SUB1	1
#define COMMON_ERROR_NUM2_SUB2	2
#define COMMON_ERROR_NUM3	0x01003//发送和接收脉冲不符
#define COMMON_ERROR_NUM3_SUB0	0 
#define COMMON_ERROR_NUM4	0x01004 //急停报警
#define COMMON_ERROR_NUM5	0x01005 //外部急停报警
#define COMMON_ERROR_NUM5_SUB0	0
#define COMMON_ERROR_NUM5_SUB1	1
#define COMMON_ERROR_NUM6	0x01006 //伺服报警
#define COMMON_ERROR_NUM7	0x01007 //焊缝跟踪器连接失败
#define COMMON_ERROR_NUM8	0x01008 //软极限报警
#define COMMON_ERROR_NUM8_SUB0	0
#define COMMON_ERROR_NUM8_SUB1	1
#define COMMON_ERROR_NUM8_SUB2	2
#define COMMON_ERROR_NUM8_SUB3	3
#define COMMON_ERROR_NUM8_SUB4	4
#define COMMON_ERROR_NUM8_SUB5	5
#define COMMON_ERROR_NUM8_SUB6	6
#define COMMON_ERROR_NUM8_SUB7	7
#define COMMON_ERROR_NUM8_SUB8	8

#define COMMON_ERROR_NUM9	0x01009 //参数没有设置 
#define COMMON_ERROR_NUM9_SUB0	0
#define COMMON_ERROR_NUMA	0x0100A //参数设置错误
#define COMMON_ERROR_NUMA_SUB0	0
#define COMMON_ERROR_NUMA_SUB1	1
#define COMMON_ERROR_NUMA_SUB2	2
#define COMMON_ERROR_NUMA_SUB3	3
#define COMMON_ERROR_NUMA_SUB4	4
#define COMMON_ERROR_NUMA_SUB5	5
#define COMMON_ERROR_NUMA_SUB6	6
#define COMMON_ERROR_NUMA_SUB7	7


#define COMMON_ERROR_NUMB	0x0100B //共享内存初始化失败
#define COMMON_ERROR_NUMB_SUB0	0
#define COMMON_ERROR_NUMB_SUB1	1
#define COMMON_ERROR_NUMC	0x0100C //共享内存初始化失败
#define COMMON_ERROR_NUMC_SUB0	0
#define COMMON_ERROR_NUMC_SUB1	1
#define COMMON_ERROR_NUMC_SUB2	2
#define COMMON_ERROR_NUMC_SUB3	3
#define COMMON_ERROR_NUMC_SUB4	4
#define COMMON_ERROR_NUMC_SUB5	5
#define COMMON_ERROR_NUMC_SUB6	6

#define COMMON_ERROR_NUMD	0x0100D //没有对应功能
#define COMMON_ERROR_NUMD_SUB0	0
#define COMMON_ERROR_NUME	0x0100E //串口打开失败
#define COMMON_ERROR_NUMF	0x0100F //读取编码器失败
#define COMMON_ERROR_NUM10	0x010010 //
#define COMMON_ERROR_NUM10_SUB0	0
#define COMMON_ERROR_NUM10_SUB1	1
//译码报警
#define ANALYZE_ERROR_NUM0	0x02000 //参数类型不支持
#define ANALYZE_ERROR_NUM1	0x02001 //参数值错误
#define ANALYZE_ERROR_NUM2	0x02002 //参数变量不支持
#define ANALYZE_ERROR_NUM3	0x02003 //未知判断条件
#define ANALYZE_ERROR_NUM4	0x02004 //label标签错误
#define ANALYZE_ERROR_NUM5	0x02005 //label分配内存错
#define ANALYZE_ERROR_NUM6	0x02006 //调用文件出错
#define ANALYZE_ERROR_NUM7	0x02007 //标签格式错误,以*开头
#define ANALYZE_ERROR_NUM8	0x02008 //标签已经存在

//编译报警
#define COMPILE_ERROR_NUM0	0x03000 //未知坐标系
#define COMPILE_ERROR_NUM1	0x03001 //变量超过范围
#define COMPILE_ERROR_NUM2	0x03002 //数据计算错误
#define COMPILE_ERROR_NUM3	0x03003 //RET没有找到调用的父程序
#define COMPILE_ERROR_NUM4	0x03004 //点位数据不合法
#define COMPILE_ERROR_NUM5	0x03005 //点位数据不合法，六轴机器人２轴为０
#define COMPILE_ERROR_NUM6	0x03006 //MOVC MOVS点位重合
#define COMPILE_ERROR_NUM7	0x03007 //MOVC只有两个点
#define COMPILE_ERROR_NUM8	0x03008 //创建用户坐标系失败
#define COMPILE_ERROR_NUM8_SUB0	0
#define COMPILE_ERROR_NUM9	0x03009 //MOVC三点在一条线上
#define COMPILE_ERROR_NUMA	0x0300A //指令无效
#define COMPILE_ERROR_NUMA_SUB0	0
#define COMPILE_ERROR_NUMA_SUB1	1

//微小直线段插补功能报警
#define MOVT_ERROR_NUM0	     0x04000 //离线文件获取失败
#define MOVT_ERROR_NUM1        0x04001 //离线路径点数过少
#define MOVT_ERROR_NUM1_SUB0   0
#define MOVT_ERROR_NUM1_SUB1   1
#define MOVT_ERROR_NUM2        0x04002 //离线路径过长
#define MOVT_ERROR_NUM3        0x04003 //测量后路径点数过少
#define MOVT_ERROR_NUM3_SUB0   0
#define MOVT_ERROR_NUM3_SUB1   1
#define MOVT_ERROR_NUM4        0x04004 //测量后路径过长
#define MOVT_ERROR_NUM5        0x04005 //路径非单位矢量
#define MOVT_ERROR_NUM6        0x04006 //测量后生成的路径非单位矢量
#define MOVT_ERROR_NUM7        0x04007 //离线路径刀轴矢量和辅助矢量不是正交向量
#define MOVT_ERROR_NUM8        0x04008 //测量后生成的路径刀轴矢量和辅助矢量不是正交向量
#define MOVT_ERROR_NUM9        0x04009 //离线路径路径检查逆解错误
#define MOVT_ERROR_NUMA        0x0400A //测量后路径路径检查逆解错误
#define MOVT_ERROR_NUMB        0x0400B //插补逆解错误
#define MOVT_ERROR_NUMC        0x0400C //转换工具逆解错误
#define MOVT_ERROR_NUMD        0x0400D //逆矩阵计算错误
#define MOVT_ERROR_NUMD_SUB0   0
#define MOVT_ERROR_NUMD_SUB1   1
#define MOVT_ERROR_NUME        0x0400E //加减速计算错误
#define MOVT_ERROR_NUME_SUB0   0
#define MOVT_ERROR_NUME_SUB1   1

//标定模块报警
#define CALIBRATE_ERROR_NUM0	0x05000 //标定LPU错误
#define CALIBRATE_ERROR_NUM1	0x05001 //相邻测量点重复
#define CALIBRATE_ERROR_NUM2	0x05002 //误差过大

//数据解算报警
#define CALCULATE_ERROR_NUM0	0x06000 //数据解算无解
#define CALCULATE_ERROR_NUM0_SUB0	0
#define CALCULATE_ERROR_NUM0_SUB1	1
#define CALCULATE_ERROR_NUM0_SUB2	2
#define CALCULATE_ERROR_NUM0_SUB3	3
#define CALCULATE_ERROR_NUM0_SUB4	4
#define CALCULATE_ERROR_NUM0_SUB5	5
#define CALCULATE_ERROR_NUM0_SUB6	6
#define CALCULATE_ERROR_NUM0_SUB7	7
#define CALCULATE_ERROR_NUM0_SUB8	8
#define CALCULATE_ERROR_NUM0_SUB9	9
#define CALCULATE_ERROR_NUM0_SUBA	0xA
#define CALCULATE_ERROR_NUM0_SUBB	0xB
#define CALCULATE_ERROR_NUM0_SUBC	0xC
#define CALCULATE_ERROR_NUM0_SUBD	0xD
#define CALCULATE_ERROR_NUM0_SUBE	0xE
#define CALCULATE_ERROR_NUM0_SUBF	0xF
#define CALCULATE_ERROR_NUM0_SUB10	0x10
#define CALCULATE_ERROR_NUM0_SUB11	0x11
#define CALCULATE_ERROR_NUM0_SUB12	0x12
#define CALCULATE_ERROR_NUM0_SUB13	0x13
#define CALCULATE_ERROR_NUM0_SUB14	0x14
#define CALCULATE_ERROR_NUM0_SUB15	0x15
#define CALCULATE_ERROR_NUM0_SUB16	0x16
#define CALCULATE_ERROR_NUM0_SUB17	0x17
#define CALCULATE_ERROR_NUM0_SUB18	0x18
#define CALCULATE_ERROR_NUM0_SUB19	0x19
#define CALCULATE_ERROR_NUM1	0x06001 //计算曲线长度为０
#define CALCULATE_ERROR_NUM1_SUB0	0 
#define CALCULATE_ERROR_NUM2	0x06002 //奇异点
#define CALCULATE_ERROR_NUM2_SUB0	0
#define CALCULATE_ERROR_NUM3	0x06003 //奇异点
#define CALCULATE_ERROR_NUM3_SUB0	0
#define CALCULATE_ERROR_NUM3_SUB1	1
#define CALCULATE_ERROR_NUM4	0x06004 //速度为０
#define CALCULATE_ERROR_NUM4_SUB0	0
#define CALCULATE_ERROR_NUM5	0x06005 //示教点个数不够
//文件操作报警
#define FILE_ERROR_NUM0		0x07000 //文件已经打开
#define FILE_ERROR_NUM1		0x07001 //子程序文件调用层数超出范围
#define FILE_ERROR_NUM2		0x07002 //打开文件失败
#define FILE_ERROR_NUM2_SUB0	0
#define FILE_ERROR_NUM2_SUB1	1
#define FILE_ERROR_NUM3		0x07003 //超出文件运动指令最大行数
#define FILE_ERROR_NUM4		0x07004 //文件格式错误NOP END 没有匹配
#define FILE_ERROR_NUM5		0x07005 //文件格式错误多个NOP
#define FILE_ERROR_NUM6		0x07006 //文件格式错误多个END

//焊接报警
#define ARC_ERROR_NUM0		0x08000 //起弧失败
#define ARC_ERROR_NUM0_SUB0	0
#define ARC_ERROR_NUM1		0x08001 //没有搜索到焊缝
#define ARC_ERROR_NUM1_SUB0	0
#define ARC_ERROR_NUM1_SUB1	1
#define ARC_ERROR_NUM1_SUB2	2
#define ARC_ERROR_NUM2		0x08002 //焊机报警
#define ARC_ERROR_NUM3		0x08003 //断弧报警
#define ARC_ERROR_NUM4		0x08004 //碰撞报警
#define ARC_ERROR_NUM5		0x08005 //摆焊工艺号错误
#define ARC_ERROR_NUM5_SUB0	0
#define ARC_ERROR_NUM5_SUB1	1
#define ARC_ERROR_NUM5_SUB2	2
#define ARC_ERROR_NUM5_SUB3	3


//运行报警
#define RUNNING_ERROR_NUM0	0x09000 //运行和暂停行号不符
#define RUNNING_ERROR_NUM1	0x09001 //伺服未使能
#define RUNNING_ERROR_NUM1_SUB0	0
#define RUNNING_ERROR_NUM1_SUB1	1
#define RUNNING_ERROR_NUM1_SUB2	2
#define RUNNING_ERROR_NUM2	0x09002 //机器人状态与运动不符
#define RUNNING_ERROR_NUM2_SUB0	0
#define RUNNING_ERROR_NUM2_SUB1	1
#define RUNNING_ERROR_NUM2_SUB2	2
#define RUNNING_ERROR_NUM3	0x09003 //没有匹配的运动指令类型
#define RUNNING_ERROR_NUM4	0x09004 //无法走PL
#define RUNNING_ERROR_NUM4_SUB0	0
#define RUNNING_ERROR_NUM4_SUB1	1
#define RUNNING_ERROR_NUM4_SUB2	2
#define RUNNING_ERROR_NUM4_SUB3	3
#define RUNNING_ERROR_NUM4_SUB4	4
#define RUNNING_ERROR_NUM4_SUB5	5
#define RUNNING_ERROR_NUM5	0x09005 //机器人处于暂停状态
#define RUNNING_ERROR_NUM6	0x09006 //速度超过电机最大速
#define RUNNING_ERROR_NUM7	0x09007 //预约文件没有设置
#define RUNNING_ERROR_NUM8	0x09008 //预约功能没打开
#define RUNNING_ERROR_NUM8_SUB0	0
#define RUNNING_ERROR_NUM8_SUB1	1
//提示信息
#define INFOR_NUM0	0x0A000 //退出预约功能
#define INFOR_NUM1	0x0A001 //pause
#define INFOR_NUM1_SUB0	0//ep
#define INFOR_NUM1_SUB1	1//sd
//正逆解模块报警
#define POSINV_ERROR_NUM0     0x0B000//连杆参数2设置不合理
#define POSINV_ERROR_NUM1     0x0B001//奇异点报警

//通信模块报警
#define COMMUNICATE_ERROR_NUM0		0x0C000 //得到IP错误
#define COMMUNICATE_ERROR_NUM0_SUB0		0 //得到IP错误
#define COMMUNICATE_ERROR_NUM0_SUB1		1 //得到IP错误
#define COMMUNICATE_ERROR_NUM0_SUB2		2 //得到IP错误
#define COMMUNICATE_ERROR_NUM1		0x0C001 //创建连接错误
#define COMMUNICATE_ERROR_NUM2		0x0C002 //发送数据错误
#define COMMUNICATE_ERROR_NUM3		0x0C003 //连接超时或断网
#define COMMUNICATE_ERROR_NUM3_SUB0		0
#define COMMUNICATE_ERROR_NUM3_SUB1		1
#define COMMUNICATE_ERROR_NUM3_SUB2		2

//外部轴报警
#define EXTERNAXIS_ERROR_NUM0        0x0D000 //请先设置外部轴一
#define EXTERNAXIS_ERROR_NUM0_SUB0  0
#define EXTERNAXIS_ERROR_NUM0_SUB1  1
#define EXTERNAXIS_ERROR_NUM0_SUB2  2
#define EXTERNAXIS_ERROR_NUM0_SUB3  3
#define EXTERNAXIS_ERROR_NUM1        0x0D001//协同一轴未校验
#define EXTERNAXIS_ERROR_NUM1_SUB0  0
#define EXTERNAXIS_ERROR_NUM1_SUB1  1
#define EXTERNAXIS_ERROR_NUM2  0x0D002//协同二轴未校验
#define EXTERNAXIS_ERROR_NUM3  0x0D003//外部轴除coord外仍有变化

#define INVERSE_DYNAMIC_NUM0		0x0F000 //动力学报警
#define INVERSE_DYNAMIC_NUM0_SUB0	0
#define INVERSE_DYNAMIC_NUM0_SUB1	1

//其它报警
#define OTHER_ERROR_NUM0		0x1F000	//工具号没有设置
#define OTHER_ERROR_NUM0_SUB0	0
#define OTHER_ERROR_NUM0_SUB1	1
#define OTHER_ERROR_NUM0_SUB2	2
#define OTHER_ERROR_NUM0_SUB3	3
#define OTHER_ERROR_NUM1		0x1F001//得到工具数据错误
#define OTHER_ERROR_NUM2		0x1F002//工具数据保存错误
#define OTHER_ERROR_NUM2_SUB0	0
#define OTHER_ERROR_NUM2_SUB1	1
#define OTHER_ERROR_NUM2_SUB2	2
#define OTHER_ERROR_NUM3		0x1F003//内存分配错误
#define OTHER_ERROR_NUM3_SUB0	0
#define OTHER_ERROR_NUM3_SUB1	1
#define OTHER_ERROR_NUM3_SUB2	2
#define OTHER_ERROR_NUM3_SUB3	3
#define OTHER_ERROR_NUM3_SUB4	4
#define OTHER_ERROR_NUM3_SUB5	5
#define OTHER_ERROR_NUM3_SUB6	6

#define OTHER_ERROR_NUM4		0x1F004 //等待超时
#define OTHER_ERROR_NUM4_SUB0	0
#define OTHER_ERROR_NUM5	0x1F005 //读取干涉区失败
#define OTHER_ERROR_NUM6	0x1F006 //保存干涉区失败
#define OTHER_ERROR_NUM7	0x1F007 //设置干涉区失败
#define OTHER_ERROR_NUM8	0x1F008 //当前用户号超过最大序号
#define OTHER_ERROR_NUM9	0x1F009 //当前工具未使用
#define OTHER_ERROR_NUMA	0x1F00A //当前用户未使用
#define OTHER_ERROR_NUMB	0x1F00B //当前轴未使用
#define OTHER_ERROR_NUMC	0x1F00C //当前轴被禁止
#define OTHER_ERROR_NUMC_SUB0	0
#define OTHER_ERROR_NUMC_SUB1	1
#define OTHER_ERROR_NUMD	0x1F00D //伺服报警
#define OTHER_ERROR_NUMD_SUB0	0
#define OTHER_ERROR_NUME	0x1F00E //得到用户数据错误
#define OTHER_ERROR_NUMF	0x1F00F //用户数据保存错误
#define OTHER_ERROR_NUMF_SUB0	0
#define OTHER_ERROR_NUMF_SUB1	1
#define OTHER_ERROR_NUM10	0x1F0010//创建用户坐标失败
#define OTHER_ERROR_NUM10_SUB0	0
#define OTHER_ERROR_NUM10_SUB1	1
#define OTHER_ERROR_NUM11		0x1F0011//进入干涉区
#define OTHER_ERROR_NUM11_SUB0	0
#define OTHER_ERROR_NUM11_SUB1	1
#define OTHER_ERROR_NUM12		0x1F0012//没有准备好不能放物料
#define OTHER_ERROR_NUM12_SUB0	0
#define OTHER_ERROR_NUM12_SUB1	1
#define OTHER_ERROR_NUM12_SUB2	2
#define OTHER_ERROR_NUM13		0x1F0013//没有准备好不能取物料
#define OTHER_ERROR_NUM14		0x1F0014//抓取失败
#define OTHER_ERROR_NUM14_SUB0	0
#define OTHER_ERROR_NUM14_SUB1	1
#define OTHER_ERROR_NUM15		0x1F0015//除零错误
#define OTHER_ERROR_NUM15_SUB0	0
#define OTHER_ERROR_NUM15_SUB1	1


/***********************************PLC错误定义************************************************/
///////错误段
///////0段开始
#define PUSH_ERR						1		//指令压堆栈错误		
#define ERR_ADDRESS	 					2		//指令字错
#define OPERATER_ERR	 				3		//指令数据错误
#define Outopernumscope   				4		//操作数越界
#define Outopernumbitscope				5		//操作数位越界
#define Wrongopernumadd   				6		//操作数地址越界
#define Wrongopernums    				7		//操作数个数错误
#define Wrongconst      				8		//常量错误
#define Wrongoptionaddr 				9		//操作数错误
#define WrongGRP						10		//GRP指令错误
#define PART1afterPART2					11		//PART 1 与PART2位置错误
#define WrongPART						12		//缺少PART错误
#define ERR_NOPLCFILE					13		//找不到PLC文件	
#define WrongMOREEND					14		//指令中有多余的END
#define WrongreUsedNUM    				15		//指定的变量序号已经被使用了
#define WrongEND						16		//END不在最后一行错误
#define ERR_SAXIS_OVERTRAVEL			17		//S轴超程报警
#define ERR_LAXIS_OVERTRAVEL			18		//L轴超程报警
#define ERR_UAXIS_OVERTRAVEL			19		//U轴超程报警
#define ERR_RAXIS_OVERTRAVEL			20		//R轴超程报警
#define ERR_BAXIS_OVERTRAVEL			21		//B轴超程报警
#define ERR_TAXIS_OVERTRAVEL			22		//T轴超程报警

//错误段1
#define ERR_LEVER 						1

#define ERR_DATALIMIT					1		//变量序号超出最大或最小范围
#define ERR_COUNTVALUE					2		//指定的个数或起始数超过最大/最小值
#define ERR_VALUE						2		//译码运行时获取的值超出最大或最小范围
#define ERR_SPOT						3		//3+i(0-4)//焊接动行错误(0.大开失败，1小开失败，2闭合失败，3焊接失败)//未用，程序已被注释
#define ERR_WRITEFILE					7		//数据写入文件出错

/***********************************PLC错误定义结束**************************************/
//316-318 分为手动MOVJ,MOVL，回零三种情况，
//速度加减速度分别从0-12
/***********************************通用性错误**************************************/
///////错误段
#define COMMONLEVEL             		2		//通用性错误级别
///////错误码
#define ERR_DIVISOR_ZERO                1		//除数为零
#define ERR_DEADSTOP					2		//急停警告
#define ERR_ARRAY_BEYOND				102		//数组越界错误	//未使用
#define ERR_DATA_EXCEPTION				103		//数据异常错误	//有错误，但error.lst中没有
#define ERR_SERVO_NOT_READY				3		//伺服未准备
#define ERR_NOT_TEACH_MOD					11
#define ERR_WELD_MACHINE					12

#define ERR_OPEN_FILE_FAILSE			5		//打开文件失败
#define ERR_FILE_FORMAT_FAILSE			6		//文件格式错误
#define ERR_FILE_NOTEXIST				7		//文件不存在
#define ERR_FILE_PATH_NULL				8		//文件路径错误
#define ERR_SERVO_ALARM					9//9+i(0-5)		//伺服报警
#define ERR_MOVS_POSE_COUNTER			104 
#define ERR_MOVC_POSE_COUNTER			107 

#define ERR_SERVO_ENCODE					105 
#define ERR_UART_OPEN						106
/***********************************文件相关错误**************************************/
///////错误段
#define FILE_ERROR_LEVEL                2		//文件错误级别									
///////
#define ERR_CREATEUSERFRAME	 			15		//用户坐创建失败
#define TOOL_FILE_ERROR					16		//读取Tool文件错误
#define TOOL_FILE_ERROR_NOEXIST			0		//文件不存在
#define TOOL_FILE_ERROR_OPEN_FALSE		1		//文件打开失败
#define TOOL_FILE_ERROR_INCOMPLETE		2		//文件不完整
#define TOOL_FILE_ERROR_READ_INDEX		3		//读取索引错误																	//
#define UF_FILE_ERROR					20		//读取UF文件错误
#define UF_FILE_ERROR_NOEXIST			0		//文件不存在
#define UF_FILE_ERROR_OPEN_FALSE		1		//文件打开失败
#define UF_FILE_ERROR_INCOMPLETE		2		//文件不完整
#define UF_FILE_ERROR_READ_INDEX		3		//读取索引错误
#define AE_SYNCHRONOUS_ERROR			24//24+i(0-5)//脉冲数据与绝对脉冲数据不符，其字码代表轴号（0-5）
#define ERR_NO_LEGAL_ROOT		25
#define CALCLATE_ZERO_ERROR		26
#define DLL_LIB_ERROR		27
//-----------------------------译码错误定义-------------------------------
///////错误段
#define COMPILELEVER 					5		//编译时错误等级				
///////
#define ERR_COMDATA						1		//指令参数的数据非法
#define ERR_CTRL						2		//指令参数的数据内部控制非法
#define ERR_COMLBL						3		//未定义的指令或不存在的参数
#define ERR_PARANUM						4		//指令参数个数错误
#define ERR_MULTIPARA					5		//指令同一参数重复设置
#define ERR_NOCHAR						6		//指定的字符没有找到
#define ERR_LABEL						9		//LABEL标签格式错误或者找不到指定的LABEL标签
#define ERR_NOJOBFILE					11		//找不到JOB文件
#define ERR_NOENDORSTART 				12		//文件不存在NOP 或END
#define ERR_STACKOVER					14		//子程序调用栈不匹配(找不到调用的CALL主程序)
#define ERR_OVERSUBMAX					15		//子程序调用层数超过最大值
#define ERR_GETDATA						16		//从文件得到位置数据错误
#define ERR_ERRFNAME					17		//指定的文件名 不合法(JUMP,CALL)
#define ERR_OUT_ALARM					120		//外部设备产生的报警//控制器未使用
#define ERR_MOV_PL_DATA					121
///////错误段

/***********************************插补错误定义**************************************/
///////错误段
#define INTERPLEVEL 					5		//编译时错误等级
///////
#define ERR_MOVA_INVALID				18
#define ERR_MOVC_INVALID				19//19+i(0-4)//无效MOVC
#define ERR_MOVC_BEYOND					24//24+i(0-1)//MOVC 关节范围越界
#define ERR_MOVL_BEYOND					27//27+i(0-2)//MOVL 关节范围越界
#define ERR_MOVS_BEYOND					30//30+i(0-1)//MOVS 关节范围越界
#define ERR_SFTON_OF_NOT_MATCH			32		//SFTON SFTOF 未成对出现
#define ERR_SFT_STACK_OUT				33   	//超出SFTON SFTOF 栈空间错误
#define ERR_MOVS_INVALID				34//34+i(0-4)//无效MOVS
#define ERR_SOFTLIMIT_INVALID			46//46+i(0-5)//软极限无效
#define ERR_JOINT_BEYOND				52		//关节范围越界
#define ERR_INVAL_JOINT					309		//无效目标关节角度坐标		//控制器程序已注释
#define ERR_MOVJ_BEYOND					53//53+i(0-1)//MOVJ应用P变量(XYZ类型)越界
#define ERR_TOOL_SET						55		//工具号越界
#define ERR_USERCOORDNO_BEYOND			56		//用户数据号越界
#define ERR_CUBESLNO_BEYOND				57		//立体软极限下标越界
#define ERR_CUBESL_BEYOND				314		//超出立体软极限数据范围	//控制器程序已注释
#define ERR_RATENO_BEYOND				58		//超出倍频比数据下标越界
#define ERR_SPEED_ZERO					59//59+i(0-12)//速度为零错误
#define ERR_ACC_ZERO					72//72+i(0-12)//加速度为零错误
#define ERR_DEC_ZERO					85//85+i(0-12)//减速度为零错误
#define ERR_INTER_BUFFER_MEMERY_LACK 	319		//发送脉冲内部缓冲区空间不足	//控制器未使用
#define ERR_JOINT_NO_FETCH				320		//关节无法到达						//控制器未使用
#define ERR_CJOINT_NO_FETCH				98		//C变量关节无法到达
#define ERR_THROUGH_WORKPLACE			99		//发生穿越
#define ERR_INFERENCE_ENTER				100//100+i(0-1)//进入干涉区报警
#define ERR_SAXIS_INFERENCE_ENTER		102//102+i(0-1)//进入S轴干涉区报警
#define ERR_SINGULAR_POINT				104//104+i(0-4)		//机器人在奇异点错误
#define ERR_CURLINE_CHANGE		125
#define ERR_TIME_OUT		126
/***********************************插补错误定义结束**************************************/
 
/***********************************微段插补错误定义开始*************************************/
///////错误段
#if 0
#define MOVTINTPERRLEVEL                            5 //编译时错误等级
///////
#define ERR_MOVT_RETURN                             1//返回值错误
#define ERR_MOVT_PATH                                2//路径错误
#define ERR_MOVT_KINEMATIC                          3//运动学正逆解错误
#define ERR_MOVT_DIV0                               4//除零错误
#define ERR_MOVT_MEMORY                             5//内存错误
#endif
/*********************************微段插补错误定义结束***********************************/


#define TEST_ALLCODE_ANALYZE 					//测试指令译码(语法正确性)
#undef  TEST_ALLCODE_ANALYZE
#define TEST_ALLCODE_RUN						//测试指令译码(运行结果正确性)	
#undef TEST_ALLCODE_RUN			
#endif

/***********************************冲压工艺************************************************/
//冲压工艺报警
#define STAMP_ERROR_NUM0		0x2F000	        // ＃站报警
#define STAMP_ERROR_NUM1		0x2F001	        // 通用报警
#define STAMP_ERROR_NUM1_SUB2   2               // 加载冲压元数据失败

/***********************************数据库************************************************/
//数据库操作报警
#define DB_ERROR_NUM		    0x3F000	        // 数据库报警
#define DB_ERROR_NUM_SUB1       1               // 获取元数据对象失败
#define DB_ERROR_NUM_SUB2       2               // 获取内存地址失败
#define DB_ERROR_NUM_SUB3       3               // 获取内存值字符串失败
#define DB_ERROR_NUM_SUB4       4               // 设置内存值字符串失败

/***********************************跟踪工艺************************************************/
//跟踪工艺报警
#define TRACK_ERROR_NUM15		0x1F0015//跟踪功能错误
#define TRACK_ERROR_NUM15_SUB0	0		//跟踪标定结果错误
#define TRACK_ERROR_NUM15_SUB1	1		//获取跟踪IO检测错误
#define TRACK_ERROR_NUM15_SUB2	2		//获取跟踪工艺数据错误
#define TRACK_ERROR_NUM15_SUB3	3       //读取编码器值错误
#define TRACK_ERROR_NUM15_SUB4	4		//手动暂停或急停，检查跟踪是否有误

/***********************************视觉工艺************************************************/
//视觉工艺报警
#define VISION_ERROR_NUM15		0x1F0015//视觉功能错误
#define VISION_ERROR_NUM15_SUB0	0		//获取视觉数据错误
#define VISION_ERROR_NUM15_SUB1	1		//视觉缓冲区溢出报警
#define VISION_ERROR_NUM15_SUB2	2		//视觉初始化失败报警
#define VISION_ERROR_NUM15_SUB3	3       //读取编码器值错误
#define VISION_ERROR_NUM15_SUB4	4		//目标物体不在机器人工作空间内
#define VISION_ERROR_NUM15_SUB5	5		//手动暂停或急停，检查跟踪是否有误
#define VISION_ERROR_NUM15_SUB6	6		//目标物体超出机器人工作范围
#define VISION_ERROR_NUM15_SUB7	7       //机器人工作状态错误报警
#define VISION_ERROR_NUM15_SUB8	8       //获取初始关节值错误
#define VISION_ERROR_NUM15_SUB9	9		//视觉心跳检测出错