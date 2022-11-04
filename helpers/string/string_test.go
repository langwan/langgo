package helper_string

import "testing"

func TestUtf8StringLength(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "latin",
			args: args{str: "That's one small step for (a) man, one giant leap for mankind."},
			want: 62,
		},
		{
			name: "cn",
			args: args{str: "这是我个人的一小步，却是人类迈出一大步。"},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Utf8StringLength(tt.args.str); got != tt.want {
				t.Errorf("StringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtf8TruncateText(t *testing.T) {
	type args struct {
		text     string
		max      int
		omission string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "latin",
			args: args{text: "That's one small step for (a) man, one giant leap for mankind.", max: 15, omission: "..."},
			want: "That's one s...",
		}, {
			name: "latin",
			args: args{text: "这是我个人的一小步，却是人类迈出一大步。", max: 15, omission: "..."},
			want: "这是我个人的一小步，却是...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Utf8TruncateText(tt.args.text, tt.args.max, tt.args.omission); got != tt.want {
				t.Errorf("Utf8TruncateText() = %v, want %v", got, tt.want)
			}
		})
	}
}
