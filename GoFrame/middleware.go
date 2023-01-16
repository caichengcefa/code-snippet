func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Response.ClearBuffer()
	r.Middleware.Next()
	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err         error
		res         interface{}
		ctx         = r.Context()
		internalErr error
	)

	res = r.GetHandlerResponse()
	err = r.GetError()
	if err != nil {
		code := gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		internalErr = r.Response.WriteJson(DefaultHandlerResponse{
			Code:    code.Code(),
			Message: err.Error(),
			Data:    nil,
		})
		if internalErr != nil {
			g.Log().Error(ctx, internalErr)
		}
		return
	}
	if _, ok := res.(*apiv1.EmptyRes); ok {
		internalErr = r.Response.WriteJson(DefaultHandlerResponse{
			Code:    200,
			Message: "OK",
			Data:    nil,
		})
	} else if d, ok := res.(*apiv1.DownFileRes); ok {
		r.Response.ClearBuffer()
		r.Response.ServeFileDownload(d.Path + d.Name)
	} else {
		internalErr = r.Response.WriteJson(DefaultHandlerResponse{
			Code:    200,
			Message: "OK",
			Data:    res,
		})
	}

	if internalErr != nil {
		g.Log().Error(ctx, internalErr)
	}
}