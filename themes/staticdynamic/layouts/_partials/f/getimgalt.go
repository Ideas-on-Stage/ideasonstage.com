{{/* <!--
	
	f/getimgalt
	
	Gets alt text for image.
	
	Arguments:
	- page object or data object
	
--> */}}

{{- $alt := false -}}
{{- $data := partial "f/getdata.go" . -}}

{{- if not $data -}}
	{{- $alt = false -}}
{{- else if $data.alt -}}
	{{- $alt = $data.alt -}}
{{- else if $data.title -}}
	{{- $alt = $data.title -}}
{{- end -}}

{{- return $alt -}}
