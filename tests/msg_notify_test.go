package tests

import (
	"encoding/json"
	"fmt"
	"github.com/zhaoyunxing92/dingtalk/v2/model"
	"testing"
)

//发送工作通知
func TestSendTextWorkNotify(t *testing.T) {
	res := new(model.WorkNotify)
	res.NewTextWorkNotify("hello")
	res.UserIds = []string{"manager164"}
	notify, err := dingTalk.SendWorkNotify(res)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(notify)
}

func Benchmark_NewTextWorkNotify(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		res := new(model.WorkNotify)
		res.NewTextWorkNotify(fmt.Sprintf("test hello abc golang:%d", i))
		res.UserIds = []string{"manager164"}
		notify, err := dingTalk.SendWorkNotify(res)
		if err != nil {
			b.Fatal(err)
		}
		b.Log(notify)
	}
}

//TestSendOAWorkNotify:oa工作通知
func TestSendOAWorkNotify(t *testing.T) {

	oa := model.OA{}

	f := model.Form{}
	f.Key = "姓名:"
	f.Value = "赵云兴"

	f2 := model.Form{}
	f2.Key = "年龄:"
	f2.Value = "20"

	f3 := model.Form{}
	f3.Key = "身高:"
	f3.Value = "192"

	f4 := model.Form{}
	f4.Key = "体重:"
	f4.Value = "80KG"

	f5 := model.Form{}
	f5.Key = "爱好:"
	f5.Value = "go、java、docker、vue、吃饭"

	f6 := model.Form{}
	f6.Key = "测试:"
	f6.Value = "go、java、docker、vue"

	//设置体最多只有6个
	oa.Body.Forms = append(oa.Body.Forms, f, f2, f3, f4, f5, f6)
	oa.Body.Content = "validator用于对数据进行校验。在 Web 开发中，对用户传过来的数据我们都需要进行严格校验，防止用户的恶意请求。例如日期格式，用户年龄，性别等必须是正常的值，不能随意设置。"
	oa.Body.Title = "头部标题这会显示"
	oa.Body.Author = "赵云兴"
	oa.Body.ImageId = "@lALPDe7sx7z5xEJgzQJA"

	//设置头
	oa.Header.BgColor = "FFBBBBBB"
	oa.Header.Text = "小程序消息则不会显示"
    //pc跳转和小程序跳转
	oa.MessageUrl = "eapp://page/component/index"
	oa.PcMessageUrl = "https://ding-doc.dingtalk.com"
	// 设置状态
	oa.StatusBar.Value="处理完成"
	oa.StatusBar.BackColor="#FFF65E5E"

	res := new(model.WorkNotify)
	res.NewOAWorkNotify(oa)
	res.UserIds = []string{"manager164"}

	notify, err := dingTalk.SendWorkNotify(res)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", notify)
}

func TestSendOAWorkNotifyJson(t *testing.T) {

	oa := `{
     "msgtype": "oa",
     "oa": {
        "message_url": "http://dingtalk.com",
        "head": {
            "bgcolor": "FFBBBBBB",
            "text": "头部标题"
        },
        "body": {
            "title": "正文标题",
            "form": [
                {
                    "key": "姓名:",
                    "value": "张三"
                },
                {
                    "key": "年龄:",
                    "value": "20"
                },
                {
                    "key": "身高:",
                    "value": "1.8米"
                },
                {
                    "key": "体重:",
                    "value": "130斤"
                },
                {
                    "key": "学历:",
                    "value": "本科"
                },
                {
                    "key": "爱好:",
                    "value": "打球、听音乐"
                }
            ],
            "rich": {
                "num": "15.6",
                "unit": "元"
            },
            "content": "大段文本大段文本大段文本大段文本大段文本大段文本",
            "image": "@lADOADmaWMzazQKA",
            "file_count": "3",
            "author": "李四 "
        }
    }
}
`

	msg := new(model.OAMessage)
	err := json.Unmarshal([]byte(oa), msg)

	notify := new(model.WorkNotify)
	notify.Msg = msg

	notify.UserIds = []string{"manager164"}

	rep, err := dingTalk.SendWorkNotify(notify)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", rep)

}

func TestMarkdownWorkNotify(t *testing.T) {
	md := `>  这个是markdown，目前不知道要写啥，但是我明确的知道他支持渲染表格
# 这个是一级标题
## 二级标题测试图片

![钉钉开放平台](https://img.alicdn.com/tfs/TB1bB9QKpzqK1RjSZFoXXbfcXXa-576-96.png)

### 三级标题测试列表

* 吃饭
* 写bug

### 测试任务列表

- [x] 支持oa消息
- [x] 支持text消息
- [ ] 图片消息

### <u>下划线</u>
钉钉为企业和组织提供了很多基础办公应用例如审批、日志、视频会议等。企业可基于钉钉开放平台的能力，根据实际需要定制开发企业应用。钉钉开放平台提供了丰富的接口能力，例如通讯录管理、群会话管理、消息通知、智能工作流、待办任务、考勤等以满足企业多样的业务需求，同时基于开放平台的授权机制，降低开发者的对接成本和安全风险。`

	res := new(model.WorkNotify)
	res.UserIds = []string{"manager164"}
	res.NewMarkdownWorkNotify("markdown", md)

	rep, err := dingTalk.SendWorkNotify(res)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rep)
}

//TestCardWorkNotify:card消息
func TestCardWorkNotify(t *testing.T) {

	card := model.Card{}
	card.Title = "测试标题"
	card.SingleTitle = "详情"
	card.SingleUrl = "https://ding-doc.dingtalk.com/document#/"
	card.Content = "> 异常通知 \n * **推断场景**:同账号id命中名单变化未重新入审 \n * **异常指标**: 因以上原因导致应入审未入审的任务量"

	res := new(model.WorkNotify)
	res.UserIds = []string{"manager164"}
	res.NewCardMessage(card)

	progress, err := dingTalk.SendWorkNotify(res)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", progress)
}

//TestSendFileMessage:发送文件
func TestSendFileMessage(t *testing.T) {
	req := new(model.WorkNotify)
	req.UserIds = []string{"manager164"}
	req.NewFileMessage("@lALPDe7sx7z5xEJgzQJA")

	res, err := dingTalk.SendWorkNotify(req)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(res)
	t.Log(string(js))
}

func TestGetWorkNotifyProgress(t *testing.T) {

	progress, err := dingTalk.GetWorkNotifyProgress(292610197251)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", progress)
}

func TestGetWorkNotifyResult(t *testing.T) {

	progress, err := dingTalk.GetWorkNotifySendResult(292610197251)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", progress)
}

func TestRecallWorkNotifySendResult(t *testing.T) {

	progress, err := dingTalk.RecallWorkNotifySendResult(289041870984)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", progress)
}
