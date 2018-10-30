/**
* @api {post} /v1/robot/autosettool/cmd_gototoolpos/:num doGotoToolPos
* @apiVersion 0.1.0
* @apiName doGotoToolPos
* @apiGroup Autosettool
*
* @apiDescription doGotoToolPos
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/autosettool/cmd_gototoolpos/:num
*
* @apiUse DefaultHeader
* 
* @apiParam (Params) {string="0~6"} num  	URL Param: num to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {post} /v1/robot/autosettool/cmd_autosettoolframe doAutoSetToolFrame
* @apiVersion 0.1.0
* @apiName doAutoSetToolFrame
* @apiGroup Autosettool
*
* @apiDescription doAutoSetToolFrame
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/robot/autosettool/cmd_autosettoolframe
*
* @apiUse DefaultHeader
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/