/**
* @api {get} /v1/alarm/range?start=:start&end=:end&timestamp=:timestamp getLogs
* @apiVersion 0.1.0
* @apiName getLogs
* @apiGroup Alarm
*
* @apiDescription getLogs
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/alarm/range?start=:start&end=:end&timestamp=:timestamp
*
* @apiUse DefaultHeader
*
* @apiParam (Params) {string} start  	URL Param: start to set
* @apiParam (Params) {string} end  		URL Param: end to set
* @apiParam (Params) {string} timestamp URL Param: timestamp to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {get} /v1/alarm/level/:level?start=:start&end=:end&timestamp=:timestamp getLogsByAlarmLevel
* @apiVersion 0.1.0
* @apiName getLogsByAlarmLevel
* @apiGroup Alarm
*
* @apiDescription getLogsByAlarmLevel
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/alarm/level/:level?start=:start&end=:end&timestamp=:timestamp
*
* @apiUse DefaultHeader
*
* @apiParam (Params) {string} level  	URL Param: level to set
* @apiParam (Params) {string} start  	URL Param: start to set
* @apiParam (Params) {string} end  		URL Param: end to set
* @apiParam (Params) {string} timestamp URL Param: timestamp to set
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
*/