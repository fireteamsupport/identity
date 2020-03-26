package userroutes

import (
    "github.com/fireteamsupport/identity/internal/restserver/middleware"
    "github.com/fireteamsupport/identity/internal/initializer"
    "github.com/arturoguerra/go-logging"
    "github.com/labstack/echo/v4"
)

var log = logging.New()

type user struct {
    *initializer.Rest
}

func New(g *echo.Group, opts *initializer.Rest) error {
    u := user{opts}

    m := middleware.New(opts.JWT)

    me := g.Group("/me", m.AuthZDefault)
    me.GET("", u.GetME)
    me.PATCH("", u.PatchME)

    staff := g.Group("/")

    staff.GET(":id", u.GetId, m.AuthZ(2))
    staff.PATCH(":id", u.PatchId, m.AuthZ(3))

    return nil
}
