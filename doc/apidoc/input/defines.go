/**
* @apiDefine DefaultHeader 
* @apiHeader {string="application/json","charset=utf-8"} Content-Type=application/json;charset=utf-8
*/

/**
* @apiDefine DefaultResponsesAndExample
* @apiSuccess {String} Body  returned json message from server
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

/**
* @apiDefine DefaultJsonResponsesAndExample
* @apiSuccess {Object} Body  return json message to client
*
* @apiError (Error 5xx) InternalServerError Some error happens
*
* @apiErrorExample Response (example):
*		HTTP/1.1 500 Internal Server Error
*		{
*			"errcode":	100,
*			"errmsg":	"Some error!"		
*		}
*/