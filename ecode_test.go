package ecode

import (
	"errors"
	"log"
	"net/http"
	"testing"
)

var (
	// base error
	ParameterErr = New(1000400, "request param error", "request param error")
)

func TestEcodeWithReason(t *testing.T) {
	e := FromError(ParameterErr)
	log.Println(e.Error())   // error: code = 1000400 reason =  message = equest param error metadata = map[] cause = <nil>
	log.Println(e.Code())    // 1000400
	log.Println(e.Message()) // request param error
	log.Println("============================")

	Success = New(1, "SUCCESS", "success")

	e2 := FromError(nil)
	log.Println(e2.Error())   // error: code = 1 reason = SUCCESS message = success metadata = map[] cause = <nil>
	log.Println(e2.Code())    // 1
	log.Println(e2.Reason())  // SUCCESS
	log.Println(e2.Message()) // success
	log.Println("============================")

	sms := New(10000, "CTCC", "中国电信").WithMetadata(map[string]string{
		"name":   "jerry",
		"reason": "我是metadata",
	})
	log.Println(sms.Error())   // error: code = 10000 reason = CTCC message = 中国电信 metadata = map[name:jerry reason:我是metadata] cause = <nil>
	log.Println(sms.Code())    // 10000
	log.Println(sms.Reason())  // CTCC
	log.Println(sms.Message()) // 中国电信
	log.Println(sms.Metadata)  // map[name:jerry reason:我是metadata]
	log.Println("============================")

	mms := New(10086, "CMCC", "中国移动").WithCause(errors.New("我是原因"))
	log.Println(mms.Error())   // error: code = 10086 reason = CMCC message = 中国移动 metadata = map[] cause = 我是原因
	log.Println(mms.Code())    // 10086
	log.Println(mms.Reason())  // CMCC
	log.Println(mms.Message()) // 中国电信
	log.Println(mms.Unwrap())  // 我是原因
}

func TestIs(t *testing.T) {
	tests := []struct {
		name string
		e    *Error
		err  error
		want bool
	}{
		{
			name: "true",
			e:    New(404, "test", ""),
			err:  New(http.StatusNotFound, "test", ""),
			want: true,
		},
		{
			name: "false",
			e:    New(0, "test", ""),
			err:  errors.New("test"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ok := tt.e.Is(tt.err); ok != tt.want {
				t.Errorf("Error.Error() = %v, want %v", ok, tt.want)
			}
		})
	}
}
