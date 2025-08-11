{{/* <!-- 
	
	getHTMLAttrStyle
	
	Returns a formated style="list of styles" HTML attribute

	Arguments:
	- dictionary containing the styles, eg dict "background-color" "white" "color" "#000"
	  accepted dictionary keys are:
		- "background-color" "#color"
		- "color" "#color"
	
 --> */}}

{{- $result := "" -}}
{{- $stylecol := slice -}}
{{- $style := "" -}}
{{- $item := "" -}}

{{/* <!-- get "background-color" style --> */}}
{{- with index . "background-color" -}}
	{{- $stylecol = $stylecol | append (printf "background-color:%s" .) -}}
{{- end -}}

{{/* <!-- get "color" style --> */}}
{{- with index . "color" -}}
	{{- $stylecol = $stylecol | append (printf "color:%s" .) -}}
{{- end -}}

{{/* <!-- take collection of styles and turn them into a string --> */}}
{{- if gt (len $stylecol) 0 -}}
	{{- $style = delimit $stylecol ";" -}}
	{{- $result = printf "style=\"%s\"" $style -}}
{{- end -}}

{{- return $result -}}
