package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Article struct {
	BaseModel
	CategoryId   int    `gorm:"column:category_id" json:"CategoryId"`
	CategoryName string `gorm:"-" json:"CategoryName"`
	Title        string `gorm:"column:title" json:"Title"`
	Content      string `gorm:"column:content" json:"Content"`
	Status       int8   `gorm:"column:status" json:"Status"`
	Top          int8   `gorm:"column:top" json:"Top"`
	Depiction    string `gorm:"column:depiction" json:"Depiction"`
	Image        string `gorm:"column:image" json:"Image"`
	ReadNum      int    `gorm:"-" json:"ReadNum"`
	CommentNum   int    `gorm:"-" json:"CommentNum"`
}

type AST int8

const (
	_ AST = iota - 2
	ArticleAll
	ArticleNORMAL
	ArticleDELETE
)


func AddArticle(categoryId int, title string, content string, top int8, depiction string, image string) (err error) {

	article := Article{CategoryId: categoryId, Title: title, Content: content, Top: top, Depiction: depiction, Image: image}
	tx := dbObj.Begin()

	if top == 2 {
		if err = tx.Model(&Article{}).Where("top = ?", 2).Updates(map[string]int8{"top": 1}).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	if err = tx.Create(&article).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func EditArticle(article Article) (err error) {

	tx := dbObj.Begin()

	if article.Top == 2 {
		if err = tx.Model(&Article{}).Where("top = ?", 2).Updates(map[string]int8{"top": 1}).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	if err = tx.Model(&article).Updates(Article{
		CategoryId: article.CategoryId,
		Title:      article.Title,
		Content:    article.Content,
		Top:        article.Top,
		Depiction:  article.Depiction,
		Image:      article.Image}).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func GetArticleCountByCategoryId(categoryId int) (count int, err error) {
	if err = dbObj.Model(Article{}).Where("category_id = ? and status = ?", categoryId, 0).Count(&count).Error; err != nil {
		return
	}
	return
}

func DeleteArticle(id int) (err error) {

	if err = dbObj.Model(&Article{}).Where("id = ?", id).Update(map[string]int{"status": 1}).Error; err != nil {
		return
	}
	return
}

func GetArticleById(id int) (article Article, err error) {
	if err = dbObj.First(&article, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = NoData
		}
		return
	}
	return
}

func ExistsArticleById(id int) (exists bool, err error) {
	var count int
	if err = dbObj.Model(&Article{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return
	}

	if count > 0 {
		exists = true
	}
	return
}

func GetTopArticle() (article Article, err error) {
	if err = dbObj.Model(&Article{}).Where("status = ? AND top = ?", 0, 2).First(&article).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = NoData
		}
		return
	}
	return

}

func GetArticleListForPage(pageIndex, pageSize, categoryId int) (total int, articles []Article, err error) {

	articles = make([]Article, 0)
	db := dbObj.Model(Article{}).Where("status=?", ArticleNORMAL)
	if categoryId > 0 {
		db.Where("category_id = ?", categoryId)
	}

	if err = db.Count(&total).Error; err != nil {

		return
	}

	if total <= 0 {
		return
	}

	if categoryId == 0 {
		err = dbObj.Model(Article{}).Where("status = ?", ArticleNORMAL).Order("created_at desc ").Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&articles).Error
	} else {
		err = dbObj.Model(Article{}).Where("status = ? AND category_id = ?", ArticleNORMAL, categoryId).Order("created_at desc ").Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&articles).Error
	}

	if err != nil {
		return
	}

	categoryList, err := GetCategoryList(CategoryNORMAL)
	if err != nil {
		return
	}

	tmpCategoryList := make(map[int]string)
	for i := 0; i < len(categoryList); i++ {
		tmpCategoryList[categoryList[i].ID] = categoryList[i].Name
	}

	for i := 0; i < len(articles); i++ {
		articles[i].CategoryName = tmpCategoryList[articles[i].CategoryId]
	}

	return
}


func GetArticleByTwo(articleId int) (articles []Article, err error) {

	rows, err := dbObj.Raw(fmt.Sprintf("select id, title from tb_article where id in (select case when SIGN(id-14)>0 THEN MIN(id) when SIGN(id-%d)<0 THEN MAX(id) ELSE id end from tb_article where status = 0 GROUP BY SIGN(id-%d) ORDER BY SIGN(id-%d) ) ORDER BY id ", articleId, articleId, articleId)).Rows() // (*sql.Rows, error)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = NoData
		}
		return
	}
	defer func() {
		_ = rows.Close()
	}()
	var (
		id    int
		title string
	)
	for rows.Next() {
		if err = rows.Scan(&id, &title); err != nil {
			return
		}

		var t Article
		if id == articleId {
			continue
		}
		if id < articleId {
			t = Article{BaseModel: BaseModel{ID: id}, Title: title, Content: "s"}
		} else {
			t = Article{BaseModel: BaseModel{ID: id}, Title: title, Content: "x"}
		}
		articles = append(articles, t)
	}
	return
}
