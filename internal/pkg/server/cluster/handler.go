/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package cluster

// Handler structure for the user requests.
type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler {
	return &Handler{manager}
}
