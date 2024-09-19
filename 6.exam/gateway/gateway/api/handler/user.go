package handler

// import (
//   "encoding/json"
//   "gateway/internal/model"
//   "gateway/pkg"
//   "gateway/proto/users"
//   "io"
//   "log"
//   "math/rand"

//   "github.com/gin-gonic/gin"
//   "google.golang.org/protobuf/encoding/protojson"
// )

// // @Router           /api/users/register [post]
// // @Summary          Register a new user
// // @Description      This method is responsible for registering new users
// // @Security         BearerAuth
// // @Tags             USERS
// // @Accept           json
// // @Produce          json
// // @Param            body    body    users.UserInfo    true  "User registration details"
// // @Success          201     {object}  map[string]int  "User registered successfully"
// // @Failure          400     {object}  error            "Invalid request body or error unmarshaling"
// // @Failure          500     {object}  error            "Error storing user data or sending confirmation email"
// // @Failure          403     {object}  error            "Permission Denied"
// func (h *Handler) RegisterUser(c *gin.Context) {
//   bytes, err := io.ReadAll(c.Request.Body)
//   if err != nil {
//     log.Println("Error reading request body in RegisterUser:", err)
//     c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
//     return
//   }

//   var userRegister users.UserInfo
//   if err := protojson.Unmarshal(bytes, &userRegister); err != nil {
//     log.Println("Error unmarshaling in RegisterUser:", err)
//     c.AbortWithStatusJSON(400, gin.H{"error": "Error unmarshaling request"})
//     return
//   }

//   code := rand.Intn(10000) + 1000
//   if err := h.Redis.AddUserIntoRedis(c.Request.Context(), &userRegister, int64(code)); err != nil {
//     log.Println("Error adding user to Redis in RegisterUser:", err)
//     c.AbortWithStatusJSON(500, gin.H{"error": "Error storing user data"})
//     return
//   }

//   if err := pkg.SendEmail(userRegister.Email, pkg.UserCode(code, userRegister.Name)); err != nil {
//     log.Println("Error sending email in RegisterUser:", err)
//     c.AbortWithStatusJSON(500, gin.H{"error": "Error sending confirmation email"})
//     return
//   }

//   c.IndentedJSON(201, gin.H{"code": code})
// }

// // @Router           /api/users/verification [post]
// // @Summary          Verify user registration code
// // @Description      This method is responsible for verifying the user registration code
// // @Security         BearerAuth
// // @Tags             USERS
// // @Accept           json
// // @Produce          json
// // @Param            body    body    model.UserCode    true  "User code details"
// // @Success          201     {object}  users.UserInfo   "User created successfully"
// // @Failure          400     {object}  error            "Error verifying code"
// // @Failure          500     {object}  error            "Error creating user or sending welcome email"
// func (h *Handler) VerifyCode(c *gin.Context) {
//   var userCode model.UserCode
//   if err := json.NewDecoder(c.Request.Body).Decode(&userCode); err != nil {
//     log.Println("Error decoding request body in VerifyCode:", err)
//     c.AbortWithStatusJSON(400, gin.H{"error": "Error decoding request"})
//     return
//   }

//   userInfo, err := h.Redis.VerifyCodeAndGetUser(c.Request.Context(), userCode)
//   if err != nil {
//     log.Println("Error verifying code in VerifyCode:", err)
//     c.AbortWithStatusJSON(400, gin.H{"error": "Error verifying code"})
//     return
//   }

//   userInfo.Email = userCode.Email
//   userWithID, err := h.User.CreateUser(c.Request.Context(), userInfo)
//   if err != nil {
//     log.Println("Error creating user in VerifyCode:", err)
//     c.AbortWithStatusJSON(500, gin.H{"error": "Error creating user"})
//     return
//   }

//   if err := pkg.SendEmail(userWithID.Email, pkg.RegisterClient(userWithID.Id, userWithID.Name)); err != nil {
//     log.Println("Error sending email in VerifyCode:", err)
//     c.AbortWithStatusJSON(500, gin.H{"error": "Error sending welcome email"})
//     return
//   }

//   c.IndentedJSON(201, userWithID)
// }

// // @Router           /api/users/login [post]
// // @Summary          User login
// // @Description      This method is responsible for user login
// // @Tags             USERS
// // @Accept           json
// // @Produce          json
// // @Param            body    body    users.UserLogin    true  "User login details"
// // @Success          200     {object}  map[string]string "Login successful with token"
// // @Failure          400     {object}  error             "Error while unmarshaling"
// // @Failure          401     {object}  error             "Invalid username or password"
// // @Failure          403     {object}  error             "Permission Denied"
// func (h *Handler) UserLogin(c *gin.Context) {
//   var userLogin users.UserLogin
//   bytes, err := io.ReadAll(c.Request.Body)
//   if err != nil {
//     log.Println("Error reading request body in UserLogin:", err)
//     c.AbortWithStatusJSON(400, gin.H{"error": "No data found in body"})
//     return
//   }

//   if err := protojson.Unmarshal(bytes, &userLogin); err != nil {
//     log.Println("Error unmarshaling in UserLogin:", err)
//     c.AbortWithStatusJSON(400, gin.H{"error": "Error while unmarshaling"})
//     return
//   }

//   user, err := h.User.LoginUser(c.Request.Context(), &userLogin)
//   if err != nil {
//     c.JSON(401, gin.H{"error": "Invalid username or password"})
//     return
//   }

//   c.JSON(200, gin.H{"token": user.Token})
// }

// // @Router           /api/users/forgot-password [post]
// // @Summary          Request password reset
// // @Description      This method is responsible for sending a password reset code to the user's email
// // @Tags             USERS
// // @Accept           json
// // @Produce          json
// // @Param            body    body    users.UserID      true  "User ID for reset"
// // @Success          200     {object}  map[string]string "Code sent to your email"
// // @Failure          400     {object}  error             "Error retrieving email"
// // @Failure          500     {object}  error             "Error storing reset code or sending email"
// // @Failure          403     {object}  error             "Permission Denied"
// func (h *Handler) ForgetPassword(c *gin.Context) {
//   bytes, err := io.ReadAll(c.Request.Body)
//   if err != nil {
//     log.Println("Error reading request body in ForgetPassword:", err)
//     c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
//     return
//   }

//   var userID users.UserID
//   if err := protojson.Unmarshal(bytes, &userID); err != nil {
//     log.Println("Error unmarshaling in ForgetPassword:", err)
//     c.AbortWithStatusJSON(400, gin.H{"error": "Error unmarshaling request"})
//     return
//   }

//   userEmail, err := h.User.ForgotPassword(c.Request.Context(), &userID)
//   if err != nil {
//     log.Println("Error retrieving email in ForgetPassword:", err)
//     c.AbortWithStatusJSON(400, gin.H{"error": "Error retrieving email"})
//     return
//   }


//   code := rand.Intn(10000) + 1000
//   if err := h.Redis.ForgetPasswordIntoRedis(c.Request.Context(), userEmail.Email, int64(code)); err != nil {
//     log.Println("Error storing password reset code in Redis:", err)
//     c.AbortWithStatusJSON(500, gin.H{"error": "Error storing reset code"})
//     return
//   }

//   if err := pkg.SendEmail(userEmail.Email, pkg.UserCode(code, userEmail.Email)); err != nil {
//     log.Println("Error sending email in ForgetPassword:", err)
//     c.AbortWithStatusJSON(500, gin.H{"error": "Error sending reset email"})
//     return
//   }

//   c.IndentedJSON(200, gin.H{"message": "Code sent to your email"})
// }

// // @Router           /api/users/update-password [post]
// // @Summary          Update user password
// // @Description      This method is responsible for updating the user's password
// // @Tags             USERS
// // @Accept           json
// // @Produce          json
// // @Param            body    body    model.NewPass     true  "New password details"
// // @Success          200     {object}  users.UserInfo   "Password updated successfully"
// // @Failure          400     {object}  error             "Invalid input or user verification error"
// // @Failure          500     {object}  error             "Error while updating password"
// // @Failure          403     {object}  error             "Permission Denied"
// func (h *Handler) UpdatePassword(c *gin.Context) {
// 	var userCode model.NewPass
// 	if err := c.ShouldBindJSON(&userCode); err != nil {
// 	  log.Println("Error binding JSON in UpdatePassword:", err)
// 	  c.IndentedJSON(400, gin.H{"error": "Invalid input"})
// 	  return
// 	}
// 	if userCode.Password == "" {
// 	  c.IndentedJSON(400, gin.H{"error": "Password is required"})
// 	  return
// 	}
  
// 	userPassword, err := h.Redis.VerifyForgetPasswordUser(c.Request.Context(), userCode)
// 	if err != nil {
// 	  log.Println("Error verifying forget password user in UpdatePassword:", err)
// 	  c.IndentedJSON(400, gin.H{"error": "Error while verifying user"})
// 	  return
// 	}
  
// 	userResponse, err := h.User.UpdatePassword(c.Request.Context(), userPassword)
// 	if err != nil {
// 	  log.Println("Error updating password in UpdatePassword:", err)
// 	  c.IndentedJSON(500, gin.H{"error": "Error while updating password"})
// 	  return
// 	}
  
// 	c.IndentedJSON(200, userResponse)
//   }
  
// // @Router           /api/users/logout [post]
// // @Summary          User logout
// // @Description      This method is responsible for logging out the user
// // @Tags             USERS
// // @Accept           json
// // @Produce          json
// // @Param            body    body    users.UserID      true  "User ID for logout"
// // @Success          200     {object}  map[string]string "Logout successful"
// // @Failure          400     {object}  error             "Error while Logginf out user "
// // @Failure          403     {object}  error             "Permission Denied"
//   func (h *Handler) UserLogout(c *gin.Context) {
// 	var userID users.UserID
// 	if err := c.BindJSON(&userID); err != nil {
// 	  log.Println("Error binding JSON in UserLogout:", err)
// 	  c.IndentedJSON(400, gin.H{"error": "Error while binding JSON"})
// 	  return
// 	}
  
// 	response, err := h.User.LogOutUser(c.Request.Context(), &userID)
// 	if err != nil {
// 	  log.Println("Error logging out user in UserLogout:", err)
// 	  c.IndentedJSON(400, gin.H{"error": "Error while logging out user"})
// 	  return
// 	}
  
// 	c.JSON(200, gin.H{"message": response.Message})
//   }
  
//   // @Router           /api/users/{id} [get]
// // @Summary          Get user by ID
// // @Description      This method retrieves a user by their ID
// // @Tags             USERS
// // @Accept           json
// // @Produce          json
// // @Param            id      path    string           true  "User ID"
// // @Success          200     {object}  users.UserInfo  "User data retrieved successfully"
// // @Failure          400     {object}  error            "Error while retrieving user"
// // @Failure          403     {object}  error             "Permission Denied"
//   func (h *Handler) GetUserByID(c *gin.Context) {
// 	var userID users.UserID
// 	if err := c.BindJSON(&userID); err != nil {
// 	  log.Println("Error binding JSON in GetUserByID:", err)
// 	  c.IndentedJSON(400, gin.H{"error": "Error while binding JSON"})
// 	  return
// 	}
  
// 	user, err := h.User.GetUserById(c.Request.Context(), &userID)
// 	if err != nil {
// 	  log.Println("Error getting user by ID in GetUserByID:", err)
// 	  c.IndentedJSON(400, gin.H{"error": "Error while retrieving user"})
// 	  return
// 	}
  
// 	c.JSON(200, gin.H{"user": user})
//   }
  