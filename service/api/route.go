/**
 * Penlook Project
 *
 * Copyright (c) 2015 Penlook Development Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Author:
 *     Loi Nguyen       <loint@penlook.com>
 *     Tin Nguyen       <tinntt@penlook.com>
 */

package main

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Handler *gin.Engine
}

func (router Router) GetHandler() *gin.Engine {
	return router.Handler
}

func (router Router) Register() {
	router.Status("/status")
	router.Comment("/comment")
	router.Activity("/activity")
}

func (router Router) Status(root string) {
	route := router.Handler
	route.POST(root, postStatus)
	route.GET(root, getAllStatus)
	route.GET(root+"/:id", getStatus)
	route.PUT(root+"/:id", updateStatus)
	route.DELETE(root+"/:id", deleteStatus)
}

func (router Router) Comment(root string) {
	route := router.Handler
	route.GET(root, getAllComments)
	route.GET(root+"/:id", getComment)
}

func (router Router) Activity(root string) {
	route := router.Handler
	route.GET(root, getAllActivities)
	route.GET(root+"/:id", getActivity)
}
