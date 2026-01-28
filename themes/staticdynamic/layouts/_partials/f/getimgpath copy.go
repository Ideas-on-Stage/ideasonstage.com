{{/* <!--
	
	f/getimgpath
	
	Gets full path to picture for pages and data files.
	Allows to use a specific .thumbnail image.
	
	Arguments:
	- page object to get a thumbnail image (useful to get the thumbnail in lists of pages)
	- OR a data object with the .img property containing the link to the image
	  for example:
	  .Params (date object passed as argument)
	    .img: "/img/toto.png"
	  it also accepts the .picture and .src properties, but they are deprecated
	
	Process:
	- If page:
		- try to use thumbnail.jpg; if found, the result is set to "thumbnail.jpg" (string)
		- if not found, try to use .Params.picture
	- else try to use .img; in that case, the picture is set to the value of .img (string)
	- else try to use .picture; in that case, the picture is set to the value of .picture (string)
	- else try to use .src; in that case, the picture is set to the value of .src (string)

	- if result is a string then complete path with .url if only picture name specified,
	  eg. "pic.jpg" for url "/services/" will become "/services/pic.jpg", "/toto/pic.jgp" will be left untouched
	
	Returns:
	- path to picture,
	- or "not found" if not found.

--> */}}



{{/* <!-- INITIALIZE VARIABLES  --> */}}

{{ $type := partial "f/getdatatype" . }}
{{ $data := partial "f/getdata" . }}
{{ $imgpath := false }}
{{ $img := false }}
{{ $found := false }}
{{ $thumbtest := false }}



{{/* <!-- GET IMAGE FILE PATH --> */}}

{{/* <!-- if argument is a page object... --> */}}
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

{{/* <!-- if argument is a string... --> */}}
{{ if eq "string" $type }}
	{{/* <!-- ...then use the string as image source --> */}}
  	{{ $img = $data }}
	{{/* <!-- and remember that we found $img --> */}}
	{{ $found = true }}
{{ end }}

{{/* <!-- if $img was not found... --> */}}
{{ if not $found }}
	{{/* <!-- ...then try to find it in $data properties --> */}}
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
	{{/* <!-- ...else if the .src property is set in $data... --> */}}
	{{ else if isset $data "src" }}
		{{/* <!-- ...then use it (deprecated, use img instead) --> */}}
		{{ $img = $data.src }}
		{{/* <!-- and remember that we found $img --> */}}
		{{ $found = true }}
	{{ else }}
		{{/* <!-- ...else an image was not found anywhere, set $found to false --> */}}
		{{ $found = false }}
	{{ end }}
{{ end }}



{{/* <!-- COMPLETE IMAGE FILE PATH --> */}}

{{/* <!-- if image was found... --> */}}
{{ if $found }}
	{{/* <!-- ...then complete image path if necessary --> */}}
	{{/* <!-- get first character of string --> */}}
	{{ $firstchar := substr $img 0 1 }}
	{{/* <!-- get first four characters of string --> */}}
	{{ $firstfour := substr $img 0 4 }}
	{{/* <!-- if the first character is /... --> */}}
	{{ if eq "/" $firstchar }}
		{{/* <!-- ...then .picture contains path to img, use it --> */}}
		{{ $imgpath = $img }}
	{{/* <!-- ...else if the first four characters are https --> */}}
	{{ else if eq "http" $firstfour }}
		{{/* <!-- ...then .picture contains absolute url to img, use it --> */}}
		{{ $imgpath = $img }}
	{{/* <!-- ...else add path to picture file name  --> */}}
	{{ else }}
		{{/* <!-- if the url is defined in $data... --> */}}
		{{ if (isset $data "url") }}
			{{/* <!-- ...then use it to complete image path --> */}}
			{{ $imgpath = path.Join $data.url $img }}
		{{/* <!-- ...else if it's a page... --> */}}
		{{ else if eq "page" $type }}
			{{/* <!-- ...then use .RelPermalink (page path) to complete image path --> */}}
			{{ $imgpath = path.Join .RelPermalink $img }}
		{{ end }}
	{{ end }}
{{ else }}
	{{/* <!-- ...else no image was found, set path to false --> */}}
	{{ $imgpath = false }}
{{ end }}



{{/* <!-- RETURN PATH --> */}}

{{ return $imgpath }}
