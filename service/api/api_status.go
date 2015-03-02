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
	// "github.com/penlook/gin"
	"github.com/gin-gonic/gin"
)

func postStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get One Status",
	})
}

func getAllStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get All Status",
	})
}

func getStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get One Status",
	})
}

func updateStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get One Status",
	})
}

func deleteStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Get One Status",
	})
}
