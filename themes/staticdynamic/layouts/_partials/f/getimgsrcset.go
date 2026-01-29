{{/* <!--
	
	f/getimgsrcset

	Returns a slice with three things:
	- the "src" path for the fallback picture
	- the "srcset" paths for the list of responsive pictures
	- the sizes for the pictures
	
	Arguments:
	- image path as string
	
	Returns:
	- a slice with:
		- "src" as src attribute
		- "srcset" with list of image sources for different sizes
		- "sizes" as list of sizes
		
	Example code:
	{{ $srcset := partial "f/getimgsrcset.go" "/img/picture.jgp" }}
	<img src="{{ index $srcset 0 }}" srcset="{{ index $srcset 1 }}" sizes="{{ index $srcset 2 }}">

--> */}}

{{/* <!-- initialize variables --> */}}
{{ $imgSrc := false }}
{{ $imgSrcSet := false }}
{{ $imgSizes := slice }}
{{ $scale := 0 }}
{{ $size := "" }}

{{ $image := page.Resources.Get . }}
{{ if not $image }}
	{{ $image = resources.Get . }}
{{ end }}

{{ if $image }}
	<!-- scale pic if it is not a SVG (SVGs don't need to and cannot be scaled) -->
	{{ if ne "svg" $image.MediaType.SubType }}
		<!-- uses settings from config.toml depending on orientation -->
		{{ $scales := site.Params.images.landscape }}
		{{ if gt $image.Height $image.Width }}
			{{ $scales = site.Params.images.portrait }}
		{{ end }}
		
		<!--
		Add URL for each width to $imgSrcSet variable
		format: "/path/img_1000.jpg 1000w,/path/img_500.jpg 500w"
		Note: the first URL is used as "fallback" src in $imgSrc.
		-->
		{{ $imgSrcSet = slice }}
		{{ range $scales }}
			{{/* <!-- initialize variables --> */}}
			{{ $scale = .scale }}
			{{ $size = .size }}
			{{/* <!-- if width greater than scale... --> */}}
			{{ if gt $image.Width $scale }}
				{{/* <!-- ...then resize pic (downscale) and convert to webp --> */}}
				{{ $imgSrc = (printf "%dx webp" $scale | $image.Resize).RelPermalink }}
			{{ else }}
				{{/* <!-- ...else keep dimensions and convert to webp --> */}}
				{{ $imgSrc = (printf "%dx%d webp" $image.Width $image.Height | $image.Resize).RelPermalink }}
			{{ end }}
			{{/* <!-- add new srcset to srcset --> */}}
			{{ $imgSrcSet = $imgSrcSet | append (printf "%s %dw" $imgSrc $scale) }}
			{{/* <!-- add new size to sizes --> */}}
			{{ $imgSizes = $imgSizes | append (printf "%s" $size ) }}
		{{ end }}
		{{/* <!-- convert slice of sizes to string of sizes separated by comma --> */}}
		{{ $sizes := (delimit $imgSizes ",") }}
		{{/* <!-- convert slice of srcset to string of srcset separated by comma --> */}}
		{{ $set := delimit $imgSrcSet "," }}
		{{/* <!-- build return value with src srcset sizes --> */}}
		{{ $imgSrcSet = slice $imgSrc $set $sizes }}
	{{ else }}
		{{ $imgSrcSet = slice $image false false }}
	{{ end }}
{{ else }}
	<!-- If no resource found, fall back to using default image -->
	{{ if . }}
		{{ if ne "" . }}
			{{ $imgSrcSet = slice . false false }}
		{{ end }}
	{{ end }}
{{ end }}
{{ return $imgSrcSet }}
