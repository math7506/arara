package main

import (
        "github.com/common-nighthawk/go-figure"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)


var apiPatterns = []string{
	// Plataformas corporativas
	"https://%s.my.salesforce.com/services/data/",
	"https://%s.my.salesforce.com/services/data/v63.0/query/",
	"https://%s.my.salesforce.com/services/data/v63.0/sobjects/",
	"https://%s.my.salesforce.com/services/oauth2/userinfo",
	"https://%s.my.salesforce.com/services/data/v63.0/chatter/users",
	"https://%s.my.salesforce.com/services/data/v63.0/tooling/query/",
	"https://%s.my.salesforce.com/services/data/v63.0/limits/",
	"https://%s.my.salesforce.com/services/data/v63.0/actions/",
        "https://%s.my.salesforce.com/services/async/63.0/job/metadata",
	"https://%s.my.salesforce.com/services/async/63.0/job/",
	"https://%s.my.salesforce.com/services/data/v63.0/tooling/metadata/",
	"https://%s.my.salesforce.com/services/data/v63.0/tooling/sobjects/CustomObject/",
        "https://%s.my.salesforce.com/services/data/v63.0/analytics/reports/",
	"https://%s.my.salesforce.com/services/data/v63.0/analytics/datasets/",
	"https://%s.my.salesforce.com/services/data/v63.0/analytics/query/",
	"https://%s.my.salesforce.com/services/data/v63.0/analytics/wave/",
        "https://%s.my.salesforce.com/services/data/v63.0/sobjects/ContentVersion/",
	"https://%s.my.salesforce.com/services/data/v63.0/sobjects/ContentDocumentLink/",
	"https://%s.my.salesforce.com/services/data/v63.0/sobjects/ContentDocument/",
        "https://%s.my.salesforce.com/services/data/v63.0/functions/",
        "https://%s.my.salesforce.com/services/data/v63.0/limits/",
        "https://%s.my.salesforce.com/services/data/v63.0/sobjects/Organization/",
        "https://%s.my.salesforce.com/services/data/v63.0/tooling/sobjects/ApexClass/",
        "https://%s.my.salesforce.com/services/data/v63.0/tooling/sobjects/ApexTrigger/",
        "https://%s.my.salesforce.com/services/data/v63.0/tooling/sobjects/CustomObject/",

	// ServiceNow
	"https://%s.service-now.com/api/now/table/incident",
	"https://%s.service-now.com/api/now/table/change_request",
	"https://%s.service-now.com/api/now/v2/table/user",
	"https://%s.service-now.com/api/now/sp/search",
	"https://%s.service-now.com/api/now/table/sys_user",
	"https://%s.service-now.com/api/sn_sc/servicecatalog/items",
	"https://%s.service-now.com/api/now/table/sys_db_object",
	"https://%s.service-now.com/api/now/table/task",
        "https://%s.service-now.com/api/sn_hr_core/hr",
        "https://%s.service-now.com/api/now/table/sys_metadata",
        "https://%s.service-now.com/api/now/table/cmdb_ci_computer/",
	"https://%s.service-now.com/api/now/table/cmdb_ci_network/",
	"https://%s.service-now.com/api/now/table/cmdb_ci_storage/",
	"https://%s.service-now.com/api/now/table/cmdb_ci_database/",
        "https://%s.service-now.com/api/now/table/kb_knowledge/",
	"https://%s.service-now.com/api/now/table/kb_article/",
	"https://%s.service-now.com/api/now/table/kb_category/",
	"https://%s.service-now.com/api/now/table/kb_language/",
        "https://%s.service-now.com/api/now/table/sys_user/",
	"https://%s.service-now.com/api/now/table/sys_user/insert",
	"https://%s.service-now.com/api/now/table/sys_user/update",
	"https://%s.service-now.com/api/now/table/sys_user_group/",
	"https://%s.service-now.com/api/now/table/sys_user_role/",
        "https://%s.service-now.com/stats.do",
        "https://%s.service-now.com/replication.do/",
        "https://%s.service-now.com/threads.do/",
        "https://%s.service-now.com/status.do",
        "https://%s.service-now.com/sys_properties_list.do",
        "https://%s.service-now.com/sys_properties.do",
        "https://%s.service-now.com/sys_user_list.do",
        "https://%s.service-now.com/sys_user.do",
        "https://%s.service-now.com/sys_metadata_list.do",
        "https://%s.service-now.com/sys_db_object_list.do",
        "https://%s.service-now.com/nav_to.do",
        "https://%s.service-now.com/cache.do",
        "https://%s.service-now.com/exceptions.do",
        "https://%s.service-now.com/sys_email_list.do",
        "https://%s.service-now.com/sys_attachment.do",
        "https://%s.service-now.com/sys_script_list.do",
        "https://%s.service-now.com/syslog_transaction_list.do",
        "https://%s.service-now.com/api/now/table/sys_dictionary",
        "https://%s.service-now.com/api/now/table/sys_schema",
        "https://%s.service-now.com/api/now/table/sys_ui_view",
        
	// Azure Web Apps
	"https://%s.azurewebsites.net/api/data",
	"https://%s.azurewebsites.net/api/health",
	"https://%s.azurewebsites.net/api/status",
	"https://%s.azurewebsites.net/api/v1/users",
	"https://%s.azurewebsites.net/api/v1/info",
	"https://%s.azurewebsites.net/api/v1/ping",
	"https://%s.azurewebsites.net/api/public",
	"https://%s.azurewebsites.net/.auth/me",
        "https://%s.azurewebsites.net/debug",
        "https://%s.azurewebsites.net/admin",
        "https://%s.azurewebsites.net/.env",
        "https://%s.azurewebsites.net/api/.env",
        "https://%s.azurewebsites.net/api/config",
        "https://%s.azurewebsites.net/api/settings",
        "https://%s.azurewebsites.net/api/secrets",
        "https://%s.azurewebsites.net/api/v1/config",
        "https://%s.azurewebsites.net/api/v1/settings",
        "https://%s.azurewebsites.net/api/v1/logs",
        "https://%s.azurewebsites.net/api/v1/debug",
        "https://%s.azurewebsites.net/api/v1/internal",
        "https://%s.azurewebsites.net/api/v1/admin",
        "https://%s.azurewebsites.net/api/internal",
        "https://%s.azurewebsites.net/api/devtools",
        "https://%s.azurewebsites.net/api/graphql",
        "https://%s.azurewebsites.net/graphql",
        "https://%s.azurewebsites.net/api/v1/token",
        "https://%s.azurewebsites.net/api/token",
        "https://%s.azurewebsites.net/api/auth/token",
        "https://%s.azurewebsites.net/api/.well-known/openid-configuration",
        "https://%s.azurewebsites.net/.well-known/openid-configuration",
        "https://%s.azurewebsites.net/debug/vars",
        "https://%s.azurewebsites.net/__debug__",
        "https://%s.azurewebsites.net/api/debug/info",

	// Atlassian (Jira/Confluence)
	"https://%s.atlassian.net/rest/api/2/issue/",
	"https://%s.atlassian.net/rest/api/2/search",
	"https://%s.atlassian.net/rest/api/2/project",
	"https://%s.atlassian.net/rest/api/2/user",
	"https://%s.atlassian.net/rest/api/2/dashboard",
	"https://%s.atlassian.net/rest/api/3/myself",
	"https://%s.atlassian.net/rest/agile/1.0/board",
	"https://%s.atlassian.net/wiki/rest/api/space",
        "https://%s.atlassian.net/rest/api/2/attachment",
        "https://%s.atlassian.net/rest/api/2/field",
        "https://%s.atlassian.net/rest/api/2/auditing/record",
        "https://%s.atlassian.net/rest/api/3/issue/picker",
        "https://%s.atlassian.net/rest/api/3/notifications",
        "https://%s.atlassian.net/rest/api/3/group",
        "https://%s.atlassian.net/rest/api/3/user/search",
        "https://%s.atlassian.net/rest/api/3/role",
        "https://%s.atlassian.net/rest/api/3/project/search",
        "https://%s.atlassian.net/rest/capabilities",
        "https://%s.atlassian.net/rest/api/latest/settings/logo",
        "https://%s.atlassian.net/plugins/servlet/oauth/authorize",
        "https://%s.atlassian.net/rest/servicedeskapi/request",
        "https://%s.atlassian.net/rest/servicedeskapi/servicedesk",
        "https://%s.atlassian.net/rest/api/3/webhook",
        "https://%s.atlassian.net/rest/api/3/issue/createmeta",
        "https://%s.atlassian.net/rest/servicedeskapi/customer",
        

	// SAP SuccessFactors
	"https://%s.successfactors.com/odata/v2/User",
	"https://%s.successfactors.com/odata/v2/JobRequisition",
	"https://%s.successfactors.com/odata/v2/PerPerson",
	"https://%s.successfactors.com/odata/v2/EmpEmployment",
	"https://%s.successfactors.com/odata/v2/Background_Education",
	"https://%s.successfactors.com/odata/v2/FODepartment",
	"https://%s.successfactors.com/odata/v2/FOCompany",
	"https://%s.successfactors.com/odata/v2/Photo",
        "https://%s.ariba.com/api/data/order",
        "https://%s.ariba.com/api/data/invoice",
        "https://%s.concur.com/api/v3.0/expenseReports",
        "https://%s.concur.com/api/v3.0/trips",

	// Zendesk
	"https://%s.zendesk.com/api/v2/users.json",
	"https://%s.zendesk.com/api/v2/tickets.json",
	"https://%s.zendesk.com/api/v2/groups.json",
	"https://%s.zendesk.com/api/v2/views.json",
	"https://%s.zendesk.com/api/v2/macros.json",
	"https://%s.zendesk.com/api/v2/help_center/articles.json",
	"https://%s.zendesk.com/api/v2/organizations.json",
	"https://%s.zendesk.com/api/v2/ticket_fields.json",
        "https://%s.zendesk.com/api/v2/tickets/894523.json",
        "https://%s.zendesk.com/api/v2/tickets/1023847.json",
        

	// GitHub Enterprise
	"https://%s.githubenterprise.com/api/v3/user",
	"https://%s.githubenterprise.com/api/v3/repos",
	"https://%s.githubenterprise.com/api/v3/orgs",
	"https://%s.githubenterprise.com/api/v3/teams",
	"https://%s.githubenterprise.com/api/v3/emojis",
	"https://%s.githubenterprise.com/api/v3/issues",
	"https://%s.githubenterprise.com/api/v3/events",
	"https://%s.githubenterprise.com/api/v3/meta",
        "https://%s.githubenterprise.com/api/v3/rate_limit",
        "https://%s.githubenterprise.com/api/v3/gists/public",
        "https://%s.githubenterprise.com/api/v3/notifications",
        "https://%s.gitlab.com/api/v4/users",
        "https://%s.gitlab.com/api/v4/projects",
        "https://%s.gitlab.com/api/v4/projects?membership=true",
        "https://%s.gitlab.com/api/v4/groups",
        "https://%s.gitlab.com/api/v4/metadata",
        "https://%s.gitlab.com/api/v4/version",
        "https://%s.gitlab.com/api/v4/runners",

	// Okta
	"https://%s.okta.com/api/v1/users",
	"https://%s.okta.com/api/v1/groups",
	"https://%s.okta.com/api/v1/apps",
	"https://%s.okta.com/api/v1/sessions/me",
	"https://%s.okta.com/api/v1/events",
	"https://%s.okta.com/api/v1/logs",
	"https://%s.okta.com/api/v1/idps",
	"https://%s.okta.com/api/v1/meta/schemas/user/default",
        "https://%s.okta.com/oauth2/default/.well-known/openid-configuration",
        "https://%s.okta.com/api/v1/meta/schemas/group/default",
        "https://%s.okta.com/.well-known/openid-configuration",
        "https://%s.okta.com/app/UserHome",

	// Grafana
	"https://%s.grafana.net/api/dashboards/home",
	"https://%s.grafana.net/api/org",
	"https://%s.grafana.net/api/search",
	"https://%s.grafana.net/api/datasources",
	"https://%s.grafana.net/api/alert-notifications",
	"https://%s.grafana.net/api/plugins",
	"https://%s.grafana.net/api/snapshots",
	"https://%s.grafana.net/api/user/preferences",
        "https://%s.grafana.net/api/admin/settings",
        "https://%s.grafana.net/api/org/users",
        "https://%s.grafana.net/api/users",
        "https://%s.grafana.net/api/user",
        "https://%s.grafana.net/api/login/ping",
        "https://%s.grafana.net/api/metrics",
        "https://%s.grafana.net/api/annotations",
        "https://%s.grafana.net/api/annotations/graph",
        "https://%s.grafana.net/api/live/ws",
        "https://%s.grafana.net/api/health",
        "https://%s.grafana.net/api/frontend/settings",
        "https://%s.grafana.net/api/short-urls",
        "https://%s.grafana.net/api/login",
        "https://%s.grafana.net/api/signup",
        "https://%s.grafana.net/api/password/forgot",
        "https://%s.grafana.net/api/auth/keys",
        "https://%s.grafana.net/api/live/ws",
        "https://%s.grafana.net/api/dashboards/db",
        "https://%s.grafana.net/api/metrics/search",

	// New Relic
	"https://%s.newrelic.com/api/v2/applications.json",
	"https://%s.newrelic.com/api/v2/alerts_policies.json",
	"https://%s.newrelic.com/api/v2/servers.json",
	"https://%s.newrelic.com/api/v2/plugins.json",
	"https://%s.newrelic.com/api/v2/key_transactions.json",
	"https://%s.newrelic.com/api/v2/synthetics_monitors.json",
	"https://%s.newrelic.com/api/v2/users.json",

	// Sentry
	"https://%s.sentry.io/api/0/projects/",
	"https://%s.sentry.io/api/0/organizations/",
	"https://%s.sentry.io/api/0/issues/",
	"https://%s.sentry.io/api/0/events/",
	"https://%s.sentry.io/api/0/teams/",
	"https://%s.sentry.io/api/0/releases/",
	"https://%s.sentry.io/api/0/environments/",
	"https://%s.sentry.io/api/0/monitors/",
        "https://%s.sentry.io/sitemap.xml",
        "https://%s.sentry.io/api/0/internal/health/",       

	// Notion
	"https://%s.notion.site/api/v3/getSpaces",
	"https://%s.notion.site/api/v3/loadUserContent",
	"https://%s.notion.site/api/v3/getUserSharedPages",
	"https://%s.notion.site/api/v3/getPublicPageData",
	"https://%s.notion.site/api/v3/syncRecordValues",
	"https://%s.notion.site/api/v3/getRecordValues",
	"https://%s.notion.site/api/v3/queryCollection",
	"https://%s.notion.site/api/v3/loadPageChunk",

	// Stripe
	"https://api.%s.stripe.com/v1/customers",
	"https://api.%s.stripe.com/v1/charges",
	"https://api.%s.stripe.com/v1/subscriptions",
	"https://api.%s.stripe.com/v1/invoices",
	"https://api.%s.stripe.com/v1/refunds",
	"https://api.%s.stripe.com/v1/products",
	"https://api.%s.stripe.com/v1/plans",
	"https://api.%s.stripe.com/v1/events",

         // Intercom
	"https://%s.intercom.io/api/v2/users",
	"https://%s.intercom.io/api/v2/contacts",
	"https://%s.intercom.io/api/v2/conversations",
	"https://%s.intercom.io/api/v2/tags",
	"https://%s.intercom.io/api/v2/segments",
	"https://%s.intercom.io/assets/intercom.js",
	"https://%s.intercom.io/widget",

        // finale
        "https://%s/auth/admin/realms/master/users",
        "https://%s/auth/realms/master/protocol/openid-connect/token",
        "https://%s/auth/admin/realms/master/clients",
        "https://%s.auth0.com/api/v2/users",
        "https://%s.auth0.com/api/v2/connections",
        "https://%s.auth0.com/api/v2/clients",
        "https://%s.auth0.com/api/v2/roles",
        "https://%s.auth0.com/api/v2/resource-servers",
        "https://%s.firebaseio.com/.json",
        "https://%s.firebaseio.com/users.json",
        "https://%s.firebaseio.com/config.json",
        "https://%s.firebaseio.com/messages.json",
        "https://%s.firebaseio.com/profiles.json",
        "https://%s.firebaseio.com/orders.json",
        "https://%s.firebaseio.com/settings.json",
        "https://%s.firebaseio.com/admins.json",
        "https://%s.firebaseio.com/notifications.json",
        "https://%s.cloudfunctions.net/",
        "https://%s.web.app/",
        "https://%s.storage.googleapis.com/",
        "https://%s.apigateway.googleapis.com/",

}

func checkURL(url string) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode != http.StatusNotFound
}

func main() {
	// ✅ Banner ASCII com cor
	figure.NewFigure("arara", "doom", true).Print()

	// Adicionar espaço antes da solicitação de entrada
	fmt.Println()

fmt.Println(`

 ARARA - API Recon Automator for Reconnaissance and Analysis
 __________________________________________________________
          \
           \  🦜
            \    "Se tem endpoint voando, a Arara vai pegar!"

Resumo objetivo:
💥 Varre APIs com mais sede que estagiário atrás de bug bounty.
🎯 Encontra endpoints REST iN PeAcE.

Assinatura:
Desenvolvido por [math7]
🚀 Powered by Go!!
`)

	var company string
	fmt.Print("🔎 🦜- Digite o nome da empresa (ex: globo, nasa...): ")
	fmt.Scanln(&company)
	company = strings.ToLower(company)

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = "🚀 Procurando APIs conhecidas... "
	s.Start()

	time.Sleep(2 * time.Second)
	s.Stop()

	fmt.Println("\n📡 Resultados:")

	for _, pattern := range apiPatterns {
		url := fmt.Sprintf(pattern, company)
		ok := checkURL(url)

		if ok {
			color.Green("✅ Encontrada: %s", url)
		} else {
			color.Red("❌ Não responde: %s", url)
		}
		time.Sleep(250 * time.Millisecond)
	}

	color.Cyan("\n🔍 Scan finalizado para: %s", company)
}

