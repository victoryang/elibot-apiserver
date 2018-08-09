/**
* @apiDefine RequestValue
*
* @apiParam (params) {string} value    Body Content: value
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
* @apiUse DefaultHeaderAndExample
*
* @apiUse DefaultResponsesAndExample
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
* @apiParam (params) {string} key  URL Param: key to get
* @apiUse DefaultHeaderAndExample
*
* @apiUse DefaultResponsesAndExample
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
* @apiParam (params) {string} key  URL Param: key to set
* @apiUse RequestValue
* @apiUse DefaultHeaderAndExample
*
* @apiUse DefaultResponsesAndExample
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
* curl -X GET http://192.168.1.253:9000/v1/reboot
*
* @apiUse DefaultHeaderAndExample
*
* @apiUse DefaultResponsesAndExample
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
* @apiUse DefaultHeaderAndExample
*
* @apiUse DefaultResponsesAndExample
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
* @apiParam (params) {string} date  URL Param: date to set
* @apiUse DefaultHeaderAndExample
*
* @apiUse DefaultResponsesAndExample
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
* @apiUse DefaultHeaderAndExample
*
* @apiUse DefaultResponsesAndExample
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
* curl -X PUT http://192.168.1.253:9000/v1/settings/ip/:ip
*
* @apiParam (params) {string} ip  URL Param: ip to set
* @apiUse DefaultHeaderAndExample
*
* @apiUse DefaultResponsesAndExample
*/