package handler

import (
	"gateway/proto/notification"
	"log"

	"github.com/gin-gonic/gin"
)

// @Router           /api/notifications [get]
// @Summary          Get all notifications
// @Description      This method retrieves all notifications for the user
// @Tags             NOTIFICATIONS
// @Security         BearerAuth
// @Produce          json
// @Success          200     {array}  notification.Notification  "List of notifications"
// @Failure          500     {object}  error                       "Unable to get notifications"
// @Failure          403     {object}  error            "Permission Denied"

func (h *Handler) GetNotification(c *gin.Context) {
	notifyList, err := h.Notification.GetNotfication(c.Request.Context(), &notification.NotifyEmpty{})
	if err != nil {
		log.Println("Error get notification")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to get notification"})
		return
	}
	c.IndentedJSON(200, notifyList)
}
// @Router           /api/notifications/unread [get]
// @Summary          Get unread notifications
// @Description      This method retrieves all unread notifications for the user
// @Tags             NOTIFICATIONS
// @Security         BearerAuth
// @Produce          json
// @Success          200     {array}  notification.NotifyList  "List of unread notifications"
// @Failure          500     {object}  error                       "Unable to get unread notifications"
// @Failure          403     {object}  error            "Permission Denied"
func (h *Handler) GetUnreadNotifications(c *gin.Context) {
	notifyList, err := h.Notification.GetUnreadNotfications(c.Request.Context(), &notification.NotifyEmpty{})
	if err != nil {
		log.Println("Error get notification")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to get notification"})
		return
	}
	c.IndentedJSON(200, notifyList)
}

