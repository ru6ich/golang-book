package sub6

const templ = `{{.TotalCount}} тем:
{{range.Items}}-------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | days Ago}} days
{{end}}`
