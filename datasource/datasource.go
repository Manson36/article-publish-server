package datasource

import (
	"github.com/article-publish-server/datamodels"
)

func init() {
	PqDB.initDB()
	RDS.initDB()

	_ = datamodels.GetModelList()
	//pqDB.AutoMigrate(datamodels.GetModelList()...)
	//if err := pqDB.BulkAddCommentToColumn(datamodels.GetModelList()...); err != nil {
	//	log.Fatal(err.Error())
	//}
}
