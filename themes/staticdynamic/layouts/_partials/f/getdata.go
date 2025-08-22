{{/* <!--
	
	f/getdata

	Returns a "harmonized" data structure, removing .Params if necessary.
	e.g. the structure:
	.Params
		.title
		.description
	will become
	.title
	.description

	the structure	
	.title
	.description
	will remeain 
	.title
	.description

	Arguments:
	- .Params.data or .data
	
	Returns:
	- .data

--> */}}

{{/* <!-- initialize variables --> */}}
{{- $data := . }}
{{- $type := partial "f/getdatatype.go" $data -}}
{{- $result := dict -}}

{{/* <!-- if page or shortcode, set data to .Params --> */}}
{{- if or (eq "page" $type) (eq "shortcode" $type) -}}
	{{- with $data.Params -}}
		{{- $data = . -}}
	{{- end -}}
	{{- $type = "dict" -}}
{{- else if eq "params" $type -}}
	{{- $type = "dict" -}}
{{- end -}}

{{- if eq "dict" $type -}}
	{{/* <!-- if dict, copy data to new modifiable structure --> */}}
	{{/* <!-- (.Params is read-only, which may make manipulating its values difficult --> */}}
	{{- range $key, $value := $data -}}
		{{- $result = merge $result (dict $key $value) -}}
	{{- end -}}
{{- else -}}
	{{/* <!-- else simply return data --> */}}
	{{- $result = $data -}}
{{- end -}}

{{- return $result -}}
