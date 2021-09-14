package routes

import (
	"project/pbm/constants"
	"project/pbm/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	//-----------Bos
	e.POST("/bos/register", controllers.RegisterBos)
	e.POST("/bos/login", controllers.Login)

	//-----------Operasional
	e.POST("/operasional/register", controllers.RegisterOperasional)
	e.POST("/operasional/login", controllers.Login)

	//-----------Pengedar
	e.POST("/pengedar/register", controllers.RegisterPengedar)
	e.POST("/pengedar/login", controllers.Login)

	//--------------------------AUTHORIZED ONLY--------------------------//
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	//-----------Bos only
	r.GET("/bos/profile", controllers.GetProfile)
	r.GET("/bos/logout", controllers.Logout)

	//-----------Operasional only
	r.GET("/operasional/profile", controllers.GetProfile)
	r.GET("/operasional/logout", controllers.Logout)
	//Search pokemon from pokeapi
	r.GET("/search/pokemons", controllers.SearchAllPokemon) // operasional only
	r.GET("/search/pokemon", controllers.SearchPokemon)     // operasional only

	//-----------Pengedar only
	r.GET("/pengedar/profile", controllers.GetProfile)
	r.GET("/pengedar/logout", controllers.Logout)

	//-----------Pokemon
	r.GET("/pokemons", controllers.GetAllPokemon)        // operasional and pengedar only
	r.POST("/pokemon", controllers.CreatePokemon)        // operasional only
	r.GET("/pokemon/id", controllers.GetPokemonById)     // operasional only
	r.GET("/pokemon/name", controllers.GetPokemonByName) // operasional and pengedar only
	r.PUT("/pokemon", controllers.UpdatePokemon)         // operasional only
	r.DELETE("/pokemon", controllers.DeletePokemon)      // operasional only

	//Transaction
	r.POST("/transaction", controllers.CreateTransaction)                 // pengedar only
	r.GET("/transactions/success", controllers.GetAllTransactionSuccess)  // bos only
	r.GET("/transactions/cancelled", controllers.GetAllTransactionFailed) // bos only
	r.GET("/transaction", controllers.GetTransaction)                     // bos only
	r.PUT("/transaction", controllers.CancelTransaction)                  // pengedar only
}
