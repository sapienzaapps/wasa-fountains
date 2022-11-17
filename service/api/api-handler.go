package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/fountains/", rt.wrap(rt.listFountains))
	rt.router.POST("/fountains/", rt.wrap(rt.createFountain))
	rt.router.PUT("/fountains/:id", rt.wrap(rt.updateFountain))
	rt.router.DELETE("/fountains/:id", rt.wrap(rt.deleteFountain))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
