package golang

const fileTpl = `
/**
  ******************************************************************************
  * Copyright (C), zwwx .
  * Author             :  yuansudong
  * Version            :  1.0
  * Date               :  2021-04-20 
  ******************************************************************************
  * @attention
  *			该板块的代码皆属于生成代码,请不要手动改动它.我在想是不是要把变量弄得看不懂.
  *			以此种方式,避免人为修改.            
  * Copyright (c) 2020 zwwx. All rights reserved.
  *
  ******************************************************************************
  */

package {{ pkg . }}

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"

	{{ range $path, $pkg := enumPackages (externalEnums .) }}
		{{ $pkg }} "{{ $path }}"
	{{ end }}
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}

	{{ range (externalEnums .) }}
		_ = {{ pkg . }}.{{ name . }}(0)
	{{ end }}
)

{{- if fileneeds . "uuid" }}
// define the regex for a UUID once up-front
var _{{ snakeCase .File.InputPath.BaseName }}_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")
{{ end }}

{{ range .AllMessages }}
	{{ template "msg" . }}
{{ end }}
`
