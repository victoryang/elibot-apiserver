#ifndef _IO_BUS_H_
#define _IO_BUS_H_

#include <stdint.h>
#include <list.h>

#define	VIRTUAL_DEVICE_MAGIC			0x494F4253	//虚拟设备魔数定义(IOBS)

#define	IOBUS_USER_IO_MAX				256	//系统支持用户定义IO的最大数目
#define	IOBUS_SYS_SIGNAL_MAX			16	// 系统支持系统信号的最大数目
#define	IOBUS_VIRTUAL_DEVICE_NUM_MAX	16	//系统支持添加扩展设备的最大数目
#define IO_ADDR_DEF_MAX 4

#define		IOBUS_BITS_MAX 1280
#define		IOBUS_REGS_MAX 512

#define		_DEBUG    0 					// 1开启调试信息，0关闭调试信息


enum io_type {
    IO_TYPE_BITS_IN = 0x00,				// 数字输入IO
    IO_TYPE_BITS_OUT = 0x01,			// 数字输出IO
    IO_TYPE_REGISTER_IN = 0x02,			// 寄存器输入IO
    IO_TYPE_REGISTER_OUT = 0x03,		// 寄存器输出IO    
    IO_TYPE_MAX = 0x04,					// IO类型最大值
};


enum device_atapter_type {
    ROBOT_DEVICE_APAPTER = 0x00,		// 机器人设备适配器
    MODBUS_TCP_DEVICE_APAPTER = 0x01,	// Mobus TCP设备适配器
    MODBUS_RTU_DEVICE_APAPTER = 0x02,	// Modbus RTU设备适配器
    DEVICE_APAPTER_TYPE_MAX = 0x03,		// 设备适配器最大编号
};

enum robot_signal_type {
	ROBOT_SIGNAL_IN_EMERGENCY_STOP = 0x0001,				// 外部急停
	ROBOT_SIGNAL_IN_EXT_TSTART = 0x0002, 					// 外部启动信号
	ROBOT_SIGNAL_IN_EXT_PAUSE = 0x0003, 					// 外部暂停信号
	ROBOT_SIGNAL_IN_EXT_SERVO_ON = 0x0004,					// 外部伺服使能信号
	ROBOT_SIGNAL_IN_EXT_CLEAR_ALARM = 0x0005,				// 外部清除报警信号

	ROBOT_SIGNAL_OUT_SERVO_STATUS = 0x1001,					// 伺服使能状态
	ROBOT_SIGNAL_OUT_REACH_ORIGIN_STATUS = 0x1002,			// 到达工作原点状态
	ROBOT_SIGNAL_OUT_TEACH_MODE_STATUS = 0x1003,			// 机器人示教模式状态
	ROBOT_SIGNAL_OUT_AUTO_MODE_STATUS = 0x1004,				// 机器人自动模式状态
	ROBOT_SIGNAL_OUT_REMOTE_MODE_STATUS = 0x1005,			// 机器人远程模式状态	
	ROBOT_SIGNAL_OUT_STOP_STATUS = 0x1006,					// 机器人停止状态	
	ROBOT_SIGNAL_OUT_PAUSE_STATUS = 0x1007,					// 机器人暂停状态	
	ROBOT_SIGNAL_OUT_RUN_STATUS = 0x1008,					// 机器人运行状态		
	ROBOT_SIGNAL_OUT_ERROR_STATUS = 0x1009,					// 机器人出错状态	
};


typedef  struct  user_io_def_t{
	char     	io_name[32];			// IO名称
	char     	device_name[32];		// 设备名称
	uint16_t    register_addr;			// 寄存器地址
	uint8_t     io_type;				// IO类型
	uint8_t     res_1;					// 预留
	uint8_t    	*data_position;			// 预留指向数据缓存位置的指针
	uint8_t    	*data_position_bak;		// 预留指向备用数据缓存位置的指针
	struct list_head list;	
}user_io_def;

//=========================================================
// 用户定义IO的集合
//=========================================================
typedef  struct user_io_collection_t{ 
	struct list_head bits_input_list;
	struct list_head bits_output_list;
	struct list_head regs_input_list;
	struct list_head regs_output_list;

	uint32_t  		io_num;							// 用户定义IO的数目
	user_io_def  	user_io[IOBUS_USER_IO_MAX];		// 用户定义IO结构体数组
}user_io_collection;


typedef struct sys_signal_map_phy_t{
	char		io_name[32];		// 用户IO名称
	uint16_t	signal_id;			// 系统信号标识
	uint8_t     signal_enable;		// 信号启用标志：0未启用；1启用
	uint8_t	res;				// 预留
}sys_signal_map_phy;

//=========================================================
// 系统信号与IO映射关系描述集合的物理存储结构
//=========================================================
typedef  struct  sys_signal_map_collection_phy_t{ 
	uint32_t			num;									// 系统信号与用户IO关联定义的数目
	sys_signal_map_phy  signal_map[IOBUS_SYS_SIGNAL_MAX];		// 系统信号与用户定义IO关联的结构体数组
}sys_signal_map_collection;

//=========================================================
// 设备IO地址描述
//=========================================================
typedef  struct  device_io_addr_def_t{ 
	uint16_t    start_addr;		//开始地址
	uint16_t    end_addr;		//结束地址
	uint8_t     io_type;		// IO地址类型
	uint8_t     res[3];			//  预留
}device_io_addr_def;

//=========================================================
//MODBUS_TCP适配器参数描述的物理结构
//=========================================================
typedef  struct modbus_tcp_adapter_param_phy_t{ 
	char				ip[16];			// IP地址
	uint16_t			port;           // IP端口号
	uint8_t				device_addr;	// 设备地址
    uint8_t				num;			// IO定义数目
    device_io_addr_def	*io_def;		// IO定义集合
}modbus_tcp_adapter_param_phy;


//=========================================================
//MODBUS_RTU适配器参数描述的物理结构
//=========================================================
typedef  struct modbus_rtu_adapter_param_phy_t{ 
	uint32_t			baud_rate;		// 波特率
	uint8_t				parity;         // 奇偶校验
	uint8_t				data_bits;		// 数据位
	uint8_t				stop_bits;		// 停止位
	uint8_t				device_addr;	// 设备地址
    uint32_t			num;			// IO定义数目
    device_io_addr_def	*io_def;		// IO定义集合
}modbus_rtu_adapter_param_phy;


//============================================================================
//ROBOT 适配器参数描述的物理结构
//============================================================================
typedef  struct robot_adapter_param_phy_t{ 
	uint8_t				device_addr;		// 设备地址
    uint8_t				num;				// IO地址定义数目
    uint8_t 			res[2];				// 预留  
    device_io_addr_def	*io_def;			// IO定义集合	
}robot_adapter_param_phy;


//=========================================================
// 设备地址映射
//=========================================================
typedef struct virtual_device_addr_map_t{ 
	uint8_t     map_flag;				// 地址映射标识，
	uint8_t     io_type_mask;			// IO定义是否有效，对应bit为1有效，否则无效
	uint16_t    res;					// 预留
	uint16_t    digital_in_start;		// 数字输入开始
	uint16_t    digital_in_end;			// 数字输入结束
	uint16_t    digital_out_start;		// 数字输出开始
	uint16_t    digital_out_end;		// 数字输出结束
	uint16_t    register_in_start;		// 寄存器输入开始
	uint16_t    register_in_end;		// 寄存器输入结束
	uint16_t    register_out_start;		// 寄存器输出开始
	uint16_t    register_out_end;		// 寄存器输出结束
}virtual_device_addr_map;

//=========================================================
// 设备适配器参数
//=========================================================
#if 1
typedef union device_adapter_param_t{
	robot_adapter_param_phy  robot_adapter;			// 设备设配器参数
	modbus_tcp_adapter_param_phy  tcp_adapter;		// 设备设配器参数
	modbus_rtu_adapter_param_phy  rtu_adapter;		// 设备设配器参数
}device_adapter_param;

typedef struct virtual_device_phy_t{ 
	uint32_t    magic;									// 魔数
	char     	device_name[32];						// 设备名称
	uint16_t    adapter_type;							// 适配器类型
	uint16_t    scan_frequency;						    // 扫描频率
	uint8_t		retry_times;							// 故障重试次数
	uint8_t     device_order;							// 设备序号
	uint8_t     device_status;							// 设备状态,0正常,1报警
	uint8_t     io_addr_def_num;						// IO地址定义数目	
	uint8_t		device_enable;							// 设备使能标志，1使能，0禁用	
	device_io_addr_def  io_addr_def[IO_ADDR_DEF_MAX];	// IO地址定义数组
	virtual_device_addr_map  addr_map;					// 设备地址映射
    device_adapter_param adapter_param;
}virtual_device_phy;
#else
typedef struct virtual_device_phy_t{ 
	uint32_t    magic;									// 魔数
	char     	device_name[32];						// 设备名称
	uint16_t    adapter_type;							// 适配器类型
	uint16_t    scan_frequency;						    // 扫描频率
	uint8_t		retry_times;							// 故障重试次数
	uint8_t     device_order;							// 设备序号
	uint8_t     device_status;							// 设备状态,0正常,1报警
	uint8_t     io_addr_def_num;						// IO地址定义数目	
	device_io_addr_def  io_addr_def[IO_ADDR_DEF_MAX];	// IO地址定义数组
	virtual_device_addr_map  addr_map;					// 设备地址映射
	union device_adapter_param_t{
		modbus_tcp_adapter_param_phy  tcp_adapter;		// 设备设配器参数
		modbus_rtu_adapter_param_phy  rtu_adapter;		// 设备设配器参数
	}adapter_param;
}virtual_device_phy;

#endif
//=========================================================
// 虚拟设备集合
//=========================================================	
typedef struct virtual_device_collection_phy_t{ 
	uint8_t  num;                                          				 	//设备数目
	virtual_device_phy  virtual_device[IOBUS_VIRTUAL_DEVICE_NUM_MAX];   	//虚拟设备描述
}virtual_device_collection_phy;

//=========================================================
// IO总线
//=========================================================
typedef struct io_bus_t{
	uint8_t		bits_input[IOBUS_BITS_MAX];
	uint8_t		bits_output[IOBUS_BITS_MAX];
	uint16_t	regs_input[IOBUS_REGS_MAX];
	uint16_t	regs_output[IOBUS_REGS_MAX];
	user_io_collection user_ios;
	sys_signal_map_collection io_signal_in;
	sys_signal_map_collection io_signal_out;
	virtual_device_collection_phy devices;
}io_bus;

io_bus *get_io_bus(void);
int get_iobus_io_var(char *name);


#endif
