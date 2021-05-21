package structures

type DownloadSession struct {
	FileId uint32 `gorm:"column:file_id"`
	Md5 string
	SessionKey string `gorm:"column:session_key"`
}

func (DownloadSession) TableName() string {
	return "download_queue"
}

func (ds DownloadSession) ResponseData() map[string]interface{} {
	return map[string]interface{} {"md5": ds.Md5, "session_key": ds.SessionKey}
}