#ifndef ARCWELDING_H
#define ARCWELDING_H
#ifdef __cplusplus
extern "C" {
#endif

#define ARC_PARA_NUM		8

struct arc_para{
    int index;                        //参数文件号
    char note[32];                 //注释
    double arcCurrent;       //焊接电流
    double arcVoltage;        //焊接电压
    double startCurrent;     //起弧电流
    double startVoltage;      //起弧电压
    double startTime;          //起弧时间
    double endCurrent;       //收弧电流
    double endVoltage;       //收弧电压
    double endTime;           //收弧时间
    double antiStickCurrent; //防粘丝电流
    double antistickVoltage; //防粘丝电压
    double antistickTime;   //防粘丝时间
    int antiRoolback;        //防回退功能
};

struct weld_para{
    double restartDist;        //再启动距离
    double restartSpeed;    //再启动速度
    double arcCheckTime;   //电弧检测时间
    double arcConfirmTime; //电弧检测确认时间
    double arcExhaustedCheckTime; //电弧耗尽检测时间
    double scrapintPaintDist;  //刮搽距离
    double scrapintPaintSpeed; //刮搽返回速度
    double prepareAspirationTime;  //预备送气时间
    double delayAspirationTime;    //延迟送气时间

    int checkCurrentInterruption;      //焊接中断弧检测
    int checkAlarmInterruption;        //电源中断弧检测
    int checkWaterInterruption;        //水冷异常检测
    int restartUpAction;                       //再启动动作
    int scrapintPaintAction;               //刮搽启动
    int antiCollisionDetection;            //防碰撞检测
	double diandongTime;		//点动送丝时间
    //电流曲线
    double outputCurrent[2];                //输出的电流
    double correspondCurrent[2];        //对应的电流
    //电压曲线
    double outputVoltage[2];                //输出的电压
    double correspondVoltage[2];        //对应的电压

	double realCur2Volt;                    //ctrl cur to correspond volt
    double realVolt2Volt;                    //ctrl volt to correspond volt
    struct arc_para arcPara[ARC_PARA_NUM];             //弧焊文件
};
#ifdef __cplusplus
}
#endif

#endif // ARCWELDING_H
