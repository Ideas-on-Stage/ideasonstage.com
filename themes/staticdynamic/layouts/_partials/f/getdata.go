{{/* <!--
	
	f/getdata

	Returns a "harmonized" data structure:
	- remove page or .Params if necessary.
	- other data types (string, int, etc.) are left untouched.
	
	This is especially useful when calling a partial with different data objects (a page, a shortcode, a dict, etc.)
	and you don't want to spend time testing every possible case.
	With f/getdata, you know you can proceed with a standard data structure.
	
	e.g. the page structure:
		page
			.Params
				.title
				.description
	will become a dict:
		.title
		.description
	
	the .Params structure:
		.Params
			.title
			.description
	will become a dict:
		.title
		.description

	the dict structure:
		.title
		.description
	will remain a dict:
		.title
		.description

	Arguments:
	- some data:
		- page object
		- or .Params dict
		- or dict
		- or any other data type such as string, int...
	
	Returns:
	- dict data for page, .Params or dict
	- the original data for other data types.
	
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
