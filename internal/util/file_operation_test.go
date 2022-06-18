package util

import (
	"log"
	"os"
	"testing"
)

func createTempFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil

}

func TestPrepareFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want1   string
		wantErr bool
	}{
		{
			name: "invalid testcase - empty filepath",
			args: args{
				"",
			},
			want1:   "",
			wantErr: true,
		},
		{
			name: "valid testcase",
			args: args{
				"test.txt",
			},
			want1:   "test.txt",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var filePath string

			// If fileName is given then create a temp file and get the filePath of the same
			if tt.args.fileName != "" {
				dir, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}

				_, err = createTempFile(tt.args.fileName)
				if err != nil {
					log.Fatal(err)
				}

				filePath = dir + `\` + tt.args.fileName
			}

			defer os.Remove(tt.args.fileName)

			_, got1, err := PrepareFile(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrepareFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got1 != tt.want1 {
				t.Errorf("PrepareFile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
