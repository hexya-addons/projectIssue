package project_issue

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/hexya/src/models/types/dates"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

func init() {
	h.ProjectIssue().DeclareModel()

	//    _mail_post_access = 'read'
	h.ProjectIssue().Methods().GetDefaultStageId().DeclareMethod(
		`GetDefaultStageId`,
		func(rs m.ProjectIssueSet) {
			//        project_id = self.env.context.get('default_project_id')
			//        if not project_id:
			//            return False
			//        return self.stage_find(project_id, [('fold', '=', False)])
		})
	h.ProjectIssue().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:   "Issue",
			Required: true,
		},
		"Active": models.BooleanField{
			Default: models.DefaultValue(true),
		},
		"DaysSinceCreation": models.IntegerField{
			Compute: h.ProjectIssue().Methods().ComputeInactivityDays(),
			String:  "Days since creation date",
			Help:    "Difference in days between creation date and current date",
		},
		"DateDeadline": models.DateField{
			String: "Deadline",
		},
		"PartnerId": models.Many2OneField{
			RelationModel: h.Partner(),
			String:        "Contact",
			Index:         true,
		},
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
			Default:       func(env models.Environment) interface{} { return env.Uid().company_id },
		},
		"Description": models.TextField{
			String: "Private Note",
		},
		"KanbanState": models.SelectionField{
			Selection: types.Selection{
				"normal":  "Normal",
				"blocked": "Blocked",
				"done":    "Ready for next stage",
			},
			String: "Kanban State",
			//track_visibility='onchange'
			Required: true,
			Default:  models.DefaultValue("normal"),
			Help: "An Issue's kanban state indicates special situations affecting it:" +
				"" +
				"                                           * Normal is" +
				"the default situation" +
				"" +
				"                                           * Blocked indicates" +
				"something is preventing the progress of this issue" +
				"" +
				"                                           * Ready for" +
				"next stage indicates the issue is ready to be pulled to the next stage",
		},
		"EmailFrom": models.CharField{
			String: "Email",
			Help:   "These people will receive email.",
			Index:  true,
		},
		"EmailCc": models.CharField{
			String: "Watchers Emails",
			Help: "These email addresses will be added to the CC field of all inbound" +
				"        and outbound emails for this record before being" +
				"sent. Separate multiple email addresses with a comma",
		},
		"DateOpen": models.DateTimeField{
			String:   "Assigned",
			ReadOnly: true,
			Index:    true,
		},
		"DateClosed": models.DateTimeField{
			String:   "Closed",
			ReadOnly: true,
			Index:    true,
		},
		"Date": models.DateTimeField{
			String: "Date",
		},
		"DateLastStageUpdate": models.DateTimeField{
			String:  "Last Stage Update",
			Index:   true,
			Default: func(env models.Environment) interface{} { return dates.Now() },
		},
		"Channel": models.CharField{
			String: "Channel",
			Help:   "Communication channel.",
		},
		"TagIds": models.Many2ManyField{
			RelationModel: h.ProjectTags(),
			String:        "Tags",
		},
		"Priority": models.SelectionField{
			Selection: types.Selection{
				"0": "Low",
				"1": "Normal",
				"2": "High",
			},
			String:  "Priority",
			Index:   true,
			Default: models.DefaultValue("0"),
		},
		"StageId": models.Many2OneField{
			RelationModel: h.ProjectTaskType(),
			String:        "Stage",
			//track_visibility='onchange'
			Index:  true,
			Filter: q.ProjectIds().Equals(project_id),
			NoCopy: true,
			//group_expand='_read_group_stage_ids'
			Default: models.DefaultValue(_get_default_stage_id),
		},
		"ProjectId": models.Many2OneField{
			RelationModel: h.ProjectProject(),
			String:        "Project",
			//track_visibility='onchange'
			Index: true,
		},
		"Duration": models.FloatField{
			String: "Duration",
		},
		"TaskId": models.Many2OneField{
			RelationModel: h.ProjectTask(),
			String:        "Task",
			Filter:        q.ProjectId().Equals(project_id),
			Help: "You can link this issue to an existing task or directly" +
				"create a new one from here",
		},
		"DayOpen": models.FloatField{
			Compute: h.ProjectIssue().Methods().ComputeDay(),
			String:  "Days to Assign",
			Stored:  true,
		},
		"DayClose": models.FloatField{
			Compute: h.ProjectIssue().Methods().ComputeDay(),
			String:  "Days to Close",
			Stored:  true,
		},
		"UserId": models.Many2OneField{
			RelationModel: h.User(),
			String:        "Assigned to",
			Index:         true,
			//track_visibility='onchange'
			Default: func(env models.Environment) interface{} { return env.uid },
		},
		"WorkingHoursOpen": models.FloatField{
			Compute: h.ProjectIssue().Methods().ComputeDay(),
			String:  "Working Hours to assign the Issue",
			Stored:  true,
		},
		"WorkingHoursClose": models.FloatField{
			Compute: h.ProjectIssue().Methods().ComputeDay(),
			String:  "Working Hours to close the Issue",
			Stored:  true,
		},
		"InactivityDays": models.IntegerField{
			Compute: h.ProjectIssue().Methods().ComputeInactivityDays(),
			String:  "Days since last action",
			Help:    "Difference in days between last action and current date",
		},
		"Color": models.IntegerField{
			String: "Color Index",
		},
		"UserEmail": models.CharField{
			Related:  `UserId.Email`,
			String:   "User Email",
			ReadOnly: true,
		},
		"DateActionLast": models.DateTimeField{
			String:   "Last Action",
			ReadOnly: true,
		},
		"DateActionNext": models.DateTimeField{
			String:   "Next Action",
			ReadOnly: true,
		},
		"LegendBlocked": models.CharField{
			Related:  `StageId.LegendBlocked`,
			String:   "Kanban Blocked Explanation",
			ReadOnly: true,
		},
		"LegendDone": models.CharField{
			Related:  `StageId.LegendDone`,
			String:   "Kanban Valid Explanation",
			ReadOnly: true,
		},
		"LegendNormal": models.CharField{
			Related:  `StageId.LegendNormal`,
			String:   "Kanban Ongoing Explanation",
			ReadOnly: true,
		},
	})
	h.ProjectIssue().Methods().ReadGroupStageIds().DeclareMethod(
		`ReadGroupStageIds`,
		func(rs m.ProjectIssueSet, stages interface{}, domain interface{}, order interface{}) {
			//        search_domain = [('id', 'in', stages.ids)]
			//        if 'default_project_id' in self.env.context:
			//            search_domain = [
			//                '|', ('project_ids', '=', self.env.context['default_project_id'])] + search_domain
			//        return stages.search(search_domain, order=order)
		})
	h.ProjectIssue().Methods().ComputeDay().DeclareMethod(
		`ComputeDay`,
		func(rs h.ProjectIssueSet) h.ProjectIssueData {
			//        for issue in self:
			//            # if the working hours on the project are not defined, use default ones (8 -> 12 and 13 -> 17 * 5)
			//            calendar = issue.project_id.resource_calendar_id
			//
			//            dt_create_date = fields.Datetime.from_string(issue.create_date)
			//            if issue.date_open:
			//                dt_date_open = fields.Datetime.from_string(issue.date_open)
			//                issue.day_open = (
			//                    dt_date_open - dt_create_date).total_seconds() / (24.0 * 3600)
			//                issue.working_hours_open = calendar.get_working_hours(dt_create_date, dt_date_open,
			//                                                                      compute_leaves=True, resource_id=False, default_interval=(8, 16))
			//
			//            if issue.date_closed:
			//                dt_date_closed = fields.Datetime.from_string(issue.date_closed)
			//                issue.day_close = (
			//                    dt_date_closed - dt_create_date).total_seconds() / (24.0 * 3600)
			//                issue.working_hours_close = calendar.get_working_hours(dt_create_date, dt_date_closed,
			//                                                                       compute_leaves=True, resource_id=False, default_interval=(8, 16))
		})
	h.ProjectIssue().Methods().ComputeInactivityDays().DeclareMethod(
		`ComputeInactivityDays`,
		func(rs h.ProjectIssueSet) h.ProjectIssueData {
			//        current_datetime = fields.Datetime.from_string(fields.Datetime.now())
			//        for issue in self:
			//            dt_create_date = fields.Datetime.from_string(
			//                issue.create_date) or current_datetime
			//            issue.days_since_creation = (
			//                current_datetime - dt_create_date).days
			//
			//            if issue.date_action_last:
			//                issue.inactivity_days = (
			//                    current_datetime - fields.Datetime.from_string(issue.date_action_last)).days
			//            elif issue.date_last_stage_update:
			//                issue.inactivity_days = (
			//                    current_datetime - fields.Datetime.from_string(issue.date_last_stage_update)).days
			//            else:
			//                issue.inactivity_days = (
			//                    current_datetime - dt_create_date).days
		})
	h.ProjectIssue().Methods().OnchangePartnerId().DeclareMethod(
		` This function sets partner email address based on partner
        `,
		func(rs m.ProjectIssueSet) {
			//        self.email_from = self.partner_id.email
		})
	h.ProjectIssue().Methods().OnchangeProjectId().DeclareMethod(
		`OnchangeProjectId`,
		func(rs m.ProjectIssueSet) {
			//        if self.project_id:
			//            if not self.partner_id and not self.email_from:
			//                self.partner_id = self.project_id.partner_id.id
			//                self.email_from = self.project_id.partner_id.email
			//            self.stage_id = self.stage_find(
			//                self.project_id.id, [('fold', '=', False)])
			//        else:
			//            self.partner_id = False
			//            self.email_from = False
			//            self.stage_id = False
		})
	h.ProjectIssue().Methods().OnchangeTaskId().DeclareMethod(
		`OnchangeTaskId`,
		func(rs m.ProjectIssueSet) {
			//        self.user_id = self.task_id.user_id
		})
	h.ProjectIssue().Methods().Copy().Extend(
		`Copy`,
		func(rs m.ProjectIssueSet, defaultName models.RecordData) {
			//        if default is None:
			//            default = {}
			//        default.update(name=_('%s (copy)') % (self.name))
			//        return super(ProjectIssue, self).copy(default=default)
		})
	h.ProjectIssue().Methods().Create().Extend(
		`Create`,
		func(rs m.ProjectIssueSet, vals models.RecordData) {
			//        context = dict(self.env.context)
			//        if vals.get('project_id') and not self.env.context.get('default_project_id'):
			//            context['default_project_id'] = vals.get('project_id')
			//        if vals.get('user_id') and not vals.get('date_open'):
			//            vals['date_open'] = fields.Datetime.now()
			//        if 'stage_id' in vals:
			//            vals.update(self.update_date_closed(vals['stage_id']))
			//        context['mail_create_nolog'] = True
			//        return super(ProjectIssue, self.with_context(context)).create(vals)
		})
	h.ProjectIssue().Methods().Write().Extend(
		`Write`,
		func(rs m.ProjectIssueSet, vals models.RecordData) {
			//        if 'stage_id' in vals:
			//            vals.update(self.update_date_closed(vals['stage_id']))
			//            vals['date_last_stage_update'] = fields.Datetime.now()
			//            if 'kanban_state' not in vals:
			//                vals['kanban_state'] = 'normal'
			//        if vals.get('user_id') and 'date_open' not in vals:
			//            vals['date_open'] = fields.Datetime.now()
			//        return super(ProjectIssue, self).write(vals)
		})
	h.ProjectIssue().Methods().GetEmptyListHelp().DeclareMethod(
		`GetEmptyListHelp`,
		func(rs m.ProjectIssueSet, help interface{}) {
			//        return super(ProjectIssue, self.with_context(
			//            empty_list_help_model='project.project',
			//            empty_list_help_id=self.env.context.get('default_project_id'),
			//            empty_list_help_document_name=_("issues")
			//        )).get_empty_list_help(help)
		})
	h.ProjectIssue().Methods().UpdateDateClosed().DeclareMethod(
		`UpdateDateClosed`,
		func(rs m.ProjectIssueSet, stage_id interface{}) {
			//        project_task_type = self.env['project.task.type'].browse(stage_id)
			//        if project_task_type.fold:
			//            return {'date_closed': fields.Datetime.now()}
			//        return {'date_closed': False}
		})
	h.ProjectIssue().Methods().StageFind().DeclareMethod(
		` Override of the base.stage method
            Parameter of the stage search taken from the issue:
            - project_id: if set, stages must belong to this project or
              be a default case
        `,
		func(rs m.ProjectIssueSet, project_id interface{}, domain interface{}, order interface{}) {
			//        search_domain = list(domain) if domain else []
			//        if project_id:
			//            search_domain += [('project_ids', '=', project_id)]
			//        project_task_type = self.env['project.task.type'].search(
			//            search_domain, order=order, limit=1)
			//        return project_task_type.id
		})
	h.ProjectIssue().Methods().TrackTemplate().DeclareMethod(
		`TrackTemplate`,
		func(rs m.ProjectIssueSet, tracking interface{}) {
			//        res = super(ProjectIssue, self)._track_template(tracking)
			//        test_issue = self[0]
			//        changes, tracking_value_ids = tracking[test_issue.id]
			//        if 'stage_id' in changes and test_issue.stage_id.mail_template_id:
			//            res['stage_id'] = (test_issue.stage_id.mail_template_id, {
			//                               'composition_mode': 'mass_mail'})
			//        return res
		})
	h.ProjectIssue().Methods().TrackSubtype().DeclareMethod(
		`TrackSubtype`,
		func(rs m.ProjectIssueSet, init_values interface{}) {
			//        self.ensure_one()
			//        if 'kanban_state' in init_values and self.kanban_state == 'blocked':
			//            return 'project_issue.mt_issue_blocked'
			//        elif 'kanban_state' in init_values and self.kanban_state == 'done':
			//            return 'project_issue.mt_issue_ready'
			//        elif 'user_id' in init_values and self.user_id:  # assigned -> new
			//            return 'project_issue.mt_issue_new'
			//        elif 'stage_id' in init_values and self.stage_id and self.stage_id.sequence <= 1:  # start stage -> new
			//            return 'project_issue.mt_issue_new'
			//        elif 'stage_id' in init_values:
			//            return 'project_issue.mt_issue_stage'
			//        return super(ProjectIssue, self)._track_subtype(init_values)
		})
	h.ProjectIssue().Methods().NotificationRecipients().DeclareMethod(
		`
        `,
		func(rs m.ProjectIssueSet, message interface{}, groups interface{}) {
			//        groups = super(ProjectIssue, self)._notification_recipients(
			//            message, groups)
			//        self.ensure_one()
			//        if not self.user_id:
			//            take_action = self._notification_link_helper('assign')
			//            project_actions = [{'url': take_action, 'title': _('I take it')}]
			//        else:
			//            new_action_id = self.env.ref(
			//                'project_issue.project_issue_categ_act0').id
			//            new_action = self._notification_link_helper(
			//                'new', action_id=new_action_id)
			//            project_actions = [{'url': new_action, 'title': _('New Issue')}]
			//        new_group = (
			//            'group_project_user', lambda partner: bool(partner.user_ids) and any(user.has_group('project.group_project_user') for user in partner.user_ids), {
			//                'actions': project_actions,
			//            })
			//        return [new_group] + groups
		})
	h.ProjectIssue().Methods().MessageGetReplyTo().DeclareMethod(
		` Override to get the reply_to of the parent project. `,
		func(rs m.ProjectIssueSet, res_ids interface{}, defaultName interface{}) {
			//        issues = self.browse(res_ids)
			//        project_ids = set(issues.mapped('project_id').ids)
			//        aliases = self.env['project.project'].message_get_reply_to(
			//            list(project_ids), default=default)
			//        return dict((issue.id, aliases.get(issue.project_id and issue.project_id.id or 0, False)) for issue in issues)
		})
	h.ProjectIssue().Methods().MessageGetSuggestedRecipients().DeclareMethod(
		`MessageGetSuggestedRecipients`,
		func(rs m.ProjectIssueSet) {
			//        recipients = super(
			//            ProjectIssue, self).message_get_suggested_recipients()
			//        try:
			//            for issue in self:
			//                if issue.partner_id:
			//                    issue._message_add_suggested_recipient(
			//                        recipients, partner=issue.partner_id, reason=_('Customer'))
			//                elif issue.email_from:
			//                    issue._message_add_suggested_recipient(
			//                        recipients, email=issue.email_from, reason=_('Customer Email'))
			//        except AccessError:  # no read access rights -> just ignore suggested recipients because this imply modifying followers
			//            pass
			//        return recipients
		})
	h.ProjectIssue().Methods().EmailSplit().DeclareMethod(
		`EmailSplit`,
		func(rs m.ProjectIssueSet, msg interface{}) {
			//        email_list = tools.email_split(
			//            (msg.get('to') or '') + ',' + (msg.get('cc') or ''))
			//        return filter(lambda x: x.split('@')[0] not in self.mapped('project_id.alias_name'), email_list)
		})
	h.ProjectIssue().Methods().MessageNew().DeclareMethod(
		` Overrides mail_thread message_new that is called by the mailgateway
            through message_process.
            This override updates the document according to the email.
        `,
		func(rs m.ProjectIssueSet, msg interface{}, custom_values interface{}) {
			//        create_context = dict(self.env.context or {})
			//        create_context['default_user_id'] = False
			//        defaults = {
			//            'name':  msg.get('subject') or _("No Subject"),
			//            'email_from': msg.get('from'),
			//            'email_cc': msg.get('cc'),
			//            'partner_id': msg.get('author_id', False),
			//        }
			//        if custom_values:
			//            defaults.update(custom_values)
			//        res_id = super(ProjectIssue, self.with_context(
			//            create_context)).message_new(msg, custom_values=defaults)
			//        issue = self.browse(res_id)
			//        email_list = issue.email_split(msg)
			//        partner_ids = filter(None, issue._find_partner_from_emails(email_list))
			//        issue.message_subscribe(partner_ids)
			//        return res_id
		})
	h.ProjectIssue().Methods().MessageUpdate().DeclareMethod(
		` Override to update the issue according to the email. `,
		func(rs m.ProjectIssueSet, msg interface{}, update_vals interface{}) {
			//        email_list = self.email_split(msg)
			//        partner_ids = filter(None, self._find_partner_from_emails(email_list))
			//        self.message_subscribe(partner_ids)
			//        return super(ProjectIssue, self).message_update(msg, update_vals=update_vals)
		})
	h.ProjectIssue().Methods().MessagePost().DeclareMethod(
		` Overrides mail_thread message_post so that we can set
the date of last action field when
            a new message is posted on the issue.
        `,
		func(rs m.ProjectIssueSet, subtype interface{}) {
			//        self.ensure_one()
			//        mail_message = super(ProjectIssue, self).message_post(
			//            subtype=subtype, **kwargs)
			//        if subtype:
			//            self.sudo().write({'date_action_last': fields.Datetime.now()})
			//        return mail_message
		})
	h.ProjectIssue().Methods().MessageGetEmailValues().DeclareMethod(
		`MessageGetEmailValues`,
		func(rs m.ProjectIssueSet, notif_mail interface{}) {
			//        self.ensure_one()
			//        res = super(ProjectIssue, self).message_get_email_values(
			//            notif_mail=notif_mail)
			//        headers = {}
			//        if res.get('headers'):
			//            try:
			//                headers.update(safe_eval(res['headers']))
			//            except Exception:
			//                pass
			//        if self.project_id:
			//            current_objects = filter(None, headers.get(
			//                'X-Odoo-Objects', '').split(','))
			//            current_objects.insert(
			//                0, 'project.project-%s, ' % self.project_id.id)
			//            headers['X-Odoo-Objects'] = ','.join(current_objects)
			//        if self.tag_ids:
			//            headers['X-Odoo-Tags'] = ','.join(self.tag_ids.mapped('name'))
			//        res['headers'] = repr(headers)
			//        return res
		})
}
