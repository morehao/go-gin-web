package obj{{.PackagePascalName}}

type {{.StructName}}BaseInfo struct {
{{- range .ModelFields}}
{{- if .IsPrimaryKey}}
    {{- continue}}
{{- end}}
{{- if or (eq .FieldName "deleted_at") (eq .FieldName "deleted_by") }}
    {{- continue}}
{{- end}}
{{- if eq .FieldType "time.Time"}}
    {{.FieldName}} int64 `json:"{{.FieldLowerCaseName}}" form:"{{.FieldLowerCaseName}}"` // {{.Comment}}
{{- else}}
    {{.FieldName}} {{.FieldType}} `json:"{{.FieldLowerCaseName}}" form:"{{.FieldLowerCaseName}}"` // {{.Comment}}
{{- end}}
{{- end}}
}
