package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var token = "K08Xx4cMpvzxRlme"

func main() {
	targetUrl := "https://asw-r.bokusen.net/user/harem?dmm_id=15034903&pc=1&param={%22is_harem%22:0,%22mst_scene_id%22:null}&t=1726588734655"
	payload := strings.NewReader("")
	req, _ := http.NewRequest("POST", targetUrl, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-asw-login-token", token)
	req.Header.Add("x-asw-resource-commercial-version", "1")
	req.Header.Add("x-asw-resource-game-version", "1")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()          // 这步是必要的，防止以后的内存泄漏，切记
	body, _ := io.ReadAll(response.Body) // 读取响应 body, 返回为 []byte
	var rsp Response
	err = json.Unmarshal(body, &rsp)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch := make(chan []string, 50)

	for j := 0; j < 50; j++ {
		go func() {
			for {
				select {
				case msg := <-ch:
					fmt.Println(msg)
					dirName := msg[1]
					fileNamePrefix := msg[2]
					images, _ := RequestImage(msg[0])
					//textFileName := fmt.Sprintf("%s/scene.txt", dirName)
					//if _, err := os.Stat(textFileName); err != nil {
					//	allText := strings.Join(texts, "\r\n")
					//	os.WriteFile(textFileName, []byte(allText), 0644)
					//}
					for i, imageUrl := range images {
						if !strings.Contains(imageUrl, "/ev/") {
							continue
						}
						fmt.Printf("id:%02d url:%s\n", i, imageUrl)
						fileName := fmt.Sprintf("imagess/%s/%s%03d.jpg", dirName, fileNamePrefix, i)
						if _, err := os.Stat(fileName); err == nil {
							fmt.Println("exist")
							continue
						}
						resp, err := http.Get(imageUrl)
						if err != nil {
							panic(err)
						}
						ss, _ := io.ReadAll(resp.Body)
						os.WriteFile(fileName, ss, 0644)
						resp.Body.Close()
					}
				}
			}
		}()
	}
	// 请求图片
	for _, s := range rsp.Data.HaremScene {
		var dirName string
		var fileName string
		if strings.Contains(s.Name, "初H") {
			dirName = strings.Split(s.Name, "－")[0]
			fileName = "【#初H】"
		} else {
			dirName = strings.Split(s.Name, "_")[0]
			dirName = strings.Split(dirName, "】")[1]
			fileName = strings.Split(s.Name, dirName)[0]
		}
		err = os.MkdirAll("imagess/"+dirName, 777)
		if err != nil {
			fmt.Println(err)
			return
		}
		var msg []string
		msg = append(msg, s.MstSceneId)
		msg = append(msg, dirName)
		msg = append(msg, fileName)
		ch <- msg

		//if i > 1 {
		//	break
		//}
	}

}

func RequestImage(id string) ([]string, []string) {
	t := strconv.FormatInt(time.Now().UnixMilli(), 10)
	url := "https://asw-r.bokusen.net/scripter/playback?dmm_id=15034903&pc=1&param={%22is_harem%22:1,%22mst_scene_id%22:%22" + id + "%22,%22next%22:%22hscenario/hscenario%22,%22currentTab%22:0,%22currentPage%22:1,%22scene_type%22:2}&t=" + t
	payload := strings.NewReader("")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-asw-login-token", token)
	req.Header.Add("x-asw-resource-commercial-version", "1")
	req.Header.Add("x-asw-resource-game-version", "1")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer response.Body.Close()              // 这步是必要的，防止以后的内存泄漏，切记
	body, _ := ioutil.ReadAll(response.Body) // 读取响应 body, 返回为 []byte
	var rsp EVResponse
	err = json.Unmarshal(body, &rsp)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	//fmt.Println(rsp.Data.Code)
	return rsp.Data.Code.Images, rsp.Data.Code.Articles
}
