package controller

import (
	intf "cleanapp/BookStore/interface"
	"cleanapp/BookStore/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type BookController struct {
	bookService intf.BookService
}

func NewBookController(echo *echo.Echo, bookServiceObject intf.BookService) {
	bookControllerObject := &BookController{
		bookService: bookServiceObject,
	}

	echo.GET("/printAuthor", bookControllerObject.PrintAuthor)
	echo.GET("/test", bookControllerObject.Test)
	echo.POST("/books", bookControllerObject.CreateBook)
	echo.GET("/books/:id", bookControllerObject.GetBook)
	echo.DELETE("/books/:id", bookControllerObject.DeleteBook)
}

func (controllerObj *BookController) PrintAuthor(ec echo.Context) error {

	return nil
}

func (controllerObj *BookController) Test(ec echo.Context) error {

	fmt.Println("**** Inside Book Controller ****")

	requestContext := ec.Request().Context()
	controllerObj.bookService.TestBookService(requestContext)

	return nil
}

func (controllerObj *BookController) CreateBook(ec echo.Context) error {
	book := new(model.Book)
	if err := ec.Bind(book); err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	result, err := controllerObj.bookService.CreateBook(ec.Request().Context(), book)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ec.JSON(http.StatusCreated, result)
}

func (controllerObj *BookController) GetBook(ec echo.Context) error {
	id := ec.Param("id")
	book, err := controllerObj.bookService.GetBookByID(ec.Request().Context(), id)
	if err != nil {
		return ec.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return ec.JSON(http.StatusOK, book)
}

func (controllerObj *BookController) DeleteBook(ec echo.Context) error {
	id := ec.Param("id")
	err := controllerObj.bookService.DeleteBook(ec.Request().Context(), id)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ec.JSON(http.StatusOK, map[string]string{"message": "Book deleted successfully"})
}
