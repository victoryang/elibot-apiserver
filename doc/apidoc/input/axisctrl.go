/**
* @api {put} /v1/robot/axisctrl/servo/:status setServoStatus
* @apiVersion 0.1.0
* @apiName setServoStatus
* @apiGroup Axisctrl
*
* @apiDescription setServoStatus
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/axisctrl/servo/:status
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="on","off"} status  	URL Param: status to set
* @apiParam (Params) {string[]="-f"} 	 [args] 	Body Param: args to set
*
* @apiUse SuccessStringResponse
* @apiUse BadRequestJsonResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {put} /v1/robot/axisctrl/dragteach/:status setDragteachStatus
* @apiVersion 0.1.0
* @apiName setDragteachStatus
* @apiGroup Axisctrl
*
* @apiDescription setDragteachStatus
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/axisctrl/dragteach/:status
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="on","off"} status  	   URL Param: status to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {post} /v1/robot/axisctrl/sync syncRobot
* @apiVersion 0.1.0
* @apiName syncRobot
* @apiGroup Axisctrl
*
* @apiDescription syncRobot
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/axisctrl/sync
*
* @apiUse DefaultHeader
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {put} /v1/robot/axisctrl/zeroencode/:axisno setZeroEncode
* @apiVersion 0.1.0
* @apiName setZeroEncode
* @apiGroup Axisctrl
*
* @apiDescription setZeroEncode
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/axisctrl/zeroencode/:axisno
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="0~AXIS_COUNT"} 	axisno  	   URL Param: axis no to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/