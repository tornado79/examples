// Code generated by goa v2.1.1, DO NOT EDIT.
//
// sommelier HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	sommelier "goa.design/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
	goahttp "goa.design/goa/http"
)

// BuildPickRequest instantiates a HTTP request object with method and path set
// to call the "sommelier" service "pick" endpoint
func (c *Client) BuildPickRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PickSommelierPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("sommelier", "pick", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePickRequest returns an encoder for requests sent to the sommelier pick
// server.
func EncodePickRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*sommelier.Criteria)
		if !ok {
			return goahttp.ErrInvalidType("sommelier", "pick", "*sommelier.Criteria", v)
		}
		body := NewPickRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("sommelier", "pick", err)
		}
		return nil
	}
}

// DecodePickResponse returns a decoder for responses returned by the sommelier
// pick endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodePickResponse may return the following errors:
//	- "no_criteria" (type sommelier.NoCriteria): http.StatusBadRequest
//	- "no_match" (type sommelier.NoMatch): http.StatusNotFound
//	- error: internal error
func DecodePickResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body PickResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("sommelier", "pick", err)
			}
			p := NewPickStoredBottleCollectionOK(body)
			view := "default"
			vres := sommelierviews.StoredBottleCollection{Projected: p, View: view}
			if err = sommelierviews.ValidateStoredBottleCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("sommelier", "pick", err)
			}
			res := sommelier.NewStoredBottleCollection(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body PickNoCriteriaResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("sommelier", "pick", err)
			}
			return nil, NewPickNoCriteria(body)
		case http.StatusNotFound:
			var (
				body PickNoMatchResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("sommelier", "pick", err)
			}
			return nil, NewPickNoMatch(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("sommelier", "pick", resp.StatusCode, string(body))
		}
	}
}

// unmarshalStoredBottleResponseToSommelierviewsStoredBottleView builds a value
// of type *sommelierviews.StoredBottleView from a value of type
// *StoredBottleResponse.
func unmarshalStoredBottleResponseToSommelierviewsStoredBottleView(v *StoredBottleResponse) *sommelierviews.StoredBottleView {
	res := &sommelierviews.StoredBottleView{
		ID:          v.ID,
		Name:        v.Name,
		Vintage:     v.Vintage,
		Description: v.Description,
		Rating:      v.Rating,
	}
	res.Winery = unmarshalWineryResponseToSommelierviewsWineryView(v.Winery)
	if v.Composition != nil {
		res.Composition = make([]*sommelierviews.ComponentView, len(v.Composition))
		for i, val := range v.Composition {
			res.Composition[i] = unmarshalComponentResponseToSommelierviewsComponentView(val)
		}
	}

	return res
}

// unmarshalWineryResponseToSommelierviewsWineryView builds a value of type
// *sommelierviews.WineryView from a value of type *WineryResponse.
func unmarshalWineryResponseToSommelierviewsWineryView(v *WineryResponse) *sommelierviews.WineryView {
	res := &sommelierviews.WineryView{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
		URL:     v.URL,
	}

	return res
}

// unmarshalComponentResponseToSommelierviewsComponentView builds a value of
// type *sommelierviews.ComponentView from a value of type *ComponentResponse.
func unmarshalComponentResponseToSommelierviewsComponentView(v *ComponentResponse) *sommelierviews.ComponentView {
	if v == nil {
		return nil
	}
	res := &sommelierviews.ComponentView{
		Varietal:   v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}
