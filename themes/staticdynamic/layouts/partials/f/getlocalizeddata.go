{{/* <!--
	
	f/getdata

	Standardizes data access base
	If it's a page, will use .Params as base
	If it's another object such as a data object, use . as base
	
	This allows to use for example .picture without having to wonder if it's .picture or .Params.picture
	
--> */}}

{{ $data := . }}
{{ $localizeddata := slice }}

{{ if $data }}
	{{ $localizeddata = index $data site.LanguageCode }}
{{ end }}

{{ return $localizeddata }}
