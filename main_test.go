package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handleErr(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		err  error
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"want err and code 400", args{
			w:    httptest.NewRecorder(),
			err:  errors.New("want err"),
			want: 400,
		}},
		{"want code 200", args{
			w:    httptest.NewRecorder(),
			err:  nil,
			want: 200,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleErr(tt.args.w, tt.args.err)
			w := tt.args.w.(*httptest.ResponseRecorder)
			if code := w.Result().StatusCode; code != tt.args.want {
				t.Errorf("status code want %d got %d", tt.args.want, code)
			}
		})
	}
}
