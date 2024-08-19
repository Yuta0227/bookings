package models
// TemplateData holds data sent from handlers to templates
// using this prevents the need to pass multiple data types to methods
type TemplateData struct{
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}