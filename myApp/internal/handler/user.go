package handler

import (
	"net/http"

	"github.com/yourusername/myapp/internal/service"
)

// UserHandler holds dependencies for user routes.
type UserHandler struct {
	Service service.UserService
}

// GetUser handles GET /users/{id}
func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user := h.Service.GetUser(id)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user))
}
