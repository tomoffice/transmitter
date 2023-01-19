package line

import (
	"testing"
)

func TestNotify_Trigger(t *testing.T) {
	type args struct {
		sendMsg string
	}
	tests := []struct {
		name    string
		n       *Notify
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test notify",
			n: &Notify{
				Token: "XXX",
			},
			args: args{
				sendMsg: "golang testing",
			},
			want:    "{\"status\":200,\"message\":\"ok\"}",
			wantErr: false,
		}, {
			name: "test notify error",
			n: &Notify{
				Token: "XXX",
			},
			args: args{
				sendMsg: "",
			},
			want:    "{\"status\":400,\"message\":\"message: must not be empty\"}",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.Trigger(tt.args.sendMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Notify.Trigger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Notify.Trigger() = %v, want %v", got, tt.want)
			}
		})
	}
}
