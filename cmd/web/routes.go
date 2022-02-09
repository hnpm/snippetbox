package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddlewares := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddlewares := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()
	mux.Get("/", dynamicMiddlewares.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddlewares.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddlewares.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddlewares.ThenFunc(app.showSnippet))

	mux.Get("/user/signup", dynamicMiddlewares.ThenFunc(app.signupForm))
	mux.Post("/user/signup", dynamicMiddlewares.ThenFunc(app.signup))
	mux.Get("/user/login", dynamicMiddlewares.ThenFunc(app.loginForm))
	mux.Post("/user/login", dynamicMiddlewares.ThenFunc(app.login))
	mux.Post("/user/logout", dynamicMiddlewares.Append(app.requireAuthenticatedUser).ThenFunc(app.logout))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddlewares.Then(mux)
}
