package project_issue

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

func init() {
	h.ProjectProject().DeclareModel()

	h.ProjectProject().AddFields(map[string]models.FieldDefinition{
		"IssueCount": models.IntegerField{
			Compute: h.ProjectProject().Methods().ComputeIssueCount(),
			String:  "Issues",
		},
		"IssueIds": models.One2ManyField{
			RelationModel: h.ProjectIssue(),
			ReverseFK:     "",
			String:        "Issues",
			Filter:        q.StageId().Fold().Equals(False).Or().StageId().Equals(False),
		},
		"LabelIssues": models.CharField{
			String:  "Use Issues as",
			Help:    "Customize the issues label, for example to call them cases.",
			Default: models.DefaultValue("Issues"),
		},
		"UseIssues": models.BooleanField{
			Related: `AnalyticAccountId.UseIssues`,
			Default: models.DefaultValue(true),
		},
		"IssueNeedactionCount": models.IntegerField{
			Compute: h.ProjectProject().Methods().IssueNeedactionCount(),
			String:  "Issues",
		},
	})
	h.ProjectProject().Methods().GetAliasModels().DeclareMethod(
		`GetAliasModels`,
		func(rs m.ProjectProjectSet) {
			//        res = super(Project, self)._get_alias_models()
			//        res.append(("project.issue", "Issues"))
			//        return res
		})
	h.ProjectProject().Methods().ComputeIssueCount().DeclareMethod(
		`ComputeIssueCount`,
		func(rs h.ProjectProjectSet) h.ProjectProjectData {
			//        for project in self:
			//            project.issue_count = self.env['project.issue'].search_count(
			//                [('project_id', '=', project.id), '|', ('stage_id.fold', '=', False), ('stage_id', '=', False)])
		})
	h.ProjectProject().Methods().IssueNeedactionCount().DeclareMethod(
		`IssueNeedactionCount`,
		func(rs h.ProjectProjectSet) h.ProjectProjectData {
			//        issue_data = self.env['project.issue'].read_group(
			//            [('project_id', 'in', self.ids), ('message_needaction', '=', True)], ['project_id'], ['project_id'])
			//        result = dict((data['project_id'][0], data['project_id_count'])
			//                      for data in issue_data)
			//        for project in self:
			//            project.issue_needaction_count = int(result.get(project.id, 0))
		})
	h.ProjectProject().Methods().OnChangeUseTasksOrIssues().DeclareMethod(
		`OnChangeUseTasksOrIssues`,
		func(rs m.ProjectProjectSet) {
			//        if self.use_tasks and not self.use_issues:
			//            self.alias_model = 'project.task'
			//        elif not self.use_tasks and self.use_issues:
			//            self.alias_model = 'project.issue'
		})
	h.ProjectProject().Methods().Write().Extend(
		`Write`,
		func(rs m.ProjectProjectSet, vals models.RecordData) {
			//        res = super(Project, self).write(vals)
			//        if 'active' in vals:
			//            # archiving/unarchiving a project does it on its issues, too
			//            issues = self.with_context(active_test=False).mapped('issue_ids')
			//            issues.write({'active': vals['active']})
			//        return res
		})
}
