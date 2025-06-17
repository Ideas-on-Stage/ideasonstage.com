{{/* <!--
	
	f/getdata

	Standardizes data access base
	If it's a page, will use .Params as base
	If it's another object such as a data object, use . as base
	
	This allows to use for example .picture without having to wonder if it's .picture or .Params.picture

	Input:
	- list of languages with localized date inside (see data files for examples)
	
	Output:
	- localized data (if found)
	
--> */}}

{{ $languagecode := site.LanguageCode }}
{{ $shortlangcode := index (split $languagecode "-") 0 }}
{{ $localizeddata := slice }}
{{ $data := . }}

{{ if $data }}
	{{ $localizeddata = index $data $languagecode }}
	{{ if not $localizeddata }}
		{{ $localizeddata = index $data $shortlangcode }}
	{{ end }}
{{ end }}

{{ return $localizeddata }}
