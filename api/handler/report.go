package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Debtors godoc
// @ID debtors
// @Router /report/debtors [GET]
// @Summary debtors
// @Description debtors
// @Tags Report
// @Accept json
// @Produce json
// @Success 201 {object} models.GetListDebtorResponse "GetListDebtorResponseBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetDebtors(c *gin.Context) {

	res, err := h.storage.Report().GetListDebtors(context.Background())
	if err != nil {
		log.Println("error whiling get list debtors:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// InvestorShare godoc
// @ID Intestor Share
// @Router /report/investor-share [GET]
// @Summary investor share
// @Description investor share
// @Tags Report
// @Accept json
// @Produce json
// @Success 201 {object} models.GetListInvestorShareResponse "GetListInvestorShareResponseBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetInvestorShare(c *gin.Context) {

	res, err := h.storage.Report().GetListInvestorShare(context.Background())
	if err != nil {
		log.Println("error whiling get list investor share:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// BranchShare godoc
// @ID Branch Share
// @Router /report/branch-share [GET]
// @Summary branch share
// @Description branch share
// @Tags Report
// @Accept json
// @Produce json
// @Success 201 {object} models.GetListBranchShareResponse "GetListBranchShareResponseBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetBranchShare(c *gin.Context) {

	res, err := h.storage.Report().GetListBranchShare(context.Background())
	if err != nil {
		log.Println("error whiling get list Branch share:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}
