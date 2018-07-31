#ifndef _PL_H_
#define _PL_H_

int PreMovlPositionLevel(int pType[],double startJoint[],double midJoint[], double targetJoint[], PlInterpDataS *plData, AutoMovCommDataS *movlData, AutoMovcDataS *movcData);
int PreMovlToMovsPositionLevel(int pType[],double lineStartJoint[],double parabolaStartJoint[],double parabolaMidJoint[],
			double parabolaTargetJoint[],PlInterpDataS *plData, AutoMovCommDataS *movlData, AutoMovcDataS *movcData, AutoMovsDataS *movsData);
int PreMovsToMovlPositionLevel(int pType[],double lineTargetJoint[],double parabolaStartJoint[],double parabolaMidJoint[],
			double parabolaTargetJoint[],PlInterpDataS *plData,AutoMovcDataS *movcData, AutoMovsDataS *movsData);
#endif // PL_H

