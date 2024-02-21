package handlers 

import (
    "github.com/weiseguy1/gotel/pkg/config"
    "github.com/weiseguy1/gotel/pkg/render"
    "github.com/weiseguy1/gotel/pkg/models"
    "net/http"
)


// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
    App *config.AppConfig
}

// NewRepo creates a repository
func NewRepo(a *config.AppConfig) *Repository {
    return &Repository {
        App: a,
    }
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
    Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
    remoteIP := r.RemoteAddr
    m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

    render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
    stringMap := make(map[string]string)
    stringMap["test"] = "Hello, again."

    remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
    stringMap["remote_ip"] = remoteIP

    render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
        StringMap: stringMap,
    })
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}
