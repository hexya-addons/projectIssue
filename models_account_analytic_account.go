package project_issue

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.AccountAnalyticAccount().DeclareModel()

	h.AccountAnalyticAccount().AddFields(map[string]models.FieldDefinition{
		"UseIssues": models.BooleanField{
			String: "Use Issues",
			Help:   "Check this box to manage customer activities through this project",
		},
	})
	h.AccountAnalyticAccount().Methods().Unlink().Extend(
		`Unlink`,
		func(rs m.AccountAnalyticAccountSet) {
			//        if self.env['project.issue'].sudo().search_count([('project_id.analytic_account_id', 'in', self.ids)]):
			//            raise UserError(
			//                _('Please remove existing issues in the project linked to the accounts you want to delete.'))
			//        return super(AccountAnalyticAccount, self).unlink()
		})
	h.AccountAnalyticAccount().Methods().TriggerProjectCreation().DeclareMethod(
		`TriggerProjectCreation`,
		func(rs m.AccountAnalyticAccountSet, vals interface{}) {
			//        res = super(AccountAnalyticAccount,
			//                    self)._trigger_project_creation(vals)
			//        return res or (vals.get('use_issues') and not 'project_creation_in_progress' in self.env.context)
		})
}
