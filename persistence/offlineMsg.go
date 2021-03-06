package persistence
import(
	"time"
	"github.com/xiaogan18/msgserver/queue"
)

type MsgContainer interface{
	Get(id string) (*OfflineMsg,error)
	GetUserMsg(userID string) ([]*OfflineMsg,error)
	Put(*OfflineMsg)
}
type OfflineMsg struct{
	queue.Message
	KeepLiveTime time.Time
}

func CreateMsgContainer(t string) MsgContainer{
	switch(t){
	default:
		c:=MemoryContainer{
			msgMap:make(map[string]*OfflineMsg,0),
			userMsgMap:make(map[string][]string,0),
		}
		c.gc(1000*120)
		return &c
	}
}