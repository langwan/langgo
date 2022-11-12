package helper_size

import "testing"

func TestFromHumanSize(t *testing.T) {
	type args struct {
		size string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				size: "1m",
			},
			want:    1024 * 1024,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RAMInBytes(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromHumanSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromHumanSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}
