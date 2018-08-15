package tool

import "testing"

func TestAesEncrypt(t *testing.T) {
	type args struct {
		plaintext string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", args{plaintext: "shinjuwu"}, "shinjuwu", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesEncrypt(tt.args.plaintext)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotDe, _ := AesDecrypt(got)
			if gotDe != tt.want {
				t.Errorf("AesEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
