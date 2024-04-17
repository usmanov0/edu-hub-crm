package customErrors

const (
	ErrUserNotFound                       = Err("user not found")
	ErrUserUpdateFailed                   = Err("failed to update user")
	ErrInvalidPassword                    = Err("invalid password")
	ErrEmptyFirstName                     = Err("empty first name")
	ErrEmptyLastName                      = Err("last name is empty")
	ErrEmptyMail                          = Err("email address is empty")
	ErrInvalidEmailFormat                 = Err("invalid email format")
	ErrBadCredentials                     = Err("fad credentials")
	ErrPhoneNumberIsInvalid               = Err("phone number is invalid")
	ErrValidationFailed                   = Err("validation failed")
	ErrFailedUpdateUserRole               = Err("failed to update user role")
	ErrIdScanFailed                       = Err("failed to scan id")
	ErrScanRows                           = Err("failed to scan rows")
	ErrFirstNameShouldBeDifferent         = Err("first name should be different from before")
	ErrLastNameShouldBeDifferent          = Err("last name should be different from before")
	ErrShouldBeDifferentPassword          = Err("password should be different from before")
	ErrShouldBeDifferentCourseName        = Err("course name should be different from before")
	ErrShouldBeDifferentCourseCapacity    = Err("capacity of students should be different from before")
	ErrShouldBeDifferentCourseDescription = Err("course description should be different from before")
	ErrFailedDeleteAccount                = Err("failed to delete user account")
	ErrInstructorAlreadyHere              = Err("instructor already here")
	ErrCourseNotFound                     = Err("course not found")
	ErrCouldNotUpdateStartDate            = Err("could not update start date")
	ErrCouldNotUpdateEndDate              = Err("could not update end date")
	ErrCouldNotUpdateInstructor           = Err("could not update instructor")
	ErrEnrollmentNotFound                 = Err("enrollment not found")
	ErrFailedExecuteQuery                 = Err("failed execute query")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
