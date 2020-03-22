// routes.go

package main

func initializeRoutes() {

	// Handle the index routes
	router.GET("/", showIndexPage)
	router.GET("/contact", showContactPage)
	router.GET("/login", showLoginPage)

	//Handle GET requests at article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)
}
