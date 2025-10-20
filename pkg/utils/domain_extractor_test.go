package utils

import (
	"testing"
)

func TestExtractDomain(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid URL with www",
			args: args{
				rawURL: "https://www.example.com/path?query=123",
			},
			want:    "www.example.com",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractDomain(tt.args.rawURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractDomain() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractRootDomain(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid URL with www",
			args: args{
				rawURL: "https://www.example.com/path?query=123",
			},
			want:    "example.com",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractRootDomain(tt.args.rawURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractRootDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractRootDomain() got = %v, want %v", got, tt.want)
			}
		})
	}
}
