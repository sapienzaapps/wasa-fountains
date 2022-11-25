package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/locationutils"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) listFountains(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the query string part. To do that, we need to check whether the latitude, longitude and range exists.
	// If latitude and longitude are specified, we parse them, and we filter results for them. If range is specified,
	// the value will be parsed and used as a filter. If it's not specified, 10 will be used as default (as specified in
	// the OpenAPI file).
	// If one of latitude or longitude is not specified (or both), no filter will be applied.

	var err error
	var fountains []database.Fountain
	if r.URL.Query().Has("latitude") && r.URL.Query().Has("longitude") {
		var lat, lng, filterRange float64

		if r.URL.Query().Has("range") {
			filterRange, err = strconv.ParseFloat(r.URL.Query().Get("range"), 32)
			if err != nil {
				// The value is not a valid float, reject the request
				w.WriteHeader(http.StatusBadRequest)
				return
			} else if filterRange < 1 || filterRange > 200 {
				// The value is out of range, reject the request
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		} else {
			filterRange = 10
		}

		lat, err = locationutils.ParseLatitudeToFloat(r.URL.Query().Get("latitude"))
		if err != nil {
			// The latitude is not a valid float, or it's out of range, reject the request
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		lng, err = locationutils.ParseLongitudeToFloat(r.URL.Query().Get("longitude"))
		if err != nil {
			// The longitude is not a valid float, or it's out of range, reject the request
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Request a filtered list of fountains from the DB
		fountains, err = rt.db.ListFountainsWithFilter(lat, lng, filterRange)
	} else {
		// Request an unfiltered list of fountains from the DB
		fountains, err = rt.db.ListFountains()
	}
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list fountains")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(fountains)
}
