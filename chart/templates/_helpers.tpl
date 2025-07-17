{{/* Helm helper templates for fantasy-pokedex */}}
{{- define "fantasy-pokedex.name" -}}
fantasy-pokedex
{{- end -}}

{{- define "fantasy-pokedex.fullname" -}}
{{ printf "%s" (include "fantasy-pokedex.name" .) }}
{{- end -}}
