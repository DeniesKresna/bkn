package gate

import (
	"strings"

	"github.com/DeniesKresna/bkn/models"
)

func (c *Gate) InitRoutes() {
	v1 := func(url string) string {
		return strings.TrimSpace("/api/v1" + url)
	}

	// user
	c.Post(v1("/user/add"), c.Protected(c.UserCreate, models.RoleNameAdmin))
	c.Put(v1("/user/edit/{id}"), c.Protected(c.UserUpdate, models.RoleNameAdmin))
	c.Post(v1("/user/login"), c.UserLogin)
	c.Post(v1("/user/create-regular"), c.Protected(c.UserRegularCreate, models.RoleNameAdmin))
	c.Post(v1("/user/register"), c.UserRegister)
	c.Get(v1("/user/detail/{id}"), c.Protected(c.UserGetDetailById, models.RoleNameAdmin))
	c.Put(v1("/user/password/{id}"), c.Protected(c.UserUpdatePassword, models.RoleNameAdmin))
	c.Put(v1("/user/password"), c.Protected(c.UserSelfUpdatePassword, models.RoleNameAdmin, models.RoleNameExpert, models.RoleNameUser))
	c.Put(v1("/user/img/{id}"), c.Protected(c.UserUpdateImg, models.RoleNameAdmin, models.RoleNameUser))
	c.Post(v1("/user/table"), c.Protected(c.UserIndex, models.RoleNameAdmin))
	c.Get(v1("/user/session"), c.Protected(c.UserGetSession))
	c.Post(v1("/user/verify"), c.UserVerifyToken)
	c.Get(v1("/user/profile"), c.Protected(c.UserProfile, models.RoleNameUser, models.RoleNameExpert))
	c.Put(v1("/user/edit"), c.Protected(c.UserSelfUpdate, models.RoleNameUser))
	c.Delete(v1("/user/delete/{id}"), c.Protected(c.UserDelete, models.RoleNameAdmin))
	c.Put(v1("/user/edit-additional-profile"), c.Protected(c.UserSelfAddOrUpdateProfile, models.RoleNameUser, models.RoleNameExpert)) // By Self
	c.Put(v1("/user/edit-additional-profile/{id}"), c.Protected(c.UserAddOrUpdateProfileByID, models.RoleNameAdmin))                  // By Admin
}
