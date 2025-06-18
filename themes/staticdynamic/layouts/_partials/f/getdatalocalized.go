{{/* <!--
	
	f/getdatalocalized

	Extract localized data from a list of translated data.

	Arguments:
	- Slice with:
	  0 page object with the proper front matter parameter, e.g. body-body-carousels-data: [ "a", "b", "c" ]
	  1 data type to retrieve, e.g. "body-carousel"
	
	Returns:
	- localized data (if found)
	
--> */}}

{{/* <!-- initialize variables --> */}}
{{ $localizeddata := slice }}
{{ $localizedresult := slice }}

{{/* <!-- get current site language code --> */}}
{{ $languagecode := site.LanguageCode }}

{{/* <!-- get short language code (en-GB becomes en) --> */}}
{{ $shortlangcode := index (split $languagecode "-") 0 }}

{{/* <!-- get page object from argument --> */}}
{{ $page := index . 0 }}

{{/* <!-- get block type from argument, e.g. "body-carousel" --> */}}
{{ $pageparam := index . 1 }}

{{/* <!-- get data type from block type, e.g. "body-carousel" gives "carousel" as data type --> */}}
{{ $dataparam := index (split $pageparam "-") 1 }}

{{/* <!-- use block type to retrieve front matter parameter --> */}}
{{ $datalist := index $page.Params (printf "%s%s" $pageparam "-data") }}

{{/* <!-- if front matter parameter is a single string, convert it to list --> */}}
{{ if (partial "f/isstring.go" $datalist) }}
	{{ $datalist = slice $datalist }}
{{ end }}

{{/* <!-- get list of data files matching data type --> */}}
{{ $datafilelist := index site.Data $dataparam }}

{{/* <!-- iterate over list from page front matter parameter (does nothing if list empty) --> */}}
{{ range $datalist }}

	{{/* <!-- get data file matching name configured in page front matter--> */}}
	{{ $datafile := index $datafilelist . }}

	{{/* <!-- try to get localized data matching full language code, e.g. "en-GB" --> */}}
	{{ $localizeddata = index $datafile $languagecode }}

	{{/* <!-- if not found, try to get localized data matching short language code, e.g. "en" --> */}}
	{{ if not $localizeddata }}
		{{ $localizeddata = index $datafile $shortlangcode }}
	{{ end }}

	{{/* <!-- if localized data found, add it to list of results --> */}}
	{{ if $localizeddata }}
		{{ $localizedresult = $localizedresult | append $localizeddata }}
	{{ end }}

{{ end }}

{{ return $localizedresult }}
