/**
* @api {post} /v1/db/backup DBBackupDB
* @apiVersion 0.1.0
* @apiName DBBackupDB
* @apiGroup DBBackup
*
* @apiDescription DBBackupDB
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/db/backup
*
* @apiUse DefaultHeader
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {get} /v1/db/backup DBListBackups
* @apiVersion 0.1.0
* @apiName DBListBackups
* @apiGroup DBBackup
*
* @apiDescription DBListBackups
*
* @apiExample Example usage:
* curl -X GET http://192.168.1.253:9000/v1/db/backup
*
* @apiUse DefaultHeader
*
* @apiUse DefaultJsonResponsesAndExample
*/

/**
* @api {delete} /v1/db/backup/:name DBDelBackup
* @apiVersion 0.1.0
* @apiName DBDelBackup
* @apiGroup DBBackup
*
* @apiDescription DBDelBackup
*
* @apiExample Example usage:
* curl -X DELETE http://192.168.1.253:9000/v1/db/backup/:name
*
* @apiUse DefaultHeader
* 
* @apiParam (params) {string} name  	   URL Param: db name to delete
*
* @apiUse DefaultResponsesAndExample
*/

/**
* @api {post} /v1/db/backup/:name/restore DBRestoreBackup
* @apiVersion 0.1.0
* @apiName DBRestoreBackup
* @apiGroup DBBackup
*
* @apiDescription DBRestoreBackup
*
* @apiExample Example usage:
* curl -X POST http://192.168.1.253:9000/v1/db/backup/:name/restore
*
* @apiUse DefaultHeader
* 
* @apiParam (params) {string} name  	   URL Param: db name to restore
*
* @apiUse DefaultResponsesAndExample
*/