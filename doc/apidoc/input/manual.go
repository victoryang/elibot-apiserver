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
* @apiParam (Params) {string="joint","cart","tool","user","cylinder"} mode  	   URL Param: mode to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
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
* curl -X POST http://192.168.1.253:9000/v1/robot/manual/cmd_manual/:axis
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="0~11"} axis  	   URL Param: axis to set
*
* @apiParam (Params) {string="-r"} [args] 	   Body Param: args to set
*
* @apiUse SuccessStringResponse
* @apiUse BadRequestJsonResponse
* @apiUse InternalServerErrorJsonResponse
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
* curl -X POST http://192.168.1.253:9000/v1/robot/manual/cmd_runforward
*
* @apiUse DefaultHeader
*
* @apiParam (Params) {string[]="axis0, axis1..."} args 	   Body Param: args to set
*
* @apiUse SuccessStringResponse
* @apiUse BadRequestJsonResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {post} /v1/robot/manual/cmd_runtozero/:status doRunToZero
* @apiVersion 0.1.0
* @apiName doRunToZero
* @apiGroup manual
*
* @apiDescription doRunToZero
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/manual/cmd_runtozero/:status
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="origin","zero"} status  	   URL Param: status to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
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
* curl -X POST http://192.168.1.253:9000/v1/robot/manual/cmd_stop
*
* @apiUse DefaultHeader
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/