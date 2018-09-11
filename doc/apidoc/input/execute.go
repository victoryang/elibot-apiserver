/**
* @api {post} /v1/robot/execute/cmd_run doRunCmd
* @apiVersion 0.1.0
* @apiName doRunCmd
* @apiGroup Execute
*
* @apiDescription doRunCmd
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/execute/cmd_run
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string[]="-c","-s","-f <file>","-l <line>", "-n <count>"} args  	   BODY Param: run -c, continue mode. 
*																									run -s, run one step, don't clear envirment.
*																									run -f <file> -l <line> -n <count>, run from file@line
*
* @apiUse SuccessStringResponse
* @apiUse BadRequestJsonResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {post} /v1/robot/execute/cmd_pause doPause
* @apiVersion 0.1.0
* @apiName doPause
* @apiGroup Execute
*
* @apiDescription doPause
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/execute/cmd_pause
*
* @apiUse DefaultHeader
* 
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {put} /v1/robot/execute/robotmode/:mode setRobotMode
* @apiVersion 0.1.0
* @apiName setRobotMode
* @apiGroup Execute
*
* @apiDescription setRobotMode
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/execute/robotmode/:mode
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="play","teach","remote"} mode  	URL Param: mode to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {post} /v1/robot/execute/cmd_clearalarm doClearAlarm
* @apiVersion 0.1.0
* @apiName doClearAlarm
* @apiGroup Execute
*
* @apiDescription doClearAlarm
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/execute/cmd_clearalarm
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string[]="-f"} 	[args]  	   Body Param: -f for force to ClearAlarm ,ten minites
*
* @apiUse SuccessStringResponse
* @apiUse BadRequestJsonResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {post} /v1/robot/execute/cmd_progreset doProgReset
* @apiVersion 0.1.0
* @apiName doProgReset
* @apiGroup Execute
*
* @apiDescription doProgReset
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/execute/cmd_progreset
*
* @apiUse DefaultHeader
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {put} /v1/robot/execute/speed/:data setSpeed
* @apiVersion 0.1.0
* @apiName setSpeed
* @apiGroup Execute
*
* @apiDescription setSpeed
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/execute/speed/:data
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="0~10000"} data  	   URL Param: data to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {put} /v1/robot/execute/mainprog/:progname setMainProg
* @apiVersion 0.1.0
* @apiName setMainProg
* @apiGroup Execute
*
* @apiDescription setMainProg
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/execute/mainprog/:progname
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string} progname  	   URL Param: progname to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {put} /v1/robot/execute/cyclemode/:cyclemode setCycleMode
* @apiVersion 0.1.0
* @apiName setCycleMode
* @apiGroup Execute
*
* @apiDescription setCycleMode
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/execute/cyclemode/:cyclemode
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="step","one","SERIES"} cyclemode  	   URL Param: cyclemode to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {post} /v1/robot/execute/cmd_load/:filename doLoadFile
* @apiVersion 0.1.0
* @apiName doLoadFile
* @apiGroup Execute
*
* @apiDescription doLoadFile
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/execute/cmd_load/:filename
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string} 		filename  	   URL Param: file to load
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {put} /v1/robot/execute/curline/:lineno setCurLine
* @apiVersion 0.1.0
* @apiName setCurLine
* @apiGroup Execute
*
* @apiDescription setCurLine
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/execute/curline/:lineno
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string} 	lineno  	   URL Param: line no to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {post} /v1/robot/execute/cmd_gobackorigin doGobackOrigin
* @apiVersion 0.1.0
* @apiName doGobackOrigin
* @apiGroup Execute
*
* @apiDescription doGobackOrigin
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.255:9000/v1/robot/execute/cmd_gobackorigin
*
* @apiUse DefaultHeader
* 
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/