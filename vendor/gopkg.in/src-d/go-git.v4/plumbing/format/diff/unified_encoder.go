		return nil
		if lb > c.ctxLines {
			clb = lb - c.ctxLines + 1
		} else {
			clb = 1
		}
		clb = lb
		ctxLines := c.ctxLines
		if ctxLines > len(c.afterContext) {
			ctxLines = len(c.afterContext)
		}
		c.current.AddOp(Equal, c.afterContext[:ctxLines]...)
		c.beforeContext = c.afterContext[ctxLines:]