package helper_graphic

import "testing"

func TestRangeMapper(t *testing.T) {
	type args struct {
		value float64
		ss    float64
		se    float64
		ds    float64
		de    float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "middle",
			args: args{
				value: 0,
				ss:    -1,
				se:    1,
				ds:    -10,
				de:    10,
			},
			want:    0,
			wantErr: false,
		}, {
			name: "value",
			args: args{
				value: 0.8,
				ss:    -1,
				se:    1,
				ds:    -10,
				de:    10,
			},
			want:    8,
			wantErr: false,
		}, {
			name: "value 1",
			args: args{
				value: 0.8,
				ss:    -1,
				se:    1,
				ds:    10,
				de:    -10,
			},
			want:    -8,
			wantErr: false,
		}, {
			name: "value 2",
			args: args{
				value: 0.8,
				ss:    1,
				se:    -1,
				ds:    10,
				de:    -10,
			},
			want:    8,
			wantErr: false,
		}, {
			name: "err 1",
			args: args{
				value: 0.8,
				ss:    -1,
				se:    -1,
				ds:    10,
				de:    -10,
			},
			want:    0,
			wantErr: true,
		}, {
			name: "err 2",
			args: args{
				value: 0.8,
				ss:    -1,
				se:    1,
				ds:    10,
				de:    10,
			},
			want:    0,
			wantErr: true,
		}, {
			name: "err 3",
			args: args{
				value: 2,
				ss:    -1,
				se:    1,
				ds:    10,
				de:    -10,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RangeMapper(tt.args.value, tt.args.ss, tt.args.se, tt.args.ds, tt.args.de)
			if (err != nil) != tt.wantErr {
				t.Errorf("RangeMapper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RangeMapper() got = %v, want %v", got, tt.want)
			}
		})
	}
}
