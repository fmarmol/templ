// Code generated by templ@(devel) DO NOT EDIT.

package testvoid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"

func render() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		_, err = io.WriteString(w, "<br>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<img")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " src=\"https://example.com/image.png\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<br>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<br>")
		if err != nil {
			return err
		}
		return err
	})
}

