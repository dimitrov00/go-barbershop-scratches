package types

import "github.com/gofiber/fiber/v2"

type (
	UserID int
)

type (
	UserRole   string
	UserStatus string
)

type (
	Name struct {
		FirstName String50LettersOnly `json:"first_name"`
		LastName  String50LettersOnly `json:"last_name"`
	}

	ContactInfo struct {
		Phone PhoneNumber  `json:"phone_number"`
		Email EmailAddress `json:"email_address"`
	}

	User struct {
		ID        UserID          `json:"id"`
		AvatarURL *ImageURL       `json:"avatar_url,omitempty"`
		Name      Name            `json:"name"`
		Contact   ContactInfo     `json:"contact"`
		Status    UserStatus      `json:"status"`
		Roles     []UserRole      `json:"roles"`
		Password  *HashedPassword `json:"-"`
		Audit     Audit           `json:"audit"`
	}
)

type (
	ExtractUserIDFromContextFn func(*fiber.Ctx) (UserID, error)

	CreateUserFn func(User) (User, error)

	GetUserByIDFn    func(UserID) (*User, error)
	GetUserByEmailFn func(EmailAddress) (*User, error)

	CheckIfUserExistsByEmailFn        func(EmailAddress) (bool, error)
	CheckIfUserExistsByEmailOrPhoneFn func(EmailAddress, PhoneNumber) (bool, error)

	UpdateContactInfoFn func(UserID, ContactInfo, AuditChangeProcessorFn) error
	UpdatePasswordFn    func(UserID, Password, AuditChangeProcessorFn) error
	UpdateRolesFn       func(UserID, []UserRole, AuditChangeProcessorFn) error
	UpdateStatusFn      func(UserID, UserStatus, AuditChangeProcessorFn) error
	UpdateNameFn        func(UserID, Name, AuditChangeProcessorFn) error
	UpdateAvatarFn      func(UserID, ImageURL, AuditChangeProcessorFn) error

	DeleteUserFn func(UserID) error
)

type (
	UserRegistrationInput struct {
		Email string `json:"email"`
	}

	UserConfirmationInput struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
		Code      string `json:"code"`
		Password  string `json:"password"`
	}

	UserLoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserLoginOutput struct {
		Token JWTToken `json:"token"`
		User  User     `json:"user"`
	}
)
