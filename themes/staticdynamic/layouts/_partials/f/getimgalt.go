{{/* <!--
	
	f/getimgalt
	
	Gets alt text for image.
	
	Arguments:
	- page object or data object
	
--> */}}

{{- $alt := false -}}
{{- $data := partial "f/getdata.go" . -}}

{{- if not $data -}}
	{{- $alt = "" -}}
{{- else if $data.alt -}}
	{{- $alt = $data.alt -}}
{{- else if $data.title -}}
	{{- $alt = $data.title -}}
{{- else if $data.description -}}
	{{- $alt = $data.description -}}
{{- end -}}

{{- $alt = printf "alt=\"%s\"" $alt -}}

{{- return $alt -}}
