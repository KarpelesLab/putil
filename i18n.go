package putil

import "context"

type i18nLookup interface {
	LookupSimple(ctx context.Context, token string, pStr []string, noFilter bool) string
}

func LookupI18N(ctx context.Context, token string, pStr []string, noFilter bool) any {
	if ctx != nil {
		i18n, ok := ctx.Value("_i18n").(i18nLookup)
		if ok {
			return i18n.LookupSimple(ctx, token, pStr, noFilter)
		}
	}

	if noFilter {
		token = "!" + token
	}
	return map[string]any{
		"@token": append([]string{token}, pStr...),
	}
}
