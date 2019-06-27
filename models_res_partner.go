package project_issue

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {

	h.Partner().AddFields(map[string]models.FieldDefinition{
		"IssueCount": models.IntegerField{
			Compute: h.Partner().Methods().ComputeIssueCount(),
			String:  "# Issues",
		},
	})
	h.Partner().Methods().ComputeIssueCount().DeclareMethod(
		`ComputeIssueCount`,
		func(rs h.PartnerSet) h.PartnerData {
			//        Issue = self.env['project.issue']
			//        for partner in self:
			//            partner.issue_count = Issue.search_count(
			//                [('partner_id', 'child_of', partner.commercial_partner_id.id)])
		})
}
