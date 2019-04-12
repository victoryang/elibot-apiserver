// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_IOBUS_SQL_MAPPER_H
#define MCSERVER_IOBUS_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取IO总线设备列表
#define ELIBOT_IOBUS_GET_DEVICES "elibot.iobus.getDevices"

// 获取IO总线设备预览列表
#define ELIBOT_IOBUS_GET_DEVICES_OVERVIEW "elibot.iobus.getDevicesOverview"

// 获取指定设备的基本信息
#define ELIBOT_IOBUS_GET_DEVICE_BY_ID "elibot.iobus.getDeviceById"

// 获取指定的Modbus rtu设备
#define ELIBOT_IOBUS_GET_MODBUS_RTU_DEVICE_BY_ID "elibot.iobus.getModbusRtuDeviceById"

// 获取指定的Modbus tcp设备
#define ELIBOT_IOBUS_GET_MODBUS_TCP_DEVICE_BY_ID "elibot.iobus.getModbusTcpDeviceById"

// 获取设备IO列表
#define ELIBOT_IOBUS_GET_DEVICE_IOS "elibot.iobus.getDeviceIos"

// 检查设备是否存在
#define ELIBOT_IOBUS_EXIST_DEVICE_BY_NAME "elibot.iobus.existDeviceByName"

// 检查设备是否存在
#define ELIBOT_IOBUS_EXIST_DEVICE_BY_NAME_AND_ID "elibot.iobus.existDeviceByNameAndId"

// 获取下一个设备编号
#define ELIBOT_IOBUS_GET_NEXT_DEVICE_ORDER "elibot.iobus.getNextDeviceOrder"

// 添加设备基本信息
#define ELIBOT_IOBUS_ADD_DEVICE_BASIC_INFO "elibot.iobus.addDeviceBasicInfo"

// 添加Modbus rtu设备信息
#define ELIBOT_IOBUS_ADD_MODBUS_RTU_DEVICE_INFO "elibot.iobus.addModbusRtuDeviceInfo"

// 添加Modbus tcp设备信息
#define ELIBOT_IOBUS_ADD_MODBUS_TCP_DEVICE_INFO "elibot.iobus.addModbusTcpDeviceInfo"

// 添加设备IO
#define ELIBOT_IOBUS_ADD_DEVICE_IO "elibot.iobus.addDeviceIO"

// 更新设备基本信息
#define ELIBOT_IOBUS_MODIFY_DEVICE_BASIC_INFO "elibot.iobus.modifyDeviceBasicInfo"

// 更新Modbus rtu设备信息
#define ELIBOT_IOBUS_MODIFY_MODBUS_RTU_DEVICE_INFO "elibot.iobus.modifyModbusRtuDeviceInfo"

// 更新Modbus tcp设备信息
#define ELIBOT_IOBUS_MODIFY_MODBUS_TCP_DEVICE_INFO "elibot.iobus.modifyModbusTcpDeviceInfo"

// 删除指定设备的IO
#define ELIBOT_IOBUS_DELETE_ALL_DEVICE_IO_BY_DEVICE_ID "elibot.iobus.deleteAllDeviceIODeviceId"

// 更新设备IO
#define ELIBOT_IOBUS_MODIFY_DEVICE_IO "elibot.iobus.modifyDeviceIO"

// 使能或禁用设备
#define ELIBOT_IOBUS_ENABLE_DEVICE "elibot.iobus.enableDevice"

// 设备是否被IO变量使用
#define ELIBOT_IOBUS_IS_DEVICE_IN_USE_BY_VAR "elibot.iobus.isDeviceInUseByVar"

// 是否系统设备
#define ELIBOT_IOBUS_IS_SYS_DEVICE "elibot.iobus.isSysDevice"

// 删除设备
#define ELIBOT_IOBUS_DELETE_DEVICE "elibot.iobus.deleteDevice"

// 获取IO总线IO变量列表
#define ELIBOT_IOBUS_GET_IO_VARS "elibot.iobus.getIOVars"

// 获取指定的IO变量
#define ELIBOT_IOBUS_GET_IO_VAR_BY_ID "elibot.iobus.getIOVarById"

// 检查是否存在同名的IO变量
#define ELIBOT_IOBUS_EXISTS_IO_VAR "elibot.iobus.existsIOVar"

// 检查是否存在同名的IO变量
#define ELIBOT_IOBUS_EXISTS_IO_VAR_BY_ID "elibot.iobus.existsIOVarById"

// 添加IO变量
#define ELIBOT_IOBUS_ADD_IO_VAR "elibot.iobus.addIOVar"

// 修改IO变量
#define ELIBOT_IOBUS_MODIFY_IO_VAR "elibot.iobus.modifyIOVar"

// 删除IO变量
#define ELIBOT_IOBUS_DELETE_IO_VAR "elibot.iobus.deleteIOVar"

// 获取IO变量被占用的状态
#define ELIBOT_IOBUS_GET_IO_VAR_OCCUPIED_STATUS "elibot.iobus.getIOVarOccupiedStatus"

// 获取系统输入列表
#define ELIBOT_IOBUS_GET_SYSTEM_INPUTS "elibot.iobus.getSystemInputs"

// 设置系统输入
#define ELIBOT_IOBUS_SET_SYSTEM_INPUT "elibot.iobus.setSystemInput"

// 获取系统输出列表
#define ELIBOT_IOBUS_GET_SYSTEM_OUTPUTS "elibot.iobus.getSystemOutputs"

// 设置系统输出
#define ELIBOT_IOBUS_SET_SYSTEM_OUTPUT "elibot.iobus.setSystemOutput"

// 获取IO总线映射器
// @id:映射器标识
// @dr:参数SQL映射器
sql_mapper *get_iobus_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif //MCSERVER_IOBUS_SQL_MAPPER_H
