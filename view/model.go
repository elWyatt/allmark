// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package view

type Model struct {
	Path        string
	Title       string
	Description string
	Content     string
	LanguageTag string
	Entries     []Model
}

func Error(msg string) Model {
	return Model{
		Path:    "#",
		Title:   "Error",
		Content: msg,
	}
}
