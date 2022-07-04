package finfiber

var (
	PAGE_QUERY_KEY = "page"
	SIZE_QUERY_KEY = "size"

	IS_AUTHORIZED_LOCAL_KEY = "is_authorized"
	DEFAULT_JWT_KEYS        = []string{"username", "locale"}
)

// If yout do not want to use default pagination query keys ('page' & 'size')
// Customized key mappings can be set using function SetQueryKeys
func SetQueryKeys(pageKey, sizeKey string) {
	PAGE_QUERY_KEY = pageKey
	SIZE_QUERY_KEY = sizeKey
}
