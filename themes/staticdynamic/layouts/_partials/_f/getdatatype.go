{{/* <!--
	
	f/getdatatype
	
	Returns the data type.
	Uses the result of the printf "%T" function.
	
	Arguments:
	- data
	
	Returns:
	- "type" as string (see code below for full list of returned values)
	
 --> */}}
 
{{- $type := (printf "%T" .) -}}
{{- $result := "" -}}

{{- if eq "int" $type -}}
	{{/* <!-- type is integer number --> */}}
	{{- $result = "int" -}}
{{- else if eq "float64" $type -}}
	{{/* <!-- type is floating number --> */}}
	{{- $result = "float" -}}
{{- else if eq "string" $type -}}
	{{/* <!-- type is string --> */}}
	{{- $result = "string" -}}
{{- else if eq "map[string]interface {}" $type -}}
	{{/* <!-- type is collection of dictionary entries (key, val) --> */}}
	{{- $result = "dict" -}}
{{- else if eq "[]string" $type -}}
	{{/* <!-- type is slice of strings --> */}}
	{{- $result = "slice" -}}
{{- else if eq "[]int" $type -}}
	{{/* <!-- type is slice of integers --> */}}
	{{- $result = "slice" -}}
{{- else if eq "[]interface {}" $type -}}
	{{/* <!-- type is slice of data --> */}}
	{{- $result = "slice" -}}
{{- else if eq "*hugolib.pageState" $type -}}
	{{/* <!-- type is page --> */}}
	{{- $result = "page" -}}
{{- else if eq "maps.Params" $type -}}
	{{/* <!-- type is page.Params --> */}}
	{{/* <!-- note: shortcode.Params is identified as a "dict" data type --> */}}
	{{- $result = "params" -}}
{{- else if eq "*hugolib.ShortcodeWithPage" $type -}}
	{{/* <!-- type is shortcode --> */}}
	{{- $result = "shortcode" -}}
{{- else if eq "template.HTML" $type -}}
	{{/* <!-- type is html --> */}}
	{{/* <!-- note: shortcode.Inner is identified as html --> */}}
	{{- $result = "html" -}}
{{- else if eq "<nil>" $type -}}
	{{/* <!-- type is nil (nothing) --> */}}
	{{- $result = "nil" -}}
{{- else if eq "bool" $type -}}
	{{/* <!-- type is nil (nothing) --> */}}
	{{- $result = "bool" -}}
{{- else -}}
	{{/* <!-- if not listed, return raw data type --> */}}
	{{- $result = $type -}}
{{- end -}}

{{- return $result -}}
