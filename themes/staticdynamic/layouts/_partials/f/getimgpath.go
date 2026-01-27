{{/* <!--
	
	f/getimgpath
	
	Gets full path to picture for pages and data files.
	Allows to use a specific .thumbnail image.
	
	Arguments:
	- page object or data object
	
	Process:
	- If page:
		- try to use thumbnail.jpg; if found, the result is set to "thumbnail.jpg" (string)
		- if not found, try to use .Params.picture
	- else try to use .picture; in that case, the picture is set to the value of .picture (string)
	- else try to use .img; in that case, the picture is set to the value of .img (string)
	- else try to use .src; in that case, the picture is set to the value of .src (string)

	- if result is a string then complete path with .url if only picture name specified,
	  eg. "pic.jpg" for url "/services/" will become "/services/pic.jpg", "/toto/pic.jgp" will be left untouched
	
	Returns:
	- path to picture,
	- or "not found" if not found.

--> */}}

{{/* <!-- initialize variables  --> */}}
{{ $type := partial "f/getdatatype" . }}
{{ $data := partial "f/getdata" . }}
{{ $imgpath := false }}
{{ $img := false }}
{{ $found := false }}
{{ $thumbtest := false }}

{{/* <!-- if page object... --> */}}
{{ if eq "page" $type }}
	{{/* <!-- ...then check for file named thumbnail.jpg in same directory --> */}}
	{{ $thumbtest = add .File.Dir "thumbnail.jpg" }}
	{{/* <!-- if a file named thumbnail.jpg exists... --> */}}
	{{ if fileExists $thumbtest }}
		{{/* <!-- ...then use thumbnail.jpg --> */}}
	  	{{ $img = "thumbnail.jpg" }}
		{{/* <!-- and remember that we found $img --> */}}
		{{ $found = true }}
	{{ end }}
{{ end }}

{{/* <!-- if $img was found... --> */}}
{{ if $found }}
	{{/* <!-- ...then do nothing --> */}}
{{/* <!-- else if the .img property is set in $data... --> */}}
{{ else if (isset $data "img") }}
	{{/* <!-- ...then use it --> */}}
	{{ $img = $data.img }}
	{{/* <!-- and remember that we found $img --> */}}
	{{ $found = true }}
{{/* <!-- else if the .picture property is set in $data... --> */}}
{{ else if isset $data "picture" }}
	{{/* <!-- ...then use it (deprecated, use img instead) --> */}}
	{{ $img = $data.picture }}
	{{/* <!-- and remember that we found $img --> */}}
	{{ $found = true }}
{{/* <!-- else if the .src property is set in $data... --> */}}
{{ else if isset $data "src" }}
	{{/* <!-- else it can be a data or shortcode file, use .src property --> */}}
	{{ $img = $data.src }}
	{{ $found = true }}
{{ else }}
	{{ $found = false }}
{{ end }}

{{ if $found }}
	{{/* <!-- if it's a string check if path should be completed --> */}}		
	{{ $firstchar := substr $img 0 1 }}
	{{ $firstfour := substr $img 0 4 }}
	{{ if eq "/" $firstchar }}
		{{/* <!-- if starts with / then .picture contains rel path to img --> */}}
		{{ $imgpath = $img }}
	{{ else if eq "http" $firstfour }}
		{{/* <!-- if starts with http then .picture contains abs path to img --> */}}
		{{ $imgpath = $img }}
	{{ else }}
		{{/* <!-- else add path --> */}}
		{{ if (isset $data "url") }}
			{{ $imgpath = path.Join $data.url $img }}
		{{ else }}
			{{ $imgpath = path.Join .RelPermalink $img }}
		{{ end }}
	{{ end }}
{{ else }}
	{{ $imgpath = false }}
{{ end }}

{{ return $imgpath }}
