/**
* @apiDefine RequestBodyIndexAndValue
*
* @apiParam (params) {string} index    Body Content: index
* @apiParam (params) {string} value    Body Content: value
*
* @apiParamExample {json} Request-Body (example):
* 		{
*			"index": "0",
*			"value": "0"
*		}
*/

/**
* @api {put} /v1/robot/repository/arcparam/:file_no/:md_id setArcParam
* @apiVersion 0.1.0
* @apiName setArcParam
* @apiGroup Repository
*
* @apiDescription setArcParam [md_id] [value] [file_no] [index]
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/repository/arcparam/:file_no/:md_id
*
* @apiUse DefaultHeaderAndExample
*
* @apiParam (params) {string} file_no  URL Param: file to set
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {put} /v1/robot/repository/interference/:no/:md_id setInterference
* @apiVersion 0.1.0
* @apiName setInterference
* @apiGroup Repository
*
* @apiDescription setInterference [md_id] [value] [no] [index]
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/repository/interference/:no/:md_id
*
* @apiUse DefaultHeaderAndExample
*
* @apiParam (params) {string} no  	   URL Param: no to set
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {put} /v1/robot/repository/params/:md_id 	setParam
* @apiVersion 0.1.0
* @apiName setParam
* @apiGroup Repository
*
* @apiDescription setParam [md_id] [value] [index]
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/repository/params/:md_id
*
* @apiUse DefaultHeaderAndExample
*
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {put} /v1/robot/repository/toolframes/:tool_no/:md_id/pos/:pos_no setToolFrame
* @apiVersion 0.1.0
* @apiName setToolFrame
* @apiGroup Repository
*
* @apiDescription setToolFrame [md_id] [value] [tool_no] [pos_no|index] [index]
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/repository/toolframes/:tool_no/:md_id/pos/:pos_no
*
* @apiUse DefaultHeaderAndExample
*
* @apiParam (params) {string} tool_no  URL Param: tool no to set
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiParam (params) {string} pos_no   URL Param: pos num
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {put} /v1/robot/repository/userframe/:userno/:md_id 	setUserFrame 
* @apiVersion 0.1.0
* @apiName setUserFrame
* @apiGroup Repository
*
* @apiDescription setUserFrame [md_id] [value] [userNo]
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/repository/userframe/:userno/:md_id
*
* @apiUse DefaultHeaderAndExample
*
* @apiParam (params) {string} userno   URL Param: user no
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {put} /v1/robot/repository/zeropoint/:md_id 	setZeroPoint
* @apiVersion 0.1.0
* @apiName setZeroPoint
* @apiGroup Repository
*
* @apiDescription setZeroPoint [md_id] [value] [index]
*
* @apiExample Example usage:
* curl -X PUT http://192.168.1.253:9000/v1/robot/repository/zeropoint/:md_id
*
* @apiUse DefaultHeaderAndExample
*
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultResponsesAndExample
*/