{{/* <!--
	
	f/getpicturepath
	
	Gets full path to picture for pages and data files.
	Allows to use a specific .thumbnail image.
	
	Arguments:
	- page object or data object
	
	Process:
	- If page:
		- try to use thumbnail.jpg; if found, the result is set to "thumbnail.jpg" (string)
		- if not found, try to use .Params.picture; if found, the result is set to the page object, NOT a picture name
		  returning the page object will make it easier to process the pic as a page resource to create different responsive sizes
		  see f/getpicturesrcset for more info on the creation of the responsive sizes
	- else try to use .picture; in that case, the picture is set to the value of .picture (string)
	- else try to use .src; in that case, the picture is set to the value of .src (string)

	- if result is a string then complete path with .url if only picture name specified,
	  eg. "pic.jpg" for url "/services/" will become "/services/pic.jpg", "/toto/pic.jgp" will be left untouched
	
	Returns:
	- relative path to picture,
	- or page object,
	- or "not found" if not found.

--> */}}

{{ $imgpath := "" }}
{{ $picture := "" }}
{{ $found := false }}
{{ $thumbtest := "" }}

{{ if (partial "f/ispage.go" .) }}
	{{/* <!-- then this is a page, check for thumbnail.jpg in same directory --> */}}
	{{ $thumbtest = add .File.Dir "thumbnail.jpg" }}
	{{ if fileExists $thumbtest }}
		{{/* <!-- if a file named thumbnail.jpg exists, use thumbnail.jpg --> */}}
	  	{{ $picture = "thumbnail.jpg" }}
		{{ $found = true }}
	{{ else if .Params.picture }}
		{{/* <!-- if .picture exists, use the picture property --> */}}
		{{ $picture = .Params.picture }}
		{{ $found = true }}
	{{ end }}
{{ else if (isset . "picture") }}
	{{/* <!-- else it can be a data or shortcode file, use .picture property --> */}}
	{{ $picture = .picture }}
	{{ $found = true }}
{{ else if (isset . "img") }}
	{{/* <!-- else it can be a data or shortcode file, use .img property --> */}}
	{{ $picture = .img }}
	{{ $found = true }}
{{ else if (isset . "src") }}
	{{/* <!-- else it can be a data or shortcode file, use .src property --> */}}
	{{ $picture = .src }}
	{{ $found = true }}
{{ else }}
	{{ $found = false }}
{{ end }}

{{ if $found }}
	{{ if (partial "f/isstring.go" $picture) }}
		{{/* <!-- if it's a string check if path should be completed --> */}}		
		{{ $firstchar := substr $picture 0 1 }}
		{{ $firstfour := substr $picture 0 4 }}
		{{ if eq "/" $firstchar }}
			{{/* <!-- if starts with / then .picture contains rel path to img --> */}}
			{{ $imgpath = $picture }}
		{{ else if eq "http" $firstfour }}
			{{/* <!-- if starts with http then .picture contains abs path to img --> */}}
			{{ $imgpath = $picture }}
		{{ else }}
			{{/* <!-- else .picture contains name of img, add path --> */}}
			{{ if .RelPermalink }}
				{{ $imgpath = path.Join .RelPermalink $picture }}
			{{ else if (isset . "url") }}
				{{ $imgpath = path.Join .url $picture }}
			{{ else }}
				{{ $imgpath = path.Join .RelPermalink $picture }}
			{{ end }}
		{{ end }}
	{{ else }}
		{{/* <!-- else it is a page object, return it untouched --> */}}		
		{{ $imgpath = . }}
	{{ end }}
{{ else }}
	{{ $imgpath = "not found" }}
{{ end }}

{{ return $imgpath }}
