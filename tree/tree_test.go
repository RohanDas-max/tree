package tree

import (
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {

	tests := []struct {
		name       string
		path       string
		config     Config
		wantResp   string
		wantReport Report
		wantErr    bool
	}{
		{
			name:   "success response when no flags set",
			path:   "testFolder",
			config: Config{},
			wantResp: "testFolder\n" +
				"│──dump.txt\n" +
				"└──testFolder2\n" +
				"   └──dump2.txt\n",
			wantReport: Report{},
			wantErr:    false,
		},
		{
			name: "success response when config relative path set true",
			path: "testFolder",
			config: Config{
				RelativePath: true,
				DirOnly:      false,
			},
			wantResp: "testFolder/testFolder\n" +
				"│──testFolder/dump.txt/dump.txt\n" +
				"└──testFolder/testFolder2/testFolder2\n" +
				"   └──testFolder/testFolder2/dump2.txt/dump2.txt\n",
			wantReport: Report{},
			wantErr:    false,
		},
		{
			name: "success response when config directory only set true",
			path: "testFolder",
			config: Config{
				RelativePath: false,
				DirOnly:      true,
			},
			wantResp: "testFolder\n" +
				"└──testFolder2\n",
			wantReport: Report{},
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, gotErr := tt.config.Tree(tt.path, "", "", "", Report{})

			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("got=%v, want=%v", got, tt.wantResp)
			}

			if (gotErr != nil) != tt.wantErr {
				t.Errorf("unexpected error occurred, gotErr:= %v, wantErr:=%v", gotErr, tt.wantErr)
			}
		})
	}

}
