package app

func (app *Application) SetupRoutes() {

	jwt := app.JWT(app.Config.JWT.Secret)
	app.Router.Post("/signup", jwt, app.Signup)
	app.Router.Post("/login", app.Login)
	app.Router.Get("/test", jwt, app.Authorize, app.Test)
	app.Router.Post("/logout", jwt, app.Authorize, app.Logout)

	//routes for one time passwords
	otp := app.Router.Group("/otp")
	otp.Post("/get", app.GetOTP)
	otp.Post("/verify", app.VerifyOTP)
}
