package sub6

import "html/template"

var issueList = template.Must(template.New("issuelist").Parse(`
<hl>{{.TotalCount}} T e M < / h l >
<table>
<tr style=’text-align: left’>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
ctdxa href= ’{{.HTMLURL}}1 >{{.Number}}</ax/td>
<td>{{.State}}</td>
< td x a href= * { { .User.HTMLURL}} ’ >{{ .U se r. L o g in } } < /ax /td >
c t d x a h re f= 1 { { .HTMLURL}}' >{{.T itle } } < /a x /td >
</tr>
{{end}}
</table>
`))
