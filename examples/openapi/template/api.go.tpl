{{- define "methods.name" -}}
{{- if .OperationId -}}
{{- camelcase .OperationId -}}
{{- else -}}
{{- replace "{" "" .Path | replace "}" "" | replace "/" "_" | cat .Method | camelcase -}}
{{- end -}}
{{- end -}}

{{- define "methods.argumentsList" -}}

{{- end -}}

{{- define "methods.requestName" -}}
{{- template "methods.name" . }}Request
{{- end -}}

{{- define "methods.responseName" -}}
{{- template "methods.name" . }}Response
{{- end -}}

{{- define "schema.type" -}}
	{{- if .type -}}
		{{- .type -}}
	{{- else if .ref -}}
		Ref
	{{- end -}}
{{- end -}}

package {{ .Values.go.package.name }}

type Client interface {
	{{- range $pathKey, $pathValue := .Values.openapi.paths -}}
		{{- range $methodKey, $methodValue := $pathValue }}
	{{- $methodData := dict "Path" $pathKey "Method" $methodKey "OperationId" $pathValue.operationId }}
	{{ template "methods.name" $methodData }}({{ template "methods.requestName" $methodData }}) ({{ template "methods.responseName" $methodData }}, error)
		{{- end -}}
	{{- end }}
}

{{- range $pathKey, $pathValue := .Values.openapi.paths -}}
	{{- range $methodKey, $methodValue := $pathValue }}
{{- $methodData := dict "Path" $pathKey "Method" $methodKey "OperationId" $pathValue.operationId }}

type {{ template "methods.requestName" $methodData }} struct {
	{{- range $methodValue.parameters }}
		{{ camelcase .name }} {{ .schema.type }}
	{{- end}}
}

type {{ template "methods.responseName" $methodData }} struct {

}
	{{- end -}}
{{- end }}
