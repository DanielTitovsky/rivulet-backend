package auth_transport_http

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

const googleUserInfoURL = "https://openidconnect.googleapis.com/v1/userinfo"

type GoogleUserInfoResponse struct {
	ProviderUserId    string `json:"sub"`
	ProviderUserEmail string `json:"email"`
	EmailVerified     bool   `json:"email_verified"`
	Name              string `json:"name"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
}

func (r GoogleUserInfoResponse) ToDomain() domain.OAuthUser {
	return domain.OAuthUser{
		Provider:          domain.ProviderGoogle,
		ProviderUserId:    r.ProviderUserId,
		ProviderUserEmail: r.ProviderUserEmail,
		EmailVerified:     r.EmailVerified,
		Name:              r.Name,
		GivenName:         r.GivenName,
		FamilyName:        r.FamilyName,
	}
}

func (h *AuthHttpHandler) GoogleLogin(c *gin.Context) {
	state, err := generateOAuthState()
	if err != nil {
		ctx := c.Request.Context()
		log := app_loger.FromContext(ctx)
		responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

		responseHandler.ErrorResponse(err, "Failed to generate oauth state")
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "oauthState",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   600,
	})

	redirectURL := h.googleOAuthConfig.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func (h *AuthHttpHandler) GoogleCallback(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	stateCookie, err := c.Cookie("oauthState")
	if err != nil {
		responseHandler.ErrorResponse(err, "OAuth state cookie not found")
		return
	}

	state := c.Query("state")
	if state == "" || state != stateCookie {
		responseHandler.ErrorResponse(fmt.Errorf("invalid oauth state"), "Invalid oauth state")
		return
	}

	code := c.Query("code")
	if code == "" {
		responseHandler.ErrorResponse(fmt.Errorf("oauth code is empty"), "OAuth code is empty")
		return
	}

	googleToken, err := h.googleOAuthConfig.Exchange(ctx, code)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to exchange google code")
		return
	}

	googleUser, err := h.getGoogleUser(ctx, googleToken)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get google user")
		return
	}

	user, accessToken, refreshToken, err := h.authService.OAuthLogin(ctx, googleUser.ToDomain())
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to login with google")
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "oauthState",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken.TokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   tokenCookieMaxAge(refreshToken.ExpiresAt),
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    accessToken.TokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   tokenCookieMaxAge(accessToken.ExpiresAt),
	})

	if h.frontendURL != "" {
		c.Redirect(http.StatusTemporaryRedirect, h.frontendURL)
		return
	}

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   userDTOFromDomain(user),
	})
}

func (h *AuthHttpHandler) getGoogleUser(
	ctx context.Context,
	token *oauth2.Token,
) (GoogleUserInfoResponse, error) {
	var googleUser GoogleUserInfoResponse

	client := h.googleOAuthConfig.Client(ctx, token)

	response, err := client.Get(googleUserInfoURL)
	if err != nil {
		return GoogleUserInfoResponse{}, fmt.Errorf("get google user info: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return GoogleUserInfoResponse{}, fmt.Errorf("read google user info body: %w", err)
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return GoogleUserInfoResponse{}, fmt.Errorf("google user info status: %d, body: %s", response.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &googleUser)
	if err != nil {
		return GoogleUserInfoResponse{}, fmt.Errorf("decode google user info: %w", err)
	}

	return googleUser, nil
}

func generateOAuthState() (string, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("generate random state: %w", err)
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
