// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Continuously scroll to load dynamic content
//
// The property AutoScroll is required.
type AutoScrollActionParam struct {
	AutoScroll AutoScrollActionAutoScrollUnionParam `json:"auto_scroll,omitzero" api:"required"`
	paramObj
}

func (r AutoScrollActionParam) MarshalJSON() (data []byte, err error) {
	type shadow AutoScrollActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutoScrollActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutoScrollActionAutoScrollUnionParam struct {
	OfBool                             param.Opt[bool]                        `json:",omitzero,inline"`
	OfFloat                            param.Opt[float64]                     `json:",omitzero,inline"`
	OfString                           param.Opt[string]                      `json:",omitzero,inline"`
	OfAutoScrollActionAutoScrollObject *AutoScrollActionAutoScrollObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u AutoScrollActionAutoScrollUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAutoScrollActionAutoScrollObject)
}
func (u *AutoScrollActionAutoScrollUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutoScrollActionAutoScrollUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAutoScrollActionAutoScrollObject) {
		return u.OfAutoScrollActionAutoScrollObject
	}
	return nil
}

type AutoScrollActionAutoScrollObjectParam struct {
	StepSize param.Opt[float64] `json:"step_size,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	ClickSelector AutoScrollActionAutoScrollObjectClickSelectorUnionParam `json:"click_selector,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container AutoScrollActionAutoScrollObjectContainerUnionParam `json:"container,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	DelayAfterScroll AutoScrollActionAutoScrollObjectDelayAfterScrollUnionParam `json:"delay_after_scroll,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	IdleTimeout AutoScrollActionAutoScrollObjectIdleTimeoutUnionParam `json:"idle_timeout,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	LoadingSelector AutoScrollActionAutoScrollObjectLoadingSelectorUnionParam `json:"loading_selector,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	MaxDuration AutoScrollActionAutoScrollObjectMaxDurationUnionParam `json:"max_duration,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	PauseOnSelector AutoScrollActionAutoScrollObjectPauseOnSelectorUnionParam `json:"pause_on_selector,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required AutoScrollActionAutoScrollObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip AutoScrollActionAutoScrollObjectSkipUnionParam `json:"skip,omitzero"`
	paramObj
}

func (r AutoScrollActionAutoScrollObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow AutoScrollActionAutoScrollObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutoScrollActionAutoScrollObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutoScrollActionAutoScrollObjectClickSelectorUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AutoScrollActionAutoScrollObjectClickSelectorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AutoScrollActionAutoScrollObjectClickSelectorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutoScrollActionAutoScrollObjectClickSelectorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutoScrollActionAutoScrollObjectContainerUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AutoScrollActionAutoScrollObjectContainerUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AutoScrollActionAutoScrollObjectContainerUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutoScrollActionAutoScrollObjectContainerUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutoScrollActionAutoScrollObjectDelayAfterScrollUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u AutoScrollActionAutoScrollObjectDelayAfterScrollUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *AutoScrollActionAutoScrollObjectDelayAfterScrollUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutoScrollActionAutoScrollObjectDelayAfterScrollUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutoScrollActionAutoScrollObjectIdleTimeoutUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u AutoScrollActionAutoScrollObjectIdleTimeoutUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *AutoScrollActionAutoScrollObjectIdleTimeoutUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutoScrollActionAutoScrollObjectIdleTimeoutUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutoScrollActionAutoScrollObjectLoadingSelectorUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AutoScrollActionAutoScrollObjectLoadingSelectorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AutoScrollActionAutoScrollObjectLoadingSelectorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutoScrollActionAutoScrollObjectLoadingSelectorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutoScrollActionAutoScrollObjectMaxDurationUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u AutoScrollActionAutoScrollObjectMaxDurationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *AutoScrollActionAutoScrollObjectMaxDurationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutoScrollActionAutoScrollObjectMaxDurationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutoScrollActionAutoScrollObjectPauseOnSelectorUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AutoScrollActionAutoScrollObjectPauseOnSelectorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AutoScrollActionAutoScrollObjectPauseOnSelectorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutoScrollActionAutoScrollObjectPauseOnSelectorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutoScrollActionAutoScrollObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfAutoScrollActionAutoScrollObjectRequiredString)
	OfAutoScrollActionAutoScrollObjectRequiredString param.Opt[AutoScrollActionAutoScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                           param.Opt[bool]                                           `json:",omitzero,inline"`
	paramUnion
}

func (u AutoScrollActionAutoScrollObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutoScrollActionAutoScrollObjectRequiredString, u.OfBool)
}
func (u *AutoScrollActionAutoScrollObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutoScrollActionAutoScrollObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAutoScrollActionAutoScrollObjectRequiredString) {
		return &u.OfAutoScrollActionAutoScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type AutoScrollActionAutoScrollObjectRequiredString string

const (
	AutoScrollActionAutoScrollObjectRequiredStringTrue  AutoScrollActionAutoScrollObjectRequiredString = "true"
	AutoScrollActionAutoScrollObjectRequiredStringFalse AutoScrollActionAutoScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutoScrollActionAutoScrollObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfAutoScrollActionAutoScrollObjectSkipString)
	OfAutoScrollActionAutoScrollObjectSkipString param.Opt[AutoScrollActionAutoScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                                       param.Opt[bool]                                       `json:",omitzero,inline"`
	paramUnion
}

func (u AutoScrollActionAutoScrollObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutoScrollActionAutoScrollObjectSkipString, u.OfBool)
}
func (u *AutoScrollActionAutoScrollObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutoScrollActionAutoScrollObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAutoScrollActionAutoScrollObjectSkipString) {
		return &u.OfAutoScrollActionAutoScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type AutoScrollActionAutoScrollObjectSkipString string

const (
	AutoScrollActionAutoScrollObjectSkipStringTrue  AutoScrollActionAutoScrollObjectSkipString = "true"
	AutoScrollActionAutoScrollObjectSkipStringFalse AutoScrollActionAutoScrollObjectSkipString = "false"
)

// Click on an element by selector
//
// The property Click is required.
type ClickActionParam struct {
	Click ClickActionClickUnionParam `json:"click,omitzero" api:"required"`
	paramObj
}

func (r ClickActionParam) MarshalJSON() (data []byte, err error) {
	type shadow ClickActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClickActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ClickActionClickUnionParam struct {
	OfString                 param.Opt[string]            `json:",omitzero,inline"`
	OfStringArray            []string                     `json:",omitzero,inline"`
	OfClickActionClickObject *ClickActionClickObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u ClickActionClickUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfClickActionClickObject)
}
func (u *ClickActionClickUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ClickActionClickUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfClickActionClickObject) {
		return u.OfClickActionClickObject
	}
	return nil
}

// The property Selector is required.
type ClickActionClickObjectParam struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector ClickActionClickObjectSelectorUnionParam `json:"selector,omitzero" api:"required"`
	Count    param.Opt[float64]                       `json:"count,omitzero"`
	OffsetX  param.Opt[int64]                         `json:"offset_x,omitzero"`
	OffsetY  param.Opt[int64]                         `json:"offset_y,omitzero"`
	Scroll   param.Opt[bool]                          `json:"scroll,omitzero"`
	Steps    param.Opt[float64]                       `json:"steps,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ClickActionClickObjectDelayUnionParam `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ClickActionClickObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ClickActionClickObjectSkipUnionParam `json:"skip,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	Strategy string `json:"strategy,omitzero"`
	paramObj
}

func (r ClickActionClickObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow ClickActionClickObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClickActionClickObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ClickActionClickObjectParam](
		"strategy", "linear", "ghost-cursor", "windmouse",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ClickActionClickObjectSelectorUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ClickActionClickObjectSelectorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ClickActionClickObjectSelectorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ClickActionClickObjectSelectorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ClickActionClickObjectDelayUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ClickActionClickObjectDelayUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ClickActionClickObjectDelayUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ClickActionClickObjectDelayUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ClickActionClickObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfClickActionClickObjectRequiredString)
	OfClickActionClickObjectRequiredString param.Opt[ClickActionClickObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                 param.Opt[bool]                                 `json:",omitzero,inline"`
	paramUnion
}

func (u ClickActionClickObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfClickActionClickObjectRequiredString, u.OfBool)
}
func (u *ClickActionClickObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ClickActionClickObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfClickActionClickObjectRequiredString) {
		return &u.OfClickActionClickObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ClickActionClickObjectRequiredString string

const (
	ClickActionClickObjectRequiredStringTrue  ClickActionClickObjectRequiredString = "true"
	ClickActionClickObjectRequiredStringFalse ClickActionClickObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ClickActionClickObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfClickActionClickObjectSkipString)
	OfClickActionClickObjectSkipString param.Opt[ClickActionClickObjectSkipString] `json:",omitzero,inline"`
	OfBool                             param.Opt[bool]                             `json:",omitzero,inline"`
	paramUnion
}

func (u ClickActionClickObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfClickActionClickObjectSkipString, u.OfBool)
}
func (u *ClickActionClickObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ClickActionClickObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfClickActionClickObjectSkipString) {
		return &u.OfClickActionClickObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ClickActionClickObjectSkipString string

const (
	ClickActionClickObjectSkipStringTrue  ClickActionClickObjectSkipString = "true"
	ClickActionClickObjectSkipStringFalse ClickActionClickObjectSkipString = "false"
)

// Execute JavaScript code in page context
//
// The property Eval is required.
type EvalActionParam struct {
	Eval EvalActionEvalUnionParam `json:"eval,omitzero" api:"required"`
	paramObj
}

func (r EvalActionParam) MarshalJSON() (data []byte, err error) {
	type shadow EvalActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalActionEvalUnionParam struct {
	OfString               param.Opt[string]          `json:",omitzero,inline"`
	OfEvalActionEvalObject *EvalActionEvalObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u EvalActionEvalUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfEvalActionEvalObject)
}
func (u *EvalActionEvalUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalActionEvalUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfEvalActionEvalObject) {
		return u.OfEvalActionEvalObject
	}
	return nil
}

// The property Code is required.
type EvalActionEvalObjectParam struct {
	Code string `json:"code" api:"required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required EvalActionEvalObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip EvalActionEvalObjectSkipUnionParam `json:"skip,omitzero"`
	paramObj
}

func (r EvalActionEvalObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow EvalActionEvalObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalActionEvalObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalActionEvalObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfEvalActionEvalObjectRequiredString)
	OfEvalActionEvalObjectRequiredString param.Opt[EvalActionEvalObjectRequiredString] `json:",omitzero,inline"`
	OfBool                               param.Opt[bool]                               `json:",omitzero,inline"`
	paramUnion
}

func (u EvalActionEvalObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEvalActionEvalObjectRequiredString, u.OfBool)
}
func (u *EvalActionEvalObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalActionEvalObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfEvalActionEvalObjectRequiredString) {
		return &u.OfEvalActionEvalObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type EvalActionEvalObjectRequiredString string

const (
	EvalActionEvalObjectRequiredStringTrue  EvalActionEvalObjectRequiredString = "true"
	EvalActionEvalObjectRequiredStringFalse EvalActionEvalObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalActionEvalObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfEvalActionEvalObjectSkipString)
	OfEvalActionEvalObjectSkipString param.Opt[EvalActionEvalObjectSkipString] `json:",omitzero,inline"`
	OfBool                           param.Opt[bool]                           `json:",omitzero,inline"`
	paramUnion
}

func (u EvalActionEvalObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEvalActionEvalObjectSkipString, u.OfBool)
}
func (u *EvalActionEvalObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalActionEvalObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfEvalActionEvalObjectSkipString) {
		return &u.OfEvalActionEvalObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type EvalActionEvalObjectSkipString string

const (
	EvalActionEvalObjectSkipStringTrue  EvalActionEvalObjectSkipString = "true"
	EvalActionEvalObjectSkipStringFalse EvalActionEvalObjectSkipString = "false"
)

// Make an HTTP request in browser context
//
// The property Fetch is required.
type FetchActionParam struct {
	Fetch FetchActionFetchUnionParam `json:"fetch,omitzero" api:"required" format:"uri"`
	paramObj
}

func (r FetchActionParam) MarshalJSON() (data []byte, err error) {
	type shadow FetchActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FetchActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FetchActionFetchUnionParam struct {
	OfString                 param.Opt[string]            `json:",omitzero,inline"`
	OfFetchActionFetchObject *FetchActionFetchObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u FetchActionFetchUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFetchActionFetchObject)
}
func (u *FetchActionFetchUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FetchActionFetchUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFetchActionFetchObject) {
		return u.OfFetchActionFetchObject
	}
	return nil
}

// The property URL is required.
type FetchActionFetchObjectParam struct {
	URL  string            `json:"url" api:"required" format:"uri"`
	Body param.Opt[string] `json:"body,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Headers map[string]string  `json:"headers,omitzero"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required FetchActionFetchObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip FetchActionFetchObjectSkipUnionParam `json:"skip,omitzero"`
	paramObj
}

func (r FetchActionFetchObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow FetchActionFetchObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FetchActionFetchObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FetchActionFetchObjectParam](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FetchActionFetchObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfFetchActionFetchObjectRequiredString)
	OfFetchActionFetchObjectRequiredString param.Opt[FetchActionFetchObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                 param.Opt[bool]                                 `json:",omitzero,inline"`
	paramUnion
}

func (u FetchActionFetchObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFetchActionFetchObjectRequiredString, u.OfBool)
}
func (u *FetchActionFetchObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FetchActionFetchObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFetchActionFetchObjectRequiredString) {
		return &u.OfFetchActionFetchObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type FetchActionFetchObjectRequiredString string

const (
	FetchActionFetchObjectRequiredStringTrue  FetchActionFetchObjectRequiredString = "true"
	FetchActionFetchObjectRequiredStringFalse FetchActionFetchObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FetchActionFetchObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfFetchActionFetchObjectSkipString)
	OfFetchActionFetchObjectSkipString param.Opt[FetchActionFetchObjectSkipString] `json:",omitzero,inline"`
	OfBool                             param.Opt[bool]                             `json:",omitzero,inline"`
	paramUnion
}

func (u FetchActionFetchObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFetchActionFetchObjectSkipString, u.OfBool)
}
func (u *FetchActionFetchObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FetchActionFetchObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFetchActionFetchObjectSkipString) {
		return &u.OfFetchActionFetchObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type FetchActionFetchObjectSkipString string

const (
	FetchActionFetchObjectSkipStringTrue  FetchActionFetchObjectSkipString = "true"
	FetchActionFetchObjectSkipStringFalse FetchActionFetchObjectSkipString = "false"
)

// Fill text into an input field
//
// The property Fill is required.
type FillActionParam struct {
	// Fill options with mode-specific fields. Use "type" mode for behavioral typing
	// simulation, or "paste" mode for instant paste.
	Fill FillActionFillUnionParam `json:"fill,omitzero" api:"required"`
	paramObj
}

func (r FillActionParam) MarshalJSON() (data []byte, err error) {
	type shadow FillActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FillActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FillActionFillUnionParam struct {
	OfType  *FillActionFillTypeParam  `json:",omitzero,inline"`
	OfPaste *FillActionFillPasteParam `json:",omitzero,inline"`
	paramUnion
}

func (u FillActionFillUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfType, u.OfPaste)
}
func (u *FillActionFillUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FillActionFillUnionParam) asAny() any {
	if !param.IsOmitted(u.OfType) {
		return u.OfType
	} else if !param.IsOmitted(u.OfPaste) {
		return u.OfPaste
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FillActionFillUnionParam) GetMouseMovementStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.MouseMovementStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FillActionFillUnionParam) GetTypingInterval() *FillActionFillTypeTypingIntervalUnionParam {
	if vt := u.OfType; vt != nil {
		return &vt.TypingInterval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FillActionFillUnionParam) GetTypingStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.TypingStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FillActionFillUnionParam) GetValue() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FillActionFillUnionParam) GetClickOnElement() *bool {
	if vt := u.OfType; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	} else if vt := u.OfPaste; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FillActionFillUnionParam) GetMode() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Mode)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Mode)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FillActionFillUnionParam) GetScroll() *bool {
	if vt := u.OfType; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	} else if vt := u.OfPaste; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FillActionFillUnionParam) GetTimeout() *float64 {
	if vt := u.OfType; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	} else if vt := u.OfPaste; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FillActionFillUnionParam) GetVisible() *bool {
	if vt := u.OfType; vt != nil && vt.Visible.Valid() {
		return &vt.Visible.Value
	} else if vt := u.OfPaste; vt != nil && vt.Visible.Valid() {
		return &vt.Visible.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FillActionFillUnionParam) GetSelector() (res fillActionFillUnionParamSelector) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Selector.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Selector.asAny()
	}
	return
}

// Can have the runtime types [*string], [\*[]string]
type fillActionFillUnionParamSelector struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u fillActionFillUnionParamSelector) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FillActionFillUnionParam) GetDelay() (res fillActionFillUnionParamDelay) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Delay.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Delay.asAny()
	}
	return
}

// Can have the runtime types [*float64], [*string]
type fillActionFillUnionParamDelay struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *float64:
//	case *string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u fillActionFillUnionParamDelay) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FillActionFillUnionParam) GetRequired() (res fillActionFillUnionParamRequired) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Required.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Required.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type fillActionFillUnionParamRequired struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u fillActionFillUnionParamRequired) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u FillActionFillUnionParam) GetSkip() (res fillActionFillUnionParamSkip) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Skip.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Skip.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type fillActionFillUnionParamSkip struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u fillActionFillUnionParamSkip) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[FillActionFillUnionParam](
		"mode",
		apijson.Discriminator[FillActionFillTypeParam]("type"),
		apijson.Discriminator[FillActionFillPasteParam]("paste"),
	)
}

// The properties Selector, Value are required.
type FillActionFillTypeParam struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       FillActionFillTypeSelectorUnionParam `json:"selector,omitzero" api:"required"`
	Value          string                               `json:"value" api:"required"`
	ClickOnElement param.Opt[bool]                      `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                      `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay FillActionFillTypeDelayUnionParam `json:"delay,omitzero"`
	// Any of "type".
	Mode string `json:"mode,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	MouseMovementStrategy string `json:"mouse_movement_strategy,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required FillActionFillTypeRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip FillActionFillTypeSkipUnionParam `json:"skip,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	TypingInterval FillActionFillTypeTypingIntervalUnionParam `json:"typing_interval,omitzero"`
	// Any of "simple", "distribution".
	TypingStrategy string `json:"typing_strategy,omitzero"`
	paramObj
}

func (r FillActionFillTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow FillActionFillTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FillActionFillTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FillActionFillTypeParam](
		"mode", "type",
	)
	apijson.RegisterFieldValidator[FillActionFillTypeParam](
		"mouse_movement_strategy", "linear", "ghost-cursor", "windmouse",
	)
	apijson.RegisterFieldValidator[FillActionFillTypeParam](
		"typing_strategy", "simple", "distribution",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FillActionFillTypeSelectorUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u FillActionFillTypeSelectorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *FillActionFillTypeSelectorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FillActionFillTypeSelectorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FillActionFillTypeDelayUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u FillActionFillTypeDelayUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *FillActionFillTypeDelayUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FillActionFillTypeDelayUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FillActionFillTypeRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfFillActionFillTypeRequiredString)
	OfFillActionFillTypeRequiredString param.Opt[FillActionFillTypeRequiredString] `json:",omitzero,inline"`
	OfBool                             param.Opt[bool]                             `json:",omitzero,inline"`
	paramUnion
}

func (u FillActionFillTypeRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFillActionFillTypeRequiredString, u.OfBool)
}
func (u *FillActionFillTypeRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FillActionFillTypeRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFillActionFillTypeRequiredString) {
		return &u.OfFillActionFillTypeRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type FillActionFillTypeRequiredString string

const (
	FillActionFillTypeRequiredStringTrue  FillActionFillTypeRequiredString = "true"
	FillActionFillTypeRequiredStringFalse FillActionFillTypeRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FillActionFillTypeSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfFillActionFillTypeSkipString)
	OfFillActionFillTypeSkipString param.Opt[FillActionFillTypeSkipString] `json:",omitzero,inline"`
	OfBool                         param.Opt[bool]                         `json:",omitzero,inline"`
	paramUnion
}

func (u FillActionFillTypeSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFillActionFillTypeSkipString, u.OfBool)
}
func (u *FillActionFillTypeSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FillActionFillTypeSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFillActionFillTypeSkipString) {
		return &u.OfFillActionFillTypeSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type FillActionFillTypeSkipString string

const (
	FillActionFillTypeSkipStringTrue  FillActionFillTypeSkipString = "true"
	FillActionFillTypeSkipStringFalse FillActionFillTypeSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FillActionFillTypeTypingIntervalUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u FillActionFillTypeTypingIntervalUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *FillActionFillTypeTypingIntervalUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FillActionFillTypeTypingIntervalUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The properties Mode, Selector, Value are required.
type FillActionFillPasteParam struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       FillActionFillPasteSelectorUnionParam `json:"selector,omitzero" api:"required"`
	Value          string                                `json:"value" api:"required"`
	ClickOnElement param.Opt[bool]                       `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                       `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay FillActionFillPasteDelayUnionParam `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required FillActionFillPasteRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip FillActionFillPasteSkipUnionParam `json:"skip,omitzero"`
	// This field can be elided, and will marshal its zero value as "paste".
	Mode constant.Paste `json:"mode" api:"required"`
	paramObj
}

func (r FillActionFillPasteParam) MarshalJSON() (data []byte, err error) {
	type shadow FillActionFillPasteParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FillActionFillPasteParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FillActionFillPasteSelectorUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u FillActionFillPasteSelectorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *FillActionFillPasteSelectorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FillActionFillPasteSelectorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FillActionFillPasteDelayUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u FillActionFillPasteDelayUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *FillActionFillPasteDelayUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FillActionFillPasteDelayUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FillActionFillPasteRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfFillActionFillPasteRequiredString)
	OfFillActionFillPasteRequiredString param.Opt[FillActionFillPasteRequiredString] `json:",omitzero,inline"`
	OfBool                              param.Opt[bool]                              `json:",omitzero,inline"`
	paramUnion
}

func (u FillActionFillPasteRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFillActionFillPasteRequiredString, u.OfBool)
}
func (u *FillActionFillPasteRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FillActionFillPasteRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFillActionFillPasteRequiredString) {
		return &u.OfFillActionFillPasteRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type FillActionFillPasteRequiredString string

const (
	FillActionFillPasteRequiredStringTrue  FillActionFillPasteRequiredString = "true"
	FillActionFillPasteRequiredStringFalse FillActionFillPasteRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FillActionFillPasteSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfFillActionFillPasteSkipString)
	OfFillActionFillPasteSkipString param.Opt[FillActionFillPasteSkipString] `json:",omitzero,inline"`
	OfBool                          param.Opt[bool]                          `json:",omitzero,inline"`
	paramUnion
}

func (u FillActionFillPasteSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFillActionFillPasteSkipString, u.OfBool)
}
func (u *FillActionFillPasteSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FillActionFillPasteSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFillActionFillPasteSkipString) {
		return &u.OfFillActionFillPasteSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type FillActionFillPasteSkipString string

const (
	FillActionFillPasteSkipStringTrue  FillActionFillPasteSkipString = "true"
	FillActionFillPasteSkipStringFalse FillActionFillPasteSkipString = "false"
)

// Retrieve browser cookies
//
// The property GetCookies is required.
type GetCookiesActionParam struct {
	GetCookies GetCookiesActionGetCookiesUnionParam `json:"get_cookies,omitzero" api:"required"`
	paramObj
}

func (r GetCookiesActionParam) MarshalJSON() (data []byte, err error) {
	type shadow GetCookiesActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetCookiesActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GetCookiesActionGetCookiesUnionParam struct {
	OfBool                             param.Opt[bool]                        `json:",omitzero,inline"`
	OfGetCookiesActionGetCookiesObject *GetCookiesActionGetCookiesObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u GetCookiesActionGetCookiesUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfGetCookiesActionGetCookiesObject)
}
func (u *GetCookiesActionGetCookiesUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GetCookiesActionGetCookiesUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfGetCookiesActionGetCookiesObject) {
		return u.OfGetCookiesActionGetCookiesObject
	}
	return nil
}

type GetCookiesActionGetCookiesObjectParam struct {
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required GetCookiesActionGetCookiesObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip        GetCookiesActionGetCookiesObjectSkipUnionParam `json:"skip,omitzero"`
	ExtraFields map[string]any                                 `json:"-"`
	paramObj
}

func (r GetCookiesActionGetCookiesObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow GetCookiesActionGetCookiesObjectParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *GetCookiesActionGetCookiesObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GetCookiesActionGetCookiesObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfGetCookiesActionGetCookiesObjectRequiredString)
	OfGetCookiesActionGetCookiesObjectRequiredString param.Opt[GetCookiesActionGetCookiesObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                           param.Opt[bool]                                           `json:",omitzero,inline"`
	paramUnion
}

func (u GetCookiesActionGetCookiesObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGetCookiesActionGetCookiesObjectRequiredString, u.OfBool)
}
func (u *GetCookiesActionGetCookiesObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GetCookiesActionGetCookiesObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfGetCookiesActionGetCookiesObjectRequiredString) {
		return &u.OfGetCookiesActionGetCookiesObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type GetCookiesActionGetCookiesObjectRequiredString string

const (
	GetCookiesActionGetCookiesObjectRequiredStringTrue  GetCookiesActionGetCookiesObjectRequiredString = "true"
	GetCookiesActionGetCookiesObjectRequiredStringFalse GetCookiesActionGetCookiesObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GetCookiesActionGetCookiesObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfGetCookiesActionGetCookiesObjectSkipString)
	OfGetCookiesActionGetCookiesObjectSkipString param.Opt[GetCookiesActionGetCookiesObjectSkipString] `json:",omitzero,inline"`
	OfBool                                       param.Opt[bool]                                       `json:",omitzero,inline"`
	paramUnion
}

func (u GetCookiesActionGetCookiesObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGetCookiesActionGetCookiesObjectSkipString, u.OfBool)
}
func (u *GetCookiesActionGetCookiesObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GetCookiesActionGetCookiesObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfGetCookiesActionGetCookiesObjectSkipString) {
		return &u.OfGetCookiesActionGetCookiesObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type GetCookiesActionGetCookiesObjectSkipString string

const (
	GetCookiesActionGetCookiesObjectSkipStringTrue  GetCookiesActionGetCookiesObjectSkipString = "true"
	GetCookiesActionGetCookiesObjectSkipStringFalse GetCookiesActionGetCookiesObjectSkipString = "false"
)

// Navigate to a URL
//
// The property Goto is required.
type GotoActionParam struct {
	Goto GotoActionGotoUnionParam `json:"goto,omitzero" api:"required" format:"uri"`
	paramObj
}

func (r GotoActionParam) MarshalJSON() (data []byte, err error) {
	type shadow GotoActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GotoActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GotoActionGotoUnionParam struct {
	OfString               param.Opt[string]          `json:",omitzero,inline"`
	OfGotoActionGotoObject *GotoActionGotoObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u GotoActionGotoUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfGotoActionGotoObject)
}
func (u *GotoActionGotoUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GotoActionGotoUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfGotoActionGotoObject) {
		return u.OfGotoActionGotoObject
	}
	return nil
}

// The property URL is required.
type GotoActionGotoObjectParam struct {
	URL     string            `json:"url" api:"required" format:"uri"`
	Referer param.Opt[string] `json:"referer,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required GotoActionGotoObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip GotoActionGotoObjectSkipUnionParam `json:"skip,omitzero"`
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	WaitUntil string `json:"wait_until,omitzero"`
	paramObj
}

func (r GotoActionGotoObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow GotoActionGotoObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GotoActionGotoObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GotoActionGotoObjectParam](
		"wait_until", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GotoActionGotoObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfGotoActionGotoObjectRequiredString)
	OfGotoActionGotoObjectRequiredString param.Opt[GotoActionGotoObjectRequiredString] `json:",omitzero,inline"`
	OfBool                               param.Opt[bool]                               `json:",omitzero,inline"`
	paramUnion
}

func (u GotoActionGotoObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGotoActionGotoObjectRequiredString, u.OfBool)
}
func (u *GotoActionGotoObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GotoActionGotoObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfGotoActionGotoObjectRequiredString) {
		return &u.OfGotoActionGotoObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type GotoActionGotoObjectRequiredString string

const (
	GotoActionGotoObjectRequiredStringTrue  GotoActionGotoObjectRequiredString = "true"
	GotoActionGotoObjectRequiredStringFalse GotoActionGotoObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GotoActionGotoObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfGotoActionGotoObjectSkipString)
	OfGotoActionGotoObjectSkipString param.Opt[GotoActionGotoObjectSkipString] `json:",omitzero,inline"`
	OfBool                           param.Opt[bool]                           `json:",omitzero,inline"`
	paramUnion
}

func (u GotoActionGotoObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGotoActionGotoObjectSkipString, u.OfBool)
}
func (u *GotoActionGotoObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GotoActionGotoObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfGotoActionGotoObjectSkipString) {
		return &u.OfGotoActionGotoObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type GotoActionGotoObjectSkipString string

const (
	GotoActionGotoObjectSkipStringTrue  GotoActionGotoObjectSkipString = "true"
	GotoActionGotoObjectSkipStringFalse GotoActionGotoObjectSkipString = "false"
)

// Press a keyboard key
//
// The property Press is required.
type PressActionParam struct {
	Press PressActionPressUnionParam `json:"press,omitzero" api:"required"`
	paramObj
}

func (r PressActionParam) MarshalJSON() (data []byte, err error) {
	type shadow PressActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PressActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PressActionPressUnionParam struct {
	OfString                 param.Opt[string]            `json:",omitzero,inline"`
	OfPressActionPressObject *PressActionPressObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u PressActionPressUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfPressActionPressObject)
}
func (u *PressActionPressUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PressActionPressUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfPressActionPressObject) {
		return u.OfPressActionPressObject
	}
	return nil
}

// The property Key is required.
type PressActionPressObjectParam struct {
	// Any of "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Power", "Eject",
	// "Abort", "Help", "Backspace", "Tab", "Numpad5", "NumpadEnter", "Enter", "\r",
	// "\n", "ShiftLeft", "ShiftRight", "ControlLeft", "ControlRight", "AltLeft",
	// "AltRight", "Pause", "CapsLock", "Escape", "Convert", "NonConvert", "Space",
	// "Numpad9", "PageUp", "Numpad3", "PageDown", "End", "Numpad1", "Home", "Numpad7",
	// "ArrowLeft", "Numpad4", "Numpad8", "ArrowUp", "ArrowRight", "Numpad6",
	// "Numpad2", "ArrowDown", "Select", "Open", "PrintScreen", "Insert", "Numpad0",
	// "Delete", "NumpadDecimal", "Digit0", "Digit1", "Digit2", "Digit3", "Digit4",
	// "Digit5", "Digit6", "Digit7", "Digit8", "Digit9", "KeyA", "KeyB", "KeyC",
	// "KeyD", "KeyE", "KeyF", "KeyG", "KeyH", "KeyI", "KeyJ", "KeyK", "KeyL", "KeyM",
	// "KeyN", "KeyO", "KeyP", "KeyQ", "KeyR", "KeyS", "KeyT", "KeyU", "KeyV", "KeyW",
	// "KeyX", "KeyY", "KeyZ", "MetaLeft", "MetaRight", "ContextMenu",
	// "NumpadMultiply", "NumpadAdd", "NumpadSubtract", "NumpadDivide", "F1", "F2",
	// "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F10", "F11", "F12", "F13", "F14",
	// "F15", "F16", "F17", "F18", "F19", "F20", "F21", "F22", "F23", "F24", "NumLock",
	// "ScrollLock", "AudioVolumeMute", "AudioVolumeDown", "AudioVolumeUp",
	// "MediaTrackNext", "MediaTrackPrevious", "MediaStop", "MediaPlayPause",
	// "Semicolon", "Equal", "NumpadEqual", "Comma", "Minus", "Period", "Slash",
	// "Backquote", "BracketLeft", "Backslash", "BracketRight", "Quote", "AltGraph",
	// "Props", "Cancel", "Clear", "Shift", "Control", "Alt", "Accept", "ModeChange", "
	// ", "Print", "Execute", "\u0000", "a", "b", "c", "d", "e", "f", "g", "h", "i",
	// "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y",
	// "z", "Meta", "\*", "+", "-", "/", ";", "=", ",", ".", "`", "[", "\\", "]", "'",
	// "Attn", "CrSel", "ExSel", "EraseEof", "Play", "ZoomOut", ")", "!", "@", "#",
	// "$", "%", "^", "&", "(", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
	// "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", ":",
	// "<", "\_", ">", "?", "~", "{", "|", "}", "\"", "SoftLeft", "SoftRight",
	// "Camera", "Call", "EndCall", "VolumeDown", "VolumeUp".
	Key string `json:"key,omitzero" api:"required"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay PressActionPressObjectDelayUnionParam `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required PressActionPressObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip PressActionPressObjectSkipUnionParam `json:"skip,omitzero"`
	paramObj
}

func (r PressActionPressObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow PressActionPressObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PressActionPressObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PressActionPressObjectParam](
		"key", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Power", "Eject", "Abort", "Help", "Backspace", "Tab", "Numpad5", "NumpadEnter", "Enter", "\r", "\n", "ShiftLeft", "ShiftRight", "ControlLeft", "ControlRight", "AltLeft", "AltRight", "Pause", "CapsLock", "Escape", "Convert", "NonConvert", "Space", "Numpad9", "PageUp", "Numpad3", "PageDown", "End", "Numpad1", "Home", "Numpad7", "ArrowLeft", "Numpad4", "Numpad8", "ArrowUp", "ArrowRight", "Numpad6", "Numpad2", "ArrowDown", "Select", "Open", "PrintScreen", "Insert", "Numpad0", "Delete", "NumpadDecimal", "Digit0", "Digit1", "Digit2", "Digit3", "Digit4", "Digit5", "Digit6", "Digit7", "Digit8", "Digit9", "KeyA", "KeyB", "KeyC", "KeyD", "KeyE", "KeyF", "KeyG", "KeyH", "KeyI", "KeyJ", "KeyK", "KeyL", "KeyM", "KeyN", "KeyO", "KeyP", "KeyQ", "KeyR", "KeyS", "KeyT", "KeyU", "KeyV", "KeyW", "KeyX", "KeyY", "KeyZ", "MetaLeft", "MetaRight", "ContextMenu", "NumpadMultiply", "NumpadAdd", "NumpadSubtract", "NumpadDivide", "F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F10", "F11", "F12", "F13", "F14", "F15", "F16", "F17", "F18", "F19", "F20", "F21", "F22", "F23", "F24", "NumLock", "ScrollLock", "AudioVolumeMute", "AudioVolumeDown", "AudioVolumeUp", "MediaTrackNext", "MediaTrackPrevious", "MediaStop", "MediaPlayPause", "Semicolon", "Equal", "NumpadEqual", "Comma", "Minus", "Period", "Slash", "Backquote", "BracketLeft", "Backslash", "BracketRight", "Quote", "AltGraph", "Props", "Cancel", "Clear", "Shift", "Control", "Alt", "Accept", "ModeChange", " ", "Print", "Execute", "\u0000", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "Meta", "*", "+", "-", "/", ";", "=", ",", ".", "`", "[", "\\", "]", "'", "Attn", "CrSel", "ExSel", "EraseEof", "Play", "ZoomOut", ")", "!", "@", "#", "$", "%", "^", "&", "(", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", ":", "<", "_", ">", "?", "~", "{", "|", "}", "\"", "SoftLeft", "SoftRight", "Camera", "Call", "EndCall", "VolumeDown", "VolumeUp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PressActionPressObjectDelayUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u PressActionPressObjectDelayUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *PressActionPressObjectDelayUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PressActionPressObjectDelayUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PressActionPressObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfPressActionPressObjectRequiredString)
	OfPressActionPressObjectRequiredString param.Opt[PressActionPressObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                 param.Opt[bool]                                 `json:",omitzero,inline"`
	paramUnion
}

func (u PressActionPressObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPressActionPressObjectRequiredString, u.OfBool)
}
func (u *PressActionPressObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PressActionPressObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfPressActionPressObjectRequiredString) {
		return &u.OfPressActionPressObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type PressActionPressObjectRequiredString string

const (
	PressActionPressObjectRequiredStringTrue  PressActionPressObjectRequiredString = "true"
	PressActionPressObjectRequiredStringFalse PressActionPressObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PressActionPressObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfPressActionPressObjectSkipString)
	OfPressActionPressObjectSkipString param.Opt[PressActionPressObjectSkipString] `json:",omitzero,inline"`
	OfBool                             param.Opt[bool]                             `json:",omitzero,inline"`
	paramUnion
}

func (u PressActionPressObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPressActionPressObjectSkipString, u.OfBool)
}
func (u *PressActionPressObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PressActionPressObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfPressActionPressObjectSkipString) {
		return &u.OfPressActionPressObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type PressActionPressObjectSkipString string

const (
	PressActionPressObjectSkipStringTrue  PressActionPressObjectSkipString = "true"
	PressActionPressObjectSkipStringFalse PressActionPressObjectSkipString = "false"
)

// Capture a page screenshot
//
// The property Screenshot is required.
type ScreenshotActionParam struct {
	Screenshot ScreenshotActionScreenshotUnionParam `json:"screenshot,omitzero" api:"required"`
	paramObj
}

func (r ScreenshotActionParam) MarshalJSON() (data []byte, err error) {
	type shadow ScreenshotActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScreenshotActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScreenshotActionScreenshotUnionParam struct {
	OfBool                             param.Opt[bool]                        `json:",omitzero,inline"`
	OfScreenshotActionScreenshotObject *ScreenshotActionScreenshotObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u ScreenshotActionScreenshotUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfScreenshotActionScreenshotObject)
}
func (u *ScreenshotActionScreenshotUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScreenshotActionScreenshotUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfScreenshotActionScreenshotObject) {
		return u.OfScreenshotActionScreenshotObject
	}
	return nil
}

type ScreenshotActionScreenshotObjectParam struct {
	FullPage param.Opt[bool]    `json:"full_page,omitzero"`
	Quality  param.Opt[float64] `json:"quality,omitzero"`
	// Any of "png", "jpeg", "webp".
	Format string `json:"format,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ScreenshotActionScreenshotObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ScreenshotActionScreenshotObjectSkipUnionParam `json:"skip,omitzero"`
	paramObj
}

func (r ScreenshotActionScreenshotObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow ScreenshotActionScreenshotObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScreenshotActionScreenshotObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ScreenshotActionScreenshotObjectParam](
		"format", "png", "jpeg", "webp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScreenshotActionScreenshotObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfScreenshotActionScreenshotObjectRequiredString)
	OfScreenshotActionScreenshotObjectRequiredString param.Opt[ScreenshotActionScreenshotObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                           param.Opt[bool]                                           `json:",omitzero,inline"`
	paramUnion
}

func (u ScreenshotActionScreenshotObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScreenshotActionScreenshotObjectRequiredString, u.OfBool)
}
func (u *ScreenshotActionScreenshotObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScreenshotActionScreenshotObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfScreenshotActionScreenshotObjectRequiredString) {
		return &u.OfScreenshotActionScreenshotObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ScreenshotActionScreenshotObjectRequiredString string

const (
	ScreenshotActionScreenshotObjectRequiredStringTrue  ScreenshotActionScreenshotObjectRequiredString = "true"
	ScreenshotActionScreenshotObjectRequiredStringFalse ScreenshotActionScreenshotObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScreenshotActionScreenshotObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfScreenshotActionScreenshotObjectSkipString)
	OfScreenshotActionScreenshotObjectSkipString param.Opt[ScreenshotActionScreenshotObjectSkipString] `json:",omitzero,inline"`
	OfBool                                       param.Opt[bool]                                       `json:",omitzero,inline"`
	paramUnion
}

func (u ScreenshotActionScreenshotObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScreenshotActionScreenshotObjectSkipString, u.OfBool)
}
func (u *ScreenshotActionScreenshotObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScreenshotActionScreenshotObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfScreenshotActionScreenshotObjectSkipString) {
		return &u.OfScreenshotActionScreenshotObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ScreenshotActionScreenshotObjectSkipString string

const (
	ScreenshotActionScreenshotObjectSkipStringTrue  ScreenshotActionScreenshotObjectSkipString = "true"
	ScreenshotActionScreenshotObjectSkipStringFalse ScreenshotActionScreenshotObjectSkipString = "false"
)

// Scroll the page or an element
//
// The property Scroll is required.
type ScrollActionParam struct {
	Scroll ScrollActionScrollUnionParam `json:"scroll,omitzero" api:"required"`
	paramObj
}

func (r ScrollActionParam) MarshalJSON() (data []byte, err error) {
	type shadow ScrollActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScrollActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScrollActionScrollUnionParam struct {
	OfFloat                    param.Opt[float64]             `json:",omitzero,inline"`
	OfString                   param.Opt[string]              `json:",omitzero,inline"`
	OfScrollActionScrollObject *ScrollActionScrollObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u ScrollActionScrollUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfScrollActionScrollObject)
}
func (u *ScrollActionScrollUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScrollActionScrollUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfScrollActionScrollObject) {
		return u.OfScrollActionScrollObject
	}
	return nil
}

type ScrollActionScrollObjectParam struct {
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	X       param.Opt[float64] `json:"x,omitzero"`
	Y       param.Opt[float64] `json:"y,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container ScrollActionScrollObjectContainerUnionParam `json:"container,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ScrollActionScrollObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ScrollActionScrollObjectSkipUnionParam `json:"skip,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	To ScrollActionScrollObjectToUnionParam `json:"to,omitzero"`
	paramObj
}

func (r ScrollActionScrollObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow ScrollActionScrollObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScrollActionScrollObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScrollActionScrollObjectContainerUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ScrollActionScrollObjectContainerUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ScrollActionScrollObjectContainerUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScrollActionScrollObjectContainerUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScrollActionScrollObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfScrollActionScrollObjectRequiredString)
	OfScrollActionScrollObjectRequiredString param.Opt[ScrollActionScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                   param.Opt[bool]                                   `json:",omitzero,inline"`
	paramUnion
}

func (u ScrollActionScrollObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScrollActionScrollObjectRequiredString, u.OfBool)
}
func (u *ScrollActionScrollObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScrollActionScrollObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfScrollActionScrollObjectRequiredString) {
		return &u.OfScrollActionScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ScrollActionScrollObjectRequiredString string

const (
	ScrollActionScrollObjectRequiredStringTrue  ScrollActionScrollObjectRequiredString = "true"
	ScrollActionScrollObjectRequiredStringFalse ScrollActionScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScrollActionScrollObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfScrollActionScrollObjectSkipString)
	OfScrollActionScrollObjectSkipString param.Opt[ScrollActionScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                               param.Opt[bool]                               `json:",omitzero,inline"`
	paramUnion
}

func (u ScrollActionScrollObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScrollActionScrollObjectSkipString, u.OfBool)
}
func (u *ScrollActionScrollObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScrollActionScrollObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfScrollActionScrollObjectSkipString) {
		return &u.OfScrollActionScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ScrollActionScrollObjectSkipString string

const (
	ScrollActionScrollObjectSkipStringTrue  ScrollActionScrollObjectSkipString = "true"
	ScrollActionScrollObjectSkipStringFalse ScrollActionScrollObjectSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScrollActionScrollObjectToUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ScrollActionScrollObjectToUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ScrollActionScrollObjectToUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScrollActionScrollObjectToUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Wait for a specified duration
//
// The property Wait is required.
type WaitActionParam struct {
	Wait WaitActionWaitUnionParam `json:"wait,omitzero" api:"required"`
	paramObj
}

func (r WaitActionParam) MarshalJSON() (data []byte, err error) {
	type shadow WaitActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaitActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitActionWaitUnionParam struct {
	OfFloat                param.Opt[float64]         `json:",omitzero,inline"`
	OfString               param.Opt[string]          `json:",omitzero,inline"`
	OfWaitActionWaitObject *WaitActionWaitObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u WaitActionWaitUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfWaitActionWaitObject)
}
func (u *WaitActionWaitUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitActionWaitUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfWaitActionWaitObject) {
		return u.OfWaitActionWaitObject
	}
	return nil
}

// The property Duration is required.
type WaitActionWaitObjectParam struct {
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Duration WaitActionWaitObjectDurationUnionParam `json:"duration,omitzero" api:"required"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required WaitActionWaitObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip WaitActionWaitObjectSkipUnionParam `json:"skip,omitzero"`
	paramObj
}

func (r WaitActionWaitObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow WaitActionWaitObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaitActionWaitObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitActionWaitObjectDurationUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u WaitActionWaitObjectDurationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *WaitActionWaitObjectDurationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitActionWaitObjectDurationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitActionWaitObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWaitActionWaitObjectRequiredString)
	OfWaitActionWaitObjectRequiredString param.Opt[WaitActionWaitObjectRequiredString] `json:",omitzero,inline"`
	OfBool                               param.Opt[bool]                               `json:",omitzero,inline"`
	paramUnion
}

func (u WaitActionWaitObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWaitActionWaitObjectRequiredString, u.OfBool)
}
func (u *WaitActionWaitObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitActionWaitObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWaitActionWaitObjectRequiredString) {
		return &u.OfWaitActionWaitObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type WaitActionWaitObjectRequiredString string

const (
	WaitActionWaitObjectRequiredStringTrue  WaitActionWaitObjectRequiredString = "true"
	WaitActionWaitObjectRequiredStringFalse WaitActionWaitObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitActionWaitObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWaitActionWaitObjectSkipString)
	OfWaitActionWaitObjectSkipString param.Opt[WaitActionWaitObjectSkipString] `json:",omitzero,inline"`
	OfBool                           param.Opt[bool]                           `json:",omitzero,inline"`
	paramUnion
}

func (u WaitActionWaitObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWaitActionWaitObjectSkipString, u.OfBool)
}
func (u *WaitActionWaitObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitActionWaitObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWaitActionWaitObjectSkipString) {
		return &u.OfWaitActionWaitObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type WaitActionWaitObjectSkipString string

const (
	WaitActionWaitObjectSkipStringTrue  WaitActionWaitObjectSkipString = "true"
	WaitActionWaitObjectSkipStringFalse WaitActionWaitObjectSkipString = "false"
)

// Wait for an element to appear or reach a specific state
//
// The property WaitForElement is required.
type WaitForElementActionParam struct {
	WaitForElement WaitForElementActionWaitForElementUnionParam `json:"wait_for_element,omitzero" api:"required"`
	paramObj
}

func (r WaitForElementActionParam) MarshalJSON() (data []byte, err error) {
	type shadow WaitForElementActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaitForElementActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitForElementActionWaitForElementUnionParam struct {
	OfString                                   param.Opt[string]                              `json:",omitzero,inline"`
	OfStringArray                              []string                                       `json:",omitzero,inline"`
	OfWaitForElementActionWaitForElementObject *WaitForElementActionWaitForElementObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u WaitForElementActionWaitForElementUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfWaitForElementActionWaitForElementObject)
}
func (u *WaitForElementActionWaitForElementUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitForElementActionWaitForElementUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfWaitForElementActionWaitForElementObject) {
		return u.OfWaitForElementActionWaitForElementObject
	}
	return nil
}

// The property Selector is required.
type WaitForElementActionWaitForElementObjectParam struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector WaitForElementActionWaitForElementObjectSelectorUnionParam `json:"selector,omitzero" api:"required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required WaitForElementActionWaitForElementObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip WaitForElementActionWaitForElementObjectSkipUnionParam `json:"skip,omitzero"`
	paramObj
}

func (r WaitForElementActionWaitForElementObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow WaitForElementActionWaitForElementObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaitForElementActionWaitForElementObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitForElementActionWaitForElementObjectSelectorUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u WaitForElementActionWaitForElementObjectSelectorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *WaitForElementActionWaitForElementObjectSelectorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitForElementActionWaitForElementObjectSelectorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitForElementActionWaitForElementObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWaitForElementActionWaitForElementObjectRequiredString)
	OfWaitForElementActionWaitForElementObjectRequiredString param.Opt[WaitForElementActionWaitForElementObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                   param.Opt[bool]                                                   `json:",omitzero,inline"`
	paramUnion
}

func (u WaitForElementActionWaitForElementObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWaitForElementActionWaitForElementObjectRequiredString, u.OfBool)
}
func (u *WaitForElementActionWaitForElementObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitForElementActionWaitForElementObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWaitForElementActionWaitForElementObjectRequiredString) {
		return &u.OfWaitForElementActionWaitForElementObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type WaitForElementActionWaitForElementObjectRequiredString string

const (
	WaitForElementActionWaitForElementObjectRequiredStringTrue  WaitForElementActionWaitForElementObjectRequiredString = "true"
	WaitForElementActionWaitForElementObjectRequiredStringFalse WaitForElementActionWaitForElementObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitForElementActionWaitForElementObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWaitForElementActionWaitForElementObjectSkipString)
	OfWaitForElementActionWaitForElementObjectSkipString param.Opt[WaitForElementActionWaitForElementObjectSkipString] `json:",omitzero,inline"`
	OfBool                                               param.Opt[bool]                                               `json:",omitzero,inline"`
	paramUnion
}

func (u WaitForElementActionWaitForElementObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWaitForElementActionWaitForElementObjectSkipString, u.OfBool)
}
func (u *WaitForElementActionWaitForElementObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitForElementActionWaitForElementObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWaitForElementActionWaitForElementObjectSkipString) {
		return &u.OfWaitForElementActionWaitForElementObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type WaitForElementActionWaitForElementObjectSkipString string

const (
	WaitForElementActionWaitForElementObjectSkipStringTrue  WaitForElementActionWaitForElementObjectSkipString = "true"
	WaitForElementActionWaitForElementObjectSkipStringFalse WaitForElementActionWaitForElementObjectSkipString = "false"
)

// Wait for page navigation to complete
//
// The property WaitForNavigation is required.
type WaitForNavigationActionParam struct {
	WaitForNavigation WaitForNavigationActionWaitForNavigationUnionParam `json:"wait_for_navigation,omitzero" api:"required"`
	paramObj
}

func (r WaitForNavigationActionParam) MarshalJSON() (data []byte, err error) {
	type shadow WaitForNavigationActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaitForNavigationActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitForNavigationActionWaitForNavigationUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWaitForNavigationActionWaitForNavigationString)
	OfWaitForNavigationActionWaitForNavigationString param.Opt[string]                                    `json:",omitzero,inline"`
	OfWaitForNavigationActionWaitForNavigationObject *WaitForNavigationActionWaitForNavigationObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u WaitForNavigationActionWaitForNavigationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWaitForNavigationActionWaitForNavigationString, u.OfWaitForNavigationActionWaitForNavigationObject)
}
func (u *WaitForNavigationActionWaitForNavigationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitForNavigationActionWaitForNavigationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWaitForNavigationActionWaitForNavigationString) {
		return &u.OfWaitForNavigationActionWaitForNavigationString
	} else if !param.IsOmitted(u.OfWaitForNavigationActionWaitForNavigationObject) {
		return u.OfWaitForNavigationActionWaitForNavigationObject
	}
	return nil
}

type WaitForNavigationActionWaitForNavigationString string

const (
	WaitForNavigationActionWaitForNavigationStringLoad             WaitForNavigationActionWaitForNavigationString = "load"
	WaitForNavigationActionWaitForNavigationStringDomcontentloaded WaitForNavigationActionWaitForNavigationString = "domcontentloaded"
	WaitForNavigationActionWaitForNavigationStringNetworkidle0     WaitForNavigationActionWaitForNavigationString = "networkidle0"
	WaitForNavigationActionWaitForNavigationStringNetworkidle2     WaitForNavigationActionWaitForNavigationString = "networkidle2"
)

// The property Navigation is required.
type WaitForNavigationActionWaitForNavigationObjectParam struct {
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	Navigation string `json:"navigation,omitzero" api:"required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required WaitForNavigationActionWaitForNavigationObjectRequiredUnionParam `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip WaitForNavigationActionWaitForNavigationObjectSkipUnionParam `json:"skip,omitzero"`
	paramObj
}

func (r WaitForNavigationActionWaitForNavigationObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow WaitForNavigationActionWaitForNavigationObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaitForNavigationActionWaitForNavigationObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WaitForNavigationActionWaitForNavigationObjectParam](
		"navigation", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitForNavigationActionWaitForNavigationObjectRequiredUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWaitForNavigationActionWaitForNavigationObjectRequiredString)
	OfWaitForNavigationActionWaitForNavigationObjectRequiredString param.Opt[WaitForNavigationActionWaitForNavigationObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                         `json:",omitzero,inline"`
	paramUnion
}

func (u WaitForNavigationActionWaitForNavigationObjectRequiredUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWaitForNavigationActionWaitForNavigationObjectRequiredString, u.OfBool)
}
func (u *WaitForNavigationActionWaitForNavigationObjectRequiredUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitForNavigationActionWaitForNavigationObjectRequiredUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWaitForNavigationActionWaitForNavigationObjectRequiredString) {
		return &u.OfWaitForNavigationActionWaitForNavigationObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type WaitForNavigationActionWaitForNavigationObjectRequiredString string

const (
	WaitForNavigationActionWaitForNavigationObjectRequiredStringTrue  WaitForNavigationActionWaitForNavigationObjectRequiredString = "true"
	WaitForNavigationActionWaitForNavigationObjectRequiredStringFalse WaitForNavigationActionWaitForNavigationObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitForNavigationActionWaitForNavigationObjectSkipUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfWaitForNavigationActionWaitForNavigationObjectSkipString)
	OfWaitForNavigationActionWaitForNavigationObjectSkipString param.Opt[WaitForNavigationActionWaitForNavigationObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                     param.Opt[bool]                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u WaitForNavigationActionWaitForNavigationObjectSkipUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWaitForNavigationActionWaitForNavigationObjectSkipString, u.OfBool)
}
func (u *WaitForNavigationActionWaitForNavigationObjectSkipUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitForNavigationActionWaitForNavigationObjectSkipUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWaitForNavigationActionWaitForNavigationObjectSkipString) {
		return &u.OfWaitForNavigationActionWaitForNavigationObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type WaitForNavigationActionWaitForNavigationObjectSkipString string

const (
	WaitForNavigationActionWaitForNavigationObjectSkipStringTrue  WaitForNavigationActionWaitForNavigationObjectSkipString = "true"
	WaitForNavigationActionWaitForNavigationObjectSkipStringFalse WaitForNavigationActionWaitForNavigationObjectSkipString = "false"
)
