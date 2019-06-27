package project_issue

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.ProjectConfigSettings().DeclareModel()

	h.ProjectConfigSettings().AddFields(map[string]models.FieldDefinition{
		"ModuleProjectIssueSheet": models.SelectionField{
			Selection: types.Selection{
				"": "Do not track working hours on issues",
				"": "Activate timesheets on issues",
			},
			String: "Timesheets on Issues",
			Help: "Provides timesheet support for the issues/bugs management in project." +
				"-This installs the module project_issue_sheet.",
		},
		"ModuleRatingProjectIssue": models.SelectionField{
			Selection: types.Selection{
				"": "No customer rating",
				"": "Track customer satisfaction on issues",
			},
			String: "Rating on issue",
			Help:   "This allows customers to give rating on issue",
		},
	})
}
