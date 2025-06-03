package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"github.com/Hakeera/easy_erp/config"
	"github.com/Hakeera/easy_erp/model"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

// PaginaFichaModelo exibe a página principal com a lista de fichas
// @param c echo.Context Contexto da requisição
// @return error Erro de processamento, se houver
func PaginaFichaModelo(c echo.Context) error {
	// Cria dados para o template
	data := map[string]any{
		"Title": "Gerenciamento de Fichas Técnicas",
	}

	// Carrega todos os templates necessários
	tmpl, err := template.ParseFiles(
		"view/layouts/base.html",
		"view/fichas/ficha_modelo/index.html",
	)
	if err != nil {
		log.Println("Erro ao carregar templates:", err)
		return c.String(http.StatusInternalServerError, "Erro ao carregar templates: "+err.Error())
	}

	// Executa o template base
	return tmpl.ExecuteTemplate(c.Response(), "base.html", data)
}

// ListarFichasModelo busca todas as fichas no banco de dados e as exibe na página
// Realiza a conversão dos dados JSON armazenados para objetos FichaModelo
// @param c echo.Context Contexto da requisição
// @return error Erro de processamento, se houver
func ListarFichasModelo(c echo.Context) error {
	log.Println("Iniciando Listagem de Fichas") 
	
	var registros []model.FichaModeloDB
	
	// Carregar o template corretamente
	tmpl, err := template.ParseFiles("view/fichas/ficha_modelo/lista.html")
	if err != nil {
		log.Printf("Erro ao carregar template: %v", err)
		return c.String(http.StatusInternalServerError, "Erro ao carregar template: " + err.Error())
	}
	
	// Busca os dados das fichas no banco de dados
	if err := config.GetDB().Find(&registros).Error; err != nil {
		log.Println("Erro ao buscar fichas:", err)
		return c.String(http.StatusInternalServerError, "Erro ao buscar fichas")
	}
	
	// Converte os dados JSON de cada registro em FichaModelo
	var fichas []model.FichaModelo
	for _, registro := range registros {
		var ficha model.FichaModelo
		if err := json.Unmarshal(registro.Dados, &ficha); err != nil {
			log.Printf("Erro ao deserializar ficha ID %d: %v", registro.ID, err)
			continue // Pula registros com erro de deserialização
		}
		ficha.ID = registro.ID // Atribui o ID da FichaModeloDB à FichaModelo
		fichas = append(fichas, ficha)
	}
	
	log.Printf("Total de fichas carregadas: %d", len(fichas))
	
	// Configurar o Content-Type para HTML
	c.Response().Header().Set("Content-Type", "text/html")
	
	// Renderiza o template NOMEADO com as fichas convertidas
	err = tmpl.ExecuteTemplate(c.Response(), "fichamodelo_lista", fichas)
	if err != nil {
		log.Printf("Erro ao executar template: %v", err)
		return c.String(http.StatusInternalServerError, "Erro ao renderizar template")
	}
	
	log.Println("Template renderizado com sucesso")
	return nil
}

// CriarFichaModelo processa o formulário de criação de ficha e salva os dados no banco
// Extrai os dados do formulário, incluindo arrays de grades e dimensões,
// e cria um registro JSON no banco de dados
// @param c echo.Context Contexto da requisição
// @return error Erro de processamento, se houver
func CriarFichaModelo(c echo.Context) error {
	log.Println("Iniciando Criação de Ficha Técnica")
	// Parse manual do formulário (essencial)
	c.Request().ParseForm()
	grades := c.Request().PostForm["grades_grade"]
	dimensoes := c.Request().PostForm["grades_dimensao"]
	pares := make([]model.ParGradeDimensao, 0, len(grades))
	for i := range grades {
		d, err := strconv.ParseFloat(dimensoes[i], 64)
		if err != nil {
			d = 0
		}
		pares = append(pares, model.ParGradeDimensao{
			Grade:    grades[i],
			Dimensao: d,
		})
	}
	custoCorte, _ := strconv.ParseFloat(c.FormValue("custoCorte"), 64)
	custoCostura, _ := strconv.ParseFloat(c.FormValue("custoCostura"), 64)
	custoAcabamento, _ := strconv.ParseFloat(c.FormValue("custoAcabamento"), 64)
	ficha := model.FichaModelo{
		Modelo:             c.FormValue("modelo"),
		Linhas:             c.FormValue("linhas"),
		Categoria:          c.FormValue("categoria"),
		Descricao:          c.FormValue("descricao"),
		Instrucoes:         c.FormValue("instrucoes"),
		CustoCorte:         custoCorte,
		CustoCostura:       custoCostura,
		CustoAcabamento:    custoAcabamento,
		ParesGradeDimensao: pares,
	}
	jsonData, err := json.Marshal(ficha)
	if err != nil {
		log.Println("Erro ao serializar ficha:", err)
		return c.String(http.StatusInternalServerError, "Erro ao processar ficha")
	}
	fichaDB := model.FichaModeloDB{
		Dados: datatypes.JSON(jsonData),
	}
	if err := config.GetDB().Create(&fichaDB).Error; err != nil {
		log.Println("Erro ao salvar ficha:", err)
		return c.String(http.StatusInternalServerError, "Erro ao salvar ficha")
	}
	return c.String(http.StatusOK, "Ficha Técnica salva com sucesso!")
}
