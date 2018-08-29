/**
* @api {post} /v1/files/upload UploadFile
* @apiVersion 0.1.0
* @apiName UploadFile
* @apiGroup Files
*
* @apiDescription UploadFile
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/files/upload
*
* @apiHeader {string="multipart/form-data"} 	Content-Type		Mutlpart/form-data file upload content-type
*
* @apiParam (Params) {string="fileupload"}      form_file_key  		Multipart/form-data File Field Key: for server lookup
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {get} /v1/files/jbi getJBIList
* @apiVersion 0.1.0
* @apiName getJBIList
* @apiGroup Files
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