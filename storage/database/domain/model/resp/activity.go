package resp

import "time"

type DatePlaceCourtsSubsidyDescriptionPeopleLimit struct {
	Date          time.Time
	Place         string
	CourtsAndTime string
	ClubSubsidy   int16
	Description   string
	PeopleLimit   *int16
}
