package auth

// // Scope for read/write
// const (
// 	ScopeRead            = "read"
// 	ScopeWrite           = "write"
// 	ScopeByPassAuth      = "true"
// 	AdfsKeyCacheDuration = 6 * time.Hour
// )

// var (
// 	authenticator *jwt.Authenticator
// 	routeScope    = make(map[string]string)
// )

// // New setups auth params
// func New(authority, audience string) error {
// 	authenticator = jwt.NewAuthenticator(jwt.Config{
// 		Authority:        authority,
// 		Audience:         audience,
// 		Issuer:           authority,
// 		KeyCacheDuration: AdfsKeyCacheDuration,
// 	})
// 	return nil
// }

// // RegisterRoute stores scope for each route
// func RegisterRoute(path, method, scope string) {
// 	controller := getController(path)
// 	routeScope[controller+method] = scope
// }

// // Authenticate takes care of Authentication
// func Authenticate(path, method, bearerToken string, w http.ResponseWriter) bool {
// 	controller := getController(path)
// 	if isExcludeAuthURLs(controller + method) {
// 		return true
// 	}
// 	// Check for Token and scope
// 	if err := authenticateWithScope(bearerToken, routeScope[controller+method]); err != nil {
// 		// Authentication failed
// 		w.Header().Set("Content-Type", "application/json")
// 		errBL := model.NewBLMessageFromFormat(model.EntityAuthenticationError, err.Error())
// 		w.WriteHeader(errBL.HTTPStatus)
// 		json.NewEncoder(w).Encode(errBL)
// 		return false
// 	}
// 	return true
// }

// // AuthenticateWithScope matches the claim with the accepted scope
// func authenticateWithScope(token, scope string) error {
// 	return authenticator.AuthenticateWithScope(token, scope)
// }

// // isExcludeAuthURLs checks for URLs to bypass auth
// func isExcludeAuthURLs(reqKey string) bool {
// 	if val, ok := routeScope[reqKey]; ok && val == ScopeByPassAuth {
// 		return true
// 	}
// 	return false
// }

// // getController extract controller
// func getController(path string) string {
// 	result := strings.SplitAfterN(path, "/", 6)
// 	if len(result) == 2 {
// 		return strings.Join(result, "")
// 	} else if len(result) == 4 {
// 		return strings.Join(result, "")
// 	} else if len(result) > 4 {
// 		return strings.Join(result[:5], "")
// 	}
// 	return ""
// }
