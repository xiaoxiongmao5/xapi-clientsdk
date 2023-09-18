package client

import (
	"net/url"
)

//----------------------------------
// 网易云音乐随机歌曲 调用类
//----------------------------------

func (c *Client) Api_10(param, transinfo3 string) ([]byte, error) {
	interfaceId := "10"
	//请求地址
	juheURL := "https://api.uomg.com/api/rand.music"

	//初始化参数
	params := url.Values{}

	params.Set("sort", "热歌榜")    //选填 选择输出分类[热歌榜，新歌榜，飙升榜，抖音榜，电音榜]，为空输出热歌榜
	params.Set("format", "json") //选填 选择输出格式

	//发送请求
	return c.Get(interfaceId, juheURL, params)
}
