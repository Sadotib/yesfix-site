// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.898
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Hero() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<section class=\"relative min-h-screen bg-brutalist-black text-white font-brutalist overflow-hidden z-0 pt-6 pb-2\"><!-- Background Grid --><div class=\"absolute inset-0 opacity-10 z-0 pointer-events-none\"><div class=\"grid grid-cols-6 h-full w-full\"><div class=\"border border-brutalist-yellow\"></div><div class=\"border border-brutalist-yellow\"></div><div class=\"border border-brutalist-yellow\"></div><div class=\"border border-brutalist-yellow\"></div><div class=\"border border-brutalist-yellow\"></div><div class=\"border border-brutalist-yellow\"></div></div></div><!-- Content Container --><div class=\"container relative z-10 px-4\"><!-- Bouncing Header (Desktop) --><div class=\"hidden sm:grid grid-cols-3 gap-2 text-center border-4 border-brutalist-black bg-brutalist-grey p-3 mb-4\"><div class=\"text-2xl uppercase tracking-wider\">Fast</div><div class=\"text-2xl uppercase tracking-wider\">Reliable</div><div class=\"text-2xl uppercase tracking-wider\">Local</div></div><!-- Bouncing Header (Mobile) --><div class=\"grid sm:hidden grid-cols-3 text-center border-2 border-brutalist-black bg-brutalist-grey p-2 mb-4\"><div class=\"text-sm uppercase tracking-tight\">Fast</div><div class=\"text-sm uppercase tracking-tight\">Reliable</div><div class=\"text-sm uppercase tracking-tight\">Local</div></div><!-- Main Hero Content --><div class=\"grid lg:grid-cols-2 gap-10 items-center min-h-[70vh]\"><!-- Left Text Content --><div class=\"space-y-6\"><div class=\"space-y-4\"><h2 class=\"text-5xl sm:text-5xl md:text-5xl lg:text-5xl uppercase tracking-wider leading-snug sm:leading-tight\"><span class=\"text-brutalist-yellow\">Yes</span>Fix! – The Gamusa of Services</h2><h3 class=\"text-4xl sm:text-4xl md:text-4xl lg:text-4xl uppercase tracking-wider leading-snug\"><span class=\"text-brutalist-yellow\">Dibrugarh’s Most Reliable</span><br>Solution for All<br><span class=\"relative inline-block text-brutalist-yellow\">Cleaning & Fixing Services</span></h3></div><div class=\"h-2 bg-brutalist-yellow w-full\"></div></div><!-- Right Image --><div class=\"flex justify-center\"><img src=\"public/hero.png\" alt=\"Hero\" class=\"w-full max-w-xs sm:max-w-sm md:max-w-md lg:max-w-xl object-contain\"></div></div><!-- Bottom Stats --><div class=\"grid grid-cols-1 sm:grid-cols-3 gap-6 text-center mt-6 mb-6\"><div><div class=\"text-lg sm:text-xl uppercase text-brutalist-yellow\">Verified</div><div class=\"text-base sm:text-lg uppercase\">Professionals</div></div><div><div class=\"text-lg sm:text-xl uppercase text-brutalist-yellow\">Easy</div><div class=\"text-base sm:text-lg uppercase\">Bookings</div></div><div><div class=\"text-lg sm:text-xl uppercase text-brutalist-yellow\">Proudly</div><div class=\"text-base sm:text-lg uppercase\">Made in Assam</div></div></div></div></section><!-- External Script --><script src=\"/public/js/hero.js\" defer></script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
