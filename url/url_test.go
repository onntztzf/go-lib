package url

import "testing"

func TestAppendParams(t *testing.T) {
	type args struct {
		URL       string
		newParams map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "append parameters", args: args{
			URL:       "https://www.baidu.com",
			newParams: map[string]interface{}{"param1": 1},
		}, want: "https://www.baidu.com?param1=1"},
		{name: "overwrite the original parameter", args: args{
			URL:       "https://www.baidu.com?param1=0",
			newParams: map[string]interface{}{"param1": 1},
		}, want: "https://www.baidu.com?param1=1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendParams(tt.args.URL, tt.args.newParams); got != tt.want {
				t.Errorf("AppendParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
