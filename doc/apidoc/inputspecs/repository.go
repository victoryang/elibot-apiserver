/**
* @api {put} /v2/robot/repository/arcparam/:file_no/:md_id Set arc parameters
* @apiVersion 0.1.0
* @apiName setArcParam
* @apiGroup Repository
*
* @apiDescription setArcParam [md_id] [value] [file_no] [index]
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v2/robot/repository/arcparam/:file_no/:md_id
*
* @apiHeader {string="application/json"} Content-Type=application/json;charset=utf-8 only support json body content
*
* @apiHeaderExample {json} Request-Header-Example:
* 		{
*			"Content-Type": "application/json;charset=utf-8"
*		}
*
* @apiParam (params) {string} file_no  URL Param: file to set
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiParam (params) {string} index    Body Content: index
* @apiParam (params) {string} value    Body Content: value
*
* @apiParamExample {json} Request-Body (example):
* 		{
*			"index": "0",
*			"value": "0"
*		}
*
* @apiSuccess {string} msg returned message
*
* @apiError (Error 5xx) InternalServerError Some error happens
*
* @apiErrorExample Response (example):
*		HTTP/1.1 500 Internal Server Error
*		{
*			"msg":	"internal error"
*		}
*/