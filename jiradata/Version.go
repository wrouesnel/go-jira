package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir jiradata -pkg jiradata -overwrite schemas/Project.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// Version defined from schema:
// {
//   "title": "Version",
//   "type": "object",
//   "properties": {
//     "archived": {
//       "title": "archived",
//       "type": "boolean"
//     },
//     "description": {
//       "title": "description",
//       "type": "string"
//     },
//     "expand": {
//       "title": "expand",
//       "type": "string"
//     },
//     "id": {
//       "title": "id",
//       "type": "string"
//     },
//     "moveUnfixedIssuesTo": {
//       "title": "moveUnfixedIssuesTo",
//       "type": "string"
//     },
//     "name": {
//       "title": "name",
//       "type": "string"
//     },
//     "operations": {
//       "title": "operations",
//       "type": "array",
//       "items": {
//         "title": "Simple Link",
//         "type": "object",
//         "properties": {
//           "href": {
//             "title": "href",
//             "type": "string"
//           },
//           "iconClass": {
//             "title": "iconClass",
//             "type": "string"
//           },
//           "id": {
//             "title": "id",
//             "type": "string"
//           },
//           "label": {
//             "title": "label",
//             "type": "string"
//           },
//           "styleClass": {
//             "title": "styleClass",
//             "type": "string"
//           },
//           "title": {
//             "title": "title",
//             "type": "string"
//           },
//           "weight": {
//             "title": "weight",
//             "type": "integer"
//           }
//         }
//       }
//     },
//     "overdue": {
//       "title": "overdue",
//       "type": "boolean"
//     },
//     "project": {
//       "title": "project",
//       "type": "string"
//     },
//     "projectId": {
//       "title": "projectId",
//       "type": "integer"
//     },
//     "released": {
//       "title": "released",
//       "type": "boolean"
//     },
//     "remotelinks": {
//       "title": "remotelinks",
//       "type": "array",
//       "items": {
//         "title": "Remote Entity Link",
//         "type": "object",
//         "properties": {
//           "link": {
//             "title": "link"
//           },
//           "name": {
//             "title": "name",
//             "type": "string"
//           },
//           "self": {
//             "title": "self",
//             "type": "string"
//           }
//         }
//       }
//     },
//     "self": {
//       "title": "self",
//       "type": "string"
//     },
//     "userReleaseDate": {
//       "title": "userReleaseDate",
//       "type": "string"
//     },
//     "userStartDate": {
//       "title": "userStartDate",
//       "type": "string"
//     }
//   }
// }
type Version struct {
	Archived            bool        `json:"archived,omitempty" yaml:"archived,omitempty"`
	Description         string      `json:"description,omitempty" yaml:"description,omitempty"`
	Expand              string      `json:"expand,omitempty" yaml:"expand,omitempty"`
	ID                  string      `json:"id,omitempty" yaml:"id,omitempty"`
	MoveUnfixedIssuesTo string      `json:"moveUnfixedIssuesTo,omitempty" yaml:"moveUnfixedIssuesTo,omitempty"`
	Name                string      `json:"name,omitempty" yaml:"name,omitempty"`
	Operations          Operations  `json:"operations,omitempty" yaml:"operations,omitempty"`
	Overdue             bool        `json:"overdue,omitempty" yaml:"overdue,omitempty"`
	Project             string      `json:"project,omitempty" yaml:"project,omitempty"`
	ProjectID           int         `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	Released            bool        `json:"released,omitempty" yaml:"released,omitempty"`
	Remotelinks         Remotelinks `json:"remotelinks,omitempty" yaml:"remotelinks,omitempty"`
	Self                string      `json:"self,omitempty" yaml:"self,omitempty"`
	UserReleaseDate     string      `json:"userReleaseDate,omitempty" yaml:"userReleaseDate,omitempty"`
	UserStartDate       string      `json:"userStartDate,omitempty" yaml:"userStartDate,omitempty"`
}
