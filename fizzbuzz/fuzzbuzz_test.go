package fizzbuzz

import "testing"

func TestFizzbuzz(t *testing.T) {

	tests := []struct {
		arg  int
		want string
	}{
		{
			arg:  1,
			want: "1",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("test", func(t *testing.T) {
			got := Fizzbuzz(tt.arg)
			if got != tt.want {
				t.Errorf("Fizzbuzz(%d) got %s expected %v\b", tt.arg, got, tt.want)
			}
		})
	}

}
