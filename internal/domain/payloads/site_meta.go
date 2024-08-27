package payloads

type SiteMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`

	Favicon string `json:"favicon"`
}
