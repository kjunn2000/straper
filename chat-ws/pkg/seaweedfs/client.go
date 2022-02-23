package seaweedfs

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type SeaweedfsClient interface {
	SaveFile(ctx context.Context, fileBytes []byte) (string, error)
	GetFile(ctx context.Context, fid string) ([]byte, error)
	DeleteFile(ctx context.Context, fid string) error
}

type seaweedfsClient struct {
	log *zap.Logger
}

func NewSeaweedfsClient(log *zap.Logger) seaweedfsClient {
	return seaweedfsClient{log: log}
}

func (s *seaweedfsClient) SaveFile(ctx context.Context, reader io.Reader) (string, error) {
	resp, err := http.Get("http://localhost:9333/dir/assign")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var weedMasterResponse WeedMasterResponse
	json.Unmarshal(body, &weedMasterResponse)
	if err != nil {
		return "", err
	}
	// url := "http://" + weedMasterResponse.Url + "/" + weedMasterResponse.Fid
	// Local
	url := "http://" + "localhost:8080" + "/" + weedMasterResponse.Fid
	if err := s.sendMultiPartRequest(reader, url); err != nil {
		return "", err
	}
	return weedMasterResponse.Fid, nil
}

func (s *seaweedfsClient) sendMultiPartRequest(reader io.Reader, url string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormField("file")
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, reader)
	if err != nil {
		return err
	}
	writer.Close()
	req, err := http.NewRequest("POST", url, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	_, err = client.Do(req)
	if err != nil {
		s.log.Warn("Multipart post request to seaweedfs server failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *seaweedfsClient) GetFile(ctx context.Context, fid string) ([]byte, error) {
	// fidArr := strings.Split(fid, ",")
	// resp, err := http.Get("http://localhost:9333/dir/lookup?volumeId=" + fidArr[0])
	// if err != nil {
	// 	s.log.Warn("Seaweedfs look up volume failed", zap.Error(err))
	// 	return []byte{}, err
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	s.log.Warn("Read response body failed", zap.Error(err))
	// 	return []byte{}, err
	// }
	// var weedVolumeLoopUpResponse WeedVolumeLoopUpResponse
	// json.Unmarshal(body, &weedVolumeLoopUpResponse)

	// resp, err = http.Get("http://" + weedVolumeLoopUpResponse.Locations[0].PublicUrl + "/" + fid)
	// Local
	resp, err := http.Get("http://" + "localhost:8080" + "/" + fid)
	if err != nil {
		s.log.Warn("Seaweedfs get file failed", zap.Error(err))
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.log.Warn("Read response body failed", zap.Error(err))
		return []byte{}, err
	}
	return body, nil
}

func (s *seaweedfsClient) DeleteFile(ctx context.Context, fid string) error {
	// fidArr := strings.Split(fid, ",")
	// resp, err := http.Get("http://localhost:9333/dir/lookup?volumeId=" + fidArr[0])
	// if err != nil {
	// 	s.log.Warn("Seaweedfs look up volume failed", zap.Error(err))
	// 	return err
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	s.log.Warn("Read response body failed", zap.Error(err))
	// 	return err
	// }
	// var weedVolumeLoopUpResponse WeedVolumeLoopUpResponse
	// json.Unmarshal(body, &weedVolumeLoopUpResponse)

	// client := &http.Client{}

	// req, err := http.NewRequest("DELETE", "http://"+weedVolumeLoopUpResponse.Locations[0].PublicUrl+"/"+fid, nil)

	// Local
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "http://localhost:8080/"+fid, nil)
	if err != nil {
		return err
	}

	if _, err = client.Do(req); err != nil {
		return err
	}
	return nil
}
