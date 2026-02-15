// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"context"
	"net/http"
	"slices"

	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/internal/requestconfig"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
)

// MapService contains methods and other services that help with interacting with
// the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMapService] method instead.
type MapService struct {
	Options []option.RequestOption
}

// NewMapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMapService(opts ...option.RequestOption) (r MapService) {
	r = MapService{}
	r.Options = opts
	return
}

// Create map task
func (r *MapService) Run(ctx context.Context, body MapRunParams, opts ...option.RequestOption) (res *MapRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/map"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response schema for map requests.
type MapRunResponse struct {
	// Array of mapped links with optional titles and descriptions.
	Links []MapRunResponseLink `json:"links,required"`
	// Indicates if the map request was successful.
	Success bool `json:"success,required"`
	// Unique identifier for the map task.
	TaskID string `json:"task_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Links       respjson.Field
		Success     respjson.Field
		TaskID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MapRunResponse) RawJSON() string { return r.JSON.raw }
func (r *MapRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MapRunResponseLink struct {
	URL         string `json:"url,required" format:"uri"`
	Description string `json:"description"`
	Title       string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Description respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MapRunResponseLink) RawJSON() string { return r.JSON.raw }
func (r *MapRunResponseLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MapRunParams struct {
	// Url to map.
	URL string `json:"url,required"`
	// Maximum number of links to return.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Country code for geolocation and proxy selection
	//
	// Any of "AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT",
	// "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ",
	// "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA",
	// "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU",
	// "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE",
	// "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB",
	// "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS",
	// "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL",
	// "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG",
	// "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI",
	// "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG",
	// "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV",
	// "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP",
	// "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN",
	// "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB",
	// "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR",
	// "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK",
	// "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US",
	// "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "XK", "YE",
	// "YT", "ZA", "ZM", "ZW".
	Country MapRunParamsCountry `json:"country,omitzero"`
	// Includes subdomains of the main domain in the mapping process.
	//
	// Any of "domain", "subdomain", "all".
	DomainFilter MapRunParamsDomainFilter `json:"domain_filter,omitzero"`
	// Locale for browser language and region settings
	//
	// Any of "aa-DJ", "aa-ER", "aa-ET", "af", "af-NA", "af-ZA", "ak", "ak-GH", "am",
	// "am-ET", "an-ES", "ar", "ar-AE", "ar-BH", "ar-DZ", "ar-EG", "ar-IN", "ar-IQ",
	// "ar-JO", "ar-KW", "ar-LB", "ar-LY", "ar-MA", "ar-OM", "ar-QA", "ar-SA", "ar-SD",
	// "ar-SY", "ar-TN", "ar-YE", "as", "as-IN", "asa", "asa-TZ", "ast-ES", "az",
	// "az-AZ", "az-Cyrl", "az-Cyrl-AZ", "az-Latn", "az-Latn-AZ", "be", "be-BY", "bem",
	// "bem-ZM", "ber-DZ", "ber-MA", "bez", "bez-TZ", "bg", "bg-BG", "bho-IN", "bm",
	// "bm-ML", "bn", "bn-BD", "bn-IN", "bo", "bo-CN", "bo-IN", "br-FR", "brx-IN",
	// "bs", "bs-BA", "byn-ER", "ca", "ca-AD", "ca-ES", "ca-FR", "ca-IT", "cgg",
	// "cgg-UG", "chr", "chr-US", "crh-UA", "cs", "cs-CZ", "csb-PL", "cv-RU", "cy",
	// "cy-GB", "da", "da-DK", "dav", "dav-KE", "de", "de-AT", "de-BE", "de-CH",
	// "de-DE", "de-LI", "de-LU", "dv-MV", "dz-BT", "ebu", "ebu-KE", "ee", "ee-GH",
	// "ee-TG", "el", "el-CY", "el-GR", "en", "en-AG", "en-AS", "en-AU", "en-BE",
	// "en-BW", "en-BZ", "en-CA", "en-DK", "en-GB", "en-GU", "en-HK", "en-IE", "en-IN",
	// "en-JM", "en-MH", "en-MP", "en-MT", "en-MU", "en-NA", "en-NG", "en-NZ", "en-PH",
	// "en-PK", "en-SG", "en-TT", "en-UM", "en-US", "en-VI", "en-ZA", "en-ZM", "en-ZW",
	// "eo", "es", "es-419", "es-AR", "es-BO", "es-CL", "es-CO", "es-CR", "es-CU",
	// "es-DO", "es-EC", "es-ES", "es-GQ", "es-GT", "es-HN", "es-MX", "es-NI", "es-PA",
	// "es-PE", "es-PR", "es-PY", "es-SV", "es-US", "es-UY", "es-VE", "et", "et-EE",
	// "eu", "eu-ES", "fa", "fa-AF", "fa-IR", "ff", "ff-SN", "fi", "fi-FI", "fil",
	// "fil-PH", "fo", "fo-FO", "fr", "fr-BE", "fr-BF", "fr-BI", "fr-BJ", "fr-BL",
	// "fr-CA", "fr-CD", "fr-CF", "fr-CG", "fr-CH", "fr-CI", "fr-CM", "fr-DJ", "fr-FR",
	// "fr-GA", "fr-GN", "fr-GP", "fr-GQ", "fr-KM", "fr-LU", "fr-MC", "fr-MF", "fr-MG",
	// "fr-ML", "fr-MQ", "fr-NE", "fr-RE", "fr-RW", "fr-SN", "fr-TD", "fr-TG",
	// "fur-IT", "fy-DE", "fy-NL", "ga", "ga-IE", "gd-GB", "gez-ER", "gez-ET", "gl",
	// "gl-ES", "gsw", "gsw-CH", "gu", "gu-IN", "guz", "guz-KE", "gv", "gv-GB", "ha",
	// "ha-Latn", "ha-Latn-GH", "ha-Latn-NE", "ha-Latn-NG", "ha-NG", "haw", "haw-US",
	// "he", "he-IL", "hi", "hi-IN", "hne-IN", "hr", "hr-HR", "hsb-DE", "ht-HT", "hu",
	// "hu-HU", "hy", "hy-AM", "id", "id-ID", "ig", "ig-NG", "ii", "ii-CN", "ik-CA",
	// "is", "is-IS", "it", "it-CH", "it-IT", "iu-CA", "iw-IL", "ja", "ja-JP", "jmc",
	// "jmc-TZ", "ka", "ka-GE", "kab", "kab-DZ", "kam", "kam-KE", "kde", "kde-TZ",
	// "kea", "kea-CV", "khq", "khq-ML", "ki", "ki-KE", "kk", "kk-Cyrl", "kk-Cyrl-KZ",
	// "kk-KZ", "kl", "kl-GL", "kln", "kln-KE", "km", "km-KH", "kn", "kn-IN", "ko",
	// "ko-KR", "kok", "kok-IN", "ks-IN", "ku-TR", "kw", "kw-GB", "ky-KG", "lag",
	// "lag-TZ", "lb-LU", "lg", "lg-UG", "li-BE", "li-NL", "lij-IT", "lo-LA", "lt",
	// "lt-LT", "luo", "luo-KE", "luy", "luy-KE", "lv", "lv-LV", "mag-IN", "mai-IN",
	// "mas", "mas-KE", "mas-TZ", "mer", "mer-KE", "mfe", "mfe-MU", "mg", "mg-MG",
	// "mhr-RU", "mi-NZ", "mk", "mk-MK", "ml", "ml-IN", "mn-MN", "mr", "mr-IN", "ms",
	// "ms-BN", "ms-MY", "mt", "mt-MT", "my", "my-MM", "nan-TW", "naq", "naq-NA", "nb",
	// "nb-NO", "nd", "nd-ZW", "nds-DE", "nds-NL", "ne", "ne-IN", "ne-NP", "nl",
	// "nl-AW", "nl-BE", "nl-NL", "nn", "nn-NO", "nr-ZA", "nso-ZA", "nyn", "nyn-UG",
	// "oc-FR", "om", "om-ET", "om-KE", "or", "or-IN", "os-RU", "pa", "pa-Arab",
	// "pa-Arab-PK", "pa-Guru", "pa-Guru-IN", "pa-IN", "pa-PK", "pap-AN", "pl",
	// "pl-PL", "ps", "ps-AF", "pt", "pt-BR", "pt-GW", "pt-MZ", "pt-PT", "rm", "rm-CH",
	// "ro", "ro-MD", "ro-RO", "rof", "rof-TZ", "ru", "ru-MD", "ru-RU", "ru-UA", "rw",
	// "rw-RW", "rwk", "rwk-TZ", "sa-IN", "saq", "saq-KE", "sc-IT", "sd-IN", "se-NO",
	// "seh", "seh-MZ", "ses", "ses-ML", "sg", "sg-CF", "shi", "shi-Latn",
	// "shi-Latn-MA", "shi-Tfng", "shi-Tfng-MA", "shs-CA", "si", "si-LK", "sid-ET",
	// "sk", "sk-SK", "sl", "sl-SI", "sn", "sn-ZW", "so", "so-DJ", "so-ET", "so-KE",
	// "so-SO", "sq", "sq-AL", "sq-MK", "sr", "sr-Cyrl", "sr-Cyrl-BA", "sr-Cyrl-ME",
	// "sr-Cyrl-RS", "sr-Latn", "sr-Latn-BA", "sr-Latn-ME", "sr-Latn-RS", "sr-ME",
	// "sr-RS", "ss-ZA", "st-ZA", "sv", "sv-FI", "sv-SE", "sw", "sw-KE", "sw-TZ", "ta",
	// "ta-IN", "ta-LK", "te", "te-IN", "teo", "teo-KE", "teo-UG", "tg-TJ", "th",
	// "th-TH", "ti", "ti-ER", "ti-ET", "tig-ER", "tk-TM", "tl-PH", "tn-ZA", "to",
	// "to-TO", "tr", "tr-CY", "tr-TR", "ts-ZA", "tt-RU", "tzm", "tzm-Latn",
	// "tzm-Latn-MA", "ug-CN", "uk", "uk-UA", "unm-US", "ur", "ur-IN", "ur-PK", "uz",
	// "uz-Arab", "uz-Arab-AF", "uz-Cyrl", "uz-Cyrl-UZ", "uz-Latn", "uz-Latn-UZ",
	// "uz-UZ", "ve-ZA", "vi", "vi-VN", "vun", "vun-TZ", "wa-BE", "wae-CH", "wal-ET",
	// "wo-SN", "xh-ZA", "xog", "xog-UG", "yi-US", "yo", "yo-NG", "yue-HK", "zh",
	// "zh-CN", "zh-HK", "zh-Hans", "zh-Hans-CN", "zh-Hans-HK", "zh-Hans-MO",
	// "zh-Hans-SG", "zh-Hant", "zh-Hant-HK", "zh-Hant-MO", "zh-Hant-TW", "zh-SG",
	// "zh-TW", "zu", "zu-ZA", "auto".
	Locale MapRunParamsLocale `json:"locale,omitzero"`
	// Sitemap and other methods will be used together to find URLs.
	//
	// Any of "skip", "include", "only".
	Sitemap MapRunParamsSitemap `json:"sitemap,omitzero"`
	paramObj
}

func (r MapRunParams) MarshalJSON() (data []byte, err error) {
	type shadow MapRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MapRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Country code for geolocation and proxy selection
type MapRunParamsCountry string

const (
	MapRunParamsCountryAd MapRunParamsCountry = "AD"
	MapRunParamsCountryAe MapRunParamsCountry = "AE"
	MapRunParamsCountryAf MapRunParamsCountry = "AF"
	MapRunParamsCountryAg MapRunParamsCountry = "AG"
	MapRunParamsCountryAI MapRunParamsCountry = "AI"
	MapRunParamsCountryAl MapRunParamsCountry = "AL"
	MapRunParamsCountryAm MapRunParamsCountry = "AM"
	MapRunParamsCountryAo MapRunParamsCountry = "AO"
	MapRunParamsCountryAq MapRunParamsCountry = "AQ"
	MapRunParamsCountryAr MapRunParamsCountry = "AR"
	MapRunParamsCountryAs MapRunParamsCountry = "AS"
	MapRunParamsCountryAt MapRunParamsCountry = "AT"
	MapRunParamsCountryAu MapRunParamsCountry = "AU"
	MapRunParamsCountryAw MapRunParamsCountry = "AW"
	MapRunParamsCountryAx MapRunParamsCountry = "AX"
	MapRunParamsCountryAz MapRunParamsCountry = "AZ"
	MapRunParamsCountryBa MapRunParamsCountry = "BA"
	MapRunParamsCountryBb MapRunParamsCountry = "BB"
	MapRunParamsCountryBd MapRunParamsCountry = "BD"
	MapRunParamsCountryBe MapRunParamsCountry = "BE"
	MapRunParamsCountryBf MapRunParamsCountry = "BF"
	MapRunParamsCountryBg MapRunParamsCountry = "BG"
	MapRunParamsCountryBh MapRunParamsCountry = "BH"
	MapRunParamsCountryBi MapRunParamsCountry = "BI"
	MapRunParamsCountryBj MapRunParamsCountry = "BJ"
	MapRunParamsCountryBl MapRunParamsCountry = "BL"
	MapRunParamsCountryBm MapRunParamsCountry = "BM"
	MapRunParamsCountryBn MapRunParamsCountry = "BN"
	MapRunParamsCountryBo MapRunParamsCountry = "BO"
	MapRunParamsCountryBq MapRunParamsCountry = "BQ"
	MapRunParamsCountryBr MapRunParamsCountry = "BR"
	MapRunParamsCountryBs MapRunParamsCountry = "BS"
	MapRunParamsCountryBt MapRunParamsCountry = "BT"
	MapRunParamsCountryBv MapRunParamsCountry = "BV"
	MapRunParamsCountryBw MapRunParamsCountry = "BW"
	MapRunParamsCountryBy MapRunParamsCountry = "BY"
	MapRunParamsCountryBz MapRunParamsCountry = "BZ"
	MapRunParamsCountryCa MapRunParamsCountry = "CA"
	MapRunParamsCountryCc MapRunParamsCountry = "CC"
	MapRunParamsCountryCd MapRunParamsCountry = "CD"
	MapRunParamsCountryCf MapRunParamsCountry = "CF"
	MapRunParamsCountryCg MapRunParamsCountry = "CG"
	MapRunParamsCountryCh MapRunParamsCountry = "CH"
	MapRunParamsCountryCi MapRunParamsCountry = "CI"
	MapRunParamsCountryCk MapRunParamsCountry = "CK"
	MapRunParamsCountryCl MapRunParamsCountry = "CL"
	MapRunParamsCountryCm MapRunParamsCountry = "CM"
	MapRunParamsCountryCn MapRunParamsCountry = "CN"
	MapRunParamsCountryCo MapRunParamsCountry = "CO"
	MapRunParamsCountryCr MapRunParamsCountry = "CR"
	MapRunParamsCountryCu MapRunParamsCountry = "CU"
	MapRunParamsCountryCv MapRunParamsCountry = "CV"
	MapRunParamsCountryCw MapRunParamsCountry = "CW"
	MapRunParamsCountryCx MapRunParamsCountry = "CX"
	MapRunParamsCountryCy MapRunParamsCountry = "CY"
	MapRunParamsCountryCz MapRunParamsCountry = "CZ"
	MapRunParamsCountryDe MapRunParamsCountry = "DE"
	MapRunParamsCountryDj MapRunParamsCountry = "DJ"
	MapRunParamsCountryDk MapRunParamsCountry = "DK"
	MapRunParamsCountryDm MapRunParamsCountry = "DM"
	MapRunParamsCountryDo MapRunParamsCountry = "DO"
	MapRunParamsCountryDz MapRunParamsCountry = "DZ"
	MapRunParamsCountryEc MapRunParamsCountry = "EC"
	MapRunParamsCountryEe MapRunParamsCountry = "EE"
	MapRunParamsCountryEg MapRunParamsCountry = "EG"
	MapRunParamsCountryEh MapRunParamsCountry = "EH"
	MapRunParamsCountryEr MapRunParamsCountry = "ER"
	MapRunParamsCountryEs MapRunParamsCountry = "ES"
	MapRunParamsCountryEt MapRunParamsCountry = "ET"
	MapRunParamsCountryFi MapRunParamsCountry = "FI"
	MapRunParamsCountryFj MapRunParamsCountry = "FJ"
	MapRunParamsCountryFk MapRunParamsCountry = "FK"
	MapRunParamsCountryFm MapRunParamsCountry = "FM"
	MapRunParamsCountryFo MapRunParamsCountry = "FO"
	MapRunParamsCountryFr MapRunParamsCountry = "FR"
	MapRunParamsCountryGa MapRunParamsCountry = "GA"
	MapRunParamsCountryGB MapRunParamsCountry = "GB"
	MapRunParamsCountryGd MapRunParamsCountry = "GD"
	MapRunParamsCountryGe MapRunParamsCountry = "GE"
	MapRunParamsCountryGf MapRunParamsCountry = "GF"
	MapRunParamsCountryGg MapRunParamsCountry = "GG"
	MapRunParamsCountryGh MapRunParamsCountry = "GH"
	MapRunParamsCountryGi MapRunParamsCountry = "GI"
	MapRunParamsCountryGl MapRunParamsCountry = "GL"
	MapRunParamsCountryGm MapRunParamsCountry = "GM"
	MapRunParamsCountryGn MapRunParamsCountry = "GN"
	MapRunParamsCountryGp MapRunParamsCountry = "GP"
	MapRunParamsCountryGq MapRunParamsCountry = "GQ"
	MapRunParamsCountryGr MapRunParamsCountry = "GR"
	MapRunParamsCountryGs MapRunParamsCountry = "GS"
	MapRunParamsCountryGt MapRunParamsCountry = "GT"
	MapRunParamsCountryGu MapRunParamsCountry = "GU"
	MapRunParamsCountryGw MapRunParamsCountry = "GW"
	MapRunParamsCountryGy MapRunParamsCountry = "GY"
	MapRunParamsCountryHk MapRunParamsCountry = "HK"
	MapRunParamsCountryHm MapRunParamsCountry = "HM"
	MapRunParamsCountryHn MapRunParamsCountry = "HN"
	MapRunParamsCountryHr MapRunParamsCountry = "HR"
	MapRunParamsCountryHt MapRunParamsCountry = "HT"
	MapRunParamsCountryHu MapRunParamsCountry = "HU"
	MapRunParamsCountryID MapRunParamsCountry = "ID"
	MapRunParamsCountryIe MapRunParamsCountry = "IE"
	MapRunParamsCountryIl MapRunParamsCountry = "IL"
	MapRunParamsCountryIm MapRunParamsCountry = "IM"
	MapRunParamsCountryIn MapRunParamsCountry = "IN"
	MapRunParamsCountryIo MapRunParamsCountry = "IO"
	MapRunParamsCountryIq MapRunParamsCountry = "IQ"
	MapRunParamsCountryIr MapRunParamsCountry = "IR"
	MapRunParamsCountryIs MapRunParamsCountry = "IS"
	MapRunParamsCountryIt MapRunParamsCountry = "IT"
	MapRunParamsCountryJe MapRunParamsCountry = "JE"
	MapRunParamsCountryJm MapRunParamsCountry = "JM"
	MapRunParamsCountryJo MapRunParamsCountry = "JO"
	MapRunParamsCountryJp MapRunParamsCountry = "JP"
	MapRunParamsCountryKe MapRunParamsCountry = "KE"
	MapRunParamsCountryKg MapRunParamsCountry = "KG"
	MapRunParamsCountryKh MapRunParamsCountry = "KH"
	MapRunParamsCountryKi MapRunParamsCountry = "KI"
	MapRunParamsCountryKm MapRunParamsCountry = "KM"
	MapRunParamsCountryKn MapRunParamsCountry = "KN"
	MapRunParamsCountryKp MapRunParamsCountry = "KP"
	MapRunParamsCountryKr MapRunParamsCountry = "KR"
	MapRunParamsCountryKw MapRunParamsCountry = "KW"
	MapRunParamsCountryKy MapRunParamsCountry = "KY"
	MapRunParamsCountryKz MapRunParamsCountry = "KZ"
	MapRunParamsCountryLa MapRunParamsCountry = "LA"
	MapRunParamsCountryLb MapRunParamsCountry = "LB"
	MapRunParamsCountryLc MapRunParamsCountry = "LC"
	MapRunParamsCountryLi MapRunParamsCountry = "LI"
	MapRunParamsCountryLk MapRunParamsCountry = "LK"
	MapRunParamsCountryLr MapRunParamsCountry = "LR"
	MapRunParamsCountryLs MapRunParamsCountry = "LS"
	MapRunParamsCountryLt MapRunParamsCountry = "LT"
	MapRunParamsCountryLu MapRunParamsCountry = "LU"
	MapRunParamsCountryLv MapRunParamsCountry = "LV"
	MapRunParamsCountryLy MapRunParamsCountry = "LY"
	MapRunParamsCountryMa MapRunParamsCountry = "MA"
	MapRunParamsCountryMc MapRunParamsCountry = "MC"
	MapRunParamsCountryMd MapRunParamsCountry = "MD"
	MapRunParamsCountryMe MapRunParamsCountry = "ME"
	MapRunParamsCountryMf MapRunParamsCountry = "MF"
	MapRunParamsCountryMg MapRunParamsCountry = "MG"
	MapRunParamsCountryMh MapRunParamsCountry = "MH"
	MapRunParamsCountryMk MapRunParamsCountry = "MK"
	MapRunParamsCountryMl MapRunParamsCountry = "ML"
	MapRunParamsCountryMm MapRunParamsCountry = "MM"
	MapRunParamsCountryMn MapRunParamsCountry = "MN"
	MapRunParamsCountryMo MapRunParamsCountry = "MO"
	MapRunParamsCountryMp MapRunParamsCountry = "MP"
	MapRunParamsCountryMq MapRunParamsCountry = "MQ"
	MapRunParamsCountryMr MapRunParamsCountry = "MR"
	MapRunParamsCountryMs MapRunParamsCountry = "MS"
	MapRunParamsCountryMt MapRunParamsCountry = "MT"
	MapRunParamsCountryMu MapRunParamsCountry = "MU"
	MapRunParamsCountryMv MapRunParamsCountry = "MV"
	MapRunParamsCountryMw MapRunParamsCountry = "MW"
	MapRunParamsCountryMx MapRunParamsCountry = "MX"
	MapRunParamsCountryMy MapRunParamsCountry = "MY"
	MapRunParamsCountryMz MapRunParamsCountry = "MZ"
	MapRunParamsCountryNa MapRunParamsCountry = "NA"
	MapRunParamsCountryNc MapRunParamsCountry = "NC"
	MapRunParamsCountryNe MapRunParamsCountry = "NE"
	MapRunParamsCountryNf MapRunParamsCountry = "NF"
	MapRunParamsCountryNg MapRunParamsCountry = "NG"
	MapRunParamsCountryNi MapRunParamsCountry = "NI"
	MapRunParamsCountryNl MapRunParamsCountry = "NL"
	MapRunParamsCountryNo MapRunParamsCountry = "NO"
	MapRunParamsCountryNp MapRunParamsCountry = "NP"
	MapRunParamsCountryNr MapRunParamsCountry = "NR"
	MapRunParamsCountryNu MapRunParamsCountry = "NU"
	MapRunParamsCountryNz MapRunParamsCountry = "NZ"
	MapRunParamsCountryOm MapRunParamsCountry = "OM"
	MapRunParamsCountryPa MapRunParamsCountry = "PA"
	MapRunParamsCountryPe MapRunParamsCountry = "PE"
	MapRunParamsCountryPf MapRunParamsCountry = "PF"
	MapRunParamsCountryPg MapRunParamsCountry = "PG"
	MapRunParamsCountryPh MapRunParamsCountry = "PH"
	MapRunParamsCountryPk MapRunParamsCountry = "PK"
	MapRunParamsCountryPl MapRunParamsCountry = "PL"
	MapRunParamsCountryPm MapRunParamsCountry = "PM"
	MapRunParamsCountryPn MapRunParamsCountry = "PN"
	MapRunParamsCountryPr MapRunParamsCountry = "PR"
	MapRunParamsCountryPs MapRunParamsCountry = "PS"
	MapRunParamsCountryPt MapRunParamsCountry = "PT"
	MapRunParamsCountryPw MapRunParamsCountry = "PW"
	MapRunParamsCountryPy MapRunParamsCountry = "PY"
	MapRunParamsCountryQa MapRunParamsCountry = "QA"
	MapRunParamsCountryRe MapRunParamsCountry = "RE"
	MapRunParamsCountryRo MapRunParamsCountry = "RO"
	MapRunParamsCountryRs MapRunParamsCountry = "RS"
	MapRunParamsCountryRu MapRunParamsCountry = "RU"
	MapRunParamsCountryRw MapRunParamsCountry = "RW"
	MapRunParamsCountrySa MapRunParamsCountry = "SA"
	MapRunParamsCountrySb MapRunParamsCountry = "SB"
	MapRunParamsCountrySc MapRunParamsCountry = "SC"
	MapRunParamsCountrySd MapRunParamsCountry = "SD"
	MapRunParamsCountrySe MapRunParamsCountry = "SE"
	MapRunParamsCountrySg MapRunParamsCountry = "SG"
	MapRunParamsCountrySh MapRunParamsCountry = "SH"
	MapRunParamsCountrySi MapRunParamsCountry = "SI"
	MapRunParamsCountrySj MapRunParamsCountry = "SJ"
	MapRunParamsCountrySk MapRunParamsCountry = "SK"
	MapRunParamsCountrySl MapRunParamsCountry = "SL"
	MapRunParamsCountrySm MapRunParamsCountry = "SM"
	MapRunParamsCountrySn MapRunParamsCountry = "SN"
	MapRunParamsCountrySo MapRunParamsCountry = "SO"
	MapRunParamsCountrySr MapRunParamsCountry = "SR"
	MapRunParamsCountrySS MapRunParamsCountry = "SS"
	MapRunParamsCountrySt MapRunParamsCountry = "ST"
	MapRunParamsCountrySv MapRunParamsCountry = "SV"
	MapRunParamsCountrySx MapRunParamsCountry = "SX"
	MapRunParamsCountrySy MapRunParamsCountry = "SY"
	MapRunParamsCountrySz MapRunParamsCountry = "SZ"
	MapRunParamsCountryTc MapRunParamsCountry = "TC"
	MapRunParamsCountryTd MapRunParamsCountry = "TD"
	MapRunParamsCountryTf MapRunParamsCountry = "TF"
	MapRunParamsCountryTg MapRunParamsCountry = "TG"
	MapRunParamsCountryTh MapRunParamsCountry = "TH"
	MapRunParamsCountryTj MapRunParamsCountry = "TJ"
	MapRunParamsCountryTk MapRunParamsCountry = "TK"
	MapRunParamsCountryTl MapRunParamsCountry = "TL"
	MapRunParamsCountryTm MapRunParamsCountry = "TM"
	MapRunParamsCountryTn MapRunParamsCountry = "TN"
	MapRunParamsCountryTo MapRunParamsCountry = "TO"
	MapRunParamsCountryTr MapRunParamsCountry = "TR"
	MapRunParamsCountryTt MapRunParamsCountry = "TT"
	MapRunParamsCountryTv MapRunParamsCountry = "TV"
	MapRunParamsCountryTw MapRunParamsCountry = "TW"
	MapRunParamsCountryTz MapRunParamsCountry = "TZ"
	MapRunParamsCountryUa MapRunParamsCountry = "UA"
	MapRunParamsCountryUg MapRunParamsCountry = "UG"
	MapRunParamsCountryUm MapRunParamsCountry = "UM"
	MapRunParamsCountryUs MapRunParamsCountry = "US"
	MapRunParamsCountryUy MapRunParamsCountry = "UY"
	MapRunParamsCountryUz MapRunParamsCountry = "UZ"
	MapRunParamsCountryVa MapRunParamsCountry = "VA"
	MapRunParamsCountryVc MapRunParamsCountry = "VC"
	MapRunParamsCountryVe MapRunParamsCountry = "VE"
	MapRunParamsCountryVg MapRunParamsCountry = "VG"
	MapRunParamsCountryVi MapRunParamsCountry = "VI"
	MapRunParamsCountryVn MapRunParamsCountry = "VN"
	MapRunParamsCountryVu MapRunParamsCountry = "VU"
	MapRunParamsCountryWf MapRunParamsCountry = "WF"
	MapRunParamsCountryWs MapRunParamsCountry = "WS"
	MapRunParamsCountryXk MapRunParamsCountry = "XK"
	MapRunParamsCountryYe MapRunParamsCountry = "YE"
	MapRunParamsCountryYt MapRunParamsCountry = "YT"
	MapRunParamsCountryZa MapRunParamsCountry = "ZA"
	MapRunParamsCountryZm MapRunParamsCountry = "ZM"
	MapRunParamsCountryZw MapRunParamsCountry = "ZW"
)

// Includes subdomains of the main domain in the mapping process.
type MapRunParamsDomainFilter string

const (
	MapRunParamsDomainFilterDomain    MapRunParamsDomainFilter = "domain"
	MapRunParamsDomainFilterSubdomain MapRunParamsDomainFilter = "subdomain"
	MapRunParamsDomainFilterAll       MapRunParamsDomainFilter = "all"
)

// Locale for browser language and region settings
type MapRunParamsLocale string

const (
	MapRunParamsLocaleAaDj      MapRunParamsLocale = "aa-DJ"
	MapRunParamsLocaleAaEr      MapRunParamsLocale = "aa-ER"
	MapRunParamsLocaleAaEt      MapRunParamsLocale = "aa-ET"
	MapRunParamsLocaleAf        MapRunParamsLocale = "af"
	MapRunParamsLocaleAfNa      MapRunParamsLocale = "af-NA"
	MapRunParamsLocaleAfZa      MapRunParamsLocale = "af-ZA"
	MapRunParamsLocaleAk        MapRunParamsLocale = "ak"
	MapRunParamsLocaleAkGh      MapRunParamsLocale = "ak-GH"
	MapRunParamsLocaleAm        MapRunParamsLocale = "am"
	MapRunParamsLocaleAmEt      MapRunParamsLocale = "am-ET"
	MapRunParamsLocaleAnEs      MapRunParamsLocale = "an-ES"
	MapRunParamsLocaleAr        MapRunParamsLocale = "ar"
	MapRunParamsLocaleArAe      MapRunParamsLocale = "ar-AE"
	MapRunParamsLocaleArBh      MapRunParamsLocale = "ar-BH"
	MapRunParamsLocaleArDz      MapRunParamsLocale = "ar-DZ"
	MapRunParamsLocaleArEg      MapRunParamsLocale = "ar-EG"
	MapRunParamsLocaleArIn      MapRunParamsLocale = "ar-IN"
	MapRunParamsLocaleArIq      MapRunParamsLocale = "ar-IQ"
	MapRunParamsLocaleArJo      MapRunParamsLocale = "ar-JO"
	MapRunParamsLocaleArKw      MapRunParamsLocale = "ar-KW"
	MapRunParamsLocaleArLb      MapRunParamsLocale = "ar-LB"
	MapRunParamsLocaleArLy      MapRunParamsLocale = "ar-LY"
	MapRunParamsLocaleArMa      MapRunParamsLocale = "ar-MA"
	MapRunParamsLocaleArOm      MapRunParamsLocale = "ar-OM"
	MapRunParamsLocaleArQa      MapRunParamsLocale = "ar-QA"
	MapRunParamsLocaleArSa      MapRunParamsLocale = "ar-SA"
	MapRunParamsLocaleArSd      MapRunParamsLocale = "ar-SD"
	MapRunParamsLocaleArSy      MapRunParamsLocale = "ar-SY"
	MapRunParamsLocaleArTn      MapRunParamsLocale = "ar-TN"
	MapRunParamsLocaleArYe      MapRunParamsLocale = "ar-YE"
	MapRunParamsLocaleAs        MapRunParamsLocale = "as"
	MapRunParamsLocaleAsIn      MapRunParamsLocale = "as-IN"
	MapRunParamsLocaleAsa       MapRunParamsLocale = "asa"
	MapRunParamsLocaleAsaTz     MapRunParamsLocale = "asa-TZ"
	MapRunParamsLocaleAstEs     MapRunParamsLocale = "ast-ES"
	MapRunParamsLocaleAz        MapRunParamsLocale = "az"
	MapRunParamsLocaleAzAz      MapRunParamsLocale = "az-AZ"
	MapRunParamsLocaleAzCyrl    MapRunParamsLocale = "az-Cyrl"
	MapRunParamsLocaleAzCyrlAz  MapRunParamsLocale = "az-Cyrl-AZ"
	MapRunParamsLocaleAzLatn    MapRunParamsLocale = "az-Latn"
	MapRunParamsLocaleAzLatnAz  MapRunParamsLocale = "az-Latn-AZ"
	MapRunParamsLocaleBe        MapRunParamsLocale = "be"
	MapRunParamsLocaleBeBy      MapRunParamsLocale = "be-BY"
	MapRunParamsLocaleBem       MapRunParamsLocale = "bem"
	MapRunParamsLocaleBemZm     MapRunParamsLocale = "bem-ZM"
	MapRunParamsLocaleBerDz     MapRunParamsLocale = "ber-DZ"
	MapRunParamsLocaleBerMa     MapRunParamsLocale = "ber-MA"
	MapRunParamsLocaleBez       MapRunParamsLocale = "bez"
	MapRunParamsLocaleBezTz     MapRunParamsLocale = "bez-TZ"
	MapRunParamsLocaleBg        MapRunParamsLocale = "bg"
	MapRunParamsLocaleBgBg      MapRunParamsLocale = "bg-BG"
	MapRunParamsLocaleBhoIn     MapRunParamsLocale = "bho-IN"
	MapRunParamsLocaleBm        MapRunParamsLocale = "bm"
	MapRunParamsLocaleBmMl      MapRunParamsLocale = "bm-ML"
	MapRunParamsLocaleBn        MapRunParamsLocale = "bn"
	MapRunParamsLocaleBnBd      MapRunParamsLocale = "bn-BD"
	MapRunParamsLocaleBnIn      MapRunParamsLocale = "bn-IN"
	MapRunParamsLocaleBo        MapRunParamsLocale = "bo"
	MapRunParamsLocaleBoCn      MapRunParamsLocale = "bo-CN"
	MapRunParamsLocaleBoIn      MapRunParamsLocale = "bo-IN"
	MapRunParamsLocaleBrFr      MapRunParamsLocale = "br-FR"
	MapRunParamsLocaleBrxIn     MapRunParamsLocale = "brx-IN"
	MapRunParamsLocaleBs        MapRunParamsLocale = "bs"
	MapRunParamsLocaleBsBa      MapRunParamsLocale = "bs-BA"
	MapRunParamsLocaleBynEr     MapRunParamsLocale = "byn-ER"
	MapRunParamsLocaleCa        MapRunParamsLocale = "ca"
	MapRunParamsLocaleCaAd      MapRunParamsLocale = "ca-AD"
	MapRunParamsLocaleCaEs      MapRunParamsLocale = "ca-ES"
	MapRunParamsLocaleCaFr      MapRunParamsLocale = "ca-FR"
	MapRunParamsLocaleCaIt      MapRunParamsLocale = "ca-IT"
	MapRunParamsLocaleCgg       MapRunParamsLocale = "cgg"
	MapRunParamsLocaleCggUg     MapRunParamsLocale = "cgg-UG"
	MapRunParamsLocaleChr       MapRunParamsLocale = "chr"
	MapRunParamsLocaleChrUs     MapRunParamsLocale = "chr-US"
	MapRunParamsLocaleCrhUa     MapRunParamsLocale = "crh-UA"
	MapRunParamsLocaleCs        MapRunParamsLocale = "cs"
	MapRunParamsLocaleCsCz      MapRunParamsLocale = "cs-CZ"
	MapRunParamsLocaleCsbPl     MapRunParamsLocale = "csb-PL"
	MapRunParamsLocaleCvRu      MapRunParamsLocale = "cv-RU"
	MapRunParamsLocaleCy        MapRunParamsLocale = "cy"
	MapRunParamsLocaleCyGB      MapRunParamsLocale = "cy-GB"
	MapRunParamsLocaleDa        MapRunParamsLocale = "da"
	MapRunParamsLocaleDaDk      MapRunParamsLocale = "da-DK"
	MapRunParamsLocaleDav       MapRunParamsLocale = "dav"
	MapRunParamsLocaleDavKe     MapRunParamsLocale = "dav-KE"
	MapRunParamsLocaleDe        MapRunParamsLocale = "de"
	MapRunParamsLocaleDeAt      MapRunParamsLocale = "de-AT"
	MapRunParamsLocaleDeBe      MapRunParamsLocale = "de-BE"
	MapRunParamsLocaleDeCh      MapRunParamsLocale = "de-CH"
	MapRunParamsLocaleDeDe      MapRunParamsLocale = "de-DE"
	MapRunParamsLocaleDeLi      MapRunParamsLocale = "de-LI"
	MapRunParamsLocaleDeLu      MapRunParamsLocale = "de-LU"
	MapRunParamsLocaleDvMv      MapRunParamsLocale = "dv-MV"
	MapRunParamsLocaleDzBt      MapRunParamsLocale = "dz-BT"
	MapRunParamsLocaleEbu       MapRunParamsLocale = "ebu"
	MapRunParamsLocaleEbuKe     MapRunParamsLocale = "ebu-KE"
	MapRunParamsLocaleEe        MapRunParamsLocale = "ee"
	MapRunParamsLocaleEeGh      MapRunParamsLocale = "ee-GH"
	MapRunParamsLocaleEeTg      MapRunParamsLocale = "ee-TG"
	MapRunParamsLocaleEl        MapRunParamsLocale = "el"
	MapRunParamsLocaleElCy      MapRunParamsLocale = "el-CY"
	MapRunParamsLocaleElGr      MapRunParamsLocale = "el-GR"
	MapRunParamsLocaleEn        MapRunParamsLocale = "en"
	MapRunParamsLocaleEnAg      MapRunParamsLocale = "en-AG"
	MapRunParamsLocaleEnAs      MapRunParamsLocale = "en-AS"
	MapRunParamsLocaleEnAu      MapRunParamsLocale = "en-AU"
	MapRunParamsLocaleEnBe      MapRunParamsLocale = "en-BE"
	MapRunParamsLocaleEnBw      MapRunParamsLocale = "en-BW"
	MapRunParamsLocaleEnBz      MapRunParamsLocale = "en-BZ"
	MapRunParamsLocaleEnCa      MapRunParamsLocale = "en-CA"
	MapRunParamsLocaleEnDk      MapRunParamsLocale = "en-DK"
	MapRunParamsLocaleEnGB      MapRunParamsLocale = "en-GB"
	MapRunParamsLocaleEnGu      MapRunParamsLocale = "en-GU"
	MapRunParamsLocaleEnHk      MapRunParamsLocale = "en-HK"
	MapRunParamsLocaleEnIe      MapRunParamsLocale = "en-IE"
	MapRunParamsLocaleEnIn      MapRunParamsLocale = "en-IN"
	MapRunParamsLocaleEnJm      MapRunParamsLocale = "en-JM"
	MapRunParamsLocaleEnMh      MapRunParamsLocale = "en-MH"
	MapRunParamsLocaleEnMp      MapRunParamsLocale = "en-MP"
	MapRunParamsLocaleEnMt      MapRunParamsLocale = "en-MT"
	MapRunParamsLocaleEnMu      MapRunParamsLocale = "en-MU"
	MapRunParamsLocaleEnNa      MapRunParamsLocale = "en-NA"
	MapRunParamsLocaleEnNg      MapRunParamsLocale = "en-NG"
	MapRunParamsLocaleEnNz      MapRunParamsLocale = "en-NZ"
	MapRunParamsLocaleEnPh      MapRunParamsLocale = "en-PH"
	MapRunParamsLocaleEnPk      MapRunParamsLocale = "en-PK"
	MapRunParamsLocaleEnSg      MapRunParamsLocale = "en-SG"
	MapRunParamsLocaleEnTt      MapRunParamsLocale = "en-TT"
	MapRunParamsLocaleEnUm      MapRunParamsLocale = "en-UM"
	MapRunParamsLocaleEnUs      MapRunParamsLocale = "en-US"
	MapRunParamsLocaleEnVi      MapRunParamsLocale = "en-VI"
	MapRunParamsLocaleEnZa      MapRunParamsLocale = "en-ZA"
	MapRunParamsLocaleEnZm      MapRunParamsLocale = "en-ZM"
	MapRunParamsLocaleEnZw      MapRunParamsLocale = "en-ZW"
	MapRunParamsLocaleEo        MapRunParamsLocale = "eo"
	MapRunParamsLocaleEs        MapRunParamsLocale = "es"
	MapRunParamsLocaleEs419     MapRunParamsLocale = "es-419"
	MapRunParamsLocaleEsAr      MapRunParamsLocale = "es-AR"
	MapRunParamsLocaleEsBo      MapRunParamsLocale = "es-BO"
	MapRunParamsLocaleEsCl      MapRunParamsLocale = "es-CL"
	MapRunParamsLocaleEsCo      MapRunParamsLocale = "es-CO"
	MapRunParamsLocaleEsCr      MapRunParamsLocale = "es-CR"
	MapRunParamsLocaleEsCu      MapRunParamsLocale = "es-CU"
	MapRunParamsLocaleEsDo      MapRunParamsLocale = "es-DO"
	MapRunParamsLocaleEsEc      MapRunParamsLocale = "es-EC"
	MapRunParamsLocaleEsEs      MapRunParamsLocale = "es-ES"
	MapRunParamsLocaleEsGq      MapRunParamsLocale = "es-GQ"
	MapRunParamsLocaleEsGt      MapRunParamsLocale = "es-GT"
	MapRunParamsLocaleEsHn      MapRunParamsLocale = "es-HN"
	MapRunParamsLocaleEsMx      MapRunParamsLocale = "es-MX"
	MapRunParamsLocaleEsNi      MapRunParamsLocale = "es-NI"
	MapRunParamsLocaleEsPa      MapRunParamsLocale = "es-PA"
	MapRunParamsLocaleEsPe      MapRunParamsLocale = "es-PE"
	MapRunParamsLocaleEsPr      MapRunParamsLocale = "es-PR"
	MapRunParamsLocaleEsPy      MapRunParamsLocale = "es-PY"
	MapRunParamsLocaleEsSv      MapRunParamsLocale = "es-SV"
	MapRunParamsLocaleEsUs      MapRunParamsLocale = "es-US"
	MapRunParamsLocaleEsUy      MapRunParamsLocale = "es-UY"
	MapRunParamsLocaleEsVe      MapRunParamsLocale = "es-VE"
	MapRunParamsLocaleEt        MapRunParamsLocale = "et"
	MapRunParamsLocaleEtEe      MapRunParamsLocale = "et-EE"
	MapRunParamsLocaleEu        MapRunParamsLocale = "eu"
	MapRunParamsLocaleEuEs      MapRunParamsLocale = "eu-ES"
	MapRunParamsLocaleFa        MapRunParamsLocale = "fa"
	MapRunParamsLocaleFaAf      MapRunParamsLocale = "fa-AF"
	MapRunParamsLocaleFaIr      MapRunParamsLocale = "fa-IR"
	MapRunParamsLocaleFf        MapRunParamsLocale = "ff"
	MapRunParamsLocaleFfSn      MapRunParamsLocale = "ff-SN"
	MapRunParamsLocaleFi        MapRunParamsLocale = "fi"
	MapRunParamsLocaleFiFi      MapRunParamsLocale = "fi-FI"
	MapRunParamsLocaleFil       MapRunParamsLocale = "fil"
	MapRunParamsLocaleFilPh     MapRunParamsLocale = "fil-PH"
	MapRunParamsLocaleFo        MapRunParamsLocale = "fo"
	MapRunParamsLocaleFoFo      MapRunParamsLocale = "fo-FO"
	MapRunParamsLocaleFr        MapRunParamsLocale = "fr"
	MapRunParamsLocaleFrBe      MapRunParamsLocale = "fr-BE"
	MapRunParamsLocaleFrBf      MapRunParamsLocale = "fr-BF"
	MapRunParamsLocaleFrBi      MapRunParamsLocale = "fr-BI"
	MapRunParamsLocaleFrBj      MapRunParamsLocale = "fr-BJ"
	MapRunParamsLocaleFrBl      MapRunParamsLocale = "fr-BL"
	MapRunParamsLocaleFrCa      MapRunParamsLocale = "fr-CA"
	MapRunParamsLocaleFrCd      MapRunParamsLocale = "fr-CD"
	MapRunParamsLocaleFrCf      MapRunParamsLocale = "fr-CF"
	MapRunParamsLocaleFrCg      MapRunParamsLocale = "fr-CG"
	MapRunParamsLocaleFrCh      MapRunParamsLocale = "fr-CH"
	MapRunParamsLocaleFrCi      MapRunParamsLocale = "fr-CI"
	MapRunParamsLocaleFrCm      MapRunParamsLocale = "fr-CM"
	MapRunParamsLocaleFrDj      MapRunParamsLocale = "fr-DJ"
	MapRunParamsLocaleFrFr      MapRunParamsLocale = "fr-FR"
	MapRunParamsLocaleFrGa      MapRunParamsLocale = "fr-GA"
	MapRunParamsLocaleFrGn      MapRunParamsLocale = "fr-GN"
	MapRunParamsLocaleFrGp      MapRunParamsLocale = "fr-GP"
	MapRunParamsLocaleFrGq      MapRunParamsLocale = "fr-GQ"
	MapRunParamsLocaleFrKm      MapRunParamsLocale = "fr-KM"
	MapRunParamsLocaleFrLu      MapRunParamsLocale = "fr-LU"
	MapRunParamsLocaleFrMc      MapRunParamsLocale = "fr-MC"
	MapRunParamsLocaleFrMf      MapRunParamsLocale = "fr-MF"
	MapRunParamsLocaleFrMg      MapRunParamsLocale = "fr-MG"
	MapRunParamsLocaleFrMl      MapRunParamsLocale = "fr-ML"
	MapRunParamsLocaleFrMq      MapRunParamsLocale = "fr-MQ"
	MapRunParamsLocaleFrNe      MapRunParamsLocale = "fr-NE"
	MapRunParamsLocaleFrRe      MapRunParamsLocale = "fr-RE"
	MapRunParamsLocaleFrRw      MapRunParamsLocale = "fr-RW"
	MapRunParamsLocaleFrSn      MapRunParamsLocale = "fr-SN"
	MapRunParamsLocaleFrTd      MapRunParamsLocale = "fr-TD"
	MapRunParamsLocaleFrTg      MapRunParamsLocale = "fr-TG"
	MapRunParamsLocaleFurIt     MapRunParamsLocale = "fur-IT"
	MapRunParamsLocaleFyDe      MapRunParamsLocale = "fy-DE"
	MapRunParamsLocaleFyNl      MapRunParamsLocale = "fy-NL"
	MapRunParamsLocaleGa        MapRunParamsLocale = "ga"
	MapRunParamsLocaleGaIe      MapRunParamsLocale = "ga-IE"
	MapRunParamsLocaleGdGB      MapRunParamsLocale = "gd-GB"
	MapRunParamsLocaleGezEr     MapRunParamsLocale = "gez-ER"
	MapRunParamsLocaleGezEt     MapRunParamsLocale = "gez-ET"
	MapRunParamsLocaleGl        MapRunParamsLocale = "gl"
	MapRunParamsLocaleGlEs      MapRunParamsLocale = "gl-ES"
	MapRunParamsLocaleGsw       MapRunParamsLocale = "gsw"
	MapRunParamsLocaleGswCh     MapRunParamsLocale = "gsw-CH"
	MapRunParamsLocaleGu        MapRunParamsLocale = "gu"
	MapRunParamsLocaleGuIn      MapRunParamsLocale = "gu-IN"
	MapRunParamsLocaleGuz       MapRunParamsLocale = "guz"
	MapRunParamsLocaleGuzKe     MapRunParamsLocale = "guz-KE"
	MapRunParamsLocaleGv        MapRunParamsLocale = "gv"
	MapRunParamsLocaleGvGB      MapRunParamsLocale = "gv-GB"
	MapRunParamsLocaleHa        MapRunParamsLocale = "ha"
	MapRunParamsLocaleHaLatn    MapRunParamsLocale = "ha-Latn"
	MapRunParamsLocaleHaLatnGh  MapRunParamsLocale = "ha-Latn-GH"
	MapRunParamsLocaleHaLatnNe  MapRunParamsLocale = "ha-Latn-NE"
	MapRunParamsLocaleHaLatnNg  MapRunParamsLocale = "ha-Latn-NG"
	MapRunParamsLocaleHaNg      MapRunParamsLocale = "ha-NG"
	MapRunParamsLocaleHaw       MapRunParamsLocale = "haw"
	MapRunParamsLocaleHawUs     MapRunParamsLocale = "haw-US"
	MapRunParamsLocaleHe        MapRunParamsLocale = "he"
	MapRunParamsLocaleHeIl      MapRunParamsLocale = "he-IL"
	MapRunParamsLocaleHi        MapRunParamsLocale = "hi"
	MapRunParamsLocaleHiIn      MapRunParamsLocale = "hi-IN"
	MapRunParamsLocaleHneIn     MapRunParamsLocale = "hne-IN"
	MapRunParamsLocaleHr        MapRunParamsLocale = "hr"
	MapRunParamsLocaleHrHr      MapRunParamsLocale = "hr-HR"
	MapRunParamsLocaleHsbDe     MapRunParamsLocale = "hsb-DE"
	MapRunParamsLocaleHtHt      MapRunParamsLocale = "ht-HT"
	MapRunParamsLocaleHu        MapRunParamsLocale = "hu"
	MapRunParamsLocaleHuHu      MapRunParamsLocale = "hu-HU"
	MapRunParamsLocaleHy        MapRunParamsLocale = "hy"
	MapRunParamsLocaleHyAm      MapRunParamsLocale = "hy-AM"
	MapRunParamsLocaleID        MapRunParamsLocale = "id"
	MapRunParamsLocaleIDID      MapRunParamsLocale = "id-ID"
	MapRunParamsLocaleIg        MapRunParamsLocale = "ig"
	MapRunParamsLocaleIgNg      MapRunParamsLocale = "ig-NG"
	MapRunParamsLocaleIi        MapRunParamsLocale = "ii"
	MapRunParamsLocaleIiCn      MapRunParamsLocale = "ii-CN"
	MapRunParamsLocaleIkCa      MapRunParamsLocale = "ik-CA"
	MapRunParamsLocaleIs        MapRunParamsLocale = "is"
	MapRunParamsLocaleIsIs      MapRunParamsLocale = "is-IS"
	MapRunParamsLocaleIt        MapRunParamsLocale = "it"
	MapRunParamsLocaleItCh      MapRunParamsLocale = "it-CH"
	MapRunParamsLocaleItIt      MapRunParamsLocale = "it-IT"
	MapRunParamsLocaleIuCa      MapRunParamsLocale = "iu-CA"
	MapRunParamsLocaleIwIl      MapRunParamsLocale = "iw-IL"
	MapRunParamsLocaleJa        MapRunParamsLocale = "ja"
	MapRunParamsLocaleJaJp      MapRunParamsLocale = "ja-JP"
	MapRunParamsLocaleJmc       MapRunParamsLocale = "jmc"
	MapRunParamsLocaleJmcTz     MapRunParamsLocale = "jmc-TZ"
	MapRunParamsLocaleKa        MapRunParamsLocale = "ka"
	MapRunParamsLocaleKaGe      MapRunParamsLocale = "ka-GE"
	MapRunParamsLocaleKab       MapRunParamsLocale = "kab"
	MapRunParamsLocaleKabDz     MapRunParamsLocale = "kab-DZ"
	MapRunParamsLocaleKam       MapRunParamsLocale = "kam"
	MapRunParamsLocaleKamKe     MapRunParamsLocale = "kam-KE"
	MapRunParamsLocaleKde       MapRunParamsLocale = "kde"
	MapRunParamsLocaleKdeTz     MapRunParamsLocale = "kde-TZ"
	MapRunParamsLocaleKea       MapRunParamsLocale = "kea"
	MapRunParamsLocaleKeaCv     MapRunParamsLocale = "kea-CV"
	MapRunParamsLocaleKhq       MapRunParamsLocale = "khq"
	MapRunParamsLocaleKhqMl     MapRunParamsLocale = "khq-ML"
	MapRunParamsLocaleKi        MapRunParamsLocale = "ki"
	MapRunParamsLocaleKiKe      MapRunParamsLocale = "ki-KE"
	MapRunParamsLocaleKk        MapRunParamsLocale = "kk"
	MapRunParamsLocaleKkCyrl    MapRunParamsLocale = "kk-Cyrl"
	MapRunParamsLocaleKkCyrlKz  MapRunParamsLocale = "kk-Cyrl-KZ"
	MapRunParamsLocaleKkKz      MapRunParamsLocale = "kk-KZ"
	MapRunParamsLocaleKl        MapRunParamsLocale = "kl"
	MapRunParamsLocaleKlGl      MapRunParamsLocale = "kl-GL"
	MapRunParamsLocaleKln       MapRunParamsLocale = "kln"
	MapRunParamsLocaleKlnKe     MapRunParamsLocale = "kln-KE"
	MapRunParamsLocaleKm        MapRunParamsLocale = "km"
	MapRunParamsLocaleKmKh      MapRunParamsLocale = "km-KH"
	MapRunParamsLocaleKn        MapRunParamsLocale = "kn"
	MapRunParamsLocaleKnIn      MapRunParamsLocale = "kn-IN"
	MapRunParamsLocaleKo        MapRunParamsLocale = "ko"
	MapRunParamsLocaleKoKr      MapRunParamsLocale = "ko-KR"
	MapRunParamsLocaleKok       MapRunParamsLocale = "kok"
	MapRunParamsLocaleKokIn     MapRunParamsLocale = "kok-IN"
	MapRunParamsLocaleKsIn      MapRunParamsLocale = "ks-IN"
	MapRunParamsLocaleKuTr      MapRunParamsLocale = "ku-TR"
	MapRunParamsLocaleKw        MapRunParamsLocale = "kw"
	MapRunParamsLocaleKwGB      MapRunParamsLocale = "kw-GB"
	MapRunParamsLocaleKyKg      MapRunParamsLocale = "ky-KG"
	MapRunParamsLocaleLag       MapRunParamsLocale = "lag"
	MapRunParamsLocaleLagTz     MapRunParamsLocale = "lag-TZ"
	MapRunParamsLocaleLbLu      MapRunParamsLocale = "lb-LU"
	MapRunParamsLocaleLg        MapRunParamsLocale = "lg"
	MapRunParamsLocaleLgUg      MapRunParamsLocale = "lg-UG"
	MapRunParamsLocaleLiBe      MapRunParamsLocale = "li-BE"
	MapRunParamsLocaleLiNl      MapRunParamsLocale = "li-NL"
	MapRunParamsLocaleLijIt     MapRunParamsLocale = "lij-IT"
	MapRunParamsLocaleLoLa      MapRunParamsLocale = "lo-LA"
	MapRunParamsLocaleLt        MapRunParamsLocale = "lt"
	MapRunParamsLocaleLtLt      MapRunParamsLocale = "lt-LT"
	MapRunParamsLocaleLuo       MapRunParamsLocale = "luo"
	MapRunParamsLocaleLuoKe     MapRunParamsLocale = "luo-KE"
	MapRunParamsLocaleLuy       MapRunParamsLocale = "luy"
	MapRunParamsLocaleLuyKe     MapRunParamsLocale = "luy-KE"
	MapRunParamsLocaleLv        MapRunParamsLocale = "lv"
	MapRunParamsLocaleLvLv      MapRunParamsLocale = "lv-LV"
	MapRunParamsLocaleMagIn     MapRunParamsLocale = "mag-IN"
	MapRunParamsLocaleMaiIn     MapRunParamsLocale = "mai-IN"
	MapRunParamsLocaleMas       MapRunParamsLocale = "mas"
	MapRunParamsLocaleMasKe     MapRunParamsLocale = "mas-KE"
	MapRunParamsLocaleMasTz     MapRunParamsLocale = "mas-TZ"
	MapRunParamsLocaleMer       MapRunParamsLocale = "mer"
	MapRunParamsLocaleMerKe     MapRunParamsLocale = "mer-KE"
	MapRunParamsLocaleMfe       MapRunParamsLocale = "mfe"
	MapRunParamsLocaleMfeMu     MapRunParamsLocale = "mfe-MU"
	MapRunParamsLocaleMg        MapRunParamsLocale = "mg"
	MapRunParamsLocaleMgMg      MapRunParamsLocale = "mg-MG"
	MapRunParamsLocaleMhrRu     MapRunParamsLocale = "mhr-RU"
	MapRunParamsLocaleMiNz      MapRunParamsLocale = "mi-NZ"
	MapRunParamsLocaleMk        MapRunParamsLocale = "mk"
	MapRunParamsLocaleMkMk      MapRunParamsLocale = "mk-MK"
	MapRunParamsLocaleMl        MapRunParamsLocale = "ml"
	MapRunParamsLocaleMlIn      MapRunParamsLocale = "ml-IN"
	MapRunParamsLocaleMnMn      MapRunParamsLocale = "mn-MN"
	MapRunParamsLocaleMr        MapRunParamsLocale = "mr"
	MapRunParamsLocaleMrIn      MapRunParamsLocale = "mr-IN"
	MapRunParamsLocaleMs        MapRunParamsLocale = "ms"
	MapRunParamsLocaleMsBn      MapRunParamsLocale = "ms-BN"
	MapRunParamsLocaleMsMy      MapRunParamsLocale = "ms-MY"
	MapRunParamsLocaleMt        MapRunParamsLocale = "mt"
	MapRunParamsLocaleMtMt      MapRunParamsLocale = "mt-MT"
	MapRunParamsLocaleMy        MapRunParamsLocale = "my"
	MapRunParamsLocaleMyMm      MapRunParamsLocale = "my-MM"
	MapRunParamsLocaleNanTw     MapRunParamsLocale = "nan-TW"
	MapRunParamsLocaleNaq       MapRunParamsLocale = "naq"
	MapRunParamsLocaleNaqNa     MapRunParamsLocale = "naq-NA"
	MapRunParamsLocaleNb        MapRunParamsLocale = "nb"
	MapRunParamsLocaleNbNo      MapRunParamsLocale = "nb-NO"
	MapRunParamsLocaleNd        MapRunParamsLocale = "nd"
	MapRunParamsLocaleNdZw      MapRunParamsLocale = "nd-ZW"
	MapRunParamsLocaleNdsDe     MapRunParamsLocale = "nds-DE"
	MapRunParamsLocaleNdsNl     MapRunParamsLocale = "nds-NL"
	MapRunParamsLocaleNe        MapRunParamsLocale = "ne"
	MapRunParamsLocaleNeIn      MapRunParamsLocale = "ne-IN"
	MapRunParamsLocaleNeNp      MapRunParamsLocale = "ne-NP"
	MapRunParamsLocaleNl        MapRunParamsLocale = "nl"
	MapRunParamsLocaleNlAw      MapRunParamsLocale = "nl-AW"
	MapRunParamsLocaleNlBe      MapRunParamsLocale = "nl-BE"
	MapRunParamsLocaleNlNl      MapRunParamsLocale = "nl-NL"
	MapRunParamsLocaleNn        MapRunParamsLocale = "nn"
	MapRunParamsLocaleNnNo      MapRunParamsLocale = "nn-NO"
	MapRunParamsLocaleNrZa      MapRunParamsLocale = "nr-ZA"
	MapRunParamsLocaleNsoZa     MapRunParamsLocale = "nso-ZA"
	MapRunParamsLocaleNyn       MapRunParamsLocale = "nyn"
	MapRunParamsLocaleNynUg     MapRunParamsLocale = "nyn-UG"
	MapRunParamsLocaleOcFr      MapRunParamsLocale = "oc-FR"
	MapRunParamsLocaleOm        MapRunParamsLocale = "om"
	MapRunParamsLocaleOmEt      MapRunParamsLocale = "om-ET"
	MapRunParamsLocaleOmKe      MapRunParamsLocale = "om-KE"
	MapRunParamsLocaleOr        MapRunParamsLocale = "or"
	MapRunParamsLocaleOrIn      MapRunParamsLocale = "or-IN"
	MapRunParamsLocaleOsRu      MapRunParamsLocale = "os-RU"
	MapRunParamsLocalePa        MapRunParamsLocale = "pa"
	MapRunParamsLocalePaArab    MapRunParamsLocale = "pa-Arab"
	MapRunParamsLocalePaArabPk  MapRunParamsLocale = "pa-Arab-PK"
	MapRunParamsLocalePaGuru    MapRunParamsLocale = "pa-Guru"
	MapRunParamsLocalePaGuruIn  MapRunParamsLocale = "pa-Guru-IN"
	MapRunParamsLocalePaIn      MapRunParamsLocale = "pa-IN"
	MapRunParamsLocalePaPk      MapRunParamsLocale = "pa-PK"
	MapRunParamsLocalePapAn     MapRunParamsLocale = "pap-AN"
	MapRunParamsLocalePl        MapRunParamsLocale = "pl"
	MapRunParamsLocalePlPl      MapRunParamsLocale = "pl-PL"
	MapRunParamsLocalePs        MapRunParamsLocale = "ps"
	MapRunParamsLocalePsAf      MapRunParamsLocale = "ps-AF"
	MapRunParamsLocalePt        MapRunParamsLocale = "pt"
	MapRunParamsLocalePtBr      MapRunParamsLocale = "pt-BR"
	MapRunParamsLocalePtGw      MapRunParamsLocale = "pt-GW"
	MapRunParamsLocalePtMz      MapRunParamsLocale = "pt-MZ"
	MapRunParamsLocalePtPt      MapRunParamsLocale = "pt-PT"
	MapRunParamsLocaleRm        MapRunParamsLocale = "rm"
	MapRunParamsLocaleRmCh      MapRunParamsLocale = "rm-CH"
	MapRunParamsLocaleRo        MapRunParamsLocale = "ro"
	MapRunParamsLocaleRoMd      MapRunParamsLocale = "ro-MD"
	MapRunParamsLocaleRoRo      MapRunParamsLocale = "ro-RO"
	MapRunParamsLocaleRof       MapRunParamsLocale = "rof"
	MapRunParamsLocaleRofTz     MapRunParamsLocale = "rof-TZ"
	MapRunParamsLocaleRu        MapRunParamsLocale = "ru"
	MapRunParamsLocaleRuMd      MapRunParamsLocale = "ru-MD"
	MapRunParamsLocaleRuRu      MapRunParamsLocale = "ru-RU"
	MapRunParamsLocaleRuUa      MapRunParamsLocale = "ru-UA"
	MapRunParamsLocaleRw        MapRunParamsLocale = "rw"
	MapRunParamsLocaleRwRw      MapRunParamsLocale = "rw-RW"
	MapRunParamsLocaleRwk       MapRunParamsLocale = "rwk"
	MapRunParamsLocaleRwkTz     MapRunParamsLocale = "rwk-TZ"
	MapRunParamsLocaleSaIn      MapRunParamsLocale = "sa-IN"
	MapRunParamsLocaleSaq       MapRunParamsLocale = "saq"
	MapRunParamsLocaleSaqKe     MapRunParamsLocale = "saq-KE"
	MapRunParamsLocaleScIt      MapRunParamsLocale = "sc-IT"
	MapRunParamsLocaleSdIn      MapRunParamsLocale = "sd-IN"
	MapRunParamsLocaleSeNo      MapRunParamsLocale = "se-NO"
	MapRunParamsLocaleSeh       MapRunParamsLocale = "seh"
	MapRunParamsLocaleSehMz     MapRunParamsLocale = "seh-MZ"
	MapRunParamsLocaleSes       MapRunParamsLocale = "ses"
	MapRunParamsLocaleSesMl     MapRunParamsLocale = "ses-ML"
	MapRunParamsLocaleSg        MapRunParamsLocale = "sg"
	MapRunParamsLocaleSgCf      MapRunParamsLocale = "sg-CF"
	MapRunParamsLocaleShi       MapRunParamsLocale = "shi"
	MapRunParamsLocaleShiLatn   MapRunParamsLocale = "shi-Latn"
	MapRunParamsLocaleShiLatnMa MapRunParamsLocale = "shi-Latn-MA"
	MapRunParamsLocaleShiTfng   MapRunParamsLocale = "shi-Tfng"
	MapRunParamsLocaleShiTfngMa MapRunParamsLocale = "shi-Tfng-MA"
	MapRunParamsLocaleShsCa     MapRunParamsLocale = "shs-CA"
	MapRunParamsLocaleSi        MapRunParamsLocale = "si"
	MapRunParamsLocaleSiLk      MapRunParamsLocale = "si-LK"
	MapRunParamsLocaleSidEt     MapRunParamsLocale = "sid-ET"
	MapRunParamsLocaleSk        MapRunParamsLocale = "sk"
	MapRunParamsLocaleSkSk      MapRunParamsLocale = "sk-SK"
	MapRunParamsLocaleSl        MapRunParamsLocale = "sl"
	MapRunParamsLocaleSlSi      MapRunParamsLocale = "sl-SI"
	MapRunParamsLocaleSn        MapRunParamsLocale = "sn"
	MapRunParamsLocaleSnZw      MapRunParamsLocale = "sn-ZW"
	MapRunParamsLocaleSo        MapRunParamsLocale = "so"
	MapRunParamsLocaleSoDj      MapRunParamsLocale = "so-DJ"
	MapRunParamsLocaleSoEt      MapRunParamsLocale = "so-ET"
	MapRunParamsLocaleSoKe      MapRunParamsLocale = "so-KE"
	MapRunParamsLocaleSoSo      MapRunParamsLocale = "so-SO"
	MapRunParamsLocaleSq        MapRunParamsLocale = "sq"
	MapRunParamsLocaleSqAl      MapRunParamsLocale = "sq-AL"
	MapRunParamsLocaleSqMk      MapRunParamsLocale = "sq-MK"
	MapRunParamsLocaleSr        MapRunParamsLocale = "sr"
	MapRunParamsLocaleSrCyrl    MapRunParamsLocale = "sr-Cyrl"
	MapRunParamsLocaleSrCyrlBa  MapRunParamsLocale = "sr-Cyrl-BA"
	MapRunParamsLocaleSrCyrlMe  MapRunParamsLocale = "sr-Cyrl-ME"
	MapRunParamsLocaleSrCyrlRs  MapRunParamsLocale = "sr-Cyrl-RS"
	MapRunParamsLocaleSrLatn    MapRunParamsLocale = "sr-Latn"
	MapRunParamsLocaleSrLatnBa  MapRunParamsLocale = "sr-Latn-BA"
	MapRunParamsLocaleSrLatnMe  MapRunParamsLocale = "sr-Latn-ME"
	MapRunParamsLocaleSrLatnRs  MapRunParamsLocale = "sr-Latn-RS"
	MapRunParamsLocaleSrMe      MapRunParamsLocale = "sr-ME"
	MapRunParamsLocaleSrRs      MapRunParamsLocale = "sr-RS"
	MapRunParamsLocaleSSZa      MapRunParamsLocale = "ss-ZA"
	MapRunParamsLocaleStZa      MapRunParamsLocale = "st-ZA"
	MapRunParamsLocaleSv        MapRunParamsLocale = "sv"
	MapRunParamsLocaleSvFi      MapRunParamsLocale = "sv-FI"
	MapRunParamsLocaleSvSe      MapRunParamsLocale = "sv-SE"
	MapRunParamsLocaleSw        MapRunParamsLocale = "sw"
	MapRunParamsLocaleSwKe      MapRunParamsLocale = "sw-KE"
	MapRunParamsLocaleSwTz      MapRunParamsLocale = "sw-TZ"
	MapRunParamsLocaleTa        MapRunParamsLocale = "ta"
	MapRunParamsLocaleTaIn      MapRunParamsLocale = "ta-IN"
	MapRunParamsLocaleTaLk      MapRunParamsLocale = "ta-LK"
	MapRunParamsLocaleTe        MapRunParamsLocale = "te"
	MapRunParamsLocaleTeIn      MapRunParamsLocale = "te-IN"
	MapRunParamsLocaleTeo       MapRunParamsLocale = "teo"
	MapRunParamsLocaleTeoKe     MapRunParamsLocale = "teo-KE"
	MapRunParamsLocaleTeoUg     MapRunParamsLocale = "teo-UG"
	MapRunParamsLocaleTgTj      MapRunParamsLocale = "tg-TJ"
	MapRunParamsLocaleTh        MapRunParamsLocale = "th"
	MapRunParamsLocaleThTh      MapRunParamsLocale = "th-TH"
	MapRunParamsLocaleTi        MapRunParamsLocale = "ti"
	MapRunParamsLocaleTiEr      MapRunParamsLocale = "ti-ER"
	MapRunParamsLocaleTiEt      MapRunParamsLocale = "ti-ET"
	MapRunParamsLocaleTigEr     MapRunParamsLocale = "tig-ER"
	MapRunParamsLocaleTkTm      MapRunParamsLocale = "tk-TM"
	MapRunParamsLocaleTlPh      MapRunParamsLocale = "tl-PH"
	MapRunParamsLocaleTnZa      MapRunParamsLocale = "tn-ZA"
	MapRunParamsLocaleTo        MapRunParamsLocale = "to"
	MapRunParamsLocaleToTo      MapRunParamsLocale = "to-TO"
	MapRunParamsLocaleTr        MapRunParamsLocale = "tr"
	MapRunParamsLocaleTrCy      MapRunParamsLocale = "tr-CY"
	MapRunParamsLocaleTrTr      MapRunParamsLocale = "tr-TR"
	MapRunParamsLocaleTsZa      MapRunParamsLocale = "ts-ZA"
	MapRunParamsLocaleTtRu      MapRunParamsLocale = "tt-RU"
	MapRunParamsLocaleTzm       MapRunParamsLocale = "tzm"
	MapRunParamsLocaleTzmLatn   MapRunParamsLocale = "tzm-Latn"
	MapRunParamsLocaleTzmLatnMa MapRunParamsLocale = "tzm-Latn-MA"
	MapRunParamsLocaleUgCn      MapRunParamsLocale = "ug-CN"
	MapRunParamsLocaleUk        MapRunParamsLocale = "uk"
	MapRunParamsLocaleUkUa      MapRunParamsLocale = "uk-UA"
	MapRunParamsLocaleUnmUs     MapRunParamsLocale = "unm-US"
	MapRunParamsLocaleUr        MapRunParamsLocale = "ur"
	MapRunParamsLocaleUrIn      MapRunParamsLocale = "ur-IN"
	MapRunParamsLocaleUrPk      MapRunParamsLocale = "ur-PK"
	MapRunParamsLocaleUz        MapRunParamsLocale = "uz"
	MapRunParamsLocaleUzArab    MapRunParamsLocale = "uz-Arab"
	MapRunParamsLocaleUzArabAf  MapRunParamsLocale = "uz-Arab-AF"
	MapRunParamsLocaleUzCyrl    MapRunParamsLocale = "uz-Cyrl"
	MapRunParamsLocaleUzCyrlUz  MapRunParamsLocale = "uz-Cyrl-UZ"
	MapRunParamsLocaleUzLatn    MapRunParamsLocale = "uz-Latn"
	MapRunParamsLocaleUzLatnUz  MapRunParamsLocale = "uz-Latn-UZ"
	MapRunParamsLocaleUzUz      MapRunParamsLocale = "uz-UZ"
	MapRunParamsLocaleVeZa      MapRunParamsLocale = "ve-ZA"
	MapRunParamsLocaleVi        MapRunParamsLocale = "vi"
	MapRunParamsLocaleViVn      MapRunParamsLocale = "vi-VN"
	MapRunParamsLocaleVun       MapRunParamsLocale = "vun"
	MapRunParamsLocaleVunTz     MapRunParamsLocale = "vun-TZ"
	MapRunParamsLocaleWaBe      MapRunParamsLocale = "wa-BE"
	MapRunParamsLocaleWaeCh     MapRunParamsLocale = "wae-CH"
	MapRunParamsLocaleWalEt     MapRunParamsLocale = "wal-ET"
	MapRunParamsLocaleWoSn      MapRunParamsLocale = "wo-SN"
	MapRunParamsLocaleXhZa      MapRunParamsLocale = "xh-ZA"
	MapRunParamsLocaleXog       MapRunParamsLocale = "xog"
	MapRunParamsLocaleXogUg     MapRunParamsLocale = "xog-UG"
	MapRunParamsLocaleYiUs      MapRunParamsLocale = "yi-US"
	MapRunParamsLocaleYo        MapRunParamsLocale = "yo"
	MapRunParamsLocaleYoNg      MapRunParamsLocale = "yo-NG"
	MapRunParamsLocaleYueHk     MapRunParamsLocale = "yue-HK"
	MapRunParamsLocaleZh        MapRunParamsLocale = "zh"
	MapRunParamsLocaleZhCn      MapRunParamsLocale = "zh-CN"
	MapRunParamsLocaleZhHk      MapRunParamsLocale = "zh-HK"
	MapRunParamsLocaleZhHans    MapRunParamsLocale = "zh-Hans"
	MapRunParamsLocaleZhHansCn  MapRunParamsLocale = "zh-Hans-CN"
	MapRunParamsLocaleZhHansHk  MapRunParamsLocale = "zh-Hans-HK"
	MapRunParamsLocaleZhHansMo  MapRunParamsLocale = "zh-Hans-MO"
	MapRunParamsLocaleZhHansSg  MapRunParamsLocale = "zh-Hans-SG"
	MapRunParamsLocaleZhHant    MapRunParamsLocale = "zh-Hant"
	MapRunParamsLocaleZhHantHk  MapRunParamsLocale = "zh-Hant-HK"
	MapRunParamsLocaleZhHantMo  MapRunParamsLocale = "zh-Hant-MO"
	MapRunParamsLocaleZhHantTw  MapRunParamsLocale = "zh-Hant-TW"
	MapRunParamsLocaleZhSg      MapRunParamsLocale = "zh-SG"
	MapRunParamsLocaleZhTw      MapRunParamsLocale = "zh-TW"
	MapRunParamsLocaleZu        MapRunParamsLocale = "zu"
	MapRunParamsLocaleZuZa      MapRunParamsLocale = "zu-ZA"
	MapRunParamsLocaleAuto      MapRunParamsLocale = "auto"
)

// Sitemap and other methods will be used together to find URLs.
type MapRunParamsSitemap string

const (
	MapRunParamsSitemapSkip    MapRunParamsSitemap = "skip"
	MapRunParamsSitemapInclude MapRunParamsSitemap = "include"
	MapRunParamsSitemapOnly    MapRunParamsSitemap = "only"
)
