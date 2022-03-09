# App Name
{{- define "flamingops-ui.name" -}}
{{- default .Chart.Name }}
{{- end }}

# Chart Name
{{- define "flamingops-ui.chart" -}}
{{- default .Chart.Name .Chart.Version | trimSuffix "@" }}
{{- end }}

# Chart Labels
{{- define "flamingops-ui.labels" -}}
helm.sh/chart: {{ include "flamingops-ui.chart" . }}
{{ include "flamingops-ui.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

# Chart selector labels
{{- define "flamingops-ui.selectorLabels" -}}
app.kubernetes.io/name: {{ include "flamingops-ui.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
