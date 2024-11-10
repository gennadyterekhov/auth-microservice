package inits

//func InitServerSuite[T interfaces.WithServer](genericSuite T, srv *httptest.Server) {
//	fmt.Println("InitServerSuite ")
//	fmt.Println()
//
//	InitDbSuite(genericSuite)
//	InitFactorySuite(genericSuite)
//
//	if srv == nil {
//		servs := services.New(genericSuite.GetRepository())
//		controllersStruct := controllers.NewControllers(servs, serializers.New())
//
//		genericSuite.SetServer(httptest.NewServer(swagrouter.NewRouter(controllersStruct).Router))
//	} else {
//		genericSuite.SetServer(srv)
//	}
//}
