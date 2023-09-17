package config

const (
	SuccessMsg             = "Success"
	AcceptedMsg            = "Accepted"
	UnauthorizedMsg        = "Unauthorized"
	IssueWithJwt           = "Issue with JWT token variable"
	ServerErrMsg           = "Internal Server Error"
	ValidationErrMsg       = "validation Error"
	PermissionDeniedErrMsg = "permission Denied"
	AccessDenied           = "access denied"
)

const (
	MissingParentId   = "parent ID is missing"
	MissingField      = "%v is missing"
	MissingContentId  = "content ID is missing"
	MissingActionType = "action Type is missing"
)

const (
	InvalidField      = "%v is invalid"
	InvalidPageLimit  = "Page Size is invalid"
	InvalidId         = "id is invalid"
	InvalidMediaType  = "media_type is invalid"
	InvalidMediaCount = "media count invalid"
)

const (
	EmptyForm = "Form is empty"
)

const (
	SuccessStatus = 0
	ErrStatus     = 1
)

const (
	InvalidParams = "invalid parameters"
	SelfRequest   = "cannot moderate your own post"
)

const (
	MissingContentType                           = "content Type is missing"
	MissingOrInvalidInteractionType              = "missing or invalid interaction type"
	MissingOrInvalidOptionId                     = "missing or invalid option id"
	MissingComment                               = "comment is missing"
	MissingParentCommentID                       = "parent comment id is missing"
	MissingOrInvalidContentUUIDOrInteractionType = "error in binding json : content_uuid and interaction_type(0-13) required"
	InteractionRequestStructMismatch             = "error in unmarshalling json : Request body mismatch with interaction struct"
	InteractionListRequestStructMismatch         = "error in unmarshalling json : Request body mismatch with interaction list request struct"
	ContentCreationRequestStructMismatch         = "error in unmarshalling json : Request body mismatch with content creation request struct"
	InvalidInteractionType                       = "invalid interaction type, must be an integer from 0 to 13"
	InteractionListNotAccessible                 = "interaction list not accessible for this interaction"
	MissingOptions                               = "options for a poll field must be atleast 2 and atmost 4"
	MissingMediaURL                              = "media_url is missing or empty string"
	MissingMediaThumbnail                        = "media_thumbnail is missing or empty string"
	InvalidChallengeGroup                        = "invalid challenge group"
	InvalidReportReason                          = "report reason missing"
	MissingCommentID                             = "comment id is missing"
	BoostLikesRequestStructMismatch              = "error in unmarshalling json : Request body mismatch with BoostLikesRequest struct"
	EnoughLikesOnContent                         = "cannot boost likes on the given content, already got enough likes"
	ActionTypeAndWebUrlAndScreenIdMisMatch       = "action_type and web_url and screen_id is missing"
	InvalidRewardId                              = "invalid reward id"
)

const (
	ContentDoesNotExist             = "content does not exist"
	LikeAlreadyExists               = "like already exists"
	LikeDoesNotExist                = "like does not exist"
	ContentIsASharedPost            = "content is already a shared post"
	UserIsContentOwner              = "user is the Content Owner"
	UserIsNotContentOwner           = "user is not the Content Owner"
	VoteAlreadyExists               = "vote already exists"
	ParentCommentDoesNotExist       = "parent comment does not exist"
	CommentDoesNotExist             = "comment does not exist"
	ReportAlreadyExists             = "report already exists"
	BookMarkAlreadyExists           = "bookmark already exists"
	BookMarkDoesNotExist            = "bookmark does not exist"
	LikeOnCommentExist              = "like on comment exist"
	UserDoesNotExist                = "user does not exist"
	LikeOnCommentDoesNotExist       = "like on commment does not exist"
	PinAlreadyExist                 = "content is already pinned"
	ConnectionDoesNotExist          = "cannot unfollow this user, connection didn't exists"
	ConnectionNotActive             = "connection not active"
	FollowerUserDoesNotExist        = "follower user does not exists"
	FollowedUserDoesNotExist        = "followed user does not exists"
	InteractionBoosterUsersNotExist = "interaction booster users does not exist"
)

const (
	MaxNumberOfUserTaggingLimitExceeded = "max number of user tagging limit exceeded"
)

const (
	NoCategory = "category not available"
	NoTemplate = "no template associated with given category id"
	NoTag      = "tags not found"
)

const (
	InvalidReportId             = "invalid report id"
	RolePermissionContradiction = "this role does not have the permission"
)

const (
	RedisKeyNotFound = "redis key nil or does not exist"
)
