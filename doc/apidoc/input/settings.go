/**
* @apiDefine RequestValue
*
* @apiParam (Params) {string} value    Body Content: value
*
* @apiParamExample {json} Request-Body (example):
* 		{
*			"value": "0"
*		}
*/

/**
* @api {get} /v1/settings/kv getAllKV
* @apiVersion 0.1.0
* @apiName getAllKV
* @apiGroup Settings
*
* @apiDescription getAllKV
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/settings/kv
*
* @apiUse DefaultHeader
*
* @apiUse SuccessJsonResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {get} /v1/settings/kv/:key getKV
* @apiVersion 0.1.0
* @apiName getKV
* @apiGroup Settings
*
* @apiDescription getKV
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/settings/kv/:key
*
* @apiParam (Params) {string} key  URL Param: key to get
* @apiUse DefaultHeader
*
* @apiUse SuccessJsonResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {post} /v1/settings/kv/:key setKV
* @apiVersion 0.1.0
* @apiName setKV
* @apiGroup Settings
*
* @apiDescription setKV
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/settings/kv/:key
*
* @apiParam (Params) {string} key  URL Param: key to set
* @apiUse RequestValue
* @apiUse DefaultHeader
*
* @apiUse SuccessJsonResponse
* @apiUse BadRequestJsonResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {get} /v1/settings/reboot rebootSystem
* @apiVersion 0.1.0
* @apiName rebootSystem
* @apiGroup Settings
*
* @apiDescription rebootSystem
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/settings/reboot
*
* @apiUse DefaultHeader
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {get} /v1/settings/date getSystemDate
* @apiVersion 0.1.0
* @apiName getSystemDate
* @apiGroup Settings
*
* @apiDescription getSystemDate
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/settings/date
*
* @apiUse DefaultHeader
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {put} /v1/settings/date/:date setSystemDate
* @apiVersion 0.1.0
* @apiName setSystemDate
* @apiGroup Settings
*
* @apiDescription setSystemDate
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/settings/date/:date
*
* @apiParam (Params) {string="1900-1-1", "12:00:00"} date  URL Param: date to set
* @apiUse DefaultHeader
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {get} /v1/settings/ip getSystemIP
* @apiVersion 0.1.0
* @apiName getSystemIP
* @apiGroup Settings
*
* @apiDescription getSystemIP
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/settings/ip
*
* @apiUse DefaultHeader
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @apiDefine RequestBodyForSettingIP
*
* @apiParam (Params) {string} address    Body Content: address
* @apiParam (Params) {string} netmask    Body Content: netmask
* @apiParam (Params) {string} network    Body Content: network
* @apiParam (Params) {string} broadcast  Body Content: broadcast
* @apiParam (Params) {string} gateway    Body Content: gateway
*
* @apiParamExample {json} Request-Body (example):
* 		{
*			"address":"192.168.1.253",
*			"netmask":"255.255.255.0",
*			"network":"192.168.1.0",
*			"broadcast":"192.168.255.255",
*			"gateway":"192.168.1.1"
*		}
*/

/**
* @api {put} /v1/settings/ip/:ip setSystemIP
* @apiVersion 0.1.0
* @apiName setSystemIP
* @apiGroup Settings
*
* @apiDescription setSystemIP
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/settings/ip
*
* @apiUse RequestBodyForSettingIP
* @apiUse DefaultHeader
*
* @apiUse SuccessStringResponse
* @apiUse BadRequestJsonResponse
* @apiUse InternalServerErrorJsonResponse
*/