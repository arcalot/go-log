package log_test

import (
	"testing"

	"go.arcalot.io/log/v2"
)

func TestMessageLabels(t *testing.T) {
	var l log.Labels = map[string]string{}
	labelStrEmpty := l.String()
	if l.String() != "" {
		t.Fatalf("Incorrect label string: %s", labelStrEmpty)
	}
	l = map[string]string{"foo": "bar"}
	labelStr1Key := l.String()
	if labelStr1Key != "foo=bar" {
		t.Fatalf("Incorrect label string: %s", labelStr1Key)
	}
	l = map[string]string{"foo": "bar", "baz": "Hello world!"}
	labelStr2Keys := l.String()
	if labelStr2Keys != "foo=bar;baz=Hello world!" && labelStr2Keys != "baz=Hello world!;foo=bar" {
		t.Fatalf("Incorrect label string 3: %s", labelStr2Keys)
	}
}
