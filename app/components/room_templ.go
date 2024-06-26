// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "strconv"
import "sort"
import "github.com/joaquinrovira/infiltra2-returns/app/util"
import "github.com/joaquinrovira/infiltra2-returns/app/model"
import "github.com/joaquinrovira/infiltra2-returns/app/routes"
import "github.com/joaquinrovira/infiltra2-returns/app/constants"

func RoomFull(user string, room *model.Room) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"contents\" hx-ext=\"sse\" sse-connect=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(routes.RoomSSESpecific(room.Name()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 12, Col: 86}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div id=\"boost-target\" class=\"contents\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(routes.RoomSpecific(room.Name()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 13, Col: 84}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"sse:RoomUpdate delay:1s\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = Room(user, room).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = ArtDecoLayout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Room(user string, room *model.Room) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grow relative flex flex-col p-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if ShouldDisplayWord(room) {
			templ_7745c5c3_Err = RoomWord(user, room).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = RoomNoWord(user, room).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if room.CountdownActive {
			templ_7745c5c3_Err = CountdownTimer(room).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = ReadyButton(user, room).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ReadyButton(user string, room *model.Room) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"contents\" hx-boost=\"true\" hx-target=\"#boost-target\" action=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 templ.SafeURL = templ.SafeURL(routes.ReadySpecific(room.Name()))
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var7)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"sticky bottom-0 p-12\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !UserReady(user, room) {
			var templ_7745c5c3_Var8 = []any{"border-8 border-double h-16 w-full transition-colors font-bold text-2xl bg-[#1b1220]"}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var8...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button type=\"submit\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var8).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" style=\"\n\t\t\t\t\t\tborder-image-slice: 1;\n\t\t\t\t\t\tborder-width: 0.25rem;\n\t\t\t\t\t  \tborder-image-source: linear-gradient(to right bottom, #8fbc6b, #a2bf6c, #b4c26f, #c5c573, #d5c879, #d5c072, #d5b76c, #d5af66, #c59a51, #b6843d, #a6702a, #965b17);\n\t\t\t\t\t\">Iniciar ronda</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			var templ_7745c5c3_Var10 = []any{"border-8 border-double h-16 w-full transition-colors font-bold text-2xl bg-[#1b1220]"}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var10...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button type=\"submit\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var10).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" style=\"\n\t\t\t\t\t\tborder-image-slice: 1;\n\t\t\t\t\t\tborder-width: 0.25rem;\n\t\t\t\t\t  \tborder-image-source: linear-gradient(to right bottom, #bc806b, #c88d6e, #d29a72, #dba876, #e3b77b, #dfb474, #dab26d, #d5af66, #c59a51, #b6843d, #a6702a, #965b17);\n\t\t\t\t\t\">Cancelar</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func RoomNoWord(user string, room *model.Room) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var12 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var12 == nil {
			templ_7745c5c3_Var12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid p-3 gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, userData := range SortedUsers(user, room) {
			templ_7745c5c3_Err = UserCard(userData.Id, userData.State, user == userData.Id).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Maybe2Grids(room *model.Room) string {
	if room.UserCount() < 4 {
		return ""
	}
	return "max-md:grid-rows-2"
}

func RoomWord(user string, room *model.Room) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var13 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var13 == nil {
			templ_7745c5c3_Var13 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var14 = []any{"py-4 px-2 gap-1 md:gap-4 grid grid-flow-col overflow-y-auto", Maybe2Grids(room)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var14...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"user-cards\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var14).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, userData := range SortedUsers(user, room) {
			templ_7745c5c3_Err = UserCardSmol(userData.Id, userData.State, user == userData.Id).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if HasSelectedWord(room) {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col grow p-8 justify-center items-center container mx-auto\"><div class=\"text-center mb-6\"><span class=\"break-words text-5xl md:text-8xl capitalize font-medium\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if selectedUser, _ := room.SelectedUser(); user == selectedUser {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("🕵️ infiltrado")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				var templ_7745c5c3_Var16 string
				templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(room.SelectedWordOrEmpty())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 92, Col: 34}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div><ol class=\"font-serif text-xl max-md:tracking-wide md:text-2xl list-decimal hyphens-auto [&amp;&gt;li]:mb-4 px-4 leading-6\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if selectedUser, _ := room.SelectedUser(); user != selectedUser {
				for _, description := range SelectedWord(room).Description {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var17 string
					templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 99, Col: 23}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li>Eres el infiltrado! <i>Disimula...</i></li><li>Te recomiendo que finjas leer esta descripción.</li><li>Si ves que los demas entran a ver la definición en la RAE, también deberías aparentarlo.</li><li>Persona introducida subrepticiamente en un grupo adversario, en territorio enemigo, etc.</li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ol><a class=\"hover:*:underline font-semibold leading-6\" href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var18 templ.SafeURL = templ.SafeURL("https://dle.rae.es/" + UserWord(user, room))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var18)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" target=\"_blank\" rel=\"noopener noreferrer\"><span class=\"inline-flex items-center justify-center\"><img class=\"h-16 w-16 md:h-20 md:w-20 object-scale-down\" src=\"/_/img/logo-rae.png\"> <span class=\"w-16\">Ver en la RAE</span></span></a></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func UserCardSmol(user string, state model.UserRoomState, isUser bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var19 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var19 == nil {
			templ_7745c5c3_Var19 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col text-center items-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var20 = []any{"rounded-full p-1 bg-gradient-to-br", BoolString(!state.Ready, "from-red-500 to-amber-400", "from-green-500 to-blue-500")}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var20...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var21 string
		templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var20).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var22 = []any{BoolString(state.Ready, "bg-amber-500/30", "bg-green-500/30"), "rounded-full flex justify-center h-12 w-12 md:h-16 md:w-16 lg:h-20 lg:w-20 xl:h-20 xl:w-20 transition-all box-border border-2 md:border-4 items-center "}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var22...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var23 string
		templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var22).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><img class=\"p-1 h-[inherit] w-[inherit] object-scale-down rounded-full\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var24 string
		templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs("/_/img/user/user_" + strconv.Itoa(util.UserIdStringToInt(user)%11) + ".png")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 127, Col: 87}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><div class=\"contents mt-1 text-sm md:text-base\"><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var25 string
		templ_7745c5c3_Var25, templ_7745c5c3_Err = templ.JoinStringErrs(util.UserName(user))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 132, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var25))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if isUser {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"text-xs md:text-sm\">Tú</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var26 = []any{"text-xs font-light opacity-50"}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var26...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var27 string
		templ_7745c5c3_Var27, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var26).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var27))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var28 string
		templ_7745c5c3_Var28, templ_7745c5c3_Err = templ.JoinStringErrs(BoolString(state.Ready, "Esperando...", ""))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 137, Col: 94}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var28))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CountdownTimer(room *model.Room) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var29 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var29 == nil {
			templ_7745c5c3_Var29 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"fixed top-0 h-full w-full bg-black/70 text-white\"><div class=\"h-dvh w-full flex justify-center items-center sticky top-0\"><div id=\"countdown\" class=\"text-9xl \">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var30 string
		templ_7745c5c3_Var30, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatFloat(constants.COUNTDOWN_TIME.Seconds(), 'f', -1, 32))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 144, Col: 111}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var30))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = countdown(int(constants.COUNTDOWN_TIME.Seconds()), "countdown").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func countdown(seconds int, targetElementId string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_countdown_c9f7`,
		Function: `function __templ_countdown_c9f7(seconds, targetElementId){var countdownElement = document.getElementById(targetElementId);
	var countdownValue = Number(seconds);
	var countdown = setInterval(function() {
		countdownValue--;
		countdownElement.textContent = countdownValue;
		if (countdownValue <= 0) {
			clearInterval(countdown);
		}
	}, 1000);
}`,
		Call:       templ.SafeScript(`__templ_countdown_c9f7`, seconds, targetElementId),
		CallInline: templ.SafeScriptInline(`__templ_countdown_c9f7`, seconds, targetElementId),
	}
}

func UserCard(user string, state model.UserRoomState, isUser bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var31 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var31 == nil {
			templ_7745c5c3_Var31 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center justify-center gap-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var32 = []any{"min-w-3 min-h-3 w-3 h-3 rounded-full", BoolString(state.Ready, "bg-[#78ba5e]", "bg-[#ba705e]")}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var32...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var33 string
		templ_7745c5c3_Var33, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var32).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var33))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><img class=\"pl-4 w-24 object-scale-down rounded-2xl\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var34 string
		templ_7745c5c3_Var34, templ_7745c5c3_Err = templ.JoinStringErrs("/_/img/user/user_" + strconv.Itoa(util.UserIdStringToInt(user)%11) + ".png")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 171, Col: 85}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var34))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <svg class=\"h-24 pl-4\" xml:space=\"preserve\" version=\"1.1\" style=\"shape-rendering:geometricPrecision;text-rendering:geometricPrecision;image-rendering:optimizeQuality;\" viewBox=\"0 0 10000 60000\" x=\"0px\" y=\"0px\" fill-rule=\"evenodd\" clip-rule=\"evenodd\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:svg=\"http://www.w3.org/2000/svg\"><g id=\"g8\" transform=\"translate(-37538.018,-12659.638)\"><path class=\"fil0\" d=\"m 42403.911,45025.259 c 24.51,-5.25 49.73,-8.2 75.34,-8.71 h 16.48 c 30.57,0.6 60.57,4.71 89.51,12.02 l 4195.35,-1678.21 1.75,-568.48 -1.76,-570.44 -4250.35,-1748.75 c -11.4,1.22 -22.91,1.95 -34.5,2.18 h -8.29 l -4257.93,1746.76 -1.75,570.25 1.75,568.67 z m 496.07,414.52 2801.31,14006.52 c 2.26,11.32 3.98,22.61 5.25,33.88 l 8.99,44.78 c 162.53,807.29 226.05,1122.86 -432.33,1742.52 -7.75,7.29 -15.72,14.2 -23.87,20.75 l -1400.96,1209.81 v 3889.23 l 2225.11,2225.11 c 156.19,156.19 156.19,409.43 0,565.62 l -3209.84,3209.84 c -156.19,156.19 -409.43,156.19 -565.62,0 l -3209.84,-3209.84 c -156.19,-156.19 -156.19,-409.43 0,-565.62 l 2169.67,-2169.67 v -3930.46 l -1549.5,-1222.9 c -9.08,-7.17 -17.73,-14.68 -26.01,-22.45 l -0.55,0.58 c -658.39,-619.66 -594.85,-935.25 -432.33,-1742.55 l 15.73,-78.33 -1.5,-0.3 2803,-14014.99 -4188.45,-1680.34 c -101.7,-40.76 -161.78,-140.99 -156.29,-244.63 l -2.23,-724.46 2.27,-737.64 c 0,-104 63.77,-193.1 154.33,-230.37 l -0.04,-0.1 4198.66,-1722.44 -2811.25,-14056.21 1.5,-0.3 -15.74,-78.34 c -162.51,-807.3 -226.04,-1122.87 432.34,-1742.53 l 0.55,0.58 c 8.28,-7.77 16.93,-15.28 26.01,-22.45 l 1549.49,-1222.89 v -3930.49 l -2169.66,-2169.66 c -156.19,-156.19 -156.19,-409.43 0,-565.62 l 3209.84,-3209.84 c 156.19,-156.19 409.43,-156.19 565.62,0 l 3209.84,3209.84 c 156.19,156.19 156.19,409.43 0,565.62 l -2225.11,2225.11 v 3889.24 l 1400.96,1209.81 c 8.15,6.55 16.12,13.46 23.87,20.75 658.38,619.66 594.86,935.22 432.33,1742.5 l -8.99,44.78 c -1.26,11.27 -2.99,22.57 -5.25,33.89 l -2807.64,14038.2 4230.17,1740.45 c 99.44,40.87 158.28,138.68 154.32,240.21 l 2.24,727.9 -2.27,737.64 c 0,104.77 -64.72,194.43 -156.35,231.19 l 0.1,0.26 z m -3750.08,-3070.94 3209.84,-1251.87 c 97.29,-38.01 200.44,-35.24 290.61,0.02 l 3209.85,1251.84 c 205.38,80.25 306.82,311.81 226.57,517.19 -42.24,108.09 -126.39,187.37 -226.58,226.55 l -3209.83,1251.88 c -97.29,38.01 -200.44,35.24 -290.61,-0.02 l -3209.85,-1251.84 c -205.38,-80.25 -306.82,-311.81 -226.57,-517.19 42.24,-108.09 126.39,-187.37 226.57,-226.56 z m 3355.15,-451.44 -2111.01,823.31 2111.01,823.31 2111.01,-823.31 z m 85.78,-22686.95 2644.22,-2644.22 -2644.22,-2644.22 -2644.22,2644.22 1321.23,1321.23 v -539.44 c 0,-220.91 179.09,-400 400,-400 220.91,0 400,179.09 400,400 v 1339.44 z m 467.54,663.7 -184.73,184.73 c -156.19,156.19 -409.43,156.19 -565.62,0 l -240.18,-240.18 v 2499.11 l 269.97,-213.07 c 147.34,-129.69 371.5,-133.91 524.07,-1.88 l 196.49,169.68 z m -1086.43,3534.04 c -33.94,39.68 -75.5,72.64 -122.42,96.61 l -1616.42,1275.71 c -330.44,313.58 -291.44,507.64 -192.39,999.71 l 15.86,78.94 1.5,-0.3 2429.42,12147.07 2429.42,-12147.07 c 2.34,-11.75 5.2,-23.26 8.5,-34.55 l 8.86,-44.12 c 99.47,-494.07 138.41,-687.68 -196.42,-1003.53 l -879.48,-759.48 v 820.85 c 0,220.91 -179.09,400 -400,400 -220.91,0 -400,-179.09 -400,-400 v -1511.7 l -469.26,-405.23 z m 1086.43,42159.08 v -2398.38 l -196.49,169.68 c -152.57,132.03 -376.73,127.81 -524.07,-1.88 l -269.96,-213.06 v 2499.08 l 240.17,-240.17 c 156.19,-156.19 409.43,-156.19 565.62,0 z m -990.52,1186.68 v 1339.45 c 0,220.91 -179.09,400 -400,400 -220.91,0 -400,-179.09 -400,-400 v -539.45 l -1321.24,1321.24 2644.22,2644.22 2644.22,-2644.22 -2644.22,-2644.22 z m -218.38,-4817.36 c 46.97,23.98 88.56,56.98 122.52,96.7 l 617.12,487.04 469.26,-405.23 v -1511.7 c 0,-220.91 179.09,-400 400,-400 220.91,0 400,179.09 400,400 v 820.85 l 879.48,-759.48 c 334.84,-315.86 295.89,-509.47 196.42,-1003.55 l -8.86,-44.1 c -3.3,-11.28 -6.16,-22.81 -8.5,-34.56 l -2429.42,-12147.07 -2429.42,12147.07 -1.5,-0.3 -15.85,78.93 c -99.06,492.08 -138.07,686.15 192.38,999.73 z\" id=\"path6\"></path></g></svg><div class=\"flex flex-col items-center justify-center\"><div class=\"ml-3 text-lg font-semibold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var35 string
		templ_7745c5c3_Var35, templ_7745c5c3_Err = templ.JoinStringErrs(util.UserName(user))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 198, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var35))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var36 string
		templ_7745c5c3_Var36, templ_7745c5c3_Err = templ.JoinStringErrs(BoolString(isUser, " (Tú)", ""))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 198, Col: 100}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var36))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var37 = []any{"ml-3 text-sm font-light opacity-50"}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var37...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var38 string
		templ_7745c5c3_Var38, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var37).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var38))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var39 string
		templ_7745c5c3_Var39, templ_7745c5c3_Err = templ.JoinStringErrs(BoolString(state.Ready, "Esperando...", ""))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/room.templ`, Line: 199, Col: 100}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var39))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func BoolString(b bool, ok string, notOk string) string {
	if b {
		return ok
	} else {
		return notOk
	}
}

func UserReady(user string, room *model.Room) bool {
	state, exists := room.User(user)
	return exists && state.Ready
}

func ShouldDisplayWord(room *model.Room) bool {
	_, userSelected := room.SelectedUser()
	_, wordSelected := room.SelectedWord()
	return userSelected && wordSelected
}

// func SortedUsers(user string, room *model.Room) []UserData {
// 	return FakeSortedUsers(user, room, 30)
// }

// func FakeSortedUsers(user string, room *model.Room, n int) []UserData {
// 	users := _SortedUsers(user, room)

// 	if len(users) > 0 {
// 		for len(users) < n {
// 			users = append(users, users[0])
// 		}
// 	}

// 	return users
// }

func SortedUsers(user string, room *model.Room) []UserData {
	userStates := room.Users()
	ids := make([]string, 0, len(userStates))
	for k := range userStates {
		ids = append(ids, k)
	}
	sort.Slice(ids, func(i, j int) bool {
		if ids[i] == user {
			return true // Move specific ID to the front
		}
		if ids[j] == user {
			return false // Move specific ID to the front
		}
		return ids[i] < ids[j] // Sort the rest
	})

	users := make([]UserData, 0, len(userStates))
	for _, id := range ids {
		users = append(users, UserData{
			Id:    id,
			State: userStates[id],
		})
	}
	return users
}

type UserData struct {
	Id    string
	State model.UserRoomState
}

func SelectedWord(room *model.Room) *model.Word {
	word, exists := room.SelectedWord()
	if !exists {
		return nil
	} else {
		return &word
	}
}
func HasSelectedWord(room *model.Room) bool {
	_, exists := room.SelectedWord()
	return exists
}

func UserWord(user string, room *model.Room) string {
	if selectedUser, _ := room.SelectedUser(); user == selectedUser {
		return "infiltrado"
	} else {
		return room.SelectedWordOrEmpty()
	}
}
