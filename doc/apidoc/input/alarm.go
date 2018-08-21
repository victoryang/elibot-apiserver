/**
* @api {get} /v1/alarm/ getAllLogs
* @apiVersion 0.1.0
* @apiName getAllLogs
* @apiGroup Alarm
*
* @apiDescription getAllLogs
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/alarm/
*
* @apiUse DefaultHeader
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {get} /v1/alarm/num getLogNumber
* @apiVersion 0.1.0
* @apiName getLogNumber
* @apiGroup Alarm
*
* @apiDescription getLogNumber
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/alarm/num
*
* @apiUse DefaultHeader
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {get} /v1/alarm/range?from=:from&end=:end getLogs
* @apiVersion 0.1.0
* @apiName getLogs
* @apiGroup Alarm
*
* @apiDescription getLogs
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/alarm/range?from=:from&end=:end
*
* @apiUse DefaultHeader
*
* @apiParam (Params) {string} from  	URL Param: from to set
* @apiParam (Params) {string} end  		URL Param: end to set
*
* @apiUse DefaultResponsesAndExample
*/