/**
* @api {put} /v1/robot/manualinterpolation/coord/:mode setCoordinateMode
* @apiVersion 0.1.0
* @apiName setCoordinateMode
* @apiGroup manualinterpolation
*
* @apiDescription setCoordinateMode
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/manualinterpolation/coord/:mode
*
* @apiUse DefaultHeaderAndExample
* 
* @apiParam (params) {string} mode  	   URL Param: mode to set
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {put} /v1/robot/manualinterpolation/manual/:key/:arg setManual
* @apiVersion 0.1.0
* @apiName setManual
* @apiGroup manualinterpolation
*
* @apiDescription setManual
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/manualinterpolation/manual/:key/:arg
*
* @apiUse DefaultHeaderAndExample
* 
* @apiParam (params) {string} key  	   URL Param: key to set
*
* @apiParam (params) {string} arg  	   URL Param: arg to set
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {put} /v1/robot/manualinterpolation/runforward runForward
* @apiVersion 0.1.0
* @apiName runForward
* @apiGroup manualinterpolation
*
* @apiDescription runForward
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/manualinterpolation/runforward
*
* @apiUse DefaultHeaderAndExample
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {put} /v1/robot/manualinterpolation/run/:zero runToZero
* @apiVersion 0.1.0
* @apiName runToZero
* @apiGroup manualinterpolation
*
* @apiDescription runToZero
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/manualinterpolation/run/:zero
*
* @apiUse DefaultHeaderAndExample
* 
* @apiParam (params) {string} zero  	   URL Param: zero to set
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {put} /v1/robot/manualinterpolation/stop robotStop
* @apiVersion 0.1.0
* @apiName robotStop
* @apiGroup manualinterpolation
*
* @apiDescription robotStop
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/manualinterpolation/stop
*
* @apiUse DefaultHeaderAndExample
*
* @apiUse DefaultResponsesAndExample
*/