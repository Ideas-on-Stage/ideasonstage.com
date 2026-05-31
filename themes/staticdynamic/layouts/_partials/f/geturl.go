{{/* <!--
	
	f/geturl.go
	
	Gets url parameter for pages and data files.
	- If page then use .Permalink
	- else try to use the .href property
	- if not found then try to use .link property
	- if not found then try to use .url property
	
	Arguments:
	- page, data object, or a structure containing a .link or .url property
	
	Returns:
	- url as a string
	- empty string if not found

--> */}}

{{- $data := partial "f/getdata" . -}}
{{- $type := partial "f/getdatatype" . -}}
{{- $result := "" -}}

{{- if eq $type "page" -}}
	{{- $result = .RelPermalink -}}
{{- else if index $data "href" -}}
	{{- $result = index $data "href" -}}
{{- else if index $data "link" -}}
	{{- $result = index $data "link" -}}
{{- else if index $data "url" -}}
	{{- $result = index $data "url" -}}
{{- end -}}

{{- return $result -}}