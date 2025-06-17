{{/* <!--
	
	f/getlocalizeddata

	For a data structure, retrieves localized entry

--> */}}

{{ $data := . }}
{{ if .Params }}
	{{ $data = .Params }}
{{ end }}
{{ return $data }}
