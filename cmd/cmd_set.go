package cmd

import (
	"math/rand"

	. "github.com/louch2010/dhaiy/common"
)

//SAdd处理
func HandleSAddCommand(client *Client) {
	request := []string{client.Reqest[0], client.Reqest[1], client.Reqest[2], ""}
	client.Reqest = request
	HandleHSetCommand(client)
}

//SCard处理
func HandleSCardCommand(client *Client) {
	HandleHLenCommand(client)
}

//SMembers处理
func HandleSMembersCommand(client *Client) {
	HandleHKeysCommand(client)
}

//SRem处理
func HandleSRemCommand(client *Client) {
	HandleHDelCommand(client)
}

//SisMember处理
func HandleSisMemberCommand(client *Client) {
	HandleHExistsCommand(client)
}

//SPop处理
func HandleSPopCommand(client *Client) {
	randSet(client, true)
}

//SRandMember处理
func HandleSRandMemberCommand(client *Client) {
	randSet(client, false)
}

func randSet(client *Client, del bool) {
	m := getMap(client)
	if m == nil {
		return
	}
	if len(m) == 0 {
		client.Response = GetCmdResponse(MESSAGE_ITEM_IS_EMPTY, "", ERROR_ITEM_IS_EMPTY, client)
		return
	}
	r := rand.Intn(len(m))
	i := 0
	for k, _ := range m {
		if i == r {
			if del {
				delete(m, k)
			}
			client.Response = GetCmdResponse(MESSAGE_SUCCESS, k, nil, client)
			return
		}
		i++
	}
	client.Response = GetCmdResponse(MESSAGE_ITEM_IS_EMPTY, "", ERROR_ITEM_IS_EMPTY, client)
}
