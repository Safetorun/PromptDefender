module github.com/safetorun/PromptDefender/badwords

replace (
	"github.com/safetorun/PromptDefender/embeddings" => "../embeddings"
"github.com/safetorun/PromptDefender/badwords" => "../badwords"
)

go 1.20
