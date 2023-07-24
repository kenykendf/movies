package reason

var (
	InternalServerError = "internal server error"
	RequestFormError    = "request format is not valid"

	MoviesListsErr    = "unable to retrieve movies list"
	MoviesListByIDErr = "unable to retrieve movies information"
	MoviesUpdateErr   = "unable to update movies information"
	MoviesDeleteErr   = "unable to delete movies"
	MoviesCreateErr   = "unable to create new movies"

	RegisterFailed      = "cannot register user"
	UserAlreadyExist    = "user already exist"
	LoginFailed         = "login failed, please check your email or password"
	SaveToken           = "cannot save refresh token" // nolint:gosec
	UserNotAuthenticate = "user does not have an authentication"
	NotAuthorized       = "You are not authorized to access this resource"

	ErrInvalidToken         = "token is invalid"
	ErrNoToken              = "request does not contain an access token"
	InvalidRefreshToken     = "invalid refresh token"
	CannotCreateAccessToken = "cannot create access token"
)
