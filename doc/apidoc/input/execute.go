/**
* @api {post} /v1/robot/execute/run/:args runCmd
* @apiVersion 0.1.0
* @apiName runCmd
* @apiGroup Execute
*
* @apiDescription runCmd
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/execute/run/:args
*
* @apiUse DefaultHeader
* 
* @apiParam (params) {string} args  	   URL Param: args to set
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {post} /v1/robot/execute/pause/:args doPause
* @apiVersion 0.1.0
* @apiName doPause
* @apiGroup Execute
*
* @apiDescription doPause
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/execute/pause/:args
*
* @apiUse DefaultHeader
* 
* @apiParam (params) {string} args  	   URL Param: args to set
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {put} /v1/robot/execute/mode/:mode setRobotMode
* @apiVersion 0.1.0
* @apiName setRobotMode
* @apiGroup Execute
*
* @apiDescription setRobotMode
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/execute/mode/:mode
*
* @apiUse DefaultHeader
* 
* @apiParam (params) {string} mode  	   URL Param: mode to set
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {post} /v1/robot/execute/alarm/:args clearAlarm
* @apiVersion 0.1.0
* @apiName clearAlarm
* @apiGroup Execute
*
* @apiDescription clearAlarm
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/execute/alarm/:args
*
* @apiUse DefaultHeader
* 
* @apiParam (params) {string} args  	   URL Param: args to set
*
* @apiUse DefaultJsonResponsesAndExample
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
* @apiParam (params) {string} data  	   URL Param: data to set
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {put} /v1/robot/execute/mainfile/:filename setMainfile
* @apiVersion 0.1.0
* @apiName setMainfile
* @apiGroup Execute
*
* @apiDescription setMainfile
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/execute/mainfile/:filename
*
* @apiUse DefaultHeader
* 
* @apiParam (params) {string} filename  	   URL Param: filename to set
*
* @apiUse DefaultJsonResponsesAndExample
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
* @apiParam (params) {string} cyclemode  	   URL Param: cyclemode to set
*
* @apiUse DefaultJsonResponsesAndExample
*/