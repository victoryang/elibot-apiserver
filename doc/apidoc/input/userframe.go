/**
* @api {put} /v1/robot/userframe/userpos/:pos_no setUserPos
* @apiVersion 0.1.0
* @apiName setUserPos
* @apiGroup Userframe
*
* @apiDescription setUserPos
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/userframe/userpos/:pos_no
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="1~3"} pos_no  	   URL Param: pos_no to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {post} /v1/robot/userframe/cmd_gotouserpos/:pos_no doGotoUserPos
* @apiVersion 0.1.0
* @apiName doGotoUserPos
* @apiGroup Userframe
*
* @apiDescription doGotoUserPos
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/userframe/cmd_gotouserpos/:pos_no
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="1~3"} pos_no  	   URL Param: pos_no to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {post} /v1/robot/userframe/cmd_caluserframe doCalUserFrame
* @apiVersion 0.1.0
* @apiName doCalUserFrame
* @apiGroup Userframe
*
* @apiDescription doCalUserFrame
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/userframe/cmd_caluserframe
*
* @apiUse DefaultHeader
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/