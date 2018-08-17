# include "mcplc.h"

cJSON* setElementToArray(volatile uint8_t* base, int array_size) {
	int i=0;
	cJSON* array = cJSON_CreateArray();
	for (;i<array_size;i++) {
		cJSON_AddItemToArray(array, cJSON_CreateNumber((double)base[i]));
	}
	return array;
}

cJSON* get_plc() {
	cJSON* item = cJSON_CreateObject();
	cJSON_AddItemToObject(item, "PLC_IN", setElementToArray(SHARE_RES(plc).PLC_IN, IO_IN_NUM));
	cJSON_AddItemToObject(item, "PLC_OUT", setElementToArray(SHARE_RES(plc).PLC_OUT, IO_OUT_NUM));
	cJSON_AddItemToObject(item, "PLC_VIN", setElementToArray(SHARE_RES(plc).PLC_VIN, IO_VIN_NUM));
	cJSON_AddItemToObject(item, "PLC_VOUT", setElementToArray(SHARE_RES(plc).PLC_VOUT, IO_VOUT_NUM));
	cJSON_AddItemToObject(item, "PLC_M", setElementToArray(SHARE_RES(plc).PLC_M, IO_M_NUM));
	return item;
}