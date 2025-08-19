{{/* <!-- 
	
	getHTMLAttrStyle
	
	Takes the arguments from the shortcode
	and returns a dictionary of formatted style="list of styles" HTML attributes.
	
	Arguments:
	- shortcode context "." with these optional parameters:
		- bgcolor="#color"
		- color="#color"
	
	Returns:
	- a dictionary with:
		- "main": style="background-color:#color"
		- "text": style="color:#color"

 --> */}}

{{/* <!-- initialize variables --> */}}
{{- $result := slice -}}

{{/* <!-- get parameters --> */}}
{{- $backgroundcolor := .Get "bgcolor" -}}
{{- $contcolor := .Get "contcolor" -}}
{{- $color := .Get "color" -}}

{{- with $backgroundcolor -}}
	{{- $backgroundcolor = (printf "style=\"background-color:%s\"" .) -}}
{{- end -}}

{{- with $contcolor -}}
	{{- $contcolor = (printf "style=\"background-color:%s\"" .) -}}
{{- end -}}

{{- with $color -}}
	{{- $color = (printf "style=\"color:%s\"" .) -}}
{{- end -}}

{{/* <!-- build and return result --> */}}
{{- $result = dict "main" $backgroundcolor "cont" $contcolor "text" $color -}}
{{- return $result -}}
