{{/* <!--
	
	f/getblocks

	get blocks of the type specified, "head" or "body"
	
	Arguments:
	- "head" or "body" as string based on which blocks to get
	
	Process:
	- If page.Params.head or page.Params.body then get the list of blocks from page.Params
	- Else if a specific .Type and .Kind block list exists in layouts.yaml, use it
	- Else use _default block list for .Kind in layouts.yaml
	
	Returns:
	- Block list as a list of strings
	  e.g. [ "title", "content" ]
	
	Versions:
	2024-09-02 corrected case where specific .Kind is defined but not specific part (eg body is defined but not head)

 --> */}}

{{ $data := page }} {{/* <!-- page object --> */}}
{{ $part := . }} {{/* <!-- string equal to "head", "body" or "footer" --> */}}

{{/* <!-- set default result to _default blocks matching $part ("head" or "body") --> */}}
{{ $result := index (index (index site.Data.layouts "_default") $data.Kind) $part }}

{{ if eq $part "bodytop" }}
	{{ $result = index site.Data.layouts._common "bodytop" }}
{{ else if eq $part "bodybottom" }}
	{{ $result = index site.Data.layouts._common "bodybottom" }}
{{ else if isset page.Params $part }}
	{{/* <!-- then use specific blocks defined in page front matter --> */}}
	{{ $result = index page.Params $part }}

{{/* <!-- else if specific blocks matching page .Type are defined in /data/layouts.yaml --> */}}
{{ else if (isset site.Data.layouts $data.Type) }}
	{{ $type := index site.Data.layouts $data.Type }}
	{{ if $type }}
		{{/* <!-- and if specific blocks matching page .Kind are defined in /data/layouts.yaml --> */}}
		{{ if (isset $type $data.Kind) }}
			{{ $kind := index $type $data.Kind }}
			{{/* <!-- 2024-09-02 added test to check if $part is set in layout.yaml --> */}}
			{{/* <!-- and if specific blocks matching "part" (head or body) are defined in /data/layouts.yaml --> */}}
			{{ if (isset $kind $part) }}
				{{/* <!-- then use specific blocks defined in /data/layouts.yaml --> */}}
				{{ $result = index (index $type $data.Kind) $part }}
			{{ end }}
		{{ end }}
	{{ end }}
{{ end }}

{{ return $result }}