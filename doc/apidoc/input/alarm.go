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
* @api {get} /v1/alarm/timestamp/:timestamp getLogsByTimeStamp
* @apiVersion 0.1.0
* @apiName getLogsByTimeStamp
* @apiGroup Alarm
*
* @apiDescription getLogsByTimeStamp
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/alarm/timestamp/:timestamp
*
* @apiUse DefaultHeader
*
* @apiParam (Params) {string} timestamp  	URL Param: timestamp to set
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {get} /v1/alarm/errno/:errno getLogsByErrNo
* @apiVersion 0.1.0
* @apiName getLogsByErrNo
* @apiGroup Alarm
*
* @apiDescription getLogsByErrNo
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/alarm/errno/:errno
*
* @apiUse DefaultHeader
*
* @apiParam (Params) {string} errno  	URL Param: errno to set
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