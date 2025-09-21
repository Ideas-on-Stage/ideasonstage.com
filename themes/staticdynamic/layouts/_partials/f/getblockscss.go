{{/* <!--
	
	f/getblockscss.go

	Get css resources matching the list of block types given as argument for the current page
	- Checks for each block if a .css file exists with same name in /assets/css
	  If yes, add it to the list of .css files
	- Then uses each word of the file name to include additional css
	  e.g. "pages-news" could include css from pages.scss and news.scss if these files exist in /assets/css
	
	Arguments:
	- list of block types
	  e.g. [ "sidebar", "article" ]
	
	Returns:
	- list of css files as a list of strings (WITHOUT path, subfolder or extension)
	  eg [ "title", "content" ] (implied that these are "title.scss" and "content.scss")
	
--> */}}

{{/* <!-- initialize variables --> */}}
{{- $basepath := "assets/css" -}}
{{- $result := slice -}}

{{/* <!-- get list of blocks passed as argument --> */}}
{{- $blocktypes := . -}}

{{/* <!-- iterate over each block type --> */}}
{{- range $blocktypes -}}
	{{/* <!-- use getblocks to get the list of blocks matching the block type --> */}}
	{{- $blocks := partial "f/getblocks.go" . -}}
	
	{{/* <!-- iterate over the list of blocks --> */}}
	{{- range $blocks -}}
		{{/* <!-- try to find .scss file matching the block name --> */}}
		{{/* <!-- e.g. "body-h1" → body-h1.scss --> */}}
		{{- $filepath := printf "%s/%s.scss" $basepath . -}}
		{{- if fileExists $filepath -}}
			{{/* <!-- if found, add it to list of css files --> */}}
			{{- $result = $result | append . -}}
		{{- end -}}
		
		{{/* <!-- if block name has at least 2 words separated by a "-" --> */}}
		{{/* <!-- then to find a css file for each word --> */}}
		{{/* <!-- e.g. "body-h1" → check if "body.scss" and "h1.scss" exist --> */}}
		{{- $split := strings.Split . "-" -}}
		{{- $len := $split | len -}}
		{{- if ge $len 2 -}}
			{{- range $split -}}
				{{- $filepathsplit := printf "%s/%s.scss" $basepath . -}}
				{{- if fileExists $filepathsplit -}}
					{{/* <!-- if found, add it to list of css files --> */}}
					{{- $result = $result | append . -}}
				{{- end -}}	
			{{- end -}}
		{{- end -}}
	{{- end -}}
{{- end -}}

{{/* <!-- deduplicate result (to avoid returning the same css file twice) --> */}}
{{- $result = $result | uniq -}}

{{/* <!-- return list of css files --> */}}
{{- return $result -}}
