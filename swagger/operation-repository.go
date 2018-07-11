package swagger

// swagger:operation PUT /v2/robot/repository/arcparam/{file_no}/{md_id} setArcParam
// Origin: setArcParam [md_id] [value] [file_no] [index]
// This allows to set arcparam with some parameters
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - "name": "request"
//   "in": "body"
//	 "description": "request body"
//   "schema": {
//      "$ref": "#/parameters/Request"
//    }
// responses:
//   "default": {
//    "description": "default response for all HTTP codes",
//    "schema": {
//      "$ref": "#/responses/Response"
//    }
//   }
//

// swagger:operation PUT /v2/robot/repository/interference/{no}/{md_id} setInterference
// Origin: setInterference [md_id] [value] [no] [index]
// This allows to set interference with some parameters
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - "name": "request"
//   "in": "body"
//	 "description": "request body"
//   "schema": {
//      "$ref": "#/parameters/Request"
//    }
// responses:
//   "default": {
//    "description": "default response for all HTTP codes",
//    "schema": {
//      "$ref": "#/responses/Response"
//    }
//   }
//
// swagger:operation PUT /v2/robot/repository/params/{md_id} setParam
// Origin: setParam [md_id] [value] [index]
// This allows to set param with some parameters
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - "name": "request"
//   "in": "body"
//	 "description": "request body"
//   "schema": {
//      "$ref": "#/parameters/Request"
//    }
// responses:
//   "default": {
//    "description": "default response for all HTTP codes",
//    "schema": {
//      "$ref": "#/responses/Response"
//    }
//   }
//
// swagger:operation PUT /v2/robot/repository/toolframes/{md_id} setToolFrame
// Origin: setToolFrame [md_id] [value] [tool_no] [pos_no|index] [index]
// This allows to set toolframe with some parameters
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - "name": "request"
//   "in": "body"
//	 "description": "request body"
//   "schema": {
//      "$ref": "#/parameters/Request"
//    }
// responses:
//   "default": {
//    "description": "default response for all HTTP codes",
//    "schema": {
//      "$ref": "#/responses/Response"
//    }
//   }
//
// swagger:operation PUT /v2/robot/repository/userframe/{userno}/{md_id} setUserFrame
// Origin: setUserFrame [md_id] [value] [userNo]
// This allows to set userframe with some parameters
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - "name": "request"
//   "in": "body"
//	 "description": "request body"
//   "schema": {
//      "$ref": "#/parameters/Request"
//    }
// responses:
//   "default": {
//    "description": "default response for all HTTP codes",
//    "schema": {
//      "$ref": "#/responses/Response"
//    }
//   }
//
// swagger:operation PUT /v2/robot/repository/zeropoint/{md_id} setZeroPoint
// Origin: setZeroPoint [md_id] [value] [index]
// This allows to set zeropoint with some parameters
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - "name": "request"
//   "in": "body"
//	 "description": "request body"
//   "schema": {
//      "$ref": "#/parameters/Request"
//    }
// responses:
//   "default": {
//    "description": "default response for all HTTP codes",
//    "schema": {
//      "$ref": "#/responses/Response"
//    }
//   }
//