package controllers

import "github.com/stepanusjanu19/goRESTAPI/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateUser))).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUsers))).Methods("GET")
	s.Router.HandleFunc("/users/{user_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUser))).Methods("GET")
	s.Router.HandleFunc("/users/{user_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{user_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteUser))).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreatePost))).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetPosts))).Methods("GET")
	s.Router.HandleFunc("/posts/{post_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetPost))).Methods("GET")
	s.Router.HandleFunc("/posts/{post_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{post_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeletePost))).Methods("DELETE")

	//Item Group routes
	s.Router.HandleFunc("/itemgroups", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateItemGroup))).Methods("POST")
	s.Router.HandleFunc("/itemgroups", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetItemGroups))).Methods("GET")
	s.Router.HandleFunc("/itemgroups/{item_group_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetItemGroup))).Methods("GET")
	s.Router.HandleFunc("/itemgroups/{item_group_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateItemGroup))).Methods("PUT")
	s.Router.HandleFunc("/itemgroups/{item_group_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteItemGroup))).Methods("DELETE")
}
