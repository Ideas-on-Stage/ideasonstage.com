{{/* <!-- 
	
	getHTMLAttrClass

	Takes the arguments from the shortcode
	and creates a dictionary of formatted class="list of classes" HTML attributes.
	returned classes can be used to format html elements from the shortcode.

	Arguments:
	- shortcode context "." containing in parameters:
		- class="list of classes"
		- columns="n" where n is the number of columns
	
	Returns:
	- a dictionary with:
		- "main": class="cl-shortcode cl-shortcode-name list of classes applicable to main container"
		- "text": class="list of classes applicable to text elements"
		- "img": class="list of classes applicable to img elements"

 --> */}}

{{/* <!-- list of values for shortcode, text and img classes --> */}}
{{/* <!-- if class is not listed, it goes into main (shortcode class) --> */}}
{{- $textclasses := slice -}}
{{- $imgclasses := slice -}}

{{/* <!-- list of possible values for img classes --> */}}
{{/* <!-- currently none --> */}}

{{/* <!-- initialize variables --> */}}
{{- $result := slice -}}
{{- $name := .Name -}}
{{- $parent := .Parent -}}
{{- $classmain := slice "cl-shortcode" -}}
{{- $classtext := slice "text" -}}
{{- $classimg := slice "img" -}}

{{/* <!-- get class elements from shortcode parameters --> */}}
{{- $class := (.Get "class") -}}
{{- $columns := (.Get "columns") -}}
{{- $type := (.Get "type") -}}

{{/* <!-- process class name --> */}}
{{- $classmain = $classmain | append (printf "cl-%s" $name) -}}

{{/* <!-- add type class --> */}}
{{- with $type -}}
	{{- $classmain = $classmain | append . -}}
{{- end -}}

{{/* <!-- add columns class --> */}}
{{- with $columns -}}
	{{- $classmain = $classmain | append (printf "columns-%s" .) -}}
{{- end -}}

{{/* <!-- sort classes to main, cont, text and img --> */}}
{{- $class = split $class " " -}}
{{- range $class -}}
	{{- if in $textclasses . -}}
		{{- $classtext = $classtext | append . -}}
	{{- else if strings.Contains . "text-" -}}
		{{- $classtext = $classtext | append . -}}
	{{- else if in $imgclasses . -}}
		{{- $classimg = $classimg | append . -}}
	{{- else if strings.Contains . "img-" -}}
		{{- $classimg = $classimg | append . -}}
	{{- else if strings.Contains . "icon-" -}}
		{{- $classimg = $classimg | append . -}}
	{{- else -}}
		{{/* <!-- otherwise simly add class to other list --> */}}
		{{- $classmain = $classmain | append . -}}
	{{- end -}}
{{- end -}}

{{/* <!-- build main class="" --> */}}
{{- $classmain = delimit $classmain " " -}}
{{- $classmain = trim $classmain " " -}}
{{- $classmain = printf "class=\"%s\"" $classmain -}}

{{/* <!-- build text class="" --> */}}
{{- $classtext = delimit $classtext " " -}}
{{- $classtext = trim $classtext " " -}}
{{- $classtext = printf "class=\"%s\"" $classtext -}}

{{/* <!-- build img class="" --> */}}
{{- $classimg = delimit $classimg " " -}}
{{- $classimg = trim $classimg " " -}}
{{- $classimg = printf "class=\"%s\"" $classimg -}}

{{/* <!-- build results dictionary --> */}}
{{- $result = dict "main" $classmain "text" $classtext "img" $classimg -}}

{{- return $result -}}
