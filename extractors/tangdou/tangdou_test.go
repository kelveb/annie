package tangdou

import (
	"testing"

	"github.com/iawia002/annie/config"
	"github.com/iawia002/annie/downloader"
	"github.com/iawia002/annie/test"
)

func TestTangDou(t *testing.T) {
	config.InfoOnly = true
	config.ThreadNumber = 9
	tests := []struct {
		name     string
		args     test.Args
		playlist bool
	}{
		{
			name: "contains video URL test directly and can get title from body's div tag",
			args: test.Args{
				URL:   "http://www.tangdou.com/v95/dAOQNgMjwT2D5w2.html",
				Title: "杨丽萍广场舞《好日子天天过》喜庆双扇扇子舞",
				Size:  87611483,
			},
		},
		{
			name: "need call share url first and get the signed video URL test and can get title from head's title tag",
			args: test.Args{
				URL:   "http://m.tangdou.com/v94/dAOMMYNjwT1T2Q2.html",
				Title: "吉美广场舞《再唱山歌给党听》民族形体舞 附教学视频在线观看",
				Size:  50710318,
			},
		},
		{
			name: "share url",
			args: test.Args{
				URL:   "https://share.tangdou.com/play.php?vid=1500667821669",
				Title: "井岗紫薇广场舞18步双人舞《采槟榔》附分解",
				Size:  26693149,
			},
		},
		{
			name: "playlist test",
			args: test.Args{
				URL:   "http://www.tangdou.com/playlist/view/1882",
				Title: "青儿广场舞《小朋友们都被接走了》原创32步流行舞",
				Size:  69448816,
			},
			playlist: true,
		},
		{
			name: "playlist test2",
			args: test.Args{
				URL:   "http://www.tangdou.com/playlist/view/2816/page/4",
				Title: "茉莉广场舞 我向草原问个好 原创藏族风民族舞附教学",
				Size:  66284484,
			},
			playlist: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				data []downloader.Data
				err  error
			)
			if tt.playlist {
				// playlist mode
				config.Playlist = true
				_, err = Extract(tt.args.URL)
				test.CheckError(t, err)
			} else {
				config.Playlist = false
				data, err = Extract(tt.args.URL)
				test.CheckError(t, err)
				test.Check(t, tt.args, data[0])
			}
		})
	}
}
