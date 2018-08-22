/**
* @api {post} /v1/files/upload UploadFile
* @apiVersion 0.1.0
* @apiName UploadFile
* @apiGroup Files
*
* @apiDescription UploadFile
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.166:9000/v1/files/upload
*
* @apiHeader {string="multipart/form-data"} Content-Type=multipart/form-data
*
* @apiParam (Params) {string="fileupload"} 	fieldname  Body Param: field name for multipart/form-data
* @apiParam (Params) {string} 	filename  	Body Param: file name for multipart/form-data
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {get} /v1/files/jbi getJBIList
* @apiVersion 0.1.0
* @apiName getJBIList
* @apiGroup DBBackup
*
* @apiDescription getJBIList
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/files/jbi
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
*/