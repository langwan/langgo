package helper_mime

import "testing"

//func TestGetFileContentByFilename(t *testing.T) {
//	mimeType, fileType, err := GetFileContentByFilename("../../testdata/sample.jpg")
//	assert.NoError(t, err)
//	assert.Equal(t, mimeType, "image/jpeg")
//	assert.Equal(t, fileType, "image")
//}

func TestGetFileContentByFilename(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name         string
		args         args
		wantMimeType string
		wantFileType string
		wantErr      bool
	}{
		{
			name: "a",
			args: args{
				filename: "../../testdata/samples/sample.ARW",
			},
			wantMimeType: "image/tiff",
			wantFileType: "raw",
			wantErr:      false,
		}, {
			name: "b",
			args: args{
				filename: "../../testdata/samples/sample.jpg",
			},
			wantMimeType: "image/jpeg",
			wantFileType: "image",
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMimeType, gotFileType, err := GetMimeType(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileContentByFilename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMimeType != tt.wantMimeType {
				t.Errorf("GetFileContentByFilename() gotMimeType = %v, want %v", gotMimeType, tt.wantMimeType)
			}
			if gotFileType != tt.wantFileType {
				t.Errorf("GetFileContentByFilename() gotFileType = %v, want %v", gotFileType, tt.wantFileType)
			}
		})
	}
}
