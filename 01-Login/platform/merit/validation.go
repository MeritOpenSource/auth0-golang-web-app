package merit

import "github.com/google/uuid"

type ValidationResult struct {
	Atom *Atom `json:"atom,omitempty"`
	// field kind ID
	FieldKindID uuid.UUID `json:"fieldKindID,omitempty"`
	Message     string    `json:"message,omitempty"`
	// The id of the template field that had the failed condition.
	TemplateFieldID uuid.NullUUID `json:"templateFieldID,omitempty"`
	// The template id where the condition originates from
	TemplateID uuid.NullUUID `json:"templateID,omitempty"`
}

type Atom struct {

	// arguments
	Arguments []string `json:"arguments"`

	// error message
	// Max Length: 2000
	ErrorMessage string `json:"errorMessage"`

	// Predicates define how to evaluate an object. Each predicate defines its own list of arguments. Each
	// predicate is presented with its arguments in the list below. The arguments must be strings, but
	// the predicate attempts to parse the arguments into the given types. A `date` format, for example, could take
	// the form `2024-05-15`. `None` means the predicate does not take arguments. `...` indicates multiple
	// arguments of the type are allowed. Arguments separated by pipes `|` indicate a choice of options. Arguments
	// in parentheses `()` are optional.
	// | Predicate | Arguments |
	// |-----------|-----------|
	// | AfterThisDate | [date] |
	// | AfterThisDatetime | [datetime] |
	// | AfterThisTimeOfDay | [time] |
	// | AfterToday | None |
	// | AfterNow | None |
	// | BeforeNow | None |
	// | BeforeNowMinusXDuration | [duration] |
	// | BeforeThisDate | [date] |
	// | BeforeThisDatetime | [datetime] |
	// | BeforeThisTimeOfDay | [time] |
	// | BeforeToday | None |
	// | BeforeTodayMinusXDays | [number] |
	// | EqualToX | [number] |
	// | EqualToDate | [date] |
	// | FieldHasValue | [any] |
	// | IsEmailDomain | None |
	// | IsFalse | None |
	// | IsTrue | None |
	// | LessThanX | [integer] |
	// | MatchesThisString | [string] |
	// | MoreThanX | [integer] |
	// | MappableTo | None |
	// | ReceivedXContainersFromTemplates | [number, UUID..., (active\|inactive\|both)]: the number of containers to match, the template IDs to match, and an optional container activeness |
	// | ReceivedContainersFromXTemplates | [number, UUID..., (active\|inactive\|both)]: the number of containers to match, the template IDs to match, and an optional container activeness |
	//
	// Enum: [AfterThisDate AfterThisDatetime AfterThisTimeOfDay AfterToday AfterNow BeforeNow BeforeNowMinusXDuration BeforeThisDate BeforeThisDatetime BeforeThisTimeOfDay BeforeToday BeforeTodayMinusXDays EqualToX EqualToDate FieldHasValue IsEmailDomain IsFalse IsTrue LessThanX MatchesThisString MoreThanX MappableTo ReceivedXContainersFromTemplates ReceivedContainersFromXTemplates]
	Predicate string `json:"predicate"`

	// target
	// Format: uuid4
	Target uuid.NullUUID `json:"target"`
}
