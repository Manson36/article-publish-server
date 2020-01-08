package repositories

import (
	"github.com/article-publish-server/datamodels"
	"github.com/article-publish-server/datasource"
	"github.com/article-publish-server/utils/qnuploader"
)

type ArticleRepository interface {
	Create(news *datamodels.Article) error
	Remove(query interface{}, args ...interface{}) error
	Update(doc map[string]interface{}, query interface{}, args ...interface{}) error
	Get(query interface{}, args ...interface{}) (*datamodels.Article, error)
	List(order, limit, offset, query interface{}, args ...interface{}) ([]datamodels.Article, error)
	Count(query interface{}, args ...interface{}) (int64, error)
	Save(article *datamodels.Article) error
}

type articleRepository struct {
	commonRepository
	uploader *qnuploader.Uploader //七牛的使用，目的、作用、方法？？？
}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{
		commonRepository: commonRepository{
			db: datasource.PqDB.DB,
		},
		uploader: qnuploader.NewUploader(nil),
	}
}

func (a articleRepository) Create(news *datamodels.Article) error {
	return a.commonRepository.Create(news)
}

func (a articleRepository) Remove(query interface{}, args ...interface{}) error {
	return a.commonRepository.Remove(&datamodels.Article{}, query, args...)
}

func (a articleRepository) Update(doc map[string]interface{}, query interface{}, args ...interface{}) error {
	return a.commonRepository.Update(&datamodels.Article{}, query, doc, args...)
}

func (a articleRepository) Get(query interface{}, args ...interface{}) (*datamodels.Article, error) {
	record, err := a.commonRepository.Get(&datamodels.Article{}, query, args...)
	if record == nil {
		return nil, err
	}

	return record.(*datamodels.Article), nil
}

func (a articleRepository) List(order, limit, offset, query interface{}, args ...interface{}) ([]datamodels.Article, error) {
	var list []datamodels.Article
	if err := a.commonRepository.List(&list, order, limit, offset, query, args...); err != nil {
		return nil, err
	}
	return list, nil
}

func (a articleRepository) Count(query interface{}, args ...interface{}) (int64, error) {
	return a.commonRepository.Count(&datamodels.Article{}, query, args...)
}

func (a articleRepository) Save(article *datamodels.Article) error {
	return a.commonRepository.Save(article)
}
