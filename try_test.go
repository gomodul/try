package try_test

import (
	"testing"

	"github.com/gomodul/try"
)

func TestDo(t *testing.T) {
	type args struct {
		fn func(attempt int) error
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should success",
			args: args{
				fn: func(attempt int) error {
					try.MaxRetries = 3
					if attempt == try.MaxRetries {
						return nil
					}
					panic(try.ErrMaxRetriesReached)
				},
			},
			wantErr: false,
		},
		{
			name: "should failed exceeded retry limit",
			args: args{
				fn: func(attempt int) error {
					try.MaxRetries = 0
					panic(try.ErrMaxRetriesReached)
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := try.Do(tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsMaxRetries(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should true",
			args: args{
				err: try.ErrMaxRetriesReached,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := try.IsMaxRetries(tt.args.err); got != tt.want {
				t.Errorf("IsMaxRetries() = %v, want %v", got, tt.want)
			}
		})
	}
}
