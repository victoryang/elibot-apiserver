/**
* @apiDefine DefaultHeaderAndExample 
* @apiHeader {string="application/json"} Content-Type=application/json;charset=utf-8 only support json body content
*
* @apiHeaderExample {json} Request-Header-Example:
* 		{
*			"Content-Type": "application/json;charset=utf-8"
*		}
*/

/**
* @apiDefine DefaultResponsesAndExample
* @apiSuccess {string} msg returned message
*
* @apiError (Error 5xx) InternalServerError Some error happens
*
* @apiErrorExample Response (example):
*		HTTP/1.1 500 Internal Server Error
*		{
*			"errcode":	100,
*			"errmsg":	"error!"		
*		}
*/