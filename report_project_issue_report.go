package project_issue

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.ProjectIssueReport().DeclareModel()

	h.ProjectIssueReport().AddFields(map[string]models.FieldDefinition{
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
			ReadOnly:      true,
		},
		"OpeningDate": models.DateTimeField{
			String:   "Date of Opening",
			ReadOnly: true,
		},
		"DateClosed": models.DateTimeField{
			String:   "Date of Closing",
			ReadOnly: true,
		},
		"DateLastStageUpdate": models.DateTimeField{
			String:   "Last Stage Update",
			ReadOnly: true,
		},
		"StageId": models.Many2OneField{
			RelationModel: h.ProjectTaskType(),
			String:        "Stage",
		},
		"NbrIssues": models.IntegerField{
			String:   "# of Issues",
			ReadOnly: true,
		},
		"WorkingHoursOpen": models.FloatField{
			String:   "Avg. Working Hours to Open",
			ReadOnly: true,
			//group_operator="avg"
		},
		"WorkingHoursClose": models.FloatField{
			String:   "Avg. Working Hours to Close",
			ReadOnly: true,
			//group_operator="avg"
		},
		"DelayOpen": models.FloatField{
			String: "Avg. Delay to Open",
			//digits=(16, 2)
			ReadOnly: true,
			//group_operator="avg"
			Help: "Number of Days to open the project issue.",
		},
		"DelayClose": models.FloatField{
			String: "Avg. Delay to Close",
			//digits=(16, 2)
			ReadOnly: true,
			//group_operator="avg"
			Help: "Number of Days to close the project issue",
		},
		"Priority": models.SelectionField{
			Selection: types.Selection{
				"0": "Low",
				"1": "Normal",
				"2": "High",
			},
		},
		"ProjectId": models.Many2OneField{
			RelationModel: h.ProjectProject(),
			String:        "Project",
			ReadOnly:      true,
		},
		"UserId": models.Many2OneField{
			RelationModel: h.User(),
			String:        "Assigned to",
			ReadOnly:      true,
		},
		"PartnerId": models.Many2OneField{
			RelationModel: h.Partner(),
			String:        "Contact",
		},
		"Email": models.IntegerField{
			String:   "# Emails",
			ReadOnly: true,
		},
	})
	h.ProjectIssueReport().Fields().CreateDate().setString("Create Date")
	h.ProjectIssueReport().Fields().CreateDate().setReadOnly(true)
	h.ProjectIssueReport().Methods().Init().DeclareMethod(
		`Init`,
		func(rs m.ProjectIssueReportSet) {
			//        tools.drop_view_if_exists(self._cr, 'project_issue_report')
			//        self._cr.execute("""
			//            CREATE OR REPLACE VIEW project_issue_report AS (
			//                SELECT
			//                    c.id as id,
			//                    c.date_open as opening_date,
			//                    c.create_date as create_date,
			//                    c.date_last_stage_update as date_last_stage_update,
			//                    c.user_id,
			//                    c.working_hours_open,
			//                    c.working_hours_close,
			//                    c.stage_id,
			//                    c.date_closed as date_closed,
			//                    c.company_id as company_id,
			//                    c.priority as priority,
			//                    c.project_id as project_id,
			//                    1 as nbr_issues,
			//                    c.partner_id,
			//                    c.day_open as delay_open,
			//                    c.day_close as delay_close,
			//                    (SELECT count(id) FROM mail_message WHERE model='project.issue' AND message_type IN ('email', 'comment') AND res_id=c.id) AS email
			//
			//                FROM
			//                    project_issue c
			//                LEFT JOIN project_task t on c.task_id = t.id
			//                WHERE c.active= 'true'
			//            )""")
		})
}
