#include "resource.h"

static void get_cRobLB_with_range(cJSON* root, int num, int start, int end) {
    unsigned int *base;
    int count;
    if (start >= LB_COUNT || num >= CALL_NEST_NUM) {
        return;
    }
    if (end > LB_COUNT) {
        end = LB_COUNT;
    }
    base = SHARE_RES(locvar)[num].cRobLB + start;
    count = end - start;
    cJSON_AddItemToObject(root, "cRobLB", cJSON_CreateIntArray((const int *)base, count));
    cJSON_AddItemToObject(root, "TotalSize", cJSON_CreateNumber(LB_COUNT));
    cJSON_AddItemToObject(root, "NestNum", cJSON_CreateNumber(CALL_NEST_NUM));
}

static void get_iRobLI_with_range(cJSON* root, int num, int start, int end) {
    int i=0;
    cJSON *array;
    if (start >= LI_COUNT || num >= CALL_NEST_NUM) {
        return;
    }
    if (end > LI_COUNT) {
        end = LI_COUNT;
    }

    array = cJSON_CreateArray();
    for (i=start; i<end; i++) {
        cJSON_AddItemToArray(array, cJSON_CreateNumber(SHARE_RES(locvar)[num].iRobLI[i]));
    }
    cJSON_AddItemToObject(root, "iRobLI", array);
    cJSON_AddItemToObject(root, "TotalSize", cJSON_CreateNumber(LI_COUNT));
    cJSON_AddItemToObject(root, "NestNum", cJSON_CreateNumber(CALL_NEST_NUM));
}

static void get_dRobLD_with_range(cJSON* root, int num, int start, int end) {
    double *base;
    int count;
    if (start >= LD_COUNT || num >= CALL_NEST_NUM) {
        return;
    }
    if (end > LD_COUNT) {
        end = LD_COUNT;
    }
    base = SHARE_RES(locvar)[num].dRobLD + start;
    count = end - start;
    cJSON_AddItemToObject(root, "dRobLD", cJSON_CreateDoubleArray(base, count));
    cJSON_AddItemToObject(root, "TotalSize", cJSON_CreateNumber(LD_COUNT));
    cJSON_AddItemToObject(root, "NestNum", cJSON_CreateNumber(CALL_NEST_NUM));
}

static void get_dRobLP_with_range(cJSON* root, int num, int start, int end) {
    cJSON *array;
    int i;
    if (start > LP_COUNT || num >= CALL_NEST_NUM) {
        return;
    }
    if (end > LP_COUNT) {
        end = LP_COUNT;
    }
    
    array = cJSON_CreateArray();
    for (i=start; i<end; i++) {
        cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(locvar)[num].dRobLP[i], AXIS_COUNT));
    }
    cJSON_AddItemToObject(root, "dRobLP", array);
    cJSON_AddItemToObject(root, "TotalSize", cJSON_CreateNumber(LP_COUNT));
    cJSON_AddItemToObject(root, "NestNum", cJSON_CreateNumber(CALL_NEST_NUM));
}

static void get_dRobLV_with_range(cJSON* root, int num, int start, int end) {
    cJSON *array;
    int i;
    if (start > LV_COUNT || num >= CALL_NEST_NUM) {
        return;
    }
    if (end > LV_COUNT) {
        end = LV_COUNT;
    }

    array = cJSON_CreateArray();
    for (i=start; i<end; i++) {
        cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(locvar)[num].dRobLV[i], VSub_COUNT));
    }
    cJSON_AddItemToObject(root, "dRobLV", array);
    cJSON_AddItemToObject(root, "TotalSize", cJSON_CreateNumber(LV_COUNT));
    cJSON_AddItemToObject(root, "NestNum", cJSON_CreateNumber(CALL_NEST_NUM));
}

getLocVarFunc locVarTable[] = {
    &get_cRobLB_with_range,
    &get_iRobLI_with_range,
    &get_dRobLD_with_range,
    &get_dRobLP_with_range,
    &get_dRobLV_with_range,
};

static cJSON* get_locvar_with_range(int datatype, int number, int start, int end) {
    getLocVarFunc gf;
    cJSON* root=NULL;
    switch (datatype) {
        case GetLocVarcRobLB:
        case GetLocVariRobLI:
        case GetLocVardRobLD:
        case GetLocVardRobLP:
        case GetLocVardRobLV:
            root = cJSON_CreateObject();
            gf = locVarTable[datatype];
            gf(root, number, start, end);
            break;
        default:
            return NULL;
    }

    return root;
}

char* elt_get_locvar(int datatype, int number, int start, int end) {
    cJSON* root;
    char* ret;
    root = get_locvar_with_range(datatype, number, start, end);
    ret = cJSON_PrintUnformatted(root);
    cJSON_Delete(root);
    return ret;
}

static void get_cRobB_with_range(cJSON* root, int start, int end) {
    unsigned int *base;
    int count;
    if (start > B_COUNT) {
        return;
    }
    if (end > B_COUNT) {
        end = B_COUNT;
    }
    base = SHARE_RES(sysvar).cRobB + start;
    count = end - start;
    cJSON_AddItemToObject(root, "cRobB", cJSON_CreateIntArray((const int *)base, count));
    cJSON_AddItemToObject(root, "TotalSize", cJSON_CreateNumber(B_COUNT));
}

static void get_iRobI_with_range(cJSON* root, int start, int end) {
    int i=0;
    cJSON *array;
    if (start > I_COUNT) {
        return;
    }
    if (end > I_COUNT) {
        end = I_COUNT;
    }

    array = cJSON_CreateArray();
    for (i=start; i<end; i++) {
        cJSON_AddItemToArray(array, cJSON_CreateNumber(SHARE_RES(sysvar).iRobI[i]));
    }
    cJSON_AddItemToObject(root, "iRobI", array);
    cJSON_AddItemToObject(root, "TotalSize", cJSON_CreateNumber(I_COUNT));
}

static void get_dRobD_with_range(cJSON* root, int start, int end) {
    double *base;
    int count;
    if (start > D_COUNT) {
        return;
    }
    if (end > D_COUNT) {
        end = D_COUNT;
    }
    base = SHARE_RES(sysvar).dRobD + start;
    count = end - start;
    cJSON_AddItemToObject(root, "dRobD", cJSON_CreateDoubleArray(base, count));
    cJSON_AddItemToObject(root, "TotalSize", cJSON_CreateNumber(D_COUNT));
}

static void get_dRobP_with_range(cJSON* root, int start, int end) {
    cJSON *array;
    int i;
    if (start > P_COUNT) {
        return;
    }
    if (end > P_COUNT) {
        end = P_COUNT;
    }
    
    array = cJSON_CreateArray();
    for (i=start; i<end; i++) {
        cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(sysvar).dRobP[i], AXIS_COUNT));
    }
    cJSON_AddItemToObject(root, "dRobP", array);
    cJSON_AddItemToObject(root, "TotalSize", cJSON_CreateNumber(P_COUNT));
}

static void get_dRobV_with_range(cJSON* root, int start, int end) {
    cJSON *array;
    int i;
    if (start > V_COUNT) {
        return;
    }
    if (end > V_COUNT) {
        end = V_COUNT;
    }

    array = cJSON_CreateArray();
    for (i=start; i<end; i++) {
        cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(sysvar).dRobV[i], VSub_COUNT));
    }
    cJSON_AddItemToObject(root, "dRobV", array);
    cJSON_AddItemToObject(root, "TotalSize", cJSON_CreateNumber(V_COUNT));
}

getSysVarFunc sysVarTable[] = {
    &get_cRobB_with_range,
    &get_iRobI_with_range,
    &get_dRobD_with_range,
    &get_dRobP_with_range,
    &get_dRobV_with_range,
};

static cJSON* get_sysvar_with_range(int datatype, int start, int end) {
    getSysVarFunc gf;
    cJSON* root=NULL;
    switch (datatype) {
        case GetSysVarcRobB:
        case GetSysVariRobI:
        case GetSysVardRobD:
        case GetSysVardRobP:
        case GetSysVardRobV:
            root = cJSON_CreateObject();
            gf = sysVarTable[datatype];
            gf(root, start, end);
            break;
        default:
            return NULL;
    }

    return root;
}

char* elt_get_sysvar(int datatype, int start, int end) {
    cJSON* root;
    char* ret;
    root = get_sysvar_with_range(datatype, start, end);
    ret = cJSON_PrintUnformatted(root);
    cJSON_Delete(root);
    return ret;
}

int elt_get_io_in(int addr)
{
    int index = 0;
    int value = 0;

    if(addr < 0) {
        return -1;
    }

    index = addr >> 3;
    if(index >= IO_IN_NUM) {
        return -1;
    }

    value = SHARE_RES(plc.PLC_IN[index]);
    value=(value>>(addr&7))&0x01;

    return value;
}

int elt_get_io_out(int addr)
{
    int index = 0;
    int value = 0;

    if(addr < 0) {
        return -1;
    }

    index = addr >> 3;
    if(index >= IO_OUT_NUM) {
        return -1;
    }

    value = SHARE_RES(plc.PLC_OUT[index]);
    value=(value>>(addr&7))&0x01;

    return value;
}

int elt_get_io_vin(int addr)
{
    int index = 0;
    int value = 0;

    if(addr < 0 || addr >= M_OUT_START) {
        return -1;
    }

    index = addr >> 3;
    value = SHARE_RES(plc.PLC_M[index]);
    value=(value>>(addr&7))&0x01;

    return value;
}

int elt_get_io_vout(int addr)
{
    int index = 0;
    int value = 0;

    if(addr < M_OUT_START) {
        return -1;
    }

    index = addr >> 3;
    if(index >= IO_M_NUM) {
        return -1;
    }

    value = SHARE_RES(plc.PLC_M[index]);
    value=(value>>(addr&7))&0x01;

    return value;
}

static cJSON* setElementToArray(volatile uint8_t* base, int array_size)
{
    int i=0;
    cJSON* array = cJSON_CreateArray();
    for (;i<array_size;i++) {
        cJSON_AddItemToArray(array, cJSON_CreateNumber((double)base[i]));
    }
    return array;
}

static cJSON* get_plc()
{
    cJSON* item = cJSON_CreateObject();
    cJSON_AddItemToObject(item, "PLC_IN", setElementToArray(SHARE_RES(plc).PLC_IN, IO_IN_NUM));
    cJSON_AddItemToObject(item, "PLC_OUT", setElementToArray(SHARE_RES(plc).PLC_OUT, IO_OUT_NUM));
    cJSON_AddItemToObject(item, "PLC_VIN", setElementToArray(SHARE_RES(plc).PLC_VIN, IO_VIN_NUM));
    cJSON_AddItemToObject(item, "PLC_VOUT", setElementToArray(SHARE_RES(plc).PLC_VOUT, IO_VOUT_NUM));
    cJSON_AddItemToObject(item, "PLC_M", setElementToArray(SHARE_RES(plc).PLC_M, IO_M_NUM));
    return item;
}

char* elt_get_plc()
{
    cJSON* root = cJSON_CreateObject();
    cJSON_AddItemToObject(root, "plc", get_plc());

    char *ret = cJSON_PrintUnformatted(root);
    cJSON_Delete(root);
    return ret;
}

char* elt_get_main_program_name()
{
    cJSON* json = cJSON_CreateString(GetMainfile());

    char *ret = cJSON_PrintUnformatted(json);
    cJSON_Delete(json);
    return ret;
}

int elt_get_manual_speed_rate()
{
    return GetManualSpeedRate();
}

int elt_get_tool_num()
{
    return GetToolNumber();
}

int elt_get_user_num()
{
    return GetUserNumber();
}

static cJSON* get_zero_encode()
{
    int i = 0;
    cJSON* item;
    item = cJSON_CreateArray();
    for(i=0; i<AXIS_COUNT; i++) {
        cJSON_AddItemToArray(item, cJSON_CreateNumber(Get_ZeroEncode(i)));
    }
    return item;
}

static cJSON* get_origin()
{
    int i = 0;
    cJSON* item;
    item = cJSON_CreateArray();
    for(i=0; i<AXIS_COUNT; i++) {
        cJSON_AddItemToArray(item, cJSON_CreateNumber(Get_RobotRes_Origin(i)));
    }
    return item;
}

static cJSON* get_nv()
{
    cJSON* item;
    item = cJSON_CreateObject();
    cJSON_AddItemToObject(item, "projectname", cJSON_CreateString(GetMainfile()));
    cJSON_AddItemToObject(item, "sys_technics", cJSON_CreateNumber(getTechnics()));
    cJSON_AddItemToObject(item, "curcoordinate", cJSON_CreateNumber(GetCurCoordinate()));
    cJSON_AddItemToObject(item, "manualspeedRate", cJSON_CreateNumber(GetManualSpeedRate()));
    cJSON_AddItemToObject(item, "cyclemode", cJSON_CreateNumber(GetCycleMode()));
    cJSON_AddItemToObject(item, "toolnum", cJSON_CreateNumber(GetToolNumber()));
    cJSON_AddItemToObject(item, "usernum", cJSON_CreateNumber(GetUserNumber()));
    cJSON_AddItemToObject(item, "zero_encode", get_zero_encode());
    cJSON_AddItemToObject(item, "system_period", cJSON_CreateNumber(GetSysPeriod()));
    cJSON_AddItemToObject(item, "system_ctrl_mode", cJSON_CreateNumber(GetSysCtrlMode()));
    cJSON_AddItemToObject(item, "origin", get_origin());
    return item;
}

char* elt_get_nv()
{
    cJSON* root = cJSON_CreateObject();
    cJSON_AddItemToObject(root, "nv", get_nv());

    char *ret = cJSON_PrintUnformatted(root);
    cJSON_Delete(root);
    return ret;
}

int elt_get_robot_state()
{
    return GetRobotState();
}

char* elt_get_cur_robot_pos()
{
    cJSON* json = cJSON_CreateDoubleArray(RobotRes_MachPos_base, AXIS_COUNT);

    char *ret = cJSON_PrintUnformatted(json);
    cJSON_Delete(json);
    return ret;
}

char* elt_get_cur_robot_pose()
{
    cJSON* json = cJSON_CreateDoubleArray(RobotRes_MachPose_base, 6);

    char *ret = cJSON_PrintUnformatted(json);
    cJSON_Delete(json);
    return ret;
}

int elt_get_cur_program_line()
{
    return GetCurLine();
}

int elt_get_servo_enabled()
{
    return GetServoReady();
}

int elt_can_motor_run()
{
    return GETMOTOR_RUN_STATE();
}

int elt_get_speed_modify_play()
{
    return NVRAM_PARA(speed_modify_play);
}

int elt_get_robot_mode()
{
    return GET_TEACH_MODE();
}

static cJSON* get_resource()
{
    cJSON* item;
    item = cJSON_CreateObject();
    cJSON_AddItemToObject(item, "autorun_cycleMode", cJSON_CreateNumber(GET_AUTORUN_CYCLEMODE()));
    cJSON_AddItemToObject(item, "teach_mode", cJSON_CreateNumber(GET_TEACH_MODE()));
    cJSON_AddItemToObject(item, "machinePos", cJSON_CreateDoubleArray(RobotRes_MachPos_base, AXIS_COUNT));
    cJSON_AddItemToObject(item, "machinePose", cJSON_CreateDoubleArray(RobotRes_MachPose_base, 6));
    cJSON_AddItemToObject(item, "abs_pulse", cJSON_CreateIntArray((const int *)SHARE_RES(abs_pulse), AXIS_COUNT));
    cJSON_AddItemToObject(item, "abz_pulse", cJSON_CreateIntArray((const int *)SHARE_RES(abz_pulse), AXIS_COUNT));
    cJSON_AddItemToObject(item, "cur_encode", cJSON_CreateIntArray((const int *)SHARE_RES(cur_encode), AXIS_COUNT));
    cJSON_AddItemToObject(item, "robotState", cJSON_CreateNumber(GetRobotState()));
    cJSON_AddItemToObject(item, "servoEnabled", cJSON_CreateNumber(GetServoReady()));
    cJSON_AddItemToObject(item, "canMotorRun", cJSON_CreateNumber(GETMOTOR_RUN_STATE()));
    cJSON_AddItemToObject(item, "currentLine", cJSON_CreateNumber(GetCurLine()));
    cJSON_AddItemToObject(item, "robotMode", cJSON_CreateNumber(GetRobotMode()));
    cJSON_AddItemToObject(item, "locvar_num", cJSON_CreateNumber(SHARE_RES(locvar_num)));
    cJSON_AddItemToObject(item, "analog_ioInput", cJSON_CreateDoubleArray(SHARE_RES(analog_ioInput), ANALOG_IN_NUM));
    cJSON_AddItemToObject(item, "servo_dirve_mode", cJSON_CreateNumber(SHARE_RES(servo_dirve_mode)));
    cJSON_AddItemToObject(item, "speed_modify_play", cJSON_CreateNumber(NVRAM_PARA(speed_modify_play)));
    return item;
}

char* elt_get_resource()
{
    cJSON* root = cJSON_CreateObject();
    cJSON_AddItemToObject(root, "resource", get_resource());

    char *ret = cJSON_PrintUnformatted(root);
    cJSON_Delete(root);
    return ret;
}

int elt_get_remote_mode_status()
{
    return (int)GetRemoteAccessState();
}

int elt_get_encryption_status()
{
    return GetEncryptState();
}

int elt_get_encryption_remain_time()
{
    return GetEncryptRemainTime();
}

char* elt_get_machine_code()
{
    return GetEncryptMachCode();
}

int elt_resource_init()
{
    int ret = init_nvram(NULL);
    if (ret < 0)
    {
        return ret;
    }

    return resource_init("/rbctrl/bin/mcserver");
}