/**
* @api {put} /v1/robot/userframe/userpos/:pos_no setUserPos
* @apiVersion 0.1.0
* @apiName setUserPos
* @apiGroup Userframe
*
* @apiDescription setUserPos
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.166:9000/v1/robot/userframe/userpos/:pos_no
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="1~3"} pos_no  	   URL Param: pos_no to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/