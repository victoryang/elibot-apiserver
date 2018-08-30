/**
* @api {post} /v1/files/upload uploadFile
* @apiVersion 0.1.0
* @apiName uploadFile
* @apiGroup Files
*
* @apiDescription uploadFile
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/files/upload
*
* @apiHeader {string="multipart/form-data"} 	Content-Type		Mutlpart/form-data file upload content-type
*
* @apiParam (Params) {string="fileupload"}      form_file_key  		Multipart/form-data File Field Key: for server lookup
*
* @apiUse SuccessStringResponse
* @apiUse InternalServerErrorJsonResponse
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
* @apiUse SuccessJsonResponse
* @apiUse InternalServerErrorJsonResponse
*/

/**
* @api {get} /v1/files/jbi/:filename downloadJBIFile
* @apiVersion 0.1.0
* @apiName downloadJBIFile
* @apiGroup Files
*
* @apiDescription downloadJBIFile
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/files/jbi/:filename
*
* @apiUse DefaultHeader
*
* @apiParam (Params) {string="*.jbi"} 	filename  	URL Param: file to get
*
* @apiUse SuccessJsonResponse
* @apiUse InternalServerErrorJsonResponse
*/