{{/* <!-- 
	
	getHTMLAttrClass
	
	Returns a formated class="list of classes" HTML attribute

	Arguments:
	- slice containing the classes, eg slice "cl-shortcode cl-shortcode-item-card columns-2"
	
 --> */}}

{{/* <!-- take collection of classes and turn it into a string separated with spaces --> */}}
{{- $classes := delimit . " " -}}
{{- $result := "" -}}

{{- if ne $classes "" -}}
	{{- $result = printf "class=\"%s\"" $classes -}}
{{- end -}}

{{- return $result -}}
