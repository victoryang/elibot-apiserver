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
* @api {get} /v1/robot/repository/arc getAllArc
* @apiVersion 0.1.0
* @apiName getAllArc
* @apiGroup Repository
*
* @apiDescription getAllArc
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/arc
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/arcparams/:file_no?group=:group getArcParams
* @apiVersion 0.1.0
* @apiName getArcParams
* @apiGroup Repository
*
* @apiDescription getArcParams
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/:file_no?group=:group
*
* @apiParam (params) {string} file_no  URL Param: file no
* @apiParam (params) {string} group    URL Param: group id
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
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
* @apiUse DefaultHeader
*
* @apiParam (params) {string} file_no  URL Param: file to set
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/bookprograms getAllBookprograms
* @apiVersion 0.1.0
* @apiName getAllBookprograms
* @apiGroup Repository
*
* @apiDescription getAllBookprograms
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/bookprograms
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/enum getAllEnum
* @apiVersion 0.1.0
* @apiName getAllEnum
* @apiGroup Repository
*
* @apiDescription getAllEnum
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/enum
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/extaxis getAllExtaxis
* @apiVersion 0.1.0
* @apiName getAllExtaxis
* @apiGroup Repository
*
* @apiDescription getAllExtaxis
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/extaxis
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/interference getAllInterference
* @apiVersion 0.1.0
* @apiName getAllInterference
* @apiGroup Repository
*
* @apiDescription getAllInterference
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/interference
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
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
* @apiUse DefaultHeader
*
* @apiParam (params) {string} no  	   URL Param: no to set
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/metadata?lang=:lang getAllMetadata
* @apiVersion 0.1.0
* @apiName getAllMetadata
* @apiGroup Repository
*
* @apiDescription getAllMetadata
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/metadata?lang=:lang
*
* @apiUse DefaultHeader
*
* @apiParam (params) {string} lang  	URL Param: lang to set
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/params getParams
* @apiVersion 0.1.0
* @apiName getParams
* @apiGroup Repository
*
* @apiDescription getParams
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/params
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/parameter/id/:md_id getParameterById
* @apiVersion 0.1.0
* @apiName getParameterById
* @apiGroup Repository
*
* @apiDescription getParameterById
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/parameter/id/:md_id
*
* @apiUse DefaultHeader
*
* @apiParam (params) {string} md_id  	   URL Param: md_id to set
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/parameter/group/:group getParameterbygroup
* @apiVersion 0.1.0
* @apiName parameterbygroup
* @apiGroup Repository
*
* @apiDescription parameterbygroup
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/parameter/group/:group
*
* @apiUse DefaultHeader
*
* @apiParam (params) {string} group  		URL Param: group to set
*
* @apiUse DefaultJsonResponsesAndExample
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
* @apiUse DefaultHeader
*
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/ref getAllRef
* @apiVersion 0.1.0
* @apiName getAllRef
* @apiGroup Repository
*
* @apiDescription getAllRef
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/ref
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/toolframe getAllToolframe
* @apiVersion 0.1.0
* @apiName getAllToolframe
* @apiGroup Repository
*
* @apiDescription getAllToolframe
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/toolframe
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
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
* @apiUse DefaultHeader
*
* @apiParam (params) {string} tool_no  URL Param: tool no to set
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiParam (params) {string} pos_no   URL Param: pos num
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/userframe getAllUserframe
* @apiVersion 0.1.0
* @apiName getAllUserframe
* @apiGroup Repository
*
* @apiDescription getAllUserframe
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/userframe
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
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
* @apiUse DefaultHeader
*
* @apiParam (params) {string} userno   URL Param: user no
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {get} /v1/robot/repository/zeropoint getAllZeroPoint
* @apiVersion 0.1.0
* @apiName getAllZeroPoint
* @apiGroup Repository
*
* @apiDescription getAllZeroPoint
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/robot/repository/zeropoint
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
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
* @apiUse DefaultHeader
*
* @apiParam (params) {string} md_id    URL Param: metadata id
* @apiUse RequestBodyIndexAndValue
*
* @apiUse DefaultJsonResponsesAndExample
*/