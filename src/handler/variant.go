package handler

import (
	"fmt"
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/model"
	"gtmx/src/view/variant"
	"math/big"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type VariantHandler struct {
	Repo *repository.CatalogRepository
}

func (h VariantHandler) HandleIndex(c echo.Context) error {
	variants, err := h.Repo.ListVariants(c.Request().Context())

	if err != nil {
		return err
	}

	return render(c, variant.IndexView(variants))
}

func (h VariantHandler) HandleShow(c echo.Context) error {
	idString := c.Param("id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return err
	}

	p, err := h.Repo.GetOneVariantById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	viewVariant, err := model.Variant{}.FromDatabase(p)

	return render(c, variant.ShowView(viewVariant))
}

func (h VariantHandler) HandleNew(c echo.Context) error {
	return render(c, variant.NewView())
}

func (h VariantHandler) HandleCreate(c echo.Context) error {
	priceString := c.FormValue("price")
	price, err := strconv.ParseFloat(priceString, 64)

	if err != nil {
		return err

	}

	p := database.Variant{
		Name:  c.FormValue("name"),
		Price: pgtype.Numeric{Int: big.NewInt(int64(price * 100)), Exp: -2, Valid: true},
	}

	insertedVariant, err := h.Repo.InsertVariant(c.Request().Context(), p)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("/variant/%d", insertedVariant.ID)

	return c.Redirect(http.StatusMovedPermanently, endpoint)
}
