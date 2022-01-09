package zwadmin

import (
	"net/http"
	"time"

	trpc "git.ezbuy.me/ezbuy/corsair/digger/service/internal/rpc"
	"github.com/gin-gonic/gin"
)

func Login(in *gin.Context) {
	resp := &trpc.UserLoginResp{
		Code: 20000,
		Data: &trpc.LoginData{
			Token: "admin-token",
		},
	}

	// SetAuthCookie(in.Writer, in.Request.Host, "admin-token")
	in.SetCookie("X-Token", "admin-token", 180, "/", "localhost", false, true)
	in.JSON(http.StatusOK, resp)
}

func UserInfo(in *gin.Context) {
	resp := &trpc.UserInfoResp{
		Code: 20000,
		Data: &trpc.UserInfoData{
			Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			Introduction: "I am a super administrator",
			Name:         "Super Admin",
			Roles: []string{
				"admin",
			},
		},
	}

	// SetAuthCookie(in.Request.Response.Write(w io.Writer))

	in.JSON(http.StatusOK, resp)
}

func Logout(in *gin.Context) {
	resp := &trpc.UserLogoutResp{
		Code: 20000,
	}

	in.JSON(http.StatusOK, resp)
}

func SetAuthCookie(rw http.ResponseWriter, host, val string) {
	cookie := &http.Cookie{
		Name:     "X-Token",
		Value:    val,
		Domain:   host,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(rw, cookie)
}
