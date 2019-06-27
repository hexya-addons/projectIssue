package project_issue

	import (
		"net/http"

		"github.com/hexya-erp/hexya/src/controllers"
		"github.com/hexya-erp/hexya/src/models"
		"github.com/hexya-erp/hexya/src/models/types"
		"github.com/hexya-erp/hexya/src/models/types/dates"
		"github.com/hexya-erp/pool/h"
		"github.com/hexya-erp/pool/q"
	)
	
func init() {
h.().DeclareModel()

h.().Methods().GetMailTemplateIdDomain().DeclareMethod(
`GetMailTemplateIdDomain`,
func(rs m.Set)  {
//        domain = super(ProjectStage, self)._get_mail_template_id_domain()
//        return ['|'] + domain + [('model', '=', 'project.issue')]
})
}