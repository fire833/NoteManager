/*
*	Copyright (C) 2022  Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package course

import "time"

type CourseContext struct {
	// Title is the title of this course that the notes are a part of.
	Title string

	// Lecturer is the professor/teacher/speaker who the notes are taken off of.
	Lecturer string

	// Author is the person who wrote these notes.
	Author string `json:"author" yaml:"author"`

	// StartDate describes the starting time for this course, and the time that note titles
	// will be determined from.
	StartDate time.Time `json:"courseStart" yaml:"courseStart"`
}
