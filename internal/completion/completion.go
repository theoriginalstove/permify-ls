package completion

// A CompletionItem represents a possible suggestion by the algo.
type CompletionItem struct {
	// Label is the primary text the user sees for the completion item.
	Label string

	// Detail is the supplemental information to present to the user.
	// Any of the prefix that has already been typed is not trimmedl..
	Detail string
}
