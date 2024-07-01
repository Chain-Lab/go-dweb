/**
  @author: decision
  @date: 2024/6/29
  @note:
**/

package managers

import (
	"context"
	"encoding/json"
	"github.com/schollz/progressbar/v3"
	"github.com/sirupsen/logrus"
	"github.io/decision2016/go-dweb/interfaces"
	"os"
	"path/filepath"
	"time"
)

type progress struct {
	Commit     string   `json:"commit"`
	Version    int      `json:"version"`
	UpdateTime int64    `json:"update_time"`
	Files      []string `json:"files"`
}

type Uploader struct {
	storage *interfaces.IFileStorage

	files  []string
	commit string
}

func NewUploader() *Uploader {
	return nil
}

// upload task
func (u *Uploader) process(ctx context.Context) error {
	total := len(u.files)
	bar := progressbar.Default(int64(total))

	for idx, file := range u.files {
		err := u.save(u.files[idx+1:])
		if err != nil {
			logrus.WithError(err).Debugln("save progress to disk failed")
			return err
		}

		filename := filepath.Base(file)
		err = (*u.storage).Upload(ctx, filename, file)
		if err != nil {
			logrus.WithError(err).Debugln("error occurred when uploading file")
			return err
		}

		err = bar.Add(1)
		if err != nil {
			logrus.WithError(err).Debugln("add progress failed")
			return err
		}
	}

	return nil
}

func (u *Uploader) save(files []string) error {
	p := progress{
		Commit:     u.commit,
		Version:    0,
		UpdateTime: time.Now().Unix(),
		Files:      files,
	}

	progressBytes, err := json.Marshal(p)
	if err != nil {
		logrus.WithError(err).Debugln("marshal progress to json failed")
		return err
	}

	err = os.WriteFile(".deploy", progressBytes, 0700)
	if err != nil {
		logrus.WithError(err).Debugln("save deploy progress failed")
		return err
	}

	return nil
}

func (u *Uploader) load() (*progress, error) {
	_, err := os.Stat(".deploy")
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	cachedBytes, err := os.ReadFile(".deploy")
	if err != nil {
		logrus.WithError(err).Debugln("read cached progress file failed")
		return nil, err
	}

	var p progress
	err = json.Unmarshal(cachedBytes, &p)
	if err != nil {
		logrus.WithError(err).Debugln("unmarshal json to progress failed")
		return nil, err
	}

	return &p, nil
}

func (u *Uploader) Setup(commit string, files []string) error {
	u.commit = commit

	p, err := u.load()
	if err != nil {
		return err
	}

	if p != nil && p.Commit == u.commit {

	}

}
