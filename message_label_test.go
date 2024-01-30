package log_test

import (
    "testing"

    "go.arcalot.io/log/v2"
)

func TestMessageLabels(t *testing.T) {
    var l log.Labels = map[string]string{}
    if l.String() != "" {
        t.Fatalf("Incorrect label string: %s", l.String())
    }
    l = map[string]string{"foo": "bar"}
    if l.String() != "foo=bar" {
        t.Fatalf("Incorrect label string: %s", l.String())
    }
    l = map[string]string{"foo": "bar", "baz": "Hello world!"}
    labelStr := l.String()
    if labelStr != "foo=bar;baz=Hello world!" && labelStr != "baz=Hello world!;foo=bar" {
        t.Fatalf("Incorrect label string 3: %s", l.String())
    }
}
