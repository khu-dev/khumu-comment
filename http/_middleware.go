package http
//
//	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
//		fmt.Println("1 Outer")
//		return func(context echo.Context) error {
//			fmt.Println("1 Inner")
//			return handlerFunc(context)
//		}
//	})
//	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
//		fmt.Println("2 Outer")
//		return func(context echo.Context) error {
//			fmt.Println("2 Inner")
//			return handlerFunc(context)
//		}
//	})
//
//	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
//		fmt.Println("3 Outer")
//		return func(context echo.Context) error {
//			fmt.Println("3 Inner")
//			return handlerFunc(context)
//		}
//	})
//
//	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
//		// 그냥 이 부분에는 아무것도 작성 안 하는
//
//		return func(context echo.Context) error {
//			fmt.Println("Do Something")
//			return handlerFunc(context)
//		}
//	})
//	middleware.CORS()
//	//Output
//	// 3 Outer
//	// 2 Outer
//	// 1 Outer
//	// 1 Inner
//	// 2 Inner
//	// 3 Inner
//	// 3 Outer
