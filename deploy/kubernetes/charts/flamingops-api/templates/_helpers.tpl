# App Name
{{- define "flamingops-api.name" -}}
{{- default .Chart.Name }}
{{- end }}

# Chart Name
{{- define "flamingops-api.chart" -}}
{{- default .Chart.Name .Chart.Version | trimSuffix "@" }}
{{- end }}

# Chart Labels
{{- define "flamingops-api.labels" -}}
helm.sh/chart: {{ include "flamingops-api.chart" . }}
{{ include "flamingops-api.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

# Chart selector labels
{{- define "flamingops-api.selectorLabels" -}}
app.kubernetes.io/name: {{ include "flamingops-api.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
