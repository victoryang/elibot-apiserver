/**
* @api {put} /v1/robot/manual/coord/:mode setCoordinateMode
* @apiVersion 0.1.0
* @apiName setCoordinateMode
* @apiGroup manual
*
* @apiDescription setCoordinateMode
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/manual/coord/:mode
*
* @apiUse DefaultHeader
* 
* @apiParam (params) {string="joint","cart","tool","user","cylinder"} mode  	   URL Param: mode to set
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {post} /v1/robot/manual/cmd_manual/:axis doManual
* @apiVersion 0.1.0
* @apiName doManual
* @apiGroup manual
*
* @apiDescription doManual
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/manual/cmd_manual/:axis
*
* @apiUse DefaultHeader
* 
* @apiParam (params) {string="0~11"} axis  	   URL Param: axis to set
*
* @apiParam (params) {string="-r"} args 	   Body Param: args to set
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {post} /v1/robot/manual/cmd_runforward doRunForward
* @apiVersion 0.1.0
* @apiName doRunForward
* @apiGroup manual
*
* @apiDescription doRunForward
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/manual/cmd_runforward
*
* @apiUse DefaultHeader
*
* @apiParam (params) {string[]="axis..."} args 	   Body Param: args to set
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {post} /v1/robot/manual/cmd_runzero/:status doRunToZero
* @apiVersion 0.1.0
* @apiName doRunToZero
* @apiGroup manual
*
* @apiDescription doRunToZero
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/manual/cmd_runzero/:status
*
* @apiUse DefaultHeader
* 
* @apiParam (params) {string="origin","zero"} status  	   URL Param: status to set
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {post} /v1/robot/manual/cmd_stop doRobotStop
* @apiVersion 0.1.0
* @apiName doRobotStop
* @apiGroup manual
*
* @apiDescription doRobotStop
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/manual/cmd_stop
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
*/