# App Name
{{- define "flamingops.name" -}}
{{- default .Chart.Name }}
{{- end }}

# Chart Name
{{- define "flamingops.chart" -}}
{{- default .Chart.Name .Chart.Version | trimSuffix "@" }}
{{- end }}

# Chart Labels
{{- define "flamingops.labels" -}}
helm.sh/chart: {{ include "flamingops.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

# Chart selector labels
{{- define "flamingops.go-scope" -}}
app.kubernetes.io/scope: {{ printf "%s-go" .Chart.Name }}
{{- end }}

{{- define "flamingops.react-scope" -}}
app.kubernetes.io/scope: {{ printf "%s-react" .Chart.Name }}
{{- end }}
