/**
* @api {post} /v1/robot/interference/data/:num setInterferData
* @apiVersion 0.1.0
* @apiName setInterferData
* @apiGroup Interference
*
* @apiDescription setInterferData
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/interference/data/:num
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="0~15"} num  	URL Param: num to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/