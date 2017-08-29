// Copyright 2017 King Qiu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
// https://github.com/qjw/kelly

// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package render

import (
	"net/http"
)

func writeContentType(w http.ResponseWriter, value string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header.Set("Content-Type", value)
	}
}


