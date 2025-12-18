package stat

import (
	"net/http"
	"shorten/configs"
	"shorten/pkg/middleware"
	"shorten/pkg/res"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHander struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHander{
		StatRepository: deps.StatRepository,
	}

	router.Handle("GET /stat", middleware.IsAuth(handler.GetStat(), deps.Config))
}

func (h *StatHander) GetStat() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		by := r.URL.Query().Get("by")

		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
			return
		}

		stats := h.StatRepository.GetStats(by, from, to)

		res.JSON(w, 200, stats)
	})
}
