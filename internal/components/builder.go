/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Thursday, November 27th 2025, 2:12:40 pm
 * Author: Md. Asraful Haque
 *
 */

package components

import "context"

type Builder[Result BuilderResult, Data any] interface {
	New(ctx context.Context, data Data) (Result, error)
}

// BuilderResult can be converted to a parts list.
type BuilderResult interface {
	// ToComponents converts the results to a list of components.
	ToComponents() Components
}
