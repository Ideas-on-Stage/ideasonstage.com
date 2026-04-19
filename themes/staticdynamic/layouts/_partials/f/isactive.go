{{/* <!--
	
	f/isactive
	
	Test if page is currently active or not based on the following properties:
	- .draft
	- .date
	- .dateEnd

	Arguments:
	- . as object containing .date, .dateEnd and .draft properties (all are optional)
	
	Returns:
	- false if any is true
		.draft is true
		OR today is less than .date
		OR today is greater than .dateEnd
	- true in other cases
	
	History:
	- 2026-04-06: rewrote test to make it cleaner.
	
 --> */}}

{{/* <!-- initialize status as inactive --> */}}
{{- $result := true -}}
{{- $today := (now.Format "2006-01-02") -}}

{{/* <!-- if page is NOT draft... --> */}}
{{- if .draft -}}
	{{- if eq .draft true -}}
		{{/* <!-- ...then set base status as active --> */}}
		{{- $result = false -}}
	{{- end -}}
{{- end -}}
	
{{/* <!-- if there is a date (which means start date)... --> */}}
{{- if .date -}}
	{{/* <!-- ...then if today is less than start date... --> */}}
	{{- if lt $today .date -}}
		{{/* <!-- ...then set status to inactive --> */}}
		{{- $result = false -}}
	{{- end -}}
{{- end -}}
	
{{/* <!-- if there is an end date... --> */}}
{{- if .dateEnd -}}
	{{/* <!-- ...then if today is greater than end date...  --> */}}
	{{- if gt $today .dateEnd -}}
		{{/* <!-- ...then set status as inactive --> */}}
		{{- $result = false -}}
	{{- end -}}
{{- end -}}

{{/* <!-- return result --> */}}
{{- return $result -}}
