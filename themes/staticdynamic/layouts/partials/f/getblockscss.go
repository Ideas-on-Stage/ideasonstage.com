{{/* <!--
	
	f/getblockscss.go

	Get css resources matching the list of block types given as argument for the current page
	- Checks for each block if a .css file exists with same name in /assets/css
	  If yes, add it to the list of .css files
	- Then uses each word of the file name to include additional css
	  e.g. "pages-news" could include css from pages.scss and news.scss if these files exist in /assets/css
	
	Arguments:
	- list of block types
	  e.g. [ "side", "main" ]
	
	Returns:
	- list of css files as a list of strings (WITHOUT path, subfolder or extension)
	  eg [ "title", "content" ] (implied that these are "title.scss" and "content.scss")
	
--> */}}

{{ $blocktypes := . }}
{{ $result := slice }}

{{ range $blocktypes }}
	{{ $blocks := partial "f/getblocks.go" . }}
	{{ range $blocks }}
		{{ $filepath := printf "assets/css/%s.scss" . }}
		{{ if fileExists $filepath }}
			{{ $result = $result | append . }}
		{{ end }}
		{{ $split := strings.Split . "-" }}
		{{ $len := $split | len }}
		{{ if ge $len 2 }}
			{{ range $split }}
				{{ $filepathsplit := printf "assets/css/%s.scss" . }}
				{{ if fileExists $filepathsplit }}
					{{ $result = $result | append . }}
				{{ end }}	
			{{ end }}
		{{ end }}
	{{ end }}
{{ end }}
{{- $result = $result | uniq -}}

{{ return $result }}
