{{/* <!--
	
	f/getblocks

	get blocks for the part specified, for example blocks for part "bodytop"
	- If a custom part is declared in page.Params then get the list of blocks from page.Params
	- Else if a specific .Type and .Kind block list exists in layouts.yaml, use it
	- Else use _default block list for .Kind in layouts.yaml
	
	Arguments:
	- part name for which to get the blocks
	  e.g. "bodybottom"
	
	Returns:
	- Block list as a list of strings
	  e.g. [ "title", "content" ]
	
	Versions:
	2025-06-16 added special treatment for page.Type = home to be treated as page.Kind = home
	2024-09-02 corrected case where specific .Kind is defined but not specific part (eg body is defined but not head)

--> */}}

{{ $page := page }} {{/* <!-- page object --> */}}
{{ $part := . }} {{/* <!-- string with name of part to retrieve  --> */}}
{{ $result := index (index (index site.Data.layouts "_default") $page.Kind) $part }} {{/* <!-- else use _default blocks from /data/layouts.html --> */}}


{{/* <!-- bodytop, bodybottom, headtop and headbottom are ignored if declared elsewhere than _common --> */}}
{{/* <!-- it is therefore useless to declare a bodytop part in a specific part or in a page front matter --> */}}
{{ if eq $part "bodytop" }}
	{{ $result = index site.Data.layouts._common "bodytop" }}
{{ else if eq $part "bodybottom" }}
	{{ $result = index site.Data.layouts._common "bodybottom" }}
{{ else if eq $part "headtop" }}
	{{ $result = index site.Data.layouts._common "headtop" }}
{{ else if eq $part "headbottom" }}
	{{ $result = index site.Data.layouts._common "headbottom" }}
	
{{/* <!-- if custom part exists in page front matter, use it --> */}}
{{ else if isset page.Params $part }}
	{{ $result = index page.Params $part }}

{{/* <!-- if specific blocks matching page .Type are defined in /data/layouts.yaml, use it --> */}}
{{ else if (isset site.Data.layouts $page.Type) }}
	{{ $type := index site.Data.layouts $page.Type }}
	{{ if $type }}
		{{/* <!-- and if specific blocks matching page .Kind are defined in /data/layouts.yaml --> */}}
		{{ if (isset $type $page.Kind) }}
			{{ $kind := index $type $page.Kind }}
			{{/* <!-- 2024-09-02 added test to check if $part is set in layout.yaml --> */}}
			{{/* <!-- and if specific blocks matching "part" (head or body) are defined in /data/layouts.yaml --> */}}
			{{ if (isset $kind $part) }}
				{{/* <!-- then use specific blocks defined in /data/layouts.yaml --> */}}
				{{ $result = index (index $type $page.Kind) $part }}
			{{ end }}
		{{ end }}
	{{ end }}
	
{{/* <!-- use Kind=home if declared as page.Type --> */}}	
{{ else if eq "home" $page.Type }}
	{{ $result = index (index (index site.Data.layouts "_default") "home") $part }}
{{ end }}

{{ return $result }}
