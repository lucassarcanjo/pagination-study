package api

func (s *Server) RegisterRoutes() {
	s.Gin.GET("/users-offset", s.HandleGetUsersOffset)
	s.Gin.GET("/users-keyset", s.HandleGetUsersKeyset)
}
