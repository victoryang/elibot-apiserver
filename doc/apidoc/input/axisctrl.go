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
* @apiUse DefaultJsonResponsesAndExample
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
* @apiUse DefaultJsonResponsesAndExample
*/