package gedworx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadAll(t *testing.T) {
	a := assert.New(t)
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "boyd.ged",
			args: args{filePath: "./Boyd.ged"},
			want: []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadAll(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

				a.IsType(tt.want, got)
		})
	}
}
