{{/* <!--
	
	f/isactive
	
	Test if page is currently active or not based on .date, .expirydate and .draft

	Arguments:
	- . as object containing .date, .expirydate and .draft properties (all are optional)
	
	Returns:
	- true if
		.date is today or after today or is not set
		AND .expirydate is before today or is not set
		AND .draft is not true or is not set
	- else false (either draft, before .date or .expirydate reached)
	
 --> */}}

{{ $result := false }}
{{ if and (or (eq nil .date) (ge (now.Format "2006-01-02") .date)) (or (eq nil .expirydate) (lt (now.Format "2006-01-02") .expirydate)) }}
	{{ if ne .draft true }}
		{{ $result = true }}
	{{ end }}
{{ end }}

{{ return $result }}