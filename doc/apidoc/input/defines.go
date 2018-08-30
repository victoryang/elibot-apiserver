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
*			"errcode":	101,
*			"errmsg":	"Mcserver is not available right now"		
*		}
*/

/**
* @apiDefine SuccessStringResponse
* @apiSuccess {String} Body  return String message to client
*/

/**
* @apiDefine SuccessJsonResponse
* @apiSuccess {Object} Body  return json message to client
*/

/**
* @apiDefine BadRequestJsonResponse
* @apiError (Error 400) BadRequest 				Bad Request happen
*
* @apiErrorExample Response (example):
*		HTTP/1.1 400 Bad Request
*		{
*			"errcode":	100,
*			"errmsg":	"Could not parse request body"		
*		}
*/

/**
* @apiDefine InternalServerErrorJsonResponse
* @apiError (Error 500) InternalServerError     Internal Server Error
*
* @apiErrorExample Response (example):
*		HTTP/1.1 500 Internal Server Error
*		{
*			"errcode":	101,
*			"errmsg":	"Mcserver is not available right now"		
*		}
*/