{{/* <!--
	
	f/getpicturepath
	
	Gets url parameter for pages and data files.
	
	Arguments:
	- page, data object, or a structure containing a .link or .url property
	
	Process:
	- If page:
		- use .Permalink
	- else try to use .link property
	- if not found, try to use .url property

	Returns:
	- url as a string
	- empty string if not found

--> */}}

{{ $result := "" }}
{{ if partial "f/ispage.go" . }}
	{{ $result = .RelPermalink }}
{{ else if isset . "link" }}
	{{ $result = .link }}
{{ else if isset . "url" }}
	{{ $result = .url }}
{{ end }}

{{ return $result }}