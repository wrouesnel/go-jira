package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir jiradata -pkg jiradata -overwrite schemas/AuthParams.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// AuthParams defined from schema:
// {
//   "title": "Auth Params",
//   "id": "https://docs.atlassian.com/jira/REST/schema/auth-params#",
//   "type": "object",
//   "properties": {
//     "password": {
//       "title": "password",
//       "type": "string"
//     },
//     "username": {
//       "title": "username",
//       "type": "string"
//     }
//   }
// }
type AuthParams struct {
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
