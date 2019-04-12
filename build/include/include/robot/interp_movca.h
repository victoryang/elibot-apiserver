/****************************************************************************
**
** Copyright (C) 2017 The Elite Company Ltd.
** Contact: http://www.elite.cn
**
** MOVCA 整圆路径规划算法
**
****************************************************************************/

#ifndef MCSERVER_INTERP_MOVCA_H
#define MCSERVER_INTERP_MOVCA_H

#include <robot/interpolation.h>

/**
 * 计算整圆指令MOVCA信息
 * @param ptype
 * @param P1
 * @param P2
 * @param P3
 * @param Radius
 * @param dist
 * @param direction
 * @param movcData
 * @return
 */
int CalculateMovcaInformation(int ptype[3], double P1[], double P2[], double P3[], double Radius, double dist, int direction, AutoMovcDataS *movcData);

/**
 * 获取关节角度值（起始点：type=0，目标点：type=1）
 * @param Joint
 * @param type
 */
void GetTargetJoint(double Joint[], int type);

/**
 * 计算圆弧几何特征信息（分别针对过渡半圆type=0，整圆type=1这两种情况）
 * @param movcData
 * @param type
 */
void CalculateCircular_DataS(AutoMovcDataS *movcData, int type);

/**
 * 计算圆弧插补信息（分别针对过渡半圆type=0，整圆type=1这两种情况）
 * @param md
 * @param movcData
 * @param type
 */
void CalculateCircular_interp(MoveData_t *md, AutoMovcInterpDataS *movcData, int type);

/**
 * 计算最佳转接点处速度，使整体切割效果更佳
 * @param md
 * @return
 */
double CalcTransferSpeed(MoveData_t *md);

/**
 * 根据实际切割现象，调整整圆对应的圆心角（>=2*pi），由此需要对实际目标点的关节角进行更新
 * @param movcData
 */
void UpdateTargetJoint(AutoMovcDataS *movcData);

/**
 * 针对切割领域的整圆指令译码函数
 * @param md
 * @return
 */
int MOVCA_interpolation_S(MoveData_t *md);

/**
 * 获取转接速度
 * @return
 */
double getTransferSpeed(void);

/**
 * 整圆指令预处理
 * @param md
 * @param serial
 */
 void MOVCA_PreProcess(MoveData_t *md, int serial);

#endif //MCSERVER_INTERP_MOVCA_H
