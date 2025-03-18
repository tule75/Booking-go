package constant

const (
	PrePropertiesOwner    = "properties-owner-id"
	PrePropertiesId       = "properties-id"
	PrePropertiesSearch   = "properties-search"
	PreRoomByPropertiesId = "room-by-properties-id"
	PreRoomById           = "pre-room-by-id"
	PreReviewById         = "pre-review-by-id"
	PreReviewByPropertyId = "pre-review-by-propertiy-id"
	PreReviewByUserId     = "pre-review-by-user-id"
	PreBookingById        = "pre-booking-by-id"
	PreSoftLock           = "pre-soft-lock"
	PreBookingByUserId    = "pre-booking-by-user-id"

	KeyInsertAvailability = "key-insert-availability"
	KeyUpdateAvailability = "key-update-availability"
	KeyInsertBooking      = "key-insert-booking"
	RoleAdmin             = "ADMIN"
	RoleCustomer          = "CUSTOMER"
	RoleHost              = "HOST"

	DeleteBookingSuccess = "delete booking success"
	DeleteBookingFailure = "delete booking failure"
	CancelBookingSuccess = "Cancel booking success"
	CancelBookingFailure = "Cancel booking failure"

	JSONEncodeError = "JSON Encode error"
	JSONDecodeError = "JSON Decode error"
	KafkaSuccess    = "Successfully send message to Kafka"
	KafkaFailure    = "Failure sending message to Kafka"
	RoomIsLock      = "Room is locked"

	ConsumerSuccess        = "Success do task %s with consumer"
	ConsumerFailure        = "Failure do task %s with consumer"
	FailCommitKafkaMessage = "Failed to commit Kafka message"
)
