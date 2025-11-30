/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Sunday, November 30th 2025, 6:19:50 pm
 * Author: Md. Asraful Haque
 *
 */

package utils

import (
	"bytes"
	"text/template"
)

func RenderConfTemplate(name, confTemplate string, data any) (string, error) {
	tpl, err := template.New(name).Parse(confTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)

	if err != nil {
		return "", err
	}

	return buf.String(), nil

}
