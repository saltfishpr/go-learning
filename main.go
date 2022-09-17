package main

import (
	"os"
	"text/template"
)

const templates = `
{{- /* 注释 */ -}}
{{- /* '-'放在后面，则移除后面所有的空白字符(空格、制表符、回车和换行符)，放在前面同理。 */ -}}

{{- /* 定义一个模板名为 T1 */ -}}
{{- define "T1" -}}
Name: {{ .Name }}
{{- end -}}

{{- define "T2" -}}
{{- /* 使用 with 在值存在的时候进入该语句 */ -}}
{{ with .Bio }}; Bio: {{ . }}{{end}}
{{- end -}}

{{- define "T3" -}}
{{ template "T1" .}}{{ template "T2" .}}
{{- end -}}

{{- /* 用'.'将参数传入内层模板 */ -}}
{{- template "T3" . -}}
`

type User struct {
	Name string
	Bio  string
}

func main() {
	tmpl := template.Must(template.New("users").Parse(templates))

	u1 := User{Name: "John", Bio: "a regular user"}
	if err := tmpl.ExecuteTemplate(os.Stderr, "users", u1); err != nil {
		panic(err)
	}
	println()
	u2 := User{Name: "Alice"}
	if err := tmpl.ExecuteTemplate(os.Stderr, "users", u2); err != nil {
		panic(err)
	}
}
