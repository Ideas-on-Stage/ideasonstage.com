{{/* <!--
	
	f/getimgpath
	
	Gets full path to image file.
	
	Arguments:
	- data object with .img, .picture or .src property.
	
	Process:
	- try to use .img; in that case, the picture is set to the value of .img (string)
	- else try to use .picture; in that case, the picture is set to the value of .picture (string)
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
{{- $path := partial "f/geturl" $data -}}
{{ $imgpath := false }}
{{ $img := false }}
{{ $found := false }}
{{ $thumbtest := false }}

{{/* <!-- if the .img property is set in $data... --> */}}
{{ if (isset $data "img") }}
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
	{{/* <!-- if string contains at least one /... --> */}}
	{{ if strings.Contains $img "/" }}
		{{/* <!-- ...then it already contains path to img --> */}}
		{{ $imgpath = $img }}
	{{/* <!-- ...else if path exists... --> */}}
	{{ else if $path }}
		{{/* <!-- ...then add path --> */}}
		{{ $imgpath = path.Join $path $img }}
	{{ else }}
		{{/* <!-- ...else send back image name (means that file must be in the same directory as the page) --> */}}
		{{ $imgpath = $img }}
	{{ end }}
{{ else }}
	{{ $imgpath = "ERROR: property .img .picture or .src not found" }}
{{ end }}

{{ return $imgpath }}
