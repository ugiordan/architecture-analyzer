package main

import (
	"flag"
	"reflect"
	"testing"
)

func TestReorderArgs(t *testing.T) {
	newFS := func() *flag.FlagSet {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		fs.String("output", "", "")
		fs.String("extractors", "", "")
		fs.String("org", "", "")
		return fs
	}

	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "flags before positional",
			args: []string{"--output", "out.json", "--extractors", "docker", "repo"},
			want: []string{"--output", "out.json", "--extractors", "docker", "repo"},
		},
		{
			name: "positional before flags",
			args: []string{"repo", "--output", "out.json", "--extractors", "docker"},
			want: []string{"--output", "out.json", "--extractors", "docker", "repo"},
		},
		{
			name: "mixed order",
			args: []string{"--output", "out.json", "repo", "--extractors", "docker"},
			want: []string{"--output", "out.json", "--extractors", "docker", "repo"},
		},
		{
			name: "no flags",
			args: []string{"repo"},
			want: []string{"repo"},
		},
		{
			name: "flag with equals",
			args: []string{"repo", "--output=out.json"},
			want: []string{"--output=out.json", "repo"},
		},
		{
			name: "empty args",
			args: []string{},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := newFS()
			got := reorderArgs(fs, tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
