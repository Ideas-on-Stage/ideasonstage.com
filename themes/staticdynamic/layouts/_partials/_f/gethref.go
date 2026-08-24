{{/* <!--
	
	f/gethref.go
	
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

{{- $data := partial "_f/getdata" . -}}
{{- $type := partial "_f/getdatatype" . -}}
{{- $result := "" -}}

{{/* <!-- if object is a page... --> */}}
{{- if eq $type "page" -}}
	{{/* <!-- ...then get link from front matter --> */}}
	{{- $result = .RelPermalink -}}
{{/* <!-- else if the dictionary property "href" is defined... --> */}}
{{- else if index $data "href" -}}
	{{/* <!-- ...then get link from the dict entry "href" --> */}}
	{{- $result = index $data "href" -}}
{{/* <!-- else if the dictionary property "link" is defined... --> */}}
{{- else if index $data "link" -}}
	{{/* <!-- ...then get link from the dict entry "link" --> */}}
	{{- $result = index $data "link" -}}
{{/* <!-- else if the dictionary property "url" is defined... --> */}}
{{- else if index $data "url" -}}
	{{/* <!-- ...then get link from the dict entry "url" --> */}}
	{{- $result = index $data "url" -}}
{{- end -}}

{{- return $result -}}