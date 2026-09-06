package routes

import "testing"

func TestSSRPageDataForPath(t *testing.T) {
	tests := []struct {
		path      string
		wantTitle string
	}{
		{path: "/", wantTitle: "首页｜neter-template"},
		{path: "/latest", wantTitle: "最新资源｜neter-template"},
		{path: "/missing", wantTitle: "页面不存在｜neter-template"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			page := ssrPageDataForPath(tt.path)
			if page.Title != tt.wantTitle {
				t.Errorf("title = %q, want %q", page.Title, tt.wantTitle)
			}
			if page.Description == "" || page.Keywords == "" {
				t.Errorf("metadata must not be empty: %#v", page)
			}
		})
	}
}
