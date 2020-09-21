package Controller

import (
	"github.com/rickdana/fizzbuzzApi/Model"
	"net/http"
)

type HealthCheckController struct{}

func (hcc *HealthCheckController) GetState(w http.ResponseWriter, r *http.Request) {
	state := Model.HealthStatus{Up: "OK"}

	respond(w, http.StatusOK, "application/json", state)
}
