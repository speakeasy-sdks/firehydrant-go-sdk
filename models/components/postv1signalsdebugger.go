// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Images struct {
	Src *string `json:"src,omitempty"`
	Alt *string `json:"alt,omitempty"`
}

func (o *Images) GetSrc() *string {
	if o == nil {
		return nil
	}
	return o.Src
}

func (o *Images) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

type PostV1SignalsDebuggerLinks struct {
	Href *string `json:"href,omitempty"`
	Text *string `json:"text,omitempty"`
}

func (o *PostV1SignalsDebuggerLinks) GetHref() *string {
	if o == nil {
		return nil
	}
	return o.Href
}

func (o *PostV1SignalsDebuggerLinks) GetText() *string {
	if o == nil {
		return nil
	}
	return o.Text
}

type Signals struct {
	Summary *string                      `json:"summary,omitempty"`
	Body    *string                      `json:"body,omitempty"`
	Level   *string                      `json:"level,omitempty"`
	Images  []Images                     `json:"images,omitempty"`
	Links   []PostV1SignalsDebuggerLinks `json:"links,omitempty"`
}

func (o *Signals) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *Signals) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Signals) GetLevel() *string {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *Signals) GetImages() []Images {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *Signals) GetLinks() []PostV1SignalsDebuggerLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

type PostV1SignalsDebugger struct {
	// CEL expression
	Expression string `json:"expression"`
	// List of signals to evaluate rule expression against
	Signals []Signals `json:"signals"`
}

func (o *PostV1SignalsDebugger) GetExpression() string {
	if o == nil {
		return ""
	}
	return o.Expression
}

func (o *PostV1SignalsDebugger) GetSignals() []Signals {
	if o == nil {
		return []Signals{}
	}
	return o.Signals
}
