package project_issue

import (
	"github.com/hexya-addons/base"
	"github.com/hexya-erp/pool/h"
)

//vars

var (
)


//rights
func init() {
	h.ProjectIssue().Methods().AllowAllToGroup(GroupProjectUser)
	h.ProjectIssueReport().Methods().AllowAllToGroup(GroupProjectManager)
	h.Resource.ModelResourceCalendar().Methods().AllowAllToGroup(GroupProjectManager)
	h.ProjectIssueReport().Methods().Load().AllowGroup(GroupProjectUser)
	h.ProjectIssue().Methods().Load().AllowGroup(base.GroupPortal)
}
