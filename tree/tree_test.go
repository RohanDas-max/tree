package tree

import (
	"reflect"
	"testing"
)

func TestTreeController(t *testing.T) {

	tests := []struct {
		name     string
		config   Config
		path     string
		wantResp string
		wantErr  bool
	}{
		{
			name:   "when no flags set",
			config: Config{},
			path:   "testFolder",
			wantResp: "testFolder\n" +
				"│──dump.txt\n" +
				"└──testFolder2\n" +
				"   │──dump2.txt\n" +
				"   └──testFolder3\n" +
				"      │──dump3.txt\n" +
				"      └──testFolder4\n" +
				"         └──dump4.txt\n\n" +
				"3 directories, 4 files",
			wantErr: false,
		},
		{
			name: "when flags set to relative path",
			config: Config{
				RelativePath: true,
			},
			path: "testFolder",
			wantResp: "testFolder\n" +
				"│──testFolder/dump.txt\n" +
				"└──testFolder/testFolder2\n" +
				"   │──testFolder/testFolder2/dump2.txt\n" +
				"   └──testFolder/testFolder2/testFolder3\n" +
				"      │──testFolder/testFolder2/testFolder3/dump3.txt\n" +
				"      └──testFolder/testFolder2/testFolder3/testFolder4\n" +
				"         └──testFolder/testFolder2/testFolder3/testFolder4/dump4.txt\n\n" +
				"3 directories, 4 files",
			wantErr: false,
		},
		{
			name: "when flags set to dir only",
			config: Config{
				DirOnly: true,
			},
			path: "testFolder",
			wantResp: "testFolder\n" +
				"└──testFolder2\n" +
				"   └──testFolder3\n" +
				"      └──testFolder4\n\n" +
				"3 directories, 0 files",
			wantErr: false,
		},
		{
			name: "when flags depth set to 3",
			config: Config{
				Depth: 3,
			},
			path: "testFolder",
			wantResp: "testFolder\n" +
				"│──dump.txt\n" +
				"└──testFolder2\n" +
				"   │──dump2.txt\n" +
				"   └──testFolder3\n" +
				"      │──dump3.txt\n" +
				"      └──testFolder4\n\n" +
				"3 directories, 3 files",
			wantErr: false,
		},
		{
			name:   "when flag set to dir permission",
			config: Config{Permission: true},
			path:   "testFolder",
			wantResp: "[drwxr-xr-x] testFolder\n" +
				"│──[-rw-r--r--] dump.txt\n" +
				"└──[drwxr-xr-x] testFolder2\n" +
				"   │──[-rw-r--r--] dump2.txt\n" +
				"   └──[drwxr-xr-x] testFolder3\n" +
				"      │──[-rw-r--r--] dump3.txt\n" +
				"      └──[drwxr-xr-x] testFolder4\n" +
				"         └──[-rw-r--r--] dump4.txt\n\n" +
				"3 directories, 4 files",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, gotErr := tt.config.TreeController(tt.path)

			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("got=%v, want=%v", gotResp, tt.wantResp)
			}

			if (gotErr != nil) != tt.wantErr {
				t.Errorf("unexpected error occurred, gotErr:= %v, wantErr:=%v", gotErr, tt.wantErr)
			}

		})
	}

}
