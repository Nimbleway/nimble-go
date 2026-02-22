// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"github.com/Nimbleway/nimble-go/internal/apierror"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Continuously scroll to load dynamic content
//
// This is an alias to an internal type.
type AutoScrollActionParam = shared.AutoScrollActionParam

// This is an alias to an internal type.
type AutoScrollActionAutoScrollUnionParam = shared.AutoScrollActionAutoScrollUnionParam

// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectParam = shared.AutoScrollActionAutoScrollObjectParam

// CSS selector or array of alternative selectors. Use an array when you have
// multiple possible selectors for the same element.
//
// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectClickSelectorUnionParam = shared.AutoScrollActionAutoScrollObjectClickSelectorUnionParam

// CSS selector or array of alternative selectors. Use an array when you have
// multiple possible selectors for the same element.
//
// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectContainerUnionParam = shared.AutoScrollActionAutoScrollObjectContainerUnionParam

// Duration value that accepts various formats. Supports: number (ms), string
// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
//
// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectDelayAfterScrollUnionParam = shared.AutoScrollActionAutoScrollObjectDelayAfterScrollUnionParam

// Duration value that accepts various formats. Supports: number (ms), string
// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
//
// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectIdleTimeoutUnionParam = shared.AutoScrollActionAutoScrollObjectIdleTimeoutUnionParam

// CSS selector or array of alternative selectors. Use an array when you have
// multiple possible selectors for the same element.
//
// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectLoadingSelectorUnionParam = shared.AutoScrollActionAutoScrollObjectLoadingSelectorUnionParam

// Duration value that accepts various formats. Supports: number (ms), string
// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
//
// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectMaxDurationUnionParam = shared.AutoScrollActionAutoScrollObjectMaxDurationUnionParam

// CSS selector or array of alternative selectors. Use an array when you have
// multiple possible selectors for the same element.
//
// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectPauseOnSelectorUnionParam = shared.AutoScrollActionAutoScrollObjectPauseOnSelectorUnionParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectRequiredUnionParam = shared.AutoScrollActionAutoScrollObjectRequiredUnionParam

// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectRequiredString = shared.AutoScrollActionAutoScrollObjectRequiredString

// Equals "true"
const AutoScrollActionAutoScrollObjectRequiredStringTrue = shared.AutoScrollActionAutoScrollObjectRequiredStringTrue

// Equals "false"
const AutoScrollActionAutoScrollObjectRequiredStringFalse = shared.AutoScrollActionAutoScrollObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectSkipUnionParam = shared.AutoScrollActionAutoScrollObjectSkipUnionParam

// This is an alias to an internal type.
type AutoScrollActionAutoScrollObjectSkipString = shared.AutoScrollActionAutoScrollObjectSkipString

// Equals "true"
const AutoScrollActionAutoScrollObjectSkipStringTrue = shared.AutoScrollActionAutoScrollObjectSkipStringTrue

// Equals "false"
const AutoScrollActionAutoScrollObjectSkipStringFalse = shared.AutoScrollActionAutoScrollObjectSkipStringFalse

// Click on an element by selector
//
// This is an alias to an internal type.
type ClickActionParam = shared.ClickActionParam

// This is an alias to an internal type.
type ClickActionClickUnionParam = shared.ClickActionClickUnionParam

// This is an alias to an internal type.
type ClickActionClickObjectParam = shared.ClickActionClickObjectParam

// CSS selector or array of alternative selectors. Use an array when you have
// multiple possible selectors for the same element.
//
// This is an alias to an internal type.
type ClickActionClickObjectSelectorUnionParam = shared.ClickActionClickObjectSelectorUnionParam

// Duration value that accepts various formats. Supports: number (ms), string
// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
//
// This is an alias to an internal type.
type ClickActionClickObjectDelayUnionParam = shared.ClickActionClickObjectDelayUnionParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type ClickActionClickObjectRequiredUnionParam = shared.ClickActionClickObjectRequiredUnionParam

// This is an alias to an internal type.
type ClickActionClickObjectRequiredString = shared.ClickActionClickObjectRequiredString

// Equals "true"
const ClickActionClickObjectRequiredStringTrue = shared.ClickActionClickObjectRequiredStringTrue

// Equals "false"
const ClickActionClickObjectRequiredStringFalse = shared.ClickActionClickObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type ClickActionClickObjectSkipUnionParam = shared.ClickActionClickObjectSkipUnionParam

// This is an alias to an internal type.
type ClickActionClickObjectSkipString = shared.ClickActionClickObjectSkipString

// Equals "true"
const ClickActionClickObjectSkipStringTrue = shared.ClickActionClickObjectSkipStringTrue

// Equals "false"
const ClickActionClickObjectSkipStringFalse = shared.ClickActionClickObjectSkipStringFalse

// Execute JavaScript code in page context
//
// This is an alias to an internal type.
type EvalActionParam = shared.EvalActionParam

// This is an alias to an internal type.
type EvalActionEvalUnionParam = shared.EvalActionEvalUnionParam

// This is an alias to an internal type.
type EvalActionEvalObjectParam = shared.EvalActionEvalObjectParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type EvalActionEvalObjectRequiredUnionParam = shared.EvalActionEvalObjectRequiredUnionParam

// This is an alias to an internal type.
type EvalActionEvalObjectRequiredString = shared.EvalActionEvalObjectRequiredString

// Equals "true"
const EvalActionEvalObjectRequiredStringTrue = shared.EvalActionEvalObjectRequiredStringTrue

// Equals "false"
const EvalActionEvalObjectRequiredStringFalse = shared.EvalActionEvalObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type EvalActionEvalObjectSkipUnionParam = shared.EvalActionEvalObjectSkipUnionParam

// This is an alias to an internal type.
type EvalActionEvalObjectSkipString = shared.EvalActionEvalObjectSkipString

// Equals "true"
const EvalActionEvalObjectSkipStringTrue = shared.EvalActionEvalObjectSkipStringTrue

// Equals "false"
const EvalActionEvalObjectSkipStringFalse = shared.EvalActionEvalObjectSkipStringFalse

// Make an HTTP request in browser context
//
// This is an alias to an internal type.
type FetchActionParam = shared.FetchActionParam

// This is an alias to an internal type.
type FetchActionFetchUnionParam = shared.FetchActionFetchUnionParam

// This is an alias to an internal type.
type FetchActionFetchObjectParam = shared.FetchActionFetchObjectParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type FetchActionFetchObjectRequiredUnionParam = shared.FetchActionFetchObjectRequiredUnionParam

// This is an alias to an internal type.
type FetchActionFetchObjectRequiredString = shared.FetchActionFetchObjectRequiredString

// Equals "true"
const FetchActionFetchObjectRequiredStringTrue = shared.FetchActionFetchObjectRequiredStringTrue

// Equals "false"
const FetchActionFetchObjectRequiredStringFalse = shared.FetchActionFetchObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type FetchActionFetchObjectSkipUnionParam = shared.FetchActionFetchObjectSkipUnionParam

// This is an alias to an internal type.
type FetchActionFetchObjectSkipString = shared.FetchActionFetchObjectSkipString

// Equals "true"
const FetchActionFetchObjectSkipStringTrue = shared.FetchActionFetchObjectSkipStringTrue

// Equals "false"
const FetchActionFetchObjectSkipStringFalse = shared.FetchActionFetchObjectSkipStringFalse

// Fill text into an input field
//
// This is an alias to an internal type.
type FillActionParam = shared.FillActionParam

// Fill options with mode-specific fields. Use "type" mode for behavioral typing
// simulation, or "paste" mode for instant paste.
//
// This is an alias to an internal type.
type FillActionFillUnionParam = shared.FillActionFillUnionParam

// This is an alias to an internal type.
type FillActionFillTypeParam = shared.FillActionFillTypeParam

// CSS selector or array of alternative selectors. Use an array when you have
// multiple possible selectors for the same element.
//
// This is an alias to an internal type.
type FillActionFillTypeSelectorUnionParam = shared.FillActionFillTypeSelectorUnionParam

// Duration value that accepts various formats. Supports: number (ms), string
// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
//
// This is an alias to an internal type.
type FillActionFillTypeDelayUnionParam = shared.FillActionFillTypeDelayUnionParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type FillActionFillTypeRequiredUnionParam = shared.FillActionFillTypeRequiredUnionParam

// This is an alias to an internal type.
type FillActionFillTypeRequiredString = shared.FillActionFillTypeRequiredString

// Equals "true"
const FillActionFillTypeRequiredStringTrue = shared.FillActionFillTypeRequiredStringTrue

// Equals "false"
const FillActionFillTypeRequiredStringFalse = shared.FillActionFillTypeRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type FillActionFillTypeSkipUnionParam = shared.FillActionFillTypeSkipUnionParam

// This is an alias to an internal type.
type FillActionFillTypeSkipString = shared.FillActionFillTypeSkipString

// Equals "true"
const FillActionFillTypeSkipStringTrue = shared.FillActionFillTypeSkipStringTrue

// Equals "false"
const FillActionFillTypeSkipStringFalse = shared.FillActionFillTypeSkipStringFalse

// Duration value that accepts various formats. Supports: number (ms), string
// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
//
// This is an alias to an internal type.
type FillActionFillTypeTypingIntervalUnionParam = shared.FillActionFillTypeTypingIntervalUnionParam

// This is an alias to an internal type.
type FillActionFillPasteParam = shared.FillActionFillPasteParam

// CSS selector or array of alternative selectors. Use an array when you have
// multiple possible selectors for the same element.
//
// This is an alias to an internal type.
type FillActionFillPasteSelectorUnionParam = shared.FillActionFillPasteSelectorUnionParam

// Duration value that accepts various formats. Supports: number (ms), string
// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
//
// This is an alias to an internal type.
type FillActionFillPasteDelayUnionParam = shared.FillActionFillPasteDelayUnionParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type FillActionFillPasteRequiredUnionParam = shared.FillActionFillPasteRequiredUnionParam

// This is an alias to an internal type.
type FillActionFillPasteRequiredString = shared.FillActionFillPasteRequiredString

// Equals "true"
const FillActionFillPasteRequiredStringTrue = shared.FillActionFillPasteRequiredStringTrue

// Equals "false"
const FillActionFillPasteRequiredStringFalse = shared.FillActionFillPasteRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type FillActionFillPasteSkipUnionParam = shared.FillActionFillPasteSkipUnionParam

// This is an alias to an internal type.
type FillActionFillPasteSkipString = shared.FillActionFillPasteSkipString

// Equals "true"
const FillActionFillPasteSkipStringTrue = shared.FillActionFillPasteSkipStringTrue

// Equals "false"
const FillActionFillPasteSkipStringFalse = shared.FillActionFillPasteSkipStringFalse

// Retrieve browser cookies
//
// This is an alias to an internal type.
type GetCookiesActionParam = shared.GetCookiesActionParam

// This is an alias to an internal type.
type GetCookiesActionGetCookiesUnionParam = shared.GetCookiesActionGetCookiesUnionParam

// This is an alias to an internal type.
type GetCookiesActionGetCookiesObjectParam = shared.GetCookiesActionGetCookiesObjectParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type GetCookiesActionGetCookiesObjectRequiredUnionParam = shared.GetCookiesActionGetCookiesObjectRequiredUnionParam

// This is an alias to an internal type.
type GetCookiesActionGetCookiesObjectRequiredString = shared.GetCookiesActionGetCookiesObjectRequiredString

// Equals "true"
const GetCookiesActionGetCookiesObjectRequiredStringTrue = shared.GetCookiesActionGetCookiesObjectRequiredStringTrue

// Equals "false"
const GetCookiesActionGetCookiesObjectRequiredStringFalse = shared.GetCookiesActionGetCookiesObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type GetCookiesActionGetCookiesObjectSkipUnionParam = shared.GetCookiesActionGetCookiesObjectSkipUnionParam

// This is an alias to an internal type.
type GetCookiesActionGetCookiesObjectSkipString = shared.GetCookiesActionGetCookiesObjectSkipString

// Equals "true"
const GetCookiesActionGetCookiesObjectSkipStringTrue = shared.GetCookiesActionGetCookiesObjectSkipStringTrue

// Equals "false"
const GetCookiesActionGetCookiesObjectSkipStringFalse = shared.GetCookiesActionGetCookiesObjectSkipStringFalse

// Navigate to a URL
//
// This is an alias to an internal type.
type GotoActionParam = shared.GotoActionParam

// This is an alias to an internal type.
type GotoActionGotoUnionParam = shared.GotoActionGotoUnionParam

// This is an alias to an internal type.
type GotoActionGotoObjectParam = shared.GotoActionGotoObjectParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type GotoActionGotoObjectRequiredUnionParam = shared.GotoActionGotoObjectRequiredUnionParam

// This is an alias to an internal type.
type GotoActionGotoObjectRequiredString = shared.GotoActionGotoObjectRequiredString

// Equals "true"
const GotoActionGotoObjectRequiredStringTrue = shared.GotoActionGotoObjectRequiredStringTrue

// Equals "false"
const GotoActionGotoObjectRequiredStringFalse = shared.GotoActionGotoObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type GotoActionGotoObjectSkipUnionParam = shared.GotoActionGotoObjectSkipUnionParam

// This is an alias to an internal type.
type GotoActionGotoObjectSkipString = shared.GotoActionGotoObjectSkipString

// Equals "true"
const GotoActionGotoObjectSkipStringTrue = shared.GotoActionGotoObjectSkipStringTrue

// Equals "false"
const GotoActionGotoObjectSkipStringFalse = shared.GotoActionGotoObjectSkipStringFalse

// Press a keyboard key
//
// This is an alias to an internal type.
type PressActionParam = shared.PressActionParam

// This is an alias to an internal type.
type PressActionPressUnionParam = shared.PressActionPressUnionParam

// This is an alias to an internal type.
type PressActionPressObjectParam = shared.PressActionPressObjectParam

// Duration value that accepts various formats. Supports: number (ms), string
// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
//
// This is an alias to an internal type.
type PressActionPressObjectDelayUnionParam = shared.PressActionPressObjectDelayUnionParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type PressActionPressObjectRequiredUnionParam = shared.PressActionPressObjectRequiredUnionParam

// This is an alias to an internal type.
type PressActionPressObjectRequiredString = shared.PressActionPressObjectRequiredString

// Equals "true"
const PressActionPressObjectRequiredStringTrue = shared.PressActionPressObjectRequiredStringTrue

// Equals "false"
const PressActionPressObjectRequiredStringFalse = shared.PressActionPressObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type PressActionPressObjectSkipUnionParam = shared.PressActionPressObjectSkipUnionParam

// This is an alias to an internal type.
type PressActionPressObjectSkipString = shared.PressActionPressObjectSkipString

// Equals "true"
const PressActionPressObjectSkipStringTrue = shared.PressActionPressObjectSkipStringTrue

// Equals "false"
const PressActionPressObjectSkipStringFalse = shared.PressActionPressObjectSkipStringFalse

// Capture a page screenshot
//
// This is an alias to an internal type.
type ScreenshotActionParam = shared.ScreenshotActionParam

// This is an alias to an internal type.
type ScreenshotActionScreenshotUnionParam = shared.ScreenshotActionScreenshotUnionParam

// This is an alias to an internal type.
type ScreenshotActionScreenshotObjectParam = shared.ScreenshotActionScreenshotObjectParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type ScreenshotActionScreenshotObjectRequiredUnionParam = shared.ScreenshotActionScreenshotObjectRequiredUnionParam

// This is an alias to an internal type.
type ScreenshotActionScreenshotObjectRequiredString = shared.ScreenshotActionScreenshotObjectRequiredString

// Equals "true"
const ScreenshotActionScreenshotObjectRequiredStringTrue = shared.ScreenshotActionScreenshotObjectRequiredStringTrue

// Equals "false"
const ScreenshotActionScreenshotObjectRequiredStringFalse = shared.ScreenshotActionScreenshotObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type ScreenshotActionScreenshotObjectSkipUnionParam = shared.ScreenshotActionScreenshotObjectSkipUnionParam

// This is an alias to an internal type.
type ScreenshotActionScreenshotObjectSkipString = shared.ScreenshotActionScreenshotObjectSkipString

// Equals "true"
const ScreenshotActionScreenshotObjectSkipStringTrue = shared.ScreenshotActionScreenshotObjectSkipStringTrue

// Equals "false"
const ScreenshotActionScreenshotObjectSkipStringFalse = shared.ScreenshotActionScreenshotObjectSkipStringFalse

// Scroll the page or an element
//
// This is an alias to an internal type.
type ScrollActionParam = shared.ScrollActionParam

// This is an alias to an internal type.
type ScrollActionScrollUnionParam = shared.ScrollActionScrollUnionParam

// This is an alias to an internal type.
type ScrollActionScrollObjectParam = shared.ScrollActionScrollObjectParam

// CSS selector or array of alternative selectors. Use an array when you have
// multiple possible selectors for the same element.
//
// This is an alias to an internal type.
type ScrollActionScrollObjectContainerUnionParam = shared.ScrollActionScrollObjectContainerUnionParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type ScrollActionScrollObjectRequiredUnionParam = shared.ScrollActionScrollObjectRequiredUnionParam

// This is an alias to an internal type.
type ScrollActionScrollObjectRequiredString = shared.ScrollActionScrollObjectRequiredString

// Equals "true"
const ScrollActionScrollObjectRequiredStringTrue = shared.ScrollActionScrollObjectRequiredStringTrue

// Equals "false"
const ScrollActionScrollObjectRequiredStringFalse = shared.ScrollActionScrollObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type ScrollActionScrollObjectSkipUnionParam = shared.ScrollActionScrollObjectSkipUnionParam

// This is an alias to an internal type.
type ScrollActionScrollObjectSkipString = shared.ScrollActionScrollObjectSkipString

// Equals "true"
const ScrollActionScrollObjectSkipStringTrue = shared.ScrollActionScrollObjectSkipStringTrue

// Equals "false"
const ScrollActionScrollObjectSkipStringFalse = shared.ScrollActionScrollObjectSkipStringFalse

// CSS selector or array of alternative selectors. Use an array when you have
// multiple possible selectors for the same element.
//
// This is an alias to an internal type.
type ScrollActionScrollObjectToUnionParam = shared.ScrollActionScrollObjectToUnionParam

// Wait for a specified duration
//
// This is an alias to an internal type.
type WaitActionParam = shared.WaitActionParam

// This is an alias to an internal type.
type WaitActionWaitUnionParam = shared.WaitActionWaitUnionParam

// This is an alias to an internal type.
type WaitActionWaitObjectParam = shared.WaitActionWaitObjectParam

// Duration value that accepts various formats. Supports: number (ms), string
// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
//
// This is an alias to an internal type.
type WaitActionWaitObjectDurationUnionParam = shared.WaitActionWaitObjectDurationUnionParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type WaitActionWaitObjectRequiredUnionParam = shared.WaitActionWaitObjectRequiredUnionParam

// This is an alias to an internal type.
type WaitActionWaitObjectRequiredString = shared.WaitActionWaitObjectRequiredString

// Equals "true"
const WaitActionWaitObjectRequiredStringTrue = shared.WaitActionWaitObjectRequiredStringTrue

// Equals "false"
const WaitActionWaitObjectRequiredStringFalse = shared.WaitActionWaitObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type WaitActionWaitObjectSkipUnionParam = shared.WaitActionWaitObjectSkipUnionParam

// This is an alias to an internal type.
type WaitActionWaitObjectSkipString = shared.WaitActionWaitObjectSkipString

// Equals "true"
const WaitActionWaitObjectSkipStringTrue = shared.WaitActionWaitObjectSkipStringTrue

// Equals "false"
const WaitActionWaitObjectSkipStringFalse = shared.WaitActionWaitObjectSkipStringFalse

// Wait for an element to appear or reach a specific state
//
// This is an alias to an internal type.
type WaitForElementActionParam = shared.WaitForElementActionParam

// This is an alias to an internal type.
type WaitForElementActionWaitForElementUnionParam = shared.WaitForElementActionWaitForElementUnionParam

// This is an alias to an internal type.
type WaitForElementActionWaitForElementObjectParam = shared.WaitForElementActionWaitForElementObjectParam

// CSS selector or array of alternative selectors. Use an array when you have
// multiple possible selectors for the same element.
//
// This is an alias to an internal type.
type WaitForElementActionWaitForElementObjectSelectorUnionParam = shared.WaitForElementActionWaitForElementObjectSelectorUnionParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type WaitForElementActionWaitForElementObjectRequiredUnionParam = shared.WaitForElementActionWaitForElementObjectRequiredUnionParam

// This is an alias to an internal type.
type WaitForElementActionWaitForElementObjectRequiredString = shared.WaitForElementActionWaitForElementObjectRequiredString

// Equals "true"
const WaitForElementActionWaitForElementObjectRequiredStringTrue = shared.WaitForElementActionWaitForElementObjectRequiredStringTrue

// Equals "false"
const WaitForElementActionWaitForElementObjectRequiredStringFalse = shared.WaitForElementActionWaitForElementObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type WaitForElementActionWaitForElementObjectSkipUnionParam = shared.WaitForElementActionWaitForElementObjectSkipUnionParam

// This is an alias to an internal type.
type WaitForElementActionWaitForElementObjectSkipString = shared.WaitForElementActionWaitForElementObjectSkipString

// Equals "true"
const WaitForElementActionWaitForElementObjectSkipStringTrue = shared.WaitForElementActionWaitForElementObjectSkipStringTrue

// Equals "false"
const WaitForElementActionWaitForElementObjectSkipStringFalse = shared.WaitForElementActionWaitForElementObjectSkipStringFalse

// Wait for page navigation to complete
//
// This is an alias to an internal type.
type WaitForNavigationActionParam = shared.WaitForNavigationActionParam

// This is an alias to an internal type.
type WaitForNavigationActionWaitForNavigationUnionParam = shared.WaitForNavigationActionWaitForNavigationUnionParam

// This is an alias to an internal type.
type WaitForNavigationActionWaitForNavigationString = shared.WaitForNavigationActionWaitForNavigationString

// This is an alias to an internal type.
type WaitForNavigationActionWaitForNavigationObjectParam = shared.WaitForNavigationActionWaitForNavigationObjectParam

// Whether this action is required. If true, pipeline stops on failure. Accepts
// boolean or string "true"/"false". Default: true.
//
// This is an alias to an internal type.
type WaitForNavigationActionWaitForNavigationObjectRequiredUnionParam = shared.WaitForNavigationActionWaitForNavigationObjectRequiredUnionParam

// This is an alias to an internal type.
type WaitForNavigationActionWaitForNavigationObjectRequiredString = shared.WaitForNavigationActionWaitForNavigationObjectRequiredString

// Equals "true"
const WaitForNavigationActionWaitForNavigationObjectRequiredStringTrue = shared.WaitForNavigationActionWaitForNavigationObjectRequiredStringTrue

// Equals "false"
const WaitForNavigationActionWaitForNavigationObjectRequiredStringFalse = shared.WaitForNavigationActionWaitForNavigationObjectRequiredStringFalse

// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
// false.
//
// This is an alias to an internal type.
type WaitForNavigationActionWaitForNavigationObjectSkipUnionParam = shared.WaitForNavigationActionWaitForNavigationObjectSkipUnionParam

// This is an alias to an internal type.
type WaitForNavigationActionWaitForNavigationObjectSkipString = shared.WaitForNavigationActionWaitForNavigationObjectSkipString

// Equals "true"
const WaitForNavigationActionWaitForNavigationObjectSkipStringTrue = shared.WaitForNavigationActionWaitForNavigationObjectSkipStringTrue

// Equals "false"
const WaitForNavigationActionWaitForNavigationObjectSkipStringFalse = shared.WaitForNavigationActionWaitForNavigationObjectSkipStringFalse
