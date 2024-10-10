package completion

import (
	"context"
	"strings"
	"time"

	"github.com/Permify/permify/pkg/dsl/ast"
	"github.com/Permify/permify/pkg/dsl/token"

	"github.com/theoriginalstove/permify-ls/internal/cache"
	"github.com/theoriginalstove/permify-ls/internal/file"
)

// A CompletionItem represents a possible suggestion by the algo.
type CompletionItem struct {
	// Label is the primary text the user sees for the completion item.
	Label string

	// Detail is the supplemental information to present to the user. This is often
	// the type or return type of the completion item.
	Detail string

	// InsertText is the text to insert if the item is selected. Any of the
	// prefix that has already been typed is not trimmed.
	// The insert text does not contain snippets.
	InsertText string

	Kind uint32
	Tags []uint32

	// Depth is how many levels were searched to find this completion
	Depth int

	// Score is the internal relevance score.
	// A higher score indicated this completion item is more relevant.
	Score float64
}

// completionOpts contains completion specific configuration.
type completionOpts struct {
	documentation     bool
	placeholders      bool
	snippets          bool
	postfix           bool
	budget            time.Duration
	completeRuleCalls bool
}

const (
	// standardScore is the base score for all completion items.
	standardScore float64 = 1.0
	// highScore indicates a very relevant completion item.
	highScore float64 = 10.0
	// lowScore indicates not usefule or irrelevant completion item.
	lowScore float64 = 0.01
)

// mather matches a candidate's label against the user input.
// The returned score reflects the quality of the match. A score of zero
// indicates no match, and a score of one means perfect match.
type matcher interface {
	Score(label string) (score float32)
}

// prefixMatcher implements case sensitive prefix matching
type prefixMatcher string

func (pm prefixMatcher) Score(label string) float32 {
	if strings.HasPrefix(label, string(pm)) {
		return 1
	}
	return -1
}

// insensitivePrefixMatcher implements case insensitive prefix matching.
type insensitivePrefixMatcher string

func (ipm insensitivePrefixMatcher) Score(label string) float32 {
	if strings.HasPrefix(label, string(ipm)) {
		return 1
	}
	return -1
}

// completer contains the necessary information for a single completion request.
type completer struct {
	snapshot *cache.Snapshot
	opts     *completionOpts

	// completionContext contains the information about the trigger for this
	// completion request.
	completionContext completionContext

	// fileHandle is the handle to the file associated with this completion request.
	fileHandle file.Handle

	// filename is the name of the file associated with this completion request.
	filename string

	// is the position at which the request was triggered.
	position token.PositionInfo

	// path is the path of AST Nodes enclosing the position.
	path []ast.Node

	// items is the list of completion items returned.
	items []CompletionItem

	// surrounding describes the identifier surrounding the possition.
	surrounding *Selection

	enclosingRule *ruleInfo

	// matcher matches the candidates against the surrounding prefix.
	matcher matcher

	// startTime is when we started processing the completion request. It does
	// not include any time the request spent in the queue.
	startTime time.Time
}

// ruleInfo holds info about a rule object.
type ruleInfo struct {
	// body is the rule's body.
	body *ast.Call
}

type completionContext struct {
	// triggerCharacter is the character used to trigger completion at current position, if any.
	triggerCharacter string
}

// Selection represents the cursor position and surrounding identifier.
type Selection struct {
	content  string
	position token.PositionInfo
}
