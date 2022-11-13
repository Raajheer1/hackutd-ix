package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Raajheer1/hackutd-ix/m/v2/models"
	"github.com/Raajheer1/hackutd-ix/m/v2/utils/token"
	"github.com/gin-gonic/gin"
	"gonum.org/v1/gonum/stat"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
)

type StockInput struct {
	UserID uint   `json:"user_id" binding:"required"`
	Ticker string `json:"ticker" binding:"required"`
	Shares uint   `json:"shares" binding:"required"`
}

func CreateStock(c *gin.Context) {
	var input StockInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := models.Stock{}
	s.UserID = input.UserID
	s.Ticker = input.Ticker
	s.Shares = input.Shares

	_, err := s.CreateStock()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock Added"})
}

func GetStock(c *gin.Context) {
	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stocks, err := models.GetStock(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stocks": stocks})
}

func UpdateStock(c *gin.Context) {
	var input StockInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := models.Stock{}
	s.ID = uint(id)
	s.UserID = userId
	s.Ticker = input.Ticker
	s.Shares = input.Shares

	_, err = s.UpdateStock()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock updated."})
}

func DeleteStock(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := models.Stock{}
	s.ID = uint(id)
	s.UserID = userId

	_, err = s.DeleteStock()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock deleted."})
}

type StocksAPIResponse struct {
	Data map[string]StockPrice `json:"Monthly Adjusted Time Series"`
}
type StockPrice struct {
	OpenPrice  string `json:"1. open"`
	ClosePrice string `json:"5. adjusted close"`
}

func GetStockRisk(c *gin.Context) {
	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stocks, err := models.GetStock(userId)

	stockDeviation := make(map[string]float64)
	stockData := make(map[string][]float64)
	stockWeights := make(map[string]uint)
	for _, stock := range stocks {

		resp, err := http.Get(fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY_ADJUSTED&symbol=%s&apikey=%s", stock.Ticker, os.Getenv("API_KEY")))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		var result StocksAPIResponse
		if err := json.Unmarshal(body, &result); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var monthlyReturn []float64
		for _, prices := range result.Data {
			openPrice, err := strconv.ParseFloat(prices.OpenPrice, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			closePrice, err := strconv.ParseFloat(prices.ClosePrice, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			monthlyReturn = append(monthlyReturn, (openPrice-closePrice)/openPrice)
		}

		//Annual Variance & stddev
		variance := stat.Variance(monthlyReturn, nil) * 12.0
		stddev := math.Sqrt(variance)
		stockDeviation[stock.Ticker] = stddev
		stockData[stock.Ticker] = monthlyReturn
	}

	i := 0
	var correlation [][]float64
	for ticker, data := range stockData {
		j := 0
		for ticker2, data2 := range stockData {
			if j == 0 {
				correlation = append(correlation, []float64{})
			}
			if ticker == ticker2 {
				correlation[i] = append(correlation[i], 1)
			} else {
				sum_X, sum_Y, sum_XY, squareSum_X, squareSum_Y := 0.0, 0.0, 0.0, 0.0, 0.0
				n := 0.0
				if len(data2) < len(data) {
					n = float64(len(data2))
				} else {
					n = float64(len(data))
				}
				for x := 0; x < int(n); x++ {
					sum_X += data[x]
					sum_Y += data2[x]
					sum_XY += data[x] * data2[x]

					squareSum_X += data[x] * data[x]
					squareSum_Y += data2[x] * data2[x]
				}
				correlation[i] = append(correlation[i], (n*sum_XY-sum_X*sum_Y)/(math.Sqrt((n*squareSum_X-sum_X*sum_X)*(n*squareSum_Y-sum_Y*sum_Y))))
			}
			j++
		}
		i++
	}

	var stockDeviationInOrder []float64
	var stockDeviationTranspose []float64
	for _, dev := range stockDeviation {
		stockDeviationInOrder = append(stockDeviationInOrder, dev)
		stockDeviationTranspose = append(stockDeviationTranspose, 1)
	}

	for i := len(stockDeviationInOrder) - 1; i >= 0; i-- {
		stockDeviationTranspose[len(stockDeviationInOrder)-1-i] = stockDeviationInOrder[i]
	}

	var matrix [][]float64
	for i, num := range stockDeviationInOrder {
		for j, num2 := range stockDeviationTranspose {
			if j == 0 {
				matrix = append(matrix, []float64{})
			}

			matrix[i] = append(matrix[i], num*num2)
		}
	}

	var convMatrix [][]float64
	for i, row := range matrix {
		for j := range row {
			if j == 0 {
				convMatrix = append(convMatrix, []float64{})
			}
			convMatrix[i] = append(convMatrix[i], matrix[i][j]*correlation[i][j])
		}
	}

	var weights [][]float64
	for _, weight := range stockWeights {
		weights = append(weights, []float64{float64(weight)})
	}
	tempMatrix := MulMatrix(weights, convMatrix)

	var weightTrans [][]float64
	weightTrans = append(weightTrans, []float64{})
	for _, weight := range weights {
		weightTrans[0] = append(weightTrans[0], weight[0])
	}

	final := 0.0
	for i := range weightTrans[0] {
		final += weightTrans[0][i] * tempMatrix[i][0]
	}

	c.JSON(http.StatusOK, gin.H{"risk": math.Sqrt(final)})
}

func MulMatrix(matrix1 [][]float64, matrix2 [][]float64) [][]float64 {
	var result [][]float64
	for i, row := range matrix1 {
		for j := range row {
			if j == 0 {
				result = append(result, []float64{})
			}
			result[i] = append(result[i], matrix1[i][j]*matrix2[i][j])
		}
	}
	return result
}

func GetStockReturn(c *gin.Context) {
	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stocks, err := models.GetStock(userId)

	stockData := make(map[string]float64)
	for _, stock := range stocks {
		resp, err := http.Get(fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY_ADJUSTED&symbol=%s&apikey=%s", stock.Ticker, os.Getenv("API_KEY")))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		var result StocksAPIResponse
		if err := json.Unmarshal(body, &result); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var monthlyReturn []float64
		for _, prices := range result.Data {
			openPrice, err := strconv.ParseFloat(prices.OpenPrice, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			closePrice, err := strconv.ParseFloat(prices.ClosePrice, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			monthlyReturn = append(monthlyReturn, (openPrice-closePrice)/openPrice)
		}

		avgReturn := 0.0
		count := 0.0
		for _, month := range monthlyReturn {
			avgReturn += month
			count += 1
		}

		stockData[stock.Ticker] = avgReturn / count
	}

	var returnRate float64
	var totalStockCount uint
	for _, stock := range stocks {
		totalStockCount += stock.Shares
		returnRate += float64(stock.Shares) * stockData[stock.Ticker]
	}

	fmt.Println(stockData)

	fmt.Println(returnRate)
	fmt.Println(totalStockCount)

	c.JSON(http.StatusOK, gin.H{"return": returnRate / float64(totalStockCount)})
}
