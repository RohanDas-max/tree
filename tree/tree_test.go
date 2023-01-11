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
			wantReport: Report{2, 2},
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
			wantReport: Report{2, 2},
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
			wantReport: Report{2, 0},
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotReport, gotErr := tt.config.tree(tt.path, "", "", "", &Report{})

			gr := Report{
				DirCount:  gotReport.DirCount - (gotReport.FileCount),
				FileCount: gotReport.FileCount,
			}

			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("got=%v, want=%v", got, tt.wantResp)
			}

			if gr != tt.wantReport {
				t.Errorf("gotReport=%v, wantReport=%v", gotReport, tt.wantReport)
			}

			if (gotErr != nil) != tt.wantErr {
				t.Errorf("unexpected error occurred, gotErr:= %v, wantErr:=%v", gotErr, tt.wantErr)
			}
		})
	}
}
