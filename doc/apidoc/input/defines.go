/**
* @apiDefine DefaultHeader 
* @apiHeader {string="application/json","charset=utf-8"} Content-Type=application/json;charset=utf-8
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
* @apiError (Error Status 400) BadRequest 				Bad Request
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
* @apiError (Error Status 500) InternalServerError     Internal Server Error
*
* @apiErrorExample Response (example):
*		HTTP/1.1 500 Internal Server Error
*		{
*			"errcode":	200,
*			"errmsg":	"Mcserver is not available right now"		
*		}
*/