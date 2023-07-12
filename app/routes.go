package app

func (app *Application) SetupRoutes() {

	jwt := app.Authenticate(app.Config.JWT.Secret)
	app.Router.Post("/signup", app.Signup)
	app.Router.Post("/login", app.Login)
	app.Router.Post("/logout", jwt, app.Logout)

	//routes for one time passwords
	otp := app.Router.Group("/otp")
	otp.Post("/get", app.GetOTP)
	otp.Post("/verify", app.VerifyOTP)
}
