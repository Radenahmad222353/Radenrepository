package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Menyajikan halaman utama
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{

			"name":        "Raden Ahmad",
			"description": "My name is Raden Ahmad Badaruddin. I am currently studying at one of the universities in Makassar, South Sulawesi, namely Dipa Makassar University. I am a 6th semester Informatics Engineering student who has a hobby of exploring the world of technology. With a strong foundation in programming, I am ready to pursue opportunities in the IT world. I have a little experience in developing applications such as Java, C++, Python, Kotlin, Go-lang, Ruby, and experience in creating simple websites.",
			"image":       "/static/raden.jpg", // Ganti dengan foto kamu
		})
	})

	// Menyajikan file HTML statis dan assets
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static") // Folder untuk menyimpan gambar

	// Menjalankan server
	r.Run(":8080")
}
