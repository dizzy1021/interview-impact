package util

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Page		int
	Size		int
	Sort		string
	TotalRow 	int64
	TotalPages	int
}

func NewPagination(ctx *gin.Context) Pagination {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "0"))		
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "0"))	
	sort := ctx.DefaultQuery("sort", "")

	return Pagination{
		Page: page,
		Size: pageSize,
		Sort: sort,
	}
}

func (p *Pagination) GetOffset() int {  
    return (p.GetPage() - 1) * p.GetSize() 
}   

func (p *Pagination) GetSize() int {       
	switch {
		case p.Size > MAX_PAGE_SIZE:
	  		p.Size = MAX_PAGE_SIZE
		case p.Size <= 0:
	  		p.Size = MIN_PAGE_SIZE
	}
    return p.Size  
}   

func (p *Pagination) GetPage() int {    
    if p.Page == 0 {    
        p.Page = 1  
    }   
    return p.Page   
}   

func (p *Pagination) GetSort() string { 
    if p.Sort == "" {   
        p.Sort = "created_at desc"  
    }   
    return p.Sort   
}

func Paginate(value interface{}, p *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRow int64 
    db.Model(value).Count(&totalRow)   

    p.TotalRow = totalRow  
    totalPages := int(math.Ceil(float64(totalRow) / float64(p.Size)))    
    p.TotalPages = totalPages  
 
    return func(db *gorm.DB) *gorm.DB { 
        return db.Offset(p.GetOffset()).Limit(p.GetSize()).Order(p.GetSort())   
    }   
}