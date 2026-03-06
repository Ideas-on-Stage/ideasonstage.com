{{/* <!--
	
	f/getimgalt
	
	Gets alt text for image.
	
	Arguments:
	- page object or data object
	
--> */}}

{{- $alt := false -}}
{{- $data := partial "f/getdata.go" . -}}
{{- $type := partial "f/getdatatype.go" . -}}

{{- if not $data -}}
	{{- $alt = false -}}
{{- else if $data.alt -}}
	{{- $alt = $data.alt -}}
{{- else if $data.title -}}
	{{- $alt = $data.title -}}
{{- else if $data.description -}}
	{{- $alt = $data.description -}}
{{- end -}}

{{- if $alt -}}
	{{- $alt = printf "alt=\"%s\"" $alt -}}
{{- end -}}

{{- return $alt -}}
