package userroutes

import (
    "github.com/fireteamsupport/identity/internal/restserver/middleware"
    "github.com/fireteamsupport/identity/internal/restserver/utils"
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/go-logging"
)

var log = logging.New()

type user struct {
    *restutils.Options
}

func New(g *echo.Group, opts *restutils.Options) error {
    u := user{opts}

    m := middleware.New(opts.JWTMgmt)

    me := g.Group("/me", m.AuthZDefault)
    me.GET("", u.GetME)
    me.PATCH("", u.PatchME)

    staff := g.Group("/")

    staff.GET(":id", u.GetId, m.AuthZ(2))
    staff.PATCH(":id", u.PatchId, m.AuthZ(3))

    return nil
}
