package gobluos

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// Client captures the base URL for the BluOS player.
type Client struct {
	base string
}

// NewClient creates a new Client instance for the BluOS player at the
// specified base URL.
func NewClient(base string) *Client {
	return &Client{
		base: base,
	}
}

// Get performs an HTTP GET against the specified path.
//
// Returns an error if on HTTP errors, or XML decode errors.
// It is not considered an error if the response body is empty.
//
func (c *Client) Get(path string, v interface{}) error {
	res, err := http.Get(c.base + path)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	dec := xml.NewDecoder(res.Body)
	err = dec.Decode(v)

	// Note: the body may be empty if we called /Skip or /Back
	// when playing an external playlist, like Spotify Radio.
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Status retrieves the player's status.
func (c *Client) Status() (*StatusResponse, error) {
	result := &StatusResponse{}
	if err := c.Get("/Status", result); err != nil {
		return nil, err
	}

	return result, nil
}

// Play issues the play command.
func (c *Client) Play() (*StateResponse, error) {
	result := &StateResponse{}
	if err := c.Get("/Play", result); err != nil {
		return nil, err
	}

	return result, nil
}

// Pause issues the pause command.
func (c *Client) Pause() (*StateResponse, error) {
	result := &StateResponse{}
	if err := c.Get("/Pause", result); err != nil {
		return nil, err
	}

	return result, nil
}

// Skip issues the skip command.
//
// This is a no-op if the player is streaming.
//
func (c *Client) Skip() (*PlaybackResponse, error) {
	result := &PlaybackResponse{}
	if err := c.Get("/Skip", result); err != nil {
		return nil, err
	}

	return result, nil
}

// Back issues the back command.
//
// This is a no-op if the player is streaming.
//
func (c *Client) Back() (*PlaybackResponse, error) {
	result := &PlaybackResponse{}
	if err := c.Get("/Back", result); err != nil {
		return nil, err
	}

	return result, nil
}

// Volume retrieves the player's volume.
func (c *Client) Volume() (*VolumeResponse, error) {
	result := &VolumeResponse{}
	if err := c.Get("/Volume", result); err != nil {
		return nil, err
	}

	return result, nil
}

// SetVolume sets the current volume to the specified level.
func (c *Client) SetVolume(level int) (*VolumeResponse, error) {
	result := &VolumeResponse{}
	if err := c.Get(fmt.Sprintf("/Volume?level=%d", level), result); err != nil {
		return nil, err
	}

	return result, nil
}
