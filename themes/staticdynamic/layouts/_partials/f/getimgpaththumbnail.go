{{/* <!--
	
	f/getthumbnailpath
	
	Gets full path to thumbnail for a page included in a list.
	
	Arguments:
	- page object
	
	Process:
	- try to use thumbnail.jpg; if found, the result is set to "thumbnail.jpg" (string)
	- else try to use .img; in that case, the picture is set to the value of .img (string)
	- else try to use .picture; in that case, the picture is set to the value of .picture (string)

	- if result is a string then complete path with .url if only picture name specified,
	  for example:
	  - "pic.jpg" for url "/services/" will become "/services/pic.jpg"
	  - "/toto/pic.jgp" will be left untouched as the path is already included
	
	Returns:
	- path to picture,
	- or false if not found.

--> */}}

{{/* <!-- initialize variables  --> */}}
{{ $type := partial "f/getdatatype" . }}
{{ $data := partial "f/getdata" . }}
{{ $imgpath := false }}
{{ $img := false }}
{{ $found := false }}
{{ $thumbtest := false }}

{{/* <!-- if page object, check if thumbnail.jpg exists, if yes use it --> */}}
{{ if eq "page" $type }}
	{{/* <!-- then this is a page, check for thumbnail.jpg in same directory --> */}}
	{{ $thumbtest = add .File.Dir "thumbnail.jpg" }}
	{{ if fileExists $thumbtest }}
		{{/* <!-- if a file named thumbnail.jpg exists, use thumbnail.jpg --> */}}
	  	{{ $img = "thumbnail.jpg" }}
	{{ else if (isset $data "img") }}
		{{/* <!-- if .img exists, use it --> */}}
		{{ $img = $data.img }}
	{{ else if isset $data "picture" }}
		{{/* <!-- if .picture exists, use it (deprecated in favor of img) --> */}}
		{{ $img = $data.picture }}
	{{ end }}
	
	{{/* <!-- if it's a string check if path should be completed --> */}}
	{{/* <!-- if string contains at least one /... --> */}}
	{{ if strings.Contains $img "/" }}
		{{/* <!-- ...then it already contains path to img --> */}}
		{{ $imgpath = $img }}
	{{ else }}
		{{/* <!-- else add path --> */}}
		{{ if (isset $data "url") }}
			{{ $imgpath = path.Join $data.url $img }}
		{{ else }}
			{{ $imgpath = path.Join .RelPermalink $img }}
		{{ end }}
	{{ end }}
{{- else -}}
	{{/* <!-- error: argument must be a page object --> */}}
{{- end -}}

{{ return $imgpath }}
